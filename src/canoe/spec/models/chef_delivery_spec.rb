require "rails_helper"

RSpec.describe ChefDelivery do
  before(:each) do
    @repo = GithubRepository::Fake.new
    @config = FakeChefDeliveryConfig.new(@repo)
    @delivery = ChefDelivery.new(@config)
  end

  def build_build(attributes = {})
    defaults = {
      url: "https://github.com/builds/1",
      sha: "sha1",
      branch: "master",
      state: ChefDelivery::SUCCESS,
      updated_at: Time.current
    }
    GithubRepository::Build.new(defaults.merge(attributes))
  end

  def create_current_deploy(attributes = {})
    defaults = {
      branch: "master",
      build_url: "https://github/builds/1",
      datacenter: "test",
      environment: "testing",
      hostname: "test0-chef1-1",
      sha: "sha1",
      state: ChefDelivery::SUCCESS
    }
    ChefDeploy.create!(defaults.merge(attributes))
  end

  it "noops if chef delivery is disabled in current environment" do
    server = ChefDelivery::Server.new("dfw", "disabled", "chef1")
    checkout = ChefCheckinRequest::Checkout.new("sha1", "master")
    request = ChefCheckinRequest.new(server, checkout)
    response = @delivery.checkin(request)
    assert_equal "noop", response.action
  end

  it "noops if chef delivery is not enabled for the server" do
    server = ChefDelivery::Server.new("test", "production", "disabled")
    checkout = ChefCheckinRequest::Checkout.new("sha1", "master")
    request = ChefCheckinRequest.new(server, checkout)
    response = @delivery.checkin(request)
    assert_equal "noop", response.action
  end

  it "noops if there is no available build" do
    server = ChefDelivery::Server.new("test", "production", "chef1")
    checkout = ChefCheckinRequest::Checkout.new("sha1", "master")
    request = ChefCheckinRequest.new(server, checkout)
    @repo.current_build = GithubRepository::Build.none
    response = @delivery.checkin(request)
    assert_equal "noop", response.action
  end

  it "noops if the build is red" do
    server = ChefDelivery::Server.new("test", "production", "chef1")
    checkout = ChefCheckinRequest::Checkout.new("sha1", "master")
    request = ChefCheckinRequest.new(server, checkout)
    @repo.current_build = build_build(state: ChefDelivery::FAILURE)
    response = @delivery.checkin(request)
    assert_equal "noop", response.action
  end

  it "noops if the build is pending" do
    server = ChefDelivery::Server.new("test", "production", "chef1")
    checkout = ChefCheckinRequest::Checkout.new("sha1", "master")
    request = ChefCheckinRequest.new(server, checkout)
    @repo.current_build = build_build(state: ChefDelivery::PENDING)
    response = @delivery.checkin(request)
    assert_equal "noop", response.action
  end

  it "noops if the build is not for the master branch" do
    server = ChefDelivery::Server.new("test", "production", "chef1")
    checkout = ChefCheckinRequest::Checkout.new("sha1", "master")
    request = ChefCheckinRequest.new(server, checkout)
    @repo.current_build = build_build(
      branch: "this-is-fine-dot-jpg",
      state: ChefDelivery::PENDING
    )
    response = @delivery.checkin(request)
    assert_equal "noop", response.action
  end

  it "noops and notifies once every 30 minutes if non-master branch is checked out" do
    create_current_deploy(state: ChefDelivery::SUCCESS, sha: "sha^^^")
    server = ChefDelivery::Server.new("test", "production", "pardot0-chef1")
    checkout = ChefCheckinRequest::Checkout.new("sha1", "mybranch")
    request = ChefCheckinRequest.new(server, checkout)
    @repo.current_build = build_build(sha: "deadbeef")
    response = @delivery.checkin(request, Time.current)
    assert_equal "noop", response.action
    assert_equal 1, @config.notifier.messages.size
    msg = @config.notifier.messages.pop
    assert msg.message.include?("could not be deployed")
    assert msg.message.include?("mybranch")
    assert msg.message.include?("pardot0-chef1")

    response = @delivery.checkin(request, Time.current + 15.minutes)
    assert_equal 0, @config.notifier.messages.size
    assert_equal "noop", response.action

    response = @delivery.checkin(request, Time.current + 40.minutes)
    assert_equal 1, @config.notifier.messages.size
    assert_equal "noop", response.action
  end

  it "noops if there is no build available" do
    server = ChefDelivery::Server.new("test", "production", "pardot0-chef1")
    checkout = ChefCheckinRequest::Checkout.new("sha1", "master")
    request = ChefCheckinRequest.new(server, checkout)
    @repo.current_build = GithubRepository::Build.none
    response = @delivery.checkin(request)
    assert_equal "noop", response.action
  end

  it "noops if current deploy is pending" do
    server = ChefDelivery::Server.new("test", "production", "pardot0-chef1")
    checkout = ChefCheckinRequest::Checkout.new("sha1", "master")
    request = ChefCheckinRequest.new(server, checkout)
    @repo.current_build = build_build
    create_current_deploy(state: ChefDelivery::PENDING)
    response = @delivery.checkin(request)
    assert_equal "noop", response.action
  end

  it "noops if current sha1 is already deployed" do
    server = ChefDelivery::Server.new("test", "production", "pardot0-chef1")
    checkout = ChefCheckinRequest::Checkout.new("sha1", "master")
    request = ChefCheckinRequest.new(server, checkout)
    @repo.current_build = build_build(sha: "sha1")
    create_current_deploy(state: ChefDelivery::SUCCESS, sha: "sha1")
    response = @delivery.checkin(request)
    assert_equal "noop", response.action
  end

  it "deploys if the checkout differs from the current build" do
    server = ChefDelivery::Server.new("test", "production", "pardot0-chef1")
    checkout = ChefCheckinRequest::Checkout.new("sha1", "master")
    request = ChefCheckinRequest.new(server, checkout)
    @repo.current_build = build_build(sha: "sha2")
    response = @delivery.checkin(request)
    assert_equal "deploy", response.action
  end

  it "deploys the same build twice" do
    server = ChefDelivery::Server.new("test", "production", "pardot0-chef1")
    checkout = ChefCheckinRequest::Checkout.new("sha1", "master")
    request = ChefCheckinRequest.new(server, checkout)
    @repo.current_build = build_build(sha: "sha1^^")
    response = @delivery.checkin(request)
    assert_equal "deploy", response.action
    request = ChefCompleteDeployRequest.new(response.deploy.id, true, nil)
    @delivery.complete_deploy(request)

    checkout = ChefCheckinRequest::Checkout.new("sha1~100", "master")
    request = ChefCheckinRequest.new(server, checkout)
    response = @delivery.checkin(request)
    assert_equal "deploy", response.action
  end

  it "notifies of successful deployment" do
    deploy = create_current_deploy(
      state: ChefDelivery::PENDING,
      build_url: "https://BREAD-9000"
    )
    request = ChefCompleteDeployRequest.new(deploy.id, true, nil)
    @delivery.complete_deploy(request)
    assert_equal 1, @config.notifier.messages.size
    msg = @config.notifier.messages.pop
    assert msg.message.include?("successfully deployed")
    assert msg.message.include?("#9000")
  end

  it "notifies of failed deployment" do
    deploy = create_current_deploy(state: ChefDelivery::PENDING)
    request = ChefCompleteDeployRequest.new(deploy.id, false, "boomtown")
    @delivery.complete_deploy(request)
    assert_equal 1, @config.notifier.messages.size
    msg = @config.notifier.messages.pop
    assert msg.message.include?("failed to deploy")
    assert msg.message.include?("boomtown")
  end

  it "notifies of executed knife commands" do
    server = ChefDelivery::Server.new("dfw", "dev", "chef1")
    command = %w[environment from file fail.rb]
    request = KnifeRequest.new(server, command)
    @delivery.knife(request)
    assert_equal 1, @config.notifier.messages.size
    msg = @config.notifier.messages.pop
    assert_includes msg.message, "dfw/dev"
    assert_includes msg.message, "knife #{command.join(" ")}"
  end

  it "ignores 'knife node from file' commands" do
    server = ChefDelivery::Server.new("dfw", "dev", "chef1")
    command = %w[node from file nodes/aws/node.json]
    request = KnifeRequest.new(server, command)
    @delivery.knife(request)
    assert_equal 0, @config.notifier.messages.size
  end

  it "ignores knife help commands" do
    server = ChefDelivery::Server.new("dfw", "dev", "chef1")
    command = %w[help list]
    request = KnifeRequest.new(server, command)
    @delivery.knife(request)
    assert_equal 0, @config.notifier.messages.size
  end

  it "ignores 'knife pd sync' command" do
    server = ChefDelivery::Server.new("dfw", "dev", "chef1")
    command = %w[pd sync]
    request = KnifeRequest.new(server, command)
    @delivery.knife(request)
    assert_equal 0, @config.notifier.messages.size
  end

  it "ignores 'knife search' commands" do
    server = ChefDelivery::Server.new("dfw", "dev", "chef1")
    command = %w[search hi andy]
    request = KnifeRequest.new(server, command)
    @delivery.knife(request)
    assert_equal 0, @config.notifier.messages.size
  end
end
