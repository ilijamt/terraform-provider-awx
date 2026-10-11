resource "awx_team" "team" {
  name         = "Team"
  organization = awx_organization.credentials.id
}
