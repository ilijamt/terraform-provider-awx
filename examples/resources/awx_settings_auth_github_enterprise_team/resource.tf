resource "awx_settings_auth_github_enterprise_team" "default" {
  social_auth_github_enterprise_team_id               = "example-id"
  social_auth_github_enterprise_team_url              = "https://example.com"
  social_auth_github_enterprise_team_api_url          = "https://api.example.com"
  social_auth_github_enterprise_team_key              = "key"
  social_auth_github_enterprise_team_secret           = "secret"
  social_auth_github_enterprise_team_organization_map = jsonencode(local.org_map)
  social_auth_github_enterprise_team_team_map         = jsonencode(local.team_map)
}
