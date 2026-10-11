resource "awx_settings_auth_github_team" "default" {
  social_auth_github_team_id               = "example-id"
  social_auth_github_team_key              = "key"
  social_auth_github_team_secret           = "secret"
  social_auth_github_team_organization_map = jsonencode(local.org_map)
  social_auth_github_team_team_map         = jsonencode(local.team_map)
}
