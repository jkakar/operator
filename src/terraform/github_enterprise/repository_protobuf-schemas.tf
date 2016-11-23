resource "github_repository" "protobuf-schemas" {
  name          = "protobuf-schemas"
  description   = "One stop shop for Protocol Buffer Schema Definitions "
  homepage_url  = ""
  private       = true
  has_issues    = true
  has_downloads = true
  has_wiki      = true
}

resource "github_team_repository" "protobuf-schemas_developers" {
  repository = "${github_repository.protobuf-schemas.name}"
  team_id    = "${github_team.developers.id}"
  permission = "push"
}

resource "github_team_repository" "protobuf-schemas_ops" {
  repository = "${github_repository.protobuf-schemas.name}"
  team_id    = "${github_team.ops.id}"
  permission = "push"
}

resource "github_team_repository" "protobuf-schemas_service-accounts-write-only" {
  repository = "${github_repository.protobuf-schemas.name}"
  team_id    = "${github_team.service-accounts-write-only.id}"
  permission = "push"
}

resource "github_team_repository" "protobuf-schemas_service-accounts-read-only" {
  repository = "${github_repository.protobuf-schemas.name}"
  team_id    = "${github_team.service-accounts-read-only.id}"
  permission = "pull"
}