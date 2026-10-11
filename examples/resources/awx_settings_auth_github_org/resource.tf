resource "awx_settings_auth_github_org" "default" {
  social_auth_github_org_key              = "key"
  social_auth_github_org_organization_map = jsonencode(local.org_map)
  social_auth_github_org_secret           = "secret"
  social_auth_github_org_team_map         = jsonencode(local.team_map)
}
