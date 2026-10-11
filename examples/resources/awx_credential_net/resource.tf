resource "awx_credential_net" "example" {
  name         = "Network"
  organization = awx_organization.typed.id
  username     = "service-account"
  authorize    = false
}
