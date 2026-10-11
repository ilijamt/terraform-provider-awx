resource "awx_settings_jobs" "default" {
  ad_hoc_commands = [
    "command",
    "shell",
    "yum",
    "apt",
    "apt_key",
    "apt_repository",
    "apt_rpm",
    "service",
    "group",
    "user",
    "mount",
    "ping",
    "selinux",
    "setup",
    "win_ping",
    "win_service",
    "win_updates",
    "win_group",
    "win_user",
  ]
  allow_jinja_in_extra_vars    = "template"
  ansible_fact_cache_timeout   = 0
  awx_ansible_callback_plugins = []
  awx_collections_enabled      = true
  awx_isolation_base_path      = "/tmp"
  awx_isolation_show_paths = [
    "/etc/pki/ca-trust:/etc/pki/ca-trust:O",
    "/usr/share/pki:/usr/share/pki:O",
  ]
  awx_mount_isolated_paths_on_k8s  = false
  awx_roles_enabled                = true
  awx_show_playbook_links          = true
  awx_task_env                     = jsonencode({})
  default_inventory_update_timeout = 0
  default_job_idle_timeout         = 0
  default_job_timeout              = 0
  default_project_update_timeout   = 0
  event_stdout_max_bytes_display   = 1024
  galaxy_ignore_certs              = false
  galaxy_task_env = jsonencode(
    {
      ANSIBLE_FORCE_COLOR = "false"
      GIT_SSH_COMMAND     = "ssh -o StrictHostKeyChecking=no"
    }
  )
  max_forks                = 200
  max_websocket_event_rate = 30
  project_update_vvv       = false
  schedule_max_jobs        = 10
  stdout_max_bytes_display = 1048576
}
