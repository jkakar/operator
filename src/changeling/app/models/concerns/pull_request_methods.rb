# Helpers for handling pull request events for multipass approval
module PullRequestMethods
  extend ActiveSupport::Concern

  included do
  end

  # Things exposed to the included class as class methods
  module ClassMethods
    def new_from_pull_request_team(name, pull_request)
      new team: name,
          impact: "low",
          change_type: "minor",
          impact_probability: "medium",
          backout_plan: "We revert the pull request.", # https://i.imgur.com/27P3xPe.gif
          requester: User.for_github_login(pull_request["pull_request"]["user"]["login"]),
          reference_url: pull_request["pull_request"]["html_url"],
          title: pull_request["pull_request"]["title"],
          release_id: pull_request["pull_request"]["head"]["sha"]
    end

    def find_or_initialize_by_pull_request(pull_request)
      repo = Repository.find(pull_request["repository"]["full_name"])
      multipass = find_by(reference_url: pull_request["pull_request"]["html_url"])

      unless multipass
        multipass = new_from_pull_request_team(repo.team, pull_request)
      end
      multipass.release_id = pull_request["pull_request"]["head"]["sha"]
      multipass.title = pull_request["pull_request"]["title"]
      multipass
    end
  end

  def update_for_open_or_synchronize_pull_request(pull_request)
    return unless %w{opened synchronize}.include?(pull_request["action"])
    user = User.find_by(github_login: pull_request["sender"]["login"])
    Audited::Audit.as_user(user) do
      self.audit_comment = "API: Created from webhook #{pull_request['pull_request']['html_url']}"
      if pull_request["action"] == "opened"
        self.check_commit_statuses!
      else
        self.testing = false
      end
      self.save!
    end
  end

  def flag_merge_commits_as_successful(pull_request, merge_commit_sha)
    user = User.find_by(github_login: pull_request["sender"]["login"])
    Audited::Audit.as_user(user) do
      self.release_id = merge_commit_sha
      save
    end
  end
end
