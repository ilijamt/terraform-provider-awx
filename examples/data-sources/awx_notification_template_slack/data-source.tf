data "awx_notification_template_slack" "by_name" {
  name       = awx_notification_template_slack.example.name
  depends_on = [awx_notification_template_slack.example]
}
