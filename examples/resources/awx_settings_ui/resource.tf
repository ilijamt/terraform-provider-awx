resource "awx_settings_ui" "default" {
  ui_live_updates_enabled = true
  max_ui_job_events       = 2000
  custom_logo             = ""
  custom_login_info       = ""
}
