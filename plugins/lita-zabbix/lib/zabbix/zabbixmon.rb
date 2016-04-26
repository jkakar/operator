require 'securerandom'

module Zabbix
  class Zabbixmon

    MONITOR_NAME = "zabbixmon"
    MONITOR_SHORTHAND = "zbxmon"
    INCIDENT_KEY = "#{MONITOR_NAME}-%s"
    ERR_NON_200_HTTP_CODE = "HAL9000 HTTP'd Zabbix, but the host failed to respond to an HTTP request with the appropriate status code (! HTTP 200)"
    ERR_ZBX_CLIENT_EXCEPTION = "HAL9000 attempted to use the ZabbixApi client, but an exception was thrown/handled: exception"
    ZABBIX_ITEM_NOT_FOUND = "HAL9000 failed to find the payload via the API"
    ZBXMON_KEY = "Zabbix_status"
    ZBXMON_PAYLOAD_LENGTH = 20
    
    def initialize(redis:, zbx_client:, log:, zbx_host:, zbx_username:, zbx_password:, datacenter:)
      @redis = redis
      @client = zbx_client
      @log = log
      @zbx_host = zbx_host
      @zbx_username = zbx_username
      @zbx_password = zbx_password
      @datacenter = datacenter
      @hard_failure = nil
      @soft_failures = Set.new [] # soft-fails can used to provide feedback for hard-fail
    end

    attr_accessor :hard_failure

    # assumes not paused (pausing handled by supervisor and handler and prevents this call)
    def monitor(url, num_retries = 5, retry_interval_seconds = 5, timeout_seconds = 30)
      retry_attempt_iterator = 0
      retry_sz = "retry attempt #{(retry_attempt_iterator + 1)} / #{num_retries}"
      payload = "#{SecureRandom.urlsafe_base64(ZBXMON_PAYLOAD_LENGTH)}" # make a per-use random string
      monitor_success = false
      insert_payload(payload, url, timeout_seconds)

      while (retry_attempt_iterator < num_retries) && (@hard_failure.nil?) && (!monitor_success) do
        # the state reported back from this loop is important! soft_fail = keep trying; hard_fail = stop and notify
        sleep retry_interval_seconds
        @log.debug("[#{monitor_name}] searching for #{payload} via API...")
        monitor_success = retrieve_payload payload
        @log.warn("[#{monitor_name}] FAILED to find an item that contains #{payload} from the zabbix (#{retry_sz})") unless monitor_success
        retry_attempt_iterator += 1
        retry_sz = "retry attempt #{(retry_attempt_iterator + 1)} / #{num_retries}"
      end

      if monitor_success
        @log.info("[#{monitor_name}] successfully retrieved payload (#{retry_sz})")
        @hard_failure = nil
      else
        @hard_failure ||= @soft_failures.to_a.join('; ')
        @log.error("[#{monitor_name}] has hard failed: #{@hard_failure}")
      end
    end

    def monitor_name
      MONITOR_NAME
    end

    private
    def insert_payload(payload, url, timeout_seconds)
      @log.debug("[#{monitor_name}] value generated: #{payload}")
      payload_delivery_response = deliver_zabbixmon_payload "#{url.gsub(/%datacenter%/, @datacenter)}#{payload}", timeout_seconds

      if payload_delivery_response.code =~ /20./
        @log.debug("[#{monitor_name}] Monitor Payload Delivered Successfully")
      else
        err = "#{payload_delivery_response.code} : #{payload_delivery_response.body}"
        @hard_failure = "ZabbixMon[#{@datacenter}] payload insertion failed! #{ERR_NON_200_HTTP_CODE}\n#{err}"
        @log.error("[#{monitor_name}] ZabbixMon[#{@datacenter}] payload insertion failed! #{err} ")
      end
    end

    def retrieve_payload(payload)
      begin # get zabbix item
        success = false
        zbx_items = @client.get_item_by_name_and_lastvalue(ZBXMON_KEY, payload)
      rescue => e
        @log.error("[#{monitor_name}] #{ERR_ZBX_CLIENT_EXCEPTION}".gsub('%exception%', e))
        @soft_failures.add("#{ERR_ZBX_CLIENT_EXCEPTION}".gsub('%exception%', e).gsub(@zbx_password, '**************'))
      end
      if zbx_items
        if zbx_items.length > 0  # success case
          @log.info("[#{monitor_name}] successfully observed #{payload} from #{@zbx_host} ")
          success = true
        else # fail case
          @soft_failures.add("#{ZABBIX_ITEM_NOT_FOUND}")
          @log.error("[#{monitor_name}] zbx_items='#{zbx_items}'")
        end
      else # fail case
        @soft_failures.add("#{ZABBIX_ITEM_NOT_FOUND}")
      end
      success
    end

    def deliver_zabbixmon_payload(url, timeout_seconds = 30)
      @log.debug("[#{monitor_name}] deliver_zabbixmon_payload url = #{url}")
      uri = URI url
      unless ENV['http_proxy'].nil?
        @proxy_uri = URI.parse(ENV['http_proxy'])
        @proxy_host = @proxy_uri.host
        @proxy_port = @proxy_uri.port
        @proxy_user, @proxy_pass = @proxy_uri.userinfo.split(/:/) if @proxy_uri.userinfo
      end

      unless @proxy_uri.nil?
        http = Net::HTTP.Proxy(@proxy_host, @proxy_port, @proxy_user, @proxy_pass).new(uri.host, uri.port)
        if uri.scheme == 'https'
          http.use_ssl = true
          http.verify_mode = OpenSSL::SSL::VERIFY_NONE
        end
      else
        http = Net::HTTP.new(uri.host, uri.port)
        if uri.scheme == 'https'
          http.use_ssl = true
          http.verify_mode = OpenSSL::SSL::VERIFY_NONE
        end
      end
      #TODO: figure out timeout
      # http.read_timeout = timeout_seconds.to_i
      request = Net::HTTP::Get.new uri.request_uri
      request.basic_auth @zbx_username, @zbx_password if @zbx_username
      response = http.request(request)

    rescue Timeout::Error
      @log.error("[#{monitor_name}] HTTP TIMEOUT while attempting to insert payload")
    rescue ::Lita::Handlers::Zabbix::MonitorDataInsertionFailed
      @log.error("[#{monitor_name}] has hard failed: ::Lita::Handlers::Zabbix::MonitorDataInsertionFailed")
    end
  end
end