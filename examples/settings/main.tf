terraform {
  required_providers {
    awx = {
      source = "registry.terraform.io/ilijamt/awx"
    }
  }
}

provider "awx" {}

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
  log_aggregator_action_max_disk_usage_gb   = 1
  log_aggregator_max_disk_usage_path = "/var/lib/awx"
  log_aggregator_protocol            = "https"
  log_aggregator_rsyslogd_debug      = false
  log_aggregator_tcp_timeout         = 5
  log_aggregator_verify_cert         = true
}

data "awx_settings_misc_system" "default" {}

resource "awx_settings_misc_system" "default" {
  default_execution_environment              = data.awx_execution_environment.latest.id
  activity_stream_enabled                    = true
  activity_stream_enabled_for_inventory_sync = false
  automation_analytics_gather_interval       = 14400
  automation_analytics_url                   = "https://example.com"
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
