require "json"
require "replication_fixing/alerting_manager"
require "replication_fixing/fixing_client"
require "replication_fixing/fixing_status_client"
require "replication_fixing/hostname"
require "replication_fixing/ignore_client"
require "replication_fixing/message_throttler"
require "replication_fixing/monitor_supervisor"
require "replication_fixing/pagerduty_pager"
require "replication_fixing/replication_error_sanitizer"
require "replication_fixing/shard"
require "replication_fixing/test_pager"

module Lita
  module Handlers
    class ReplicationFixing < Handler
      config :repfix_url, default: "https://repfix.pardot.com"
      config :status_room, default: "1_ops@conf.btf.hipchat.com"
      config :replication_room, default: "1_ops-replication@conf.btf.hipchat.com"
      config :pager, default: "pagerduty"
      config :pagerduty_service_key

      http.get "/replication/_ping", :ping
      http.post "/replication/errors", :create_replication_error

      # http://rubular.com/r/Gz3fLQiR5L
      route /^ignore\s+(?<shard_id>\d+)(?:\s+(?:(?<prefix>db|whoisdb)|(?<minutes>\d+))(?:\s+(?<minutes>\d+))?)?/i, :create_ignore, command: true, help: {
        "ignore SHARD_ID" => "Ignores db-SHARD_ID for 15 minutes",
        "ignore SHARD_ID PREFIX" => "Ignores PREFIX-SHARD_ID for 15 minutes (PREFIX is, e.g., db or whoisdb)",
        "ignore SHARD_ID MINUTES" => "Ignores db-SHARD_ID for MINUTES minutes",
        "ignore SHARD_ID PREFIX MINUTES" => "Ignores PREFIX-SHARD_ID for MINUTES minutes",
      }

      route /^fix\s+(?<shard_id>\d+)(?:\s+(?<prefix>db|whoisdb))?/i, :create_fix, command: true, help: {
        "fix SHARD_ID" => "Attempts to fix db-SHARD_ID",
        "fix SHARD_ID PREFIX" => "Attempts to fix PREFIX-SHARD_ID (PREFIX is, e.g., db or whoisdb)",
      }

      route /^cancelfix\s+(?<shard_id>\d+)/, :cancel_fix, command: true, help: {
        "cancelfix SHARD_ID" => "Cancels the fix for SHARD_ID",
      }

      route /^resetignore\s+(?<shard_id>\d+)(?:\s+(?<prefix>db|whoisdb))?/i, :reset_ignore, command: true, help: {
        "resetignore SHARD_ID" => "Stops ignoring db-SHARD_ID",
        "resetignore SHARD_ID PREFIX" => "Stops ignoreing PREFIX-SHARD_ID (PREFIX is, e.g., db or whoisdb)",
      }

      route /^current(?:auto)?fixes/i, :current_fixes, command: true, help: {
        "currentfixes" => "Lists ongoing replication fixes",
      }

      route /^stopfixing/i, :stop_fixing, command: true, help: {
        "stopfixing" => "Globally pauses fixing of replication errors",
      }

      route /^startfixing/i, :start_fixing, command: true, help: {
        "startfixing" => "Globally starts fixing of replication errors",
      }

      route /^checkfixing/i, :check_fixing, command: true, help: {
        "startfixing" => "Reports whether fixing is globally enabled or disabled",
      }

      def initialize(robot)
        super

        @throttler = ::ReplicationFixing::MessageThrottler.new(robot: robot, redis: redis)
        @sanitizer = ::ReplicationFixing::ReplicationErrorSanitizer.new

        @pager = \
          case config.pager.to_s
          when "pagerduty"
            ::ReplicationFixing::PagerdutyPager.new(config.pagerduty_service_key)
          when "test"
            ::ReplicationFixing::TestPager.new
          else
            raise ArgumentError, "unknown pager type: #{config.pager.to_s}"
          end

        @alerting_manager = ::ReplicationFixing::AlertingManager.new(
          pager: @pager,
          log: log,
        )

        @ignore_client = ::ReplicationFixing::IgnoreClient.new(redis)
        @fixing_status_client = ::ReplicationFixing::FixingStatusClient.new(redis)
        @fixing_client = ::ReplicationFixing::FixingClient.new(
          repfix_url: config.repfix_url,
          fixing_status_client: @fixing_status_client,
          log: log,
        )
        @monitor_supervisor = ::ReplicationFixing::MonitorSupervisor.new(
          redis: redis,
          fixing_client: @fixing_client
        )

        @status_room = ::Lita::Source.new(room: config.status_room)
        @replication_room = ::Lita::Source.new(room: config.replication_room)
      end

      on(:connected) do
        robot.join(config.status_room)
        robot.join(config.replication_room)
      end

      def ping(request, response)
        response.status = 200
        response.body << ""
      end

      def create_replication_error(request, response)
        body = request.POST
        if body["hostname"]
          begin
            hostname = ::ReplicationFixing::Hostname.new(body["hostname"])
            shard = hostname.shard

            ignoring = @ignore_client.ignoring?(shard)
            if ignoring
              log.debug("Shard is ignored: #{shard}")

              count = @ignore_client.incr_skipped_errors_count
              if (count % 200).zero?
                @throttler.send_message(@status_room, "@here FYI: Replication fixing has been stopped, but I've seen about #{result.skipped_errors_count} go by.")
                @alerting_manager.notify_replication_disabled_but_many_errors
              end
            else
              error = body["error"].to_s
              unless error.empty?
                robot.send_message(@replication_room, "#{hostname}: #{body["error"]}")
              end

              mysql_last_error = body["mysql_last_error"].to_s
              unless mysql_last_error.empty?
                sanitized_error = @sanitizer.sanitize(mysql_last_error)
                robot.send_message(@replication_room, "#{hostname}: #{sanitized_error}")
              end

              result = @fixing_client.fix(shard: hostname)
              @alerting_manager.ingest_fix_result(shard_or_hostname: hostname, result: result)

              case result
              when ::ReplicationFixing::FixingClient::NoErrorDetected
                # This generally means there was an error, but it's not a replication statement issue
                @throttler.send_message(@status_room, "/me is noticing a potential issue with #{hostname}: #{body["error"]}")
              else
                reply_with_fix_result(shard: shard, result: result)
                ensure_monitoring(shard: shard)
              end
            end

            response.status = 201
          rescue ::ReplicationFixing::Hostname::MalformedHostname
            response.status = 400
            response.body << JSON.dump("error" => "malformed hostname")
          end
        else
          response.status = 400
          response.body << JSON.dump("error" => "hostname missing")
        end
      end

      def create_ignore(response)
        shard_id = response.match_data["shard_id"].to_i
        prefix = response.match_data["prefix"] || "db"
        minutes = (response.match_data["minutes"] || "15").to_i

        shard = ::ReplicationFixing::Shard.new(prefix, shard_id)
        begin
          @ignore_client.ignore(shard, expire: minutes*60)
          response.reply_with_mention("OK, I will ignore #{shard} for #{minutes} minutes")
        rescue => e
          response.reply_with_mention("Sorry, something went wrong: #{e}")
        end
      end

      def create_fix(response)
        shard_id = response.match_data["shard_id"].to_i
        prefix = response.match_data["prefix"] || "db"
        shard = ::ReplicationFixing::Shard.new(prefix, shard_id)

        result = @fixing_client.fix(shard: shard)

        case result
        when ::ReplicationFixing::FixingClient::NoErrorDetected
          response.reply_with_mention "I didn't detect any errors detected on #{shard}"
        when ::ReplicationFixing::FixingClient::NotFixable
          response.reply_with_mention "Sorry, I'm not able to fix #{shard} right now. I need a human to resolve it."
        when ::ReplicationFixing::FixingClient::FixInProgress
          ongoing_minutes = ((Time.now - result.started_at) / 60.0).to_i
          if ongoing_minutes <= 0
            response.reply_with_mention "OK, I'm trying to fix #{shard}"
          else
            response.reply_with_mention "Hmm, I've already been trying to fix #{shard} for #{ongoing_minutes.to_i} minutes now"
          end
        when ::ReplicationFixing::FixingClient::ErrorCheckingFixability
          response.reply_with_mention "Sorry, I got an error while checking fixability: #{result.error}"
        else
          response.reply_with_mention "Sorry, I got an unknown result: #{result}"
        end

        ensure_monitoring(shard: shard)
      end

      def cancel_fix(response)
        shard_id = response.match_data["shard_id"].to_i
        prefix = "db" # TODO: Apparently there is no way to cancel a fix on a specific prefix in rep_fix
        shard = ::ReplicationFixing::Shard.new(prefix, shard_id)

        result = @fixing_client.cancel(shard: shard)

        if result.success?
          response.reply_with_mention "OK, I cancelled all the fixes for #{shard}"
        else
          response.reply_with_mention "Sorry, I wasn't able to cancel the fixes for #{shard}: #{result.message}"
        end
      end

      def reset_ignore(response)
        shard_id = response.match_data["shard_id"].to_i
        prefix = response.match_data["prefix"] || "db"

        shard = ::ReplicationFixing::Shard.new(prefix, shard_id)
        begin
          @ignore_client.reset_ignore(shard)
          response.reply_with_mention("OK, I will no longer ignore #{shard}")
        rescue => e
          response.reply_with_mention("Sorry, something went wrong: #{e}")
        end
      end

      def current_fixes(response)
        begin
          fixes = @fixing_status_client.current_fixes
          if fixes.length > 0
            response.reply_with_mention("I'm currently fixing: #{fixes.map { |f| f.shard.to_s }.join(", ")}")
          else
            response.reply_with_mention("I'm not fixing anything right now")
          end
        rescue => e
          response.reply_with_mention("Sorry, something went wrong: #{e}")
        end
      end

      def stop_fixing(response)
        begin
          @ignore_client.ignore_all
          response.reply_with_mention("OK, I've stopped fixing replication for ALL shards")
        rescue => e
          response.reply_with_mention("Sorry, something went wrong: #{e}")
        end
      end

      def start_fixing(response)
        begin
          @ignore_client.reset_ignore_all
          response.reply_with_mention("OK, I've started fixing replication")
        rescue => e
          response.reply_with_mention("Sorry, something went wrong: #{e}")
        end
      end

      def check_fixing(response)
        begin
          if @ignore_client.ignoring_all?
            response.reply_with_mention("(nope) Replication fixing is globally disabled")
          else
            response.reply_with_mention("(goodnews) Replication fixing is globally enabled")
          end
        rescue => e
          response.reply_with_mention("Sorry, something went wrong: #{e}")
        end
      end

      private
      def reply_with_fix_result(shard:, result:)
        ignoring = @ignore_client.ignoring?(shard)
        if ignoring
          log.debug("Shard is ignored: #{shard}")
        else
          case result
          when ::ReplicationFixing::FixingClient::NoErrorDetected
            @throttler.send_message(@status_room, "(successful) Replication is fixed on #{shard}")
          when ::ReplicationFixing::FixingClient::NotFixable
            @throttler.send_message(@status_room, "@all Replication is broken on #{shard}, but I'm not able to fix it")
          when ::ReplicationFixing::FixingClient::FixInProgress
            ongoing_minutes = (Time.now - result.started_at) / 60.0
            if ongoing_minutes >= 10.0
              @alerting_manager.notify_fixing_a_long_while(shard: shard, started_at: result.started_at)
              @throttler.send_message(@status_room, "@all I've been trying to fix replication on #{shard} for #{ongoing_minutes.to_i} minutes now")
            else
              @throttler.send_message(@status_room, "/me is fixing replication on #{shard} (ongoing for #{ongoing_minutes.to_i} minutes)")
            end
          when ::ReplicationFixing::FixingClient::FixableErrorOccurring
            @throttler.send_message(@status_room, "/me is noticing a fixable replication error on #{shard}")
          when ::ReplicationFixing::FixingClient::ErrorCheckingFixability
            @throttler.send_message(@status_room, "/me is getting an error while trying to check the fixability of #{shard}: #{result.error}")
          else
            log.error("Got unknown response from client: #{result}")
          end
        end
      end

      def ensure_monitoring(shard:)
        monitor = ::ReplicationFixing::Monitor.new(shard: shard, tick: 30)
        monitor.on_tick do |result|
          reply_with_fix_result(shard: shard, result: result)
        end
        monitor.on_replication_fixed do |result|
          begin
            @fixing_status_client.reset_status(shard: shard)
          rescue => e
            log.error("Unable to reset status: #{e}")
          end
        end

        @monitor_supervisor.start_exclusive_monitor(monitor)
      end

      Lita.register_handler(self)
    end
  end
end
