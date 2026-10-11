resource "awx_credential_centrify_vault_kv" "example" {
  name            = "Centrify Vault Credential Provider Lookup"
  organization    = awx_organization.typed.id
  url             = "https://service.example.com"
  client_id       = "client-id"
  client_password = var.secrets["centrify_vault_kv_client_password"]
}
