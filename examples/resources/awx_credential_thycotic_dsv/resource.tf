resource "awx_credential_thycotic_dsv" "example" {
  name          = "Thycotic DevOps Secrets Vault"
  organization  = awx_organization.typed.id
  tenant        = "tenant-id"
  tld           = "com"
  client_id     = "client-id"
  client_secret = var.secrets["thycotic_dsv_client_secret"]
}
