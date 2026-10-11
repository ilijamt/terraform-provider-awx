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
    "HTTP_X_FORWARDED_FOR"
  ]
}
