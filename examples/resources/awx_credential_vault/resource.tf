resource "awx_credential_vault" "example" {
  name           = "Vault"
  organization   = awx_organization.typed.id
  vault_password = var.secrets["vault_vault_password"]
}
