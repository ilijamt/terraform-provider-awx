resource "awx_settings_jobs" "zero_defaults" {
  event_stdout_max_bytes_display = 2048
  max_websocket_event_rate       = 60
  stdout_max_bytes_display       = 524288

  schedule_max_jobs = 25
  max_forks         = 100
}
