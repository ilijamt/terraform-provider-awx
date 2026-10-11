resource "awx_settings_auth_github_enterprise_org" "default" {
  social_auth_github_enterprise_org_url              = "https://example.com"
  social_auth_github_enterprise_org_api_url          = "https://api.example.com"
  social_auth_github_enterprise_org_name             = "Example Org name"
  social_auth_github_enterprise_org_key              = "key"
  social_auth_github_enterprise_org_organization_map = jsonencode(local.org_map)
  social_auth_github_enterprise_org_secret           = "secret"
  social_auth_github_enterprise_org_team_map         = jsonencode(local.team_map)
}
