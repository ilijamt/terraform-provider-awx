data "awx_execution_environment" "latest" {
  name = "AWX EE (latest)"
}

resource "awx_settings_ui" "default" {
  ui_live_updates_enabled = true
  max_ui_job_events       = 2000
  custom_logo             = ""
  custom_login_info       = ""
}

resource "awx_settings_oidc" "default" {
  social_auth_oidc_verify_ssl = false
}

resource "awx_settings_misc_logging" "default" {
  api_400_error_log_format        = "status {status_code} received by user {user_name} attempting to access {url_path} from {remote_addr}"
  log_aggregator_enabled          = false
  log_aggregator_individual_facts = false
  log_aggregator_level            = "INFO"
  log_aggregator_loggers = [
    "awx",
    "activity_stream",
    "job_events",
    "system_tracking",
  ]
  log_aggregator_action_max_disk_usage_gb = 1
  log_aggregator_max_disk_usage_path      = "/var/lib/awx"
  log_aggregator_protocol                 = "https"
  log_aggregator_rsyslogd_debug           = false
  log_aggregator_tcp_timeout              = 5
  log_aggregator_verify_cert              = true
}

resource "awx_settings_misc_system" "default" {
  default_execution_environment              = data.awx_execution_environment.latest.id
  activity_stream_enabled                    = true
  activity_stream_enabled_for_inventory_sync = false
  automation_analytics_gather_interval       = 14400
  automation_analytics_url                   = "https://example.com"
  tower_url_base                             = "http://awx.local"
  insights_tracking_state                    = false
  manage_organization_auth                   = true
  org_admins_can_see_all_users               = true
  proxy_ip_allowed_list                      = []
  remote_host_headers = [
    "REMOTE_ADDR",
    "REMOTE_HOST",
    "HTTP_X_FORWARDED_FOR",
  ]
}

resource "awx_settings_misc_authentication" "default" {
  allow_oauth2_for_external_users = false
  auth_basic_enabled              = true
  disable_local_auth              = false
  oauth2_provider = jsonencode({
    ACCESS_TOKEN_EXPIRE_SECONDS       = 31536000000
    AUTHORIZATION_CODE_EXPIRE_SECONDS = 600
    REFRESH_TOKEN_EXPIRE_SECONDS      = 2628000
  })
  session_cookie_age      = 1800000
  sessions_per_user       = -1
  social_auth_user_fields = []
}

data "awx_settings_misc_system" "default" {
  depends_on = [awx_settings_misc_system.default]
}
