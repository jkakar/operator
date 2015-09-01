# Server represents a server where code is deployed for a given target.
#
# At this time, `Server` specifically represents a server where code is _pulled_
# via `pull_agent`. Servers where code is _pushed_ are stored in `sync_scripts`.
# As `sync_scripts` becomes deprecated, eventually all servers where code is
# deployed will be a `Server` instance.
class Server < ActiveRecord::Base
  validates :hostname, uniqueness: true

  scope :enabled, -> { where(enabled: true) }

  belongs_to :deploy_target
  has_many :repo_server
  has_many :repos, through: :repo_server
end
