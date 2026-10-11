resource "awx_settings_misc_authentication" "default" {
  allow_oauth2_for_external_users = false
  auth_basic_enabled              = true
  disable_local_auth              = false
  oauth2_provider = jsonencode(
    {
      ACCESS_TOKEN_EXPIRE_SECONDS       = 31536000000
      AUTHORIZATION_CODE_EXPIRE_SECONDS = 600
      REFRESH_TOKEN_EXPIRE_SECONDS      = 2628000
    }
  )
  session_cookie_age      = 1800000
  sessions_per_user       = -1
  social_auth_user_fields = []
}
