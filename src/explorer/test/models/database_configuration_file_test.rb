require "test_helper"

class DatabaseConfigurationFileTest < ActiveSupport::TestCase
  setup do
    @config = DatabaseConfigurationFile.load
  end

  test "global" do
    config = @config.global(Datacenter::LOCAL)
    assert_equal "mysql", config.hostname
    assert_equal "pardot_global", config.name
    assert_equal "root", config.username
    assert_nil config.password

    assert_raise(DatabaseConfigurationFile::DatacenterNotFound) do
      @config.global("brussels")
    end
  end

  test "shard" do
    config = @config.shard(Datacenter::SEATTLE, 1)
    assert_equal "mysql", config.hostname
    assert_equal "pardot_shard1", config.name
    assert_equal "root", config.username
    assert_nil config.password

    assert_raise(DatabaseConfigurationFile::ShardNotFound) do
      @config.shard(Datacenter::SEATTLE, 3)
    end

    assert_raise(DatabaseConfigurationFile::DatacenterNotFound) do
      @config.shard("brussels", 1)
    end
  end
end
