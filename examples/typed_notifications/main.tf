terraform {
  required_providers {
    awx = {
      source = "ilijamt/awx"
    }
  }
}

provider "awx" {}

resource "awx_organization" "typed" {
  name = "Typed Notifications"
}

# Every awssns field defaults, so a template with no configuration is legal.
resource "awx_notification_template_awssns" "example" {
  name                  = "AWS SNS"
  organization          = awx_organization.typed.id
  aws_region            = "eu-west-1"
  aws_access_key_id     = "AKIAEXAMPLE"
  aws_secret_access_key = var.secrets["awssns_secret_access_key"]
  sns_topic_arn         = "arn:aws:sns:eu-west-1:123456789012:alerts"
}

resource "awx_notification_template_email" "example" {
  name         = "Email"
  organization = awx_organization.typed.id
  host         = "smtp.example.com"
  port         = 587
  username     = "service-account"
  password     = var.secrets["email_password"]
  use_tls      = true
  use_ssl      = false
  sender       = "awx@example.com"
  recipients   = ["ops@example.com"]
}

resource "awx_notification_template_grafana" "example" {
  name         = "Grafana"
  organization = awx_organization.typed.id
  grafana_url  = "https://grafana.example.com"
  grafana_key  = var.secrets["grafana_key"]
}

resource "awx_notification_template_irc" "example" {
  name         = "IRC"
  organization = awx_organization.typed.id
  server       = "irc.example.com"
  port         = 6697
  nickname     = "awx"
  password     = var.secrets["irc_password"]
  use_ssl      = true
  targets      = ["#ops"]
}

resource "awx_notification_template_mattermost" "example" {
  name                     = "Mattermost"
  organization             = awx_organization.typed.id
  mattermost_url           = "https://mattermost.example.com/hooks/example"
  mattermost_no_verify_ssl = false
}

resource "awx_notification_template_pagerduty" "example" {
  name         = "Pagerduty"
  organization = awx_organization.typed.id
  subdomain    = "example"
  token        = var.secrets["pagerduty_token"]
  service_key  = "example-service-key"
  client_name  = "awx"
}

resource "awx_notification_template_rocketchat" "example" {
  name                     = "Rocket.Chat"
  organization             = awx_organization.typed.id
  rocketchat_url           = "https://rocketchat.example.com/hooks/example"
  rocketchat_no_verify_ssl = false
}

resource "awx_notification_template_slack" "example" {
  name         = "Slack"
  organization = awx_organization.typed.id
  token        = var.secrets["slack_token"]
  channels     = ["#ops", "#alerts"]
}

resource "awx_notification_template_twilio" "example" {
  name          = "Twilio"
  organization  = awx_organization.typed.id
  account_sid   = "ACexample"
  account_token = var.secrets["twilio_account_token"]
  from_number   = "+15005550006"
  to_numbers    = ["+15005550001"]
}

# headers has no default, so AWX requires it even when empty.
resource "awx_notification_template_webhook" "example" {
  name                     = "Webhook"
  organization             = awx_organization.typed.id
  url                      = "https://hooks.example.com/awx"
  http_method              = "POST"
  headers                  = { "X-Source" = "awx" }
  username                 = "service-account"
  password                 = var.secrets["webhook_password"]
  disable_ssl_verification = false
}

data "awx_notification_template_slack" "by_name" {
  name       = awx_notification_template_slack.example.name
  depends_on = [awx_notification_template_slack.example]
}
