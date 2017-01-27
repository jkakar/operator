class PardotRepository
  ANSIBLE = "Pardot/ansible".freeze
  BLUE_MESH = "Pardot/blue-mesh".freeze
  BREAD = "Pardot/bread".freeze
  CHEF = "Pardot/chef".freeze
  INTERNAL_API = "Pardot/internal-api".freeze
  MESH = "Pardot/mesh".freeze
  MURDOC = "Pardot/murdoc".freeze
  PARDOT = "Pardot/pardot".freeze
  TEAM_OPS = "Pardot/ops".freeze
  TEAM_DEVELOPERS = "Pardot/developers".freeze
  TEST_STATUS = "Test Jobs".freeze
  FINAL_STATUS = "Final Jobs".freeze

  def initialize(nwo)
    @name_with_owner = nwo
  end

  attr_reader :name_with_owner

  def team
    if CHEF
      TEAM_OPS
    else
      TEAM_DEVELOPERS
    end
  end

  def ticket_reference_required?
    true
  end
end