resource "awx_credential_registry" "example" {
  name         = "Container Registry"
  organization = awx_organization.typed.id
  host         = "quay.io"
  verify_ssl   = true
}
