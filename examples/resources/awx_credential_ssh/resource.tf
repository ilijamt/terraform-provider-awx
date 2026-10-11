resource "awx_credential_ssh" "example" {
  name         = "Machine"
  organization = awx_organization.typed.id
}
