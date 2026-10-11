resource "awx_credential_aim" "example" {
  name         = "CyberArk Central Credential Provider Lookup"
  organization = awx_organization.typed.id
  url          = "https://service.example.com"
  app_id       = var.secrets["aim_app_id"]
  verify       = true
}
