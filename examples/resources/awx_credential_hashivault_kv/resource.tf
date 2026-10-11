resource "awx_credential_hashivault_kv" "example" {
  name         = "HashiCorp Vault Secret Lookup"
  organization = awx_organization.typed.id
  url          = "https://service.example.com"
  api_version  = "v1"
}
