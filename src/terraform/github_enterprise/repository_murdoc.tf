resource "github_repository" "murdoc" {
  name          = "murdoc"
  description   = "NurtureStudio Storm Toplogy - Murdoc makes his own rules, and so shall we"
  homepage_url  = ""
  private       = true
  has_issues    = true
  has_downloads = true
  has_wiki      = true
}

resource "github_team_repository" "murdoc_developers" {
  repository = "${github_repository.murdoc.name}"
  team_id    = "${github_team.developers.id}"
  permission = "push"
}

resource "github_team_repository" "murdoc_service-accounts-read-only" {
  repository = "${github_repository.murdoc.name}"
  team_id    = "${github_team.service-accounts-read-only.id}"
  permission = "pull"
}

resource "github_team_repository" "murdoc_service-accounts-write-only" {
  repository = "${github_repository.murdoc.name}"
  team_id    = "${github_team.service-accounts-write-only.id}"
  permission = "push"
}