resource "awx_settings_auth_google_oauth2" "default" {
  social_auth_google_oauth2_auth_extra_arguments = jsonencode({})
  social_auth_google_oauth2_key                  = "oauth2_key"
  social_auth_google_oauth2_organization_map     = jsonencode({})
  social_auth_google_oauth2_secret               = "secret"
  social_auth_google_oauth2_team_map             = jsonencode({})
  social_auth_google_oauth2_whitelisted_domains = [
    "https://example.com"
  ]
}
