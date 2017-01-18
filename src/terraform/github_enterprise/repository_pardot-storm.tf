resource "github_repository" "pardot-storm" {
  name          = "pardot-storm"
  description   = "All the things needed to create a Pardot storm topology"
  homepage_url  = ""
  private       = false
  has_issues    = true
  has_downloads = true
  has_wiki      = true
}

resource "github_team_repository" "pardot-storm_developers" {
  repository = "${github_repository.pardot-storm.name}"
  team_id    = "${github_team.developers.id}"
  permission = "push"
}

resource "github_team_repository" "pardot-storm_service-accounts-read-only" {
  repository = "${github_repository.pardot-storm.name}"
  team_id    = "${github_team.service-accounts-read-only.id}"
  permission = "pull"
}

resource "github_team_repository" "pardot-storm_service-accounts-write-only" {
  repository = "${github_repository.pardot-storm.name}"
  team_id    = "${github_team.service-accounts-write-only.id}"
  permission = "push"
}

resource "github_team_repository" "pardot-storm_engineering-managers" {
  repository = "${github_repository.pardot-storm.name}"
  team_id    = "${github_team.engineering-managers.id}"
  permission = "admin"
}

resource "github_team_repository" "pardot-storm_site-reliability-engineers" {
  repository = "${github_repository.pardot-storm.name}"
  team_id    = "${github_team.site-reliability-engineers.id}"
  permission = "admin"
}
