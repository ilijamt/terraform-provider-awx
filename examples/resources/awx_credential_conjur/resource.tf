resource "awx_credential_conjur" "example" {
  name         = "CyberArk Conjur Secrets Manager Lookup"
  organization = awx_organization.typed.id
  url          = "https://service.example.com"
  api_key      = var.secrets["conjur_api_key"]
  account      = "account"
  username     = "service-account"
}
