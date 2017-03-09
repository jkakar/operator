resource "github_repository" "breakout" {
  name          = "breakout"
  description   = ""
  homepage_url  = ""
  private       = false
  has_issues    = true
  has_downloads = true
  has_wiki      = true
}

resource "github_team_repository" "breakout_developers" {
  repository = "${github_repository.breakout.name}"
  team_id    = "${github_team.developers.id}"
  permission = "push"
}

resource "github_team_repository" "breakout_read-only-users" {
  repository = "${github_repository.breakout.name}"
  team_id    = "${github_team.read-only-users.id}"
  permission = "pull"
}

resource "github_team_repository" "breakout_service-accounts-write-only" {
  repository = "${github_repository.breakout.name}"
  team_id    = "${github_team.service-accounts-write-only.id}"
  permission = "push"
}

resource "github_team_repository" "breakout_service-accounts-administrators" {
  repository = "${github_repository.breakout.name}"
  team_id    = "${github_team.service-accounts-administrators.id}"
  permission = "admin"
}