class TerraformProject
  NAME = "terraform".freeze

  ESTATES = [
    "aws/pardot",
    "aws/pardot-atlassian",
    "aws/pardot-ci",
    "aws/pardot-qe",
    "aws/pardotops",
    "aws/pardotpublic",
  ].freeze

  def self.find!(notifier = nil)
    notifier ||= HipchatNotifier.new
    project = Project.find_by!(name: NAME)

    new(project, notifier)
  end

  def self.create!
    Project.find_or_create_by!(name: NAME) do |project|
      project.icon = "server"
      project.bamboo_project = "BREAD"
      project.bamboo_plan = "BREAD"
      project.bamboo_job = "TER"
      project.repository = "Pardot/bread"
    end
  end

  def initialize(project, notifier)
    @project = project
    @notifier = notifier
  end

  def deploy(user, estate, build)
    unless ESTATES.include?(estate)
      return TerraformDeployResponse.unknown_estate(estate)
    end

    deploy = TerraformDeploy.transaction do
      deploys = TerraformDeploy.pending(estate)

      if deploys.count > 0
        return TerraformDeployResponse.locked(deploys.first!)
      end

      TerraformDeploy.create!(
        project_id: @project.id,
        auth_user_id: user.id,
        estate_name: estate,
        branch_name: build.branch,
        commit_sha1: build.commit,
        terraform_version: build.terraform_version
      )
    end

    notification.deploy_started(deploy)

    TerraformDeployResponse.success(deploy)
  end

  def complete_deploy(deploy_id, successful)
    deploy = TerraformDeploy.find_by(id: deploy_id)

    unless deploy
      return TerraformDeployResponse.new(nil, "No such deploy: #{deploy_id.inspect}")
    end

    unless deploy.completed_at.nil?
      return TerraformDeployResponse.new(deploy, "Deploy is already complete")
    end

    deploy.update!(completed_at: Time.current, successful: successful)
    notification.deploy_complete(deploy)

    TerraformDeployResponse.success(deploy)
  end

  private

  def notification
    TerraformNotification.new(@notifier, room_ids)
  end

  def room_ids
    @project.deploy_notifications.map { |n| n.hipchat_room_id }
  end
end
