resource "awx_notification_template_pagerduty" "example" {
  name         = "Pagerduty"
  organization = awx_organization.typed.id
  subdomain    = "example"
  token        = var.secrets["pagerduty_token"]
  service_key  = "example-service-key"
  client_name  = "awx"
}
