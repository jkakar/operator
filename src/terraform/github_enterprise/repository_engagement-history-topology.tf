resource "github_repository" "engagement-history-topology" {
  name          = "engagement-history-topology"
  description   = "Pioneering marketing data on Salesforce Core"
  homepage_url  = ""
  private       = true
  has_issues    = false
  has_downloads = true
  has_wiki      = true
}

resource "github_team_repository" "engagement-history-topology_service-accounts-write-only" {
  repository = "${github_repository.engagement-history-topology.name}"
  team_id    = "${github_team.service-accounts-write-only.id}"
  permission = "push"
}

resource "github_team_repository" "engagement-history-topology_ops" {
  repository = "${github_repository.engagement-history-topology.name}"
  team_id    = "${github_team.ops.id}"
  permission = "push"
}

resource "github_team_repository" "engagement-history-topology_developers" {
  repository = "${github_repository.engagement-history-topology.name}"
  team_id    = "${github_team.developers.id}"
  permission = "push"
}

resource "github_team_repository" "engagement-history-topology_service-accounts-read-only" {
  repository = "${github_repository.engagement-history-topology.name}"
  team_id    = "${github_team.service-accounts-read-only.id}"
  permission = "pull"
}