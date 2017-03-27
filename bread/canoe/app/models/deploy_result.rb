class DeployResult < ApplicationRecord
  START = "start".freeze
  INITIATED = "initiated".freeze
  DEPLOYED = "deployed".freeze
  COMPLETED = "completed".freeze
  FAILED = "failed".freeze
  STAGES = [START, INITIATED, DEPLOYED, COMPLETED, FAILED].freeze

  belongs_to :deploy
  belongs_to :server

  validates :stage,
    presence: true,
    inclusion: { in: STAGES }

  scope :undeployed, -> { where("stage NOT IN (?)", %w[deployed completed failed]) }
  scope :incomplete, -> { where("stage NOT IN (?)", %w[completed failed]) }
  scope :completed, -> { where("stage IN (?)", %w[completed]) }
  scope :failed, -> { where("stage IN (?)", %w[failed]) }
  scope :for_server, ->(server) { where(server: server).first }
  scope :for_server_hostnames, ->(hostnames) { joins(:server).where(servers: { hostname: hostnames }) }

  scope :sort_by_server_hostname, -> { joins(:server).order("servers.hostname ASC") }

  STAGES.each do |stage|
    define_method("#{stage}?") { read_attribute(:stage) == stage }
  end

  def self.for_server_hostname(hostname)
    joins(:server).where(servers: { hostname: hostname }).first
  end

  def at_terminal_state?
    completed? || failed?
  end
end