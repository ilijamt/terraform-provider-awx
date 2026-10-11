resource "awx_credential_azure_kv" "example" {
  name         = "Microsoft Azure Key Vault"
  organization = awx_organization.typed.id
  url          = "https://service.example.com"
  client       = "client-id"
  secret       = var.secrets["azure_kv_secret"]
  tenant       = "tenant-id"
  cloud_name   = "AzureCloud"
}
