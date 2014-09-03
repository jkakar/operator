module Canoe
  module DeployLogic
    DEPLOYLOGIC_ERROR_NO_REPO   = 1
    DEPLOYLOGIC_ERROR_NO_TARGET = 2
    DEPLOYLOGIC_ERROR_NO_WHAT   = 3
    DEPLOYLOGIC_ERROR_UNABLE_TO_DEPLOY = 4
    DEPLOYLOGIC_ERROR_INVALID_WHAT = 5
    DEPLOYLOGIC_ERROR_DUPLICATE = 6

    # ----------------------------------------------------------------------
    def deploy!
      # require a repo and target
      return { error: true, reason: DEPLOYLOGIC_ERROR_NO_REPO   } if !current_repo
      return { error: true, reason: DEPLOYLOGIC_ERROR_NO_TARGET } if !current_target
      # confirm user can deploy
      if !current_target.user_can_deploy?(current_user)
        return { error: true, reason: DEPLOYLOGIC_ERROR_UNABLE_TO_DEPLOY }
      end
      # confirm again there is no active deploy
      if !current_target.active_deploy.nil?
        return { error: true, reason: DEPLOYLOGIC_ERROR_DUPLICATE }
      end

      deploy_options = { user: current_user,
                         repo: current_repo,
                         lock: (params[:lock] == "on"),
                       }

      # let's determine what we're deploying...
      %w[tag branch commit].each do |type|
        if params[type]
          deploy_options[:what] = type
          deploy_options[:what_details] = params[type]
          break
        end
      end

      # gather any servers that might be specified
      if params[:servers] == "on" && !params[:server_names].blank?
        deploy_options[:servers] = params[:server_names]
      end

      # validate that what we are deploying was included and is a real thing
      return { error: true, reason: DEPLOYLOGIC_ERROR_NO_WHAT } if deploy_options[:what].nil?
      if !valid_what?(deploy_options[:what], deploy_options[:what_details])
        return { error: true, reason: DEPLOYLOGIC_ERROR_INVALID_WHAT, what: deploy_options[:what] }
      end

      the_deploy = current_target.deploy!(deploy_options)
      if the_deploy
        { error: false, deploy: the_deploy }
      else
        # likely cause of nil response is a duplicate deploy (another guard)
        { error: true, reason: DEPLOYLOGIC_ERROR_DUPLICATE }
      end
    end

    # silly generic naming... heh
    def valid_what?(what, what_details)
      case what
      when "tag"
        tags = tags_for_current_repo
        tags.collect(&:name).include?(what_details)
      when "branch"
        branches = branches_for_current_repo
        branches.collect(&:name).include?(what_details)
      when "commit"
        commits = commits_for_current_repo
        found_commits = commits.select do |commit|
          commit.sha.match(%r{^#{what_details}})
        end
        !found_commits.empty?
      else
        false
      end
    end

  end
end
