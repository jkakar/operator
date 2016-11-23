resource "github_repository" "SFdeploy" {
  name          = "SFdeploy"
  description   = ""
  homepage_url  = ""
  private       = true
  has_issues    = true
  has_downloads = true
  has_wiki      = true
}

resource "github_team_repository" "SFdeploy_developers" {
  repository = "${github_repository.SFdeploy.name}"
  team_id    = "${github_team.developers.id}"
  permission = "push"
}