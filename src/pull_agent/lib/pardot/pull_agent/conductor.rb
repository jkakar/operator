module Pardot
  module PullAgent
    class Conductor
      # attributes usually set by environment
      attr_reader :environment

      def initialize(environment)
        @environment = environment
        @no_questions = false
      end

      def deploy!(deploy)
        fetch_strategy = environment.current_fetch_strategy
        deploy_strategy = environment.current_deploy_strategy

        if deploy_strategy.rollback?(deploy)
          return rollback!(deploy, deploy_strategy)
        end

        unless fetch_strategy.valid?(deploy)
          invalid_fetch_warning(deploy)
          return false
        end

        environment.execute_pre_fetch_hooks(deploy)
        payload_path = fetch_strategy.fetch(deploy)
        exit_for_invalid_fetch_path if payload_path.empty?

        begin
          environment.execute_post_fetch_hooks(deploy)

          environment.execute_pre_deploy_hooks(deploy)
          if success = deploy_strategy.deploy(payload_path, deploy)
            environment.execute_post_deploy_hooks(deploy)
          end

          success
        ensure
          fetch_strategy.cleanup(deploy)
        end
      end

      def rollback!(deploy, deploy_strategy)
        environment.execute_pre_deploy_hooks(deploy)
        if success = deploy_strategy.rollback(deploy)
          environment.execute_post_deploy_hooks(deploy)
        end

        success
      end

      # Executes the restart phase of the deploy. Only performed on one server per
      # deploy.
      def restart!(deploy)
        environment.execute_restart_tasks(deploy)
      end

      def dont_ask!
        @no_questions = true
      end

      def silent?
        @no_questions
      end

      def invalid_fetch_warning(deploy)
        Logger.log(:err, "Requested deploy #{deploy.inspect} was not found")
      end

      def exit_for_invalid_fetch_path
        Logger.log(:err, "No local path available. Maybe check the environment definitions.")
        # TODO: Exiting from this deep into the system is bad design. Can we do something better?
        exit 1
      end
    end
  end
end