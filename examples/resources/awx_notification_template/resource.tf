resource "awx_notification_template" "demo_webhook_notification" {
  name              = "Demo Webhook Notification"
  notification_type = "webhook"
  organization      = awx_organization.demo_organization.id
  notification_configuration = jsonencode({
    "url"                      = "http://example.com"
    "http_method"              = "POST"
    "disable_ssl_verification" = true
    "username"                 = ""
    "password"                 = ""
    "headers"                  = {}
  })
}
