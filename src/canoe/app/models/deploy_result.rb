class DeployResult < ActiveRecord::Base
  belongs_to :deploy
  belongs_to :server

  STAGES = %w[initiated deployed completed failed].freeze
  validates :stage,
    presence: true,
    inclusion: { in: STAGES }

  scope :incomplete, -> { where("stage NOT IN (?)", %w[completed failed]) }
  scope :for_server, -> (server) { where(server: server).first }
  scope :for_server_hostnames, -> (hostnames) { joins(:server).where(servers: { hostname: hostnames }) }

  scope :sort_by_server_hostname, -> { joins(:server).order("servers.hostname ASC") }

  STAGES.each do |stage|
    scope stage, -> { where(stage: stage) }
    define_method("#{stage}?") { self.stage == stage }
  end

  def self.for_server_hostname(hostname)
    joins(:server).where(servers: { hostname: hostname }).first
  end

  def at_terminal_state?
    completed? || failed?
  end
end
