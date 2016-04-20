require 'securerandom'

module Zabbix
  class Zabbixmon

    MONITOR_NAME = "zabbixmon"
    MONITOR_SHORTHAND = "zbxmon"
    INCIDENT_KEY = "#{MONITOR_NAME}-%s"
    ERR_NON_200_HTTP_CODE = "HAL9000 HTTP'd Zabbix, but the host failed to respond to an HTTP request with the appropriate status code (! HTTP 200)"
    ERR_ZBX_CLIENT_EXCEPTION = "HAL9000 attempted to use the ZabbixApi client, but an exception was thrown/handled: exception"
    ZABBIX_ITEM_NOT_FOUND = "HAL9000 searched for an iteam w/ a particluar key and value, but did not find it. This is bad."
    ZBXMON_TEST_API_ENDPOINT = 'cgi-bin/zabbix-server-check.sh'
    ZBXMON_ITEM = 'system:general'
    ZBXMON_KEY = 'zabbix_status'
    ZBXMON_PAYLOAD_LENGTH = 10
    
    def initialize(redis:, zbx_client:, log:, zbx_host:, zbx_username:, zbx_password:, datacenter:)
      @redis = redis
      @client = zbx_client
      @log = log
      @zbx_host = zbx_host
      @zbx_username = zbx_username
      @zbx_password = zbx_password
      @datacenter = datacenter
      @hard_failure = nil
    end

    attr_accessor :hard_failure

    # assumes not paused (pausing handled by supervisor and handler and prevents this call)
    def monitor(num_retries = 5, retry_interval_seconds = 5, timeout_seconds = 30)
      retry_attempt_iterator = 0
      retry_sz = "retry attempt #{(retry_attempt_iterator + 1)} / #{num_retries}"
      payload = "#{SecureRandom.urlsafe_base64(ZBXMON_PAYLOAD_LENGTH)}" # make a per-use random string
      url = "https://#{@zbx_host}/#{ZBXMON_TEST_API_ENDPOINT}?#{payload}"
      monitor_success = false
      soft_failures = Set.new [] # soft-fails can used to provide feedback for hard-fail

      @log.debug("[#{monitor_name}] value generated: #{payload}")
      payload_delivery_response_code = deliver_zabbixmon_payload url, timeout_seconds
      if payload_delivery_response_code =~ /20./
        @log.debug("[#{monitor_name}] Monitor Payload Delivered Successfully")
      else
        @hard_failure = "ZabbixMon[#{@datacenter}] payload insertion failed! : #{ERR_NON_200_HTTP_CODE}"
        @log.error("[#{monitor_name}] ZabbixMon[#{@datacenter}] payload insertion failed! ERROR = '#{ERR_NON_200_HTTP_CODE}'")
      end

      while (retry_attempt_iterator < num_retries) && (@hard_failure.nil?) && (!monitor_success) do
        # the state reported back from this loop is important! soft_fail = keep trying; hard_fail = stop and notify
        sleep retry_interval_seconds
        begin # get zabbix item
          apiresponse = @client.get_item_by_key_and_lastvalue(ZBXMON_KEY, payload)
          zbx_items = apiresponse['result']
          @log.debug("[#{monitor_name}] zabbix client 'got_item' successfully")
        rescue => e
          @log.error("[#{monitor_name}] #{ERR_ZBX_CLIENT_EXCEPTION}".gsub('%exception%', e))
          soft_failures.add("#{ERR_ZBX_CLIENT_EXCEPTION}".gsub('%exception%', e))
        end rescue StandardError # consume the exception and continue; zbx_items being nil is okay from here
        if zbx_items
          if zbx_items.length > 0  # success case
            @log.info("[#{monitor_name}] successfully observed '#{ZBXMON_KEY} : #{payload}' from #{@zbx_host} (#{retry_sz})")
            monitor_success = true
          else # fail case
            soft_failures.add("#{ZABBIX_ITEM_NOT_FOUND}")
          end
        else # fail case
          soft_failures.add("#{ZABBIX_ITEM_NOT_FOUND}")
        end
        @log.warn("[#{monitor_name}] FAILED to find #{ZBXMON_KEY} : #{payload} from the zabbix 'item' (#{retry_sz})") unless monitor_success
        retry_attempt_iterator += 1
      end

      # work is done! Establish pass/fail here
      if monitor_success
        @log.info("[#{monitor_name}] 's work is done here. There is no issue to report. (successkid)")
        @hard_failure = nil
      else
        @hard_failure ||= soft_failures.to_a.join('; ')
        @log.error("[#{monitor_name}] has hard failed: #{@hard_failure} ")
      end
    end

    def monitor_name
      MONITOR_NAME
    end

    private
    def deliver_zabbixmon_payload(url, timeout_seconds = 30)
      @log.debug("[#{monitor_name}] deliver_zabbixmon_payload url = #{url}")
      uri = URI(url)
      req = Net::HTTP::Get.new(uri)
      req.basic_auth @zbx_username, @zbx_password
      res = Net::HTTP.start(uri.hostname, uri.port, :read_timeout => timeout_seconds) {|http|
        http.request(req)
      }
      res.code
    rescue Timeout::Error
      @log.error("[#{monitor_name}] HTTP TIMEOUT while attempting to insert payload")
    rescue ::Lita::Handlers::Zabbix::MonitorDataInsertionFailed
      @log.error("[#{monitor_name}] has hard failed: ::Lita::Handlers::Zabbix::MonitorDataInsertionFailed")
    end
  end
end