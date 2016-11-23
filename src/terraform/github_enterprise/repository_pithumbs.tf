resource "github_repository" "pithumbs" {
  name          = "pithumbs"
  description   = "Home for service to quickly generate thumbnails from HTML"
  homepage_url  = ""
  private       = true
  has_issues    = true
  has_downloads = true
  has_wiki      = true
}

resource "github_team_repository" "pithumbs_developers" {
  repository = "${github_repository.pithumbs.name}"
  team_id    = "${github_team.developers.id}"
  permission = "push"
}

resource "github_team_repository" "pithumbs_service-accounts-write-only" {
  repository = "${github_repository.pithumbs.name}"
  team_id    = "${github_team.service-accounts-write-only.id}"
  permission = "push"
}

resource "github_team_repository" "pithumbs_service-accounts-read-only" {
  repository = "${github_repository.pithumbs.name}"
  team_id    = "${github_team.service-accounts-read-only.id}"
  permission = "pull"
}
