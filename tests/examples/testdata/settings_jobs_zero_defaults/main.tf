resource "awx_settings_jobs" "zero_defaults" {
  event_stdout_max_bytes_display = 0
  max_websocket_event_rate       = 0
  stdout_max_bytes_display       = 0
  default_job_timeout            = 0
  ansible_fact_cache_timeout     = 0

  schedule_max_jobs = 10
  max_forks         = 200
}
