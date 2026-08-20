resource "awx_organization" "org" {
  name = "NotificationTemplates"
}

# Every awssns field defaults, so this is the minimal legal template.
resource "awx_notification_template_awssns" "smoke" {
  name         = "AWS SNS smoke"
  organization = awx_organization.org.id
  aws_region   = "eu-west-1"
}

resource "awx_notification_template_email" "smoke" {
  name         = "Email smoke"
  organization = awx_organization.org.id
  host         = "smtp.example.com"
  port         = 587
  username     = "service-account"
  password     = "email-secret"
  use_tls      = true
  use_ssl      = false
  sender       = "awx@example.com"
  recipients   = ["ops@example.com", "oncall@example.com"]
}

resource "awx_notification_template_grafana" "smoke" {
  name         = "Grafana smoke"
  organization = awx_organization.org.id
  grafana_url  = "https://grafana.example.com"
  grafana_key  = "grafana-secret"
}

resource "awx_notification_template_irc" "smoke" {
  name         = "IRC smoke"
  organization = awx_organization.org.id
  server       = "irc.example.com"
  port         = 6697
  nickname     = "awx"
  password     = "irc-secret"
  use_ssl      = true
  targets      = ["#ops", "#alerts"]
}

resource "awx_notification_template_mattermost" "smoke" {
  name                     = "Mattermost smoke"
  organization             = awx_organization.org.id
  mattermost_url           = "https://mattermost.example.com/hooks/smoke"
  mattermost_no_verify_ssl = false
}

resource "awx_notification_template_pagerduty" "smoke" {
  name         = "Pagerduty smoke"
  organization = awx_organization.org.id
  subdomain    = "example"
  token        = "pagerduty-secret"
  service_key  = "example-service-key"
  client_name  = "awx"
}

resource "awx_notification_template_rocketchat" "smoke" {
  name                     = "Rocket.Chat smoke"
  organization             = awx_organization.org.id
  rocketchat_url           = "https://rocketchat.example.com/hooks/smoke"
  rocketchat_no_verify_ssl = false
}

resource "awx_notification_template_slack" "smoke" {
  name         = "Slack smoke"
  organization = awx_organization.org.id
  token        = "slack-secret"
  channels     = ["#ops", "#alerts"]
}

resource "awx_notification_template_twilio" "smoke" {
  name          = "Twilio smoke"
  organization  = awx_organization.org.id
  account_sid   = "ACexample"
  account_token = "twilio-secret"
  from_number   = "+15005550006"
  to_numbers    = ["+15005550001", "+15005550002"]
}

# The shape from issue #172, without the lifecycle { ignore_changes } workaround.
resource "awx_notification_template_webhook" "smoke" {
  name         = "Webhook smoke"
  description  = "Created by TestIntegration_NotificationTemplates"
  organization = awx_organization.org.id
  url          = "https://hooks.example.com/awx"
  username     = "service-account"
  password     = "webhook-secret"
  headers      = {}
}

data "awx_notification_template_webhook" "webhook_by_name" {
  name = awx_notification_template_webhook.smoke.name
}

data "awx_notification_template_irc" "irc_by_name" {
  name = awx_notification_template_irc.smoke.name
}
