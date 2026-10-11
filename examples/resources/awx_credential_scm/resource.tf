resource "awx_credential_scm" "example" {
  name         = "Source Control"
  organization = awx_organization.typed.id
}
