resource "github_repository" "pardot" {
  name          = "pardot"
  description   = "The best marketing automation app in the multiverse."
  homepage_url  = ""
  private       = true
  has_issues    = false
  has_downloads = true
  has_wiki      = false
}

resource "github_team_repository" "pardot_developers" {
  repository = "${github_repository.pardot.name}"
  team_id    = "${github_team.developers.id}"
  permission = "push"
}

resource "github_team_repository" "pardot_ops" {
  repository = "${github_repository.pardot.name}"
  team_id    = "${github_team.ops.id}"
  permission = "push"
}

resource "github_team_repository" "pardot_service-accounts-write-only" {
  repository = "${github_repository.pardot.name}"
  team_id    = "${github_team.service-accounts-write-only.id}"
  permission = "push"
}

resource "github_team_repository" "pardot_service-accounts-read-only" {
  repository = "${github_repository.pardot.name}"
  team_id    = "${github_team.service-accounts-read-only.id}"
  permission = "pull"
}

resource "github_team_repository" "pardot_tier-2-support" {
  repository = "${github_repository.pardot.name}"
  team_id    = "${github_team.tier-2-support.id}"
  permission = "push"
}

resource "github_team_repository" "pardot_core-production-security" {
  repository = "${github_repository.pardot.name}"
  team_id    = "${github_team.core-production-security.id}"
  permission = "pull"
}