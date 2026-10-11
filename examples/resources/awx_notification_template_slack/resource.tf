resource "awx_notification_template_slack" "example" {
  name         = "Slack"
  organization = awx_organization.typed.id
  token        = var.secrets["slack_token"]
  channels     = ["#ops", "#alerts"]
}
