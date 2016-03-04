require "thread"

module Zabbix
  class MaintenanceSupervisor
    REDIS_NAMESPACE = "maintenance_supervisor"
    REDIS_EXPIRATIONS_KEY = [REDIS_NAMESPACE, "maintenance_expirations"].join(":")

    GLOBAL_MUTEX = Mutex.new

    def self.get_or_create(datacenter:, redis:, client:, log:)
      if @supervisors && supervisor = @supervisors[datacenter]
        supervisor
      else
        GLOBAL_MUTEX.synchronize do
          @supervisors ||= {}
          @supervisors[datacenter] = new(redis: redis, client: client, log: log)
        end
      end
    end

    def initialize(redis:, client:, log:)
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
        @redis.hset(REDIS_EXPIRATIONS_KEY, host, until_time.to_i)
      end
    end

    def stop_maintenance(host:)
      if @client.ensure_host_not_in_zabbix_maintenance_group(host)
        @redis.hdel(REDIS_EXPIRATIONS_KEY, host) > 0
      end
    rescue ::Zabbix::Client::HostNotFound
      @redis.hdel(REDIS_EXPIRATIONS_KEY, host) > 0
      raise
    end

    def run_expirations(now: Time.now)
      expired = @redis.hgetall(REDIS_EXPIRATIONS_KEY).select { |k, v| v.to_i <= now.to_i }.keys
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

    def try_supervise
      if @supervising_lock.try_lock
        begin
          loop do
            begin
              run_expirations
            rescue => e
              @log.error("Error while running expirations: #{e}")
            end

            sleep 10
          end
        ensure
          @supervising_lock.unlock
        end
      else
        @log.debug("Supervisor already executing")
      end
    end
  end
end
