resource "github_repository" "dropwizard-example" {
  name          = "dropwizard-example"
  description   = ""
  homepage_url  = ""
  private       = true
  has_issues    = true
  has_downloads = true
  has_wiki      = true
}

resource "github_team_repository" "dropwizard-example_developers" {
  repository = "${github_repository.dropwizard-example.name}"
  team_id    = "${github_team.developers.id}"
  permission = "push"
}

resource "github_team_repository" "dropwizard-example_service-accounts-read-only" {
  repository = "${github_repository.dropwizard-example.name}"
  team_id    = "${github_team.service-accounts-read-only.id}"
  permission = "pull"
}