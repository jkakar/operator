require "thread"

module Zabbix
  class MaintenanceSupervisor
    REDIS_NAMESPACE = "maintenance_supervisor"

    GLOBAL_MUTEX = Mutex.new

    attr_accessor :on_host_maintenance_expired

    def self.get_or_create(datacenter:, redis:, client:, log:)
      if @supervisors && supervisor = @supervisors[datacenter]
        supervisor
      else
        GLOBAL_MUTEX.synchronize do
          @supervisors ||= {}
          @supervisors[datacenter] = new(datacenter: datacenter, redis: redis, client: client, log: log)
        end
      end
    end

    def initialize(datacenter:, redis:, client:, log:)
      @datacenter = datacenter
      @redis = redis
      @client = client
      @log = log

      @supervising_lock = Mutex.new
    end

    def ensure_supervising
      Thread.new { try_supervise }
      true
    end

    def start_maintenance(host:, until_time:)
      if @client.ensure_host_in_zabbix_maintenance_group(host)
        @redis.hset(redis_expirations_key, host, until_time.to_i)
      end
    end

    def stop_maintenance(host:)
      if @client.ensure_host_not_in_zabbix_maintenance_group(host)
        @redis.hdel(redis_expirations_key, host) > 0
      end
    rescue ::Zabbix::Client::HostNotFound
      @redis.hdel(redis_expirations_key, host)
      raise
    end

    def run_expirations(now: Time.now)
      expired = @redis.hgetall(redis_expirations_key).select { |k, v| v.to_i <= now.to_i }.keys
      expired.select { |host|
        begin
          stop_maintenance(host: host)
          @log.info("Brought host out of maintenance: #{host}")
        rescue ::Zabbix::Client::HostNotFound
          @log.warn("Host not found while removing it from maintenance: #{host}")
          false
        rescue => e
          @log.error("Error while removing host from maintenance: #{e}")
          false
        end
      }
    end

    private

    def redis_expirations_key
      [REDIS_NAMESPACE, "maintenance_expirations", @datacenter].join(":")
    end

    def try_supervise
      if @supervising_lock.try_lock
        begin
          loop do
            expirations = \
              begin
                run_expirations
              rescue => e
                @log.error("Error while running expirations: #{e}")
                []
              end

            expirations.each { |host| notify_host_maintenance_expired(host) }
            sleep 10
          end
        ensure
          @supervising_lock.unlock
        end
      else
        @log.debug("Supervisor already executing")
      end
    end

    def notify_host_maintenance_expired(host)
      on_host_maintenance_expired.call(host) if on_host_maintenance_expired
    rescue => e
      @log.error("Error notifying host maintenance expired: #{e}")
    end
  end
end
