resource "github_repository" "pardot-refocus" {
  name          = "pardot-refocus"
  description   = "App to feed Pardot metrics to Refocus"
  homepage_url  = ""
  private       = false
  has_issues    = true
  has_downloads = true
  has_wiki      = true
}

resource "github_team_repository" "pardot-refocus_developers" {
  repository = "${github_repository.pardot-refocus.name}"
  team_id    = "${github_team.developers.id}"
  permission = "push"
}

resource "github_team_repository" "pardot-refocus_ops" {
  repository = "${github_repository.pardot-refocus.name}"
  team_id    = "${github_team.ops.id}"
  permission = "push"
}

resource "github_team_repository" "pardot-refocus_service-accounts-write-only" {
  repository = "${github_repository.pardot-refocus.name}"
  team_id    = "${github_team.service-accounts-write-only.id}"
  permission = "push"
}