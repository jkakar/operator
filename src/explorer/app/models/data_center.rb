class DataCenter
  class NotFound < StandardError
    def initialize(name)
      super "datacenter not found: #{name.inspect}"
    end
  end

  class UnauthorizedAccountAccess < StandardError
    def initialize(account_id)
      @account_id = account_id

      super "access to account #{account_id.inspect} is not authorized"
    end

    attr_reader :account_id
  end

  ENGINEERING_ROLE = 7
  DALLAS = "dfw".freeze
  LOCAL = "local".freeze
  SEATTLE = "phx".freeze

  # Returns the current Datacenter based on the Rails configuration.
  def self.current
    @datacenter ||= DataCenter.new(
      Rails.application.config.x.datacenter,
      DatabaseConfigurationFile.load
    )
  end

  def initialize(name, config)
    @name = name
    @config = config

    if ![DALLAS, LOCAL, SEATTLE].include?(name)
      raise NotFound, name
    end
  end

  attr_reader :name

  def global
    config = @config.global(@name)

    Database.new(config)
  end

  def shard_for(account_id)
    if !access_authorized?(account_id)
      raise UnauthorizedAccountAccess, account_id
    end

    account = find_account(account_id)
    config = @config.shard(@name, account.shard_id)

    Database.new(config)
  end

  def find_account(account_id)
    global_accounts.find(account_id)
  end

  def accounts
    global_accounts.all
  end

  private

  def access_authorized?(account_id)
    query = <<-SQL.freeze
      SELECT id FROM global_account_access
      WHERE role = ? AND account_id = ? AND (expires_at IS NULL OR expires_at > NOW())
      LIMIT 1
    SQL
    results = global.execute(query, [ENGINEERING_ROLE, account_id])
    results.size == 1
  end

  def global_accounts
    @global_accounts ||= GlobalAccountsCollection.new(global)
  end
end
