resource "github_team" "daleks" {
  name        = "Daleks"
  description = "https://gus.lightning.force.com/one/one.app#/sObject/a00B0000005MPjoIAG/view"
  privacy     = "closed"
}

resource "github_team_membership" "daleks_william-castillo" {
  team_id  = "${github_team.daleks.id}"
  username = "william-castillo"
  role     = "maintainer"
}

resource "github_team_membership" "daleks_david-peterson" {
  team_id  = "${github_team.daleks.id}"
  username = "david-peterson"
  role     = "maintainer"
}

resource "github_team_membership" "daleks_brian-hays" {
  team_id  = "${github_team.daleks.id}"
  username = "brian-hays"
  role     = "maintainer"
}

resource "github_team_membership" "daleks_jan-ulrich" {
  team_id  = "${github_team.daleks.id}"
  username = "jan-ulrich"
  role     = "maintainer"


