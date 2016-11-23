resource "github_repository" "engagement-history-package" {
  name          = "engagement-history-package"
  description   = "Engagement History"
  homepage_url  = ""
  private       = false
  has_issues    = true
  has_downloads = true
  has_wiki      = true
}

resource "github_team_repository" "engagement-history-package_developers" {
  repository = "${github_repository.engagement-history-package.name}"
  team_id    = "${github_team.developers.id}"
  permission = "push"
}

resource "github_team_repository" "engagement-history-package_service-accounts-read-only" {
  repository = "${github_repository.engagement-history-package.name}"
  team_id    = "${github_team.service-accounts-read-only.id}"
  permission = "pull"
}

resource "github_team_repository" "engagement-history-package_service-accounts-write-only" {
  repository = "${github_repository.engagement-history-package.name}"
  team_id    = "${github_team.service-accounts-write-only.id}"
  permission = "push"
}