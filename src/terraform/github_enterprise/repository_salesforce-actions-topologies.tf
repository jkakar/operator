resource "github_repository" "salesforce-actions-topologies" {
  name          = "salesforce-actions-topologies"
  description   = "3rd Party Actions Storm Topology"
  homepage_url  = ""
  private       = false
  has_issues    = false
  has_downloads = true
  has_wiki      = false
}

resource "github_team_repository" "salesforce-actions-topologies_developers" {
  repository = "${github_repository.salesforce-actions-topologies.name}"
  team_id    = "${github_team.developers.id}"
  permission = "push"
}

resource "github_team_repository" "salesforce-actions-topologies_service-accounts-write-only" {
  repository = "${github_repository.salesforce-actions-topologies.name}"
  team_id    = "${github_team.service-accounts-write-only.id}"
  permission = "push"
}

resource "github_team_repository" "salesforce-actions-topologies_engineering-managers" {
  repository = "${github_repository.salesforce-actions-topologies.name}"
  team_id    = "${github_team.engineering-managers.id}"
  permission = "admin"
}

resource "github_team_repository" "salesforce-actions-topologies_site-reliability-engineers" {
  repository = "${github_repository.salesforce-actions-topologies.name}"
  team_id    = "${github_team.site-reliability-engineers.id}"
  permission = "admin"
}