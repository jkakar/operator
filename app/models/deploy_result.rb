class DeployResult < ActiveRecord::Base
  belongs_to :deploy
  belongs_to :server

  STAGES = %w(initiated deployed completed failed)
  validates :stage,
    presence: true,
    inclusion: {in: STAGES}

  scope :incomplete, -> { where("stage NOT IN (?)", ["completed", "failed"]) }
  scope :for_server, -> (server) { where(server: server).first }
  scope :for_server_hostnames, -> (hostnames) { joins(:server).where(servers: {hostname: hostnames}) }

  STAGES.each do |stage|
    scope stage, -> { where(stage: stage) }
    define_method("#{stage}?") { self.stage == stage }
  end

  def self.for_server_hostname(hostname)
    joins(:server).where(servers: {hostname: hostname}).first
  end
end
