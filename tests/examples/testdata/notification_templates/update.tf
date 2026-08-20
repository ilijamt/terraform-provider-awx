resource "awx_organization" "org" {
  name = "NotificationTemplates"
}

resource "awx_notification_template_awssns" "smoke" {
  name                  = "AWS SNS smoke"
  organization          = awx_organization.org.id
  aws_region            = "us-east-1"
  aws_secret_access_key = "awssns-secret"
}

resource "awx_notification_template_email" "smoke" {
  name         = "Email smoke (updated)"
  organization = awx_organization.org.id
  host         = "smtp2.example.com"
  port         = 465
  username     = "service-account"
  password     = "email-secret-rotated"
  use_tls      = false
  use_ssl      = true
  sender       = "awx@example.com"
  recipients   = ["ops@example.com"]
  timeout      = 60
}

resource "awx_notification_template_grafana" "smoke" {
  name         = "Grafana smoke"
  organization = awx_organization.org.id
  grafana_url  = "https://grafana2.example.com"
  grafana_key  = "grafana-secret-rotated"
}

resource "awx_notification_template_irc" "smoke" {
  name         = "IRC smoke"
  organization = awx_organization.org.id
  server       = "irc.example.com"
  port         = 6667
  nickname     = "awx-bot"
  password     = "irc-secret-rotated"
  use_ssl      = false
  targets      = ["#ops"]
}

resource "awx_notification_template_mattermost" "smoke" {
  name                     = "Mattermost smoke"
  organization             = awx_organization.org.id
  mattermost_url           = "https://mattermost.example.com/hooks/updated"
  mattermost_no_verify_ssl = true
}

resource "awx_notification_template_pagerduty" "smoke" {
  name         = "Pagerduty smoke"
  organization = awx_organization.org.id
  subdomain    = "example"
  token        = "pagerduty-secret-rotated"
  service_key  = "example-service-key"
  client_name  = "awx-prod"
}

resource "awx_notification_template_rocketchat" "smoke" {
  name                     = "Rocket.Chat smoke"
  organization             = awx_organization.org.id
  rocketchat_url           = "https://rocketchat.example.com/hooks/smoke"
  rocketchat_no_verify_ssl = true
}

resource "awx_notification_template_slack" "smoke" {
  name         = "Slack smoke"
  organization = awx_organization.org.id
  token        = "slack-secret-rotated"
  channels     = ["#ops", "#alerts", "#incidents"]
}

resource "awx_notification_template_twilio" "smoke" {
  name          = "Twilio smoke"
  organization  = awx_organization.org.id
  account_sid   = "ACupdated"
  account_token = "twilio-secret-rotated"
  from_number   = "+15005550006"
  to_numbers    = ["+15005550009"]
}

resource "awx_notification_template_webhook" "smoke" {
  name                     = "Webhook smoke (updated)"
  description              = "Updated by TestIntegration_NotificationTemplates"
  organization             = awx_organization.org.id
  url                      = "https://hooks.example.com/awx-v2"
  username                 = "service-account"
  password                 = "webhook-secret-rotated"
  headers                  = { "X-Source" = "awx" }
  http_method              = "PUT"
  disable_ssl_verification = true
}

data "awx_notification_template_webhook" "webhook_by_name" {
  name = awx_notification_template_webhook.smoke.name
}

data "awx_notification_template_irc" "irc_by_name" {
  name = awx_notification_template_irc.smoke.name
}
