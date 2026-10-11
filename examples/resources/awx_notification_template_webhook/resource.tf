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
