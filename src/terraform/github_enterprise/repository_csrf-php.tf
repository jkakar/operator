resource "github_repository" "csrf-php" {
  name          = "csrf-php"
  description   = ""
  homepage_url  = ""
  private       = false
  has_issues    = true
  has_downloads = true
  has_wiki      = true
}

resource "github_team_repository" "csrf-php_developers" {
  repository = "${github_repository.csrf-php.name}"
  team_id    = "${github_team.developers.id}"
  permission = "push"
}

resource "github_team_repository" "csrf-php_service-accounts-read-only" {
  repository = "${github_repository.csrf-php.name}"
  team_id    = "${github_team.service-accounts-read-only.id}"
  permission = "pull"
}

resource "github_team_repository" "csrf-php_read-only-users" {
  repository = "${github_repository.csrf-php.name}"
  team_id    = "${github_team.read-only-users.id}"
  permission = "pull"
}

resource "github_team_repository" "csrf-php_site-reliability-engineers" {
  repository = "${github_repository.csrf-php.name}"
  team_id    = "${github_team.site-reliability-engineers.id}"
  permission = "admin"
}

resource "github_team_repository" "csrf-php_engineering-managers" {
  repository = "${github_repository.csrf-php.name}"
  team_id    = "${github_team.engineering-managers.id}"
  permission = "admin"
}
