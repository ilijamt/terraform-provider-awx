resource "awx_notification_template_grafana" "example" {
  name         = "Grafana"
  organization = awx_organization.typed.id
  grafana_url  = "https://grafana.example.com"
  grafana_key  = var.secrets["grafana_key"]
}
