# Worker for handling incoming pull requests webhooks
class PullRequestHandler < ActiveJob::Base
  def default_branch?(pull_request)
    base_branch    = pull_request["pull_request"]["base"]["ref"]
    default_branch = pull_request["repository"]["default_branch"]
    base_branch == default_branch
  end

  def perform(_, data)
    pull_request = JSON.parse(data)
    if Changeling.config.pardot?
      multipass = Multipass.find_or_initialize_by_pull_request(pull_request)
      multipass.save!
      multipass.synchronize
    else
      case pull_request["action"]
      when "opened", "synchronize"
        multipass = Multipass.find_or_initialize_by_pull_request(pull_request)
        multipass.update_for_open_or_synchronize_pull_request(pull_request)
      when "closed"
        return unless pull_request["pull_request"]["merged"] && default_branch?(pull_request)
        multipass = Multipass.find_or_initialize_by_pull_request(pull_request)

        commit_sha = pull_request["pull_request"]["merge_commit_sha"]
        multipass.flag_merge_commits_as_successful(pull_request, commit_sha)
      end
    end
  end
end