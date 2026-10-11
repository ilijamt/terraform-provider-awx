resource "awx_credential_hashivault_ssh" "example" {
  name         = "HashiCorp Vault Signed SSH"
  organization = awx_organization.typed.id
  url          = "https://service.example.com"
}
