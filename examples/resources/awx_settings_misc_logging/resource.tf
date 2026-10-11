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
