require "replication_fixing/shard"

module ReplicationFixing
  # Add top-level class documentation comment here.
  class Hostname
    MalformedHostname = Class.new(StandardError)

    attr_reader :hostname, :shard, :cluster_id

    def initialize(hostname)
      @hostname = hostname
      parse_hostname
    end

    def to_s
      @hostname
    end

    def ==(other)
      other.is_a?(Hostname) && hostname == other.hostname
    end

    def eql?(other)
      other.is_a?(Hostname) && hostname == other.hostname
    end

    def hash
      @hostname.hash
    end

    def prefix
      shard.prefix
    end

    def shard_id
      shard.shard_id
    end

    def datacenter
      shard.datacenter
    end

    private

    def parse_hostname
      if /\Apardot0-(?<type>dbshard|whoisdb)(?<cluster_id>\d+)-(?<shard_id>\d+)
          -(?<datacenter>[^-]+)\z/x =~ @hostname
        prefix = \
          case type
          when "whoisdb"
            "whoisdb"
          else
            "db"
          end

        shard_id = shard_id.to_i

        @shard = Shard.new(prefix, shard_id, datacenter)
        @cluster_id = cluster_id.to_i
      else
        raise MalformedHostname
      end
    end
  end
end
