resource "github_repository" "wave" {
  name          = "wave"
  description   = "Salesforce Wave project"
  homepage_url  = ""
  private       = false
  has_issues    = true
  has_downloads = true
  has_wiki      = true
}

resource "github_team_repository" "wave_site-reliability-engineers" {
  repository = "${github_repository.wave.name}"
  team_id    = "${github_team.site-reliability-engineers.id}"
  permission = "admin"
}

resource "github_team_repository" "wave_engineering-managers" {
  repository = "${github_repository.wave.name}"
  team_id    = "${github_team.engineering-managers.id}"
  permission = "admin"
}
