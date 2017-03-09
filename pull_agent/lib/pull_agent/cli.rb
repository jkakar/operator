module PullAgent
  class CLI
    attr_reader :environment, :project

    def initialize(args = ARGV)
      @arguments = args
      parse_arguments!

      GlobalConfiguration.load(@environment).merge_into_environment
    end

    def checkin
      ENV["LOG_LEVEL"] = "7"
      Instrumentation.setup("pull-agent", @environment, log_stream: Logger)

      if @project == "chef"
        return checkin_chef
      end

      deploy = Canoe.latest_deploy(@environment, @project)
      Logger.context[:deploy_id] = deploy.id
      Logger.context[:project] = @project
      Logger.context[:environment] = @environment

      if deploy.applies_to_this_server?
        if deploy.action.nil?
          Logger.log(:debug, "Nothing to do for this deploy at this time")
        else
          Logger.log(:info, "Starting '#{deploy.action}' phase")
          deployer = DeployerRegistry.fetch(@project).new(@environment, deploy)
          case deploy.action
          when "deploy"
            deployer.deploy
          when "restart"
            deployer.restart
          else
            Logger.log(:error, "Unknown deployment action: #{@deploy.action}")
            return false
          end
          Logger.log(:info, "Finished '#{deploy.action}' phase")
          Canoe.notify_server(@environment, deploy)
        end
      else
        Logger.log(:debug, "The deploy does not apply to this server")
      end
    end

    def checkin_chef
      PullAgent::Deployers::Chef.new(@environment).deploy
    end

    def self.knife(args)
      if args.size < 2
        raise ArgumentError, "Usage: pull-agent-knife <environment> <command...>"
      end

      environment = args.shift
      hostname = ShellHelper.hostname
      datacenter =
        if environment == "dev"
          "local"
        else
          hostname.split("-")[3]
        end

      GlobalConfiguration.load(environment).merge_into_environment

      request = {
        payload: JSON.dump(
          command: args,
          server: {
            datacenter: datacenter,
            environment: environment,
            hostname: hostname
          }
        )
      }

      Canoe.knife(request)
    end

    private

    def parse_arguments!
      # environment and project (repository name) are required
      if @arguments.size != 2
        raise ArgumentError, usage
      else
        @environment, @project = @arguments
      end
    end

    def usage
      readme = File.expand_path("../../../../README.md", __FILE__)
      if File.exist?(readme)
        File.read(readme)
      else
        "Please refer to the README for usage information"
      end
    end
  end
end