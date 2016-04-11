class Database
  GLOBAL = "global".freeze
  SHARD = "shard".freeze

  class QueryError < StandardError
  end

  def initialize(user, config)
    @user = user
    @config = config
  end

  def hostname
    @config.hostname
  end

  def name
    @config.name
  end

  def tables
    activerecord_connection.tables
  end

  def execute(sql, params = [])
    query = Query.new(self, connection, @user, sql)
    query.execute(params)
  end

  private

  def activerecord_connection
    @activerecord_connection ||=
      ActiveRecord::ConnectionAdapters::Mysql2Adapter.new(
        connection,
        nil,
        nil,
        {}
      )
  end

  def connection
    @connection ||= establish_connection
  end

  def establish_connection
    connection = Mysql2::Client.new(
      host: @config.hostname,
      port: @config.port,
      username: @config.username,
      password: @config.password,
      database: @config.name,
    )
    connection.query_options.merge!(symbolize_keys: true)
    connection
  end
end
