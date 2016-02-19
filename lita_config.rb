require_relative "lib/lita/adapters/nothing"

Lita.configure do |config|
  config.robot.name = "Hal9000"
  config.robot.alias = "!"
  config.robot.locale = :en

  # The severity of messages to log. Options are:
  # :debug, :info, :warn, :error, :fatal
  # Messages at the selected level and above will be logged.
  config.robot.log_level = :info

  # An array of user IDs that are considered administrators. These users
  # the ability to add and remove other users from authorization groups.
  # What is considered a user ID will change depending on which adapter you use.
  # config.robot.admins = ["1", "2"]

  config.robot.adapter = (ENV.fetch("LITA_ADAPTER", "shell")).to_sym

  config.http.host = "0.0.0.0"
  config.http.port = 8080

  ## Example: Set options for the chosen adapter.
  # config.adapter.username = "myname"
  # config.adapter.password = "secret"
  config.adapters.hipchat.server = ENV.fetch("HIPCHAT_SERVER", "hipchat.dev.pardot.com")
  config.adapters.hipchat.jid = ENV.fetch("HIPCHAT_JID", "1_342@chat.btf.hipchat.com")
  config.adapters.hipchat.password = ENV.fetch("HIPCHAT_PASSWORD", "")
  config.adapters.hipchat.debug = true

  # Replication fixing
  config.handlers.replication_fixing.pagerduty_service_key = ENV.fetch("PAGERDUTY_SERVICE_KEY", "")

  ## Example: Set options for the Redis connection.
  config.redis[:host] = ENV.fetch("REDIS_HOST", "127.0.0.1")
  config.redis[:port] = ENV.fetch("REDIS_PORT", "6379").to_i

  ## Example: Set configuration for any loaded handlers. See the handler's
  ## documentation for options.
  # config.handlers.some_handler.some_config_key = "value"
end
