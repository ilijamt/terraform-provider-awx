resource "awx_settings_auth_azuread_oauth2" "default" {
  social_auth_azuread_oauth2_key              = "key"
  social_auth_azuread_oauth2_secret           = "access"
  social_auth_azuread_oauth2_organization_map = jsonencode({})
  social_auth_azuread_oauth2_team_map         = jsonencode({})
}
