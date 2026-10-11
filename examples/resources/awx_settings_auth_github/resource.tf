resource "awx_settings_auth_github" "default" {
  social_auth_github_key              = "key"
  social_auth_github_organization_map = jsonencode(local.org_map)
  social_auth_github_secret           = "secret"
  social_auth_github_team_map         = jsonencode(local.team_map)
}
