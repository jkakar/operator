require "syslog"

module Pardot
  module PullAgent
    module Logger
      PRIORITIES = {
        debug: Syslog::LOG_DEBUG,
        info: Syslog::LOG_INFO,
        notice: Syslog::LOG_NOTICE,
        warn: Syslog::LOG_WARNING,
        warning: Syslog::LOG_WARNING,
        alert: Syslog::LOG_ALERT,
        err: Syslog::LOG_ERR,
        error: Syslog::LOG_ERR,
        crit: Syslog::LOG_CRIT,
      }.freeze

      class Context
        include Enumerable

        def initialize
          @values = {}
        end

        def [](key)
          @values[key]
        end

        def []=(key, value)
          @values[key] = value
        end

        def each(&blk)
          @values.each(&blk)
        end

        def to_s
          @values.map { |k, v| "#{k}=#{v}" }.join(" ")
        end
      end

      def self.context
        @context ||= Context.new
      end

      def self.log(our_priority, message)
        context_str = context.to_s
        message = "[#{context_str}] #{message}" unless context_str.empty?

        puts "[%s] %s" % [our_priority, message] unless ENV['CRON']

        Syslog.open("pull-agent") do
          Syslog.log(PRIORITIES.fetch(our_priority), message)
        end
      end
    end
  end
end
