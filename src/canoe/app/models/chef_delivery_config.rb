class ChefDeliveryConfig
  PRODUCTION = "production".freeze
  DEV = "dev".freeze

  AWS = "ue1.aws".freeze
  DFW = "dfw".freeze
  PHX = "phx".freeze

  ENABLED = [AWS, PHX, DFW].freeze

  BREAD_ROOM = 42
  OPS_ROOM = 6

  def enabled?(server)
    case server.environment
    when DEV
      true
    when PRODUCTION
      ENABLED.include?(server.datacenter)
    else
      false
    end
  end

  def knife_notifications_enabled?(_server)
    true
  end

  def repo_name
    ENV.fetch("CANOE_CHEF_REPO", "Pardot/chef")
  end

  def master_branch
    "master"
  end

  def max_lock_age
    30.minutes
  end

  def github_url
    Project::GITHUB_URL
  end

  def github_repo
    @github_repo ||= GithubRepository.new(
      Octokit::Client.new(
        api_endpoint: "#{github_url}/api/v3",
        access_token: github_token
      ),
      repo_name
    )
  end

  def chat_room_id(server)
    case server.datacenter
    when AWS
      BREAD_ROOM
    else
      Integer(ENV.fetch("CANOE_CHEF_CHAT_ROOM_ID", OPS_ROOM))
    end
  end

  class HipchatNotifier
    def self.notify_room(room_id, message, color = nil)
      Hipchat.notify_room(room_id, message, false, color)
    end
  end

  def notifier
    @notifier ||= HipchatNotifier
  end

  private

  def github_token
    ENV.fetch("GITHUB_PASSWORD")
  end
end
