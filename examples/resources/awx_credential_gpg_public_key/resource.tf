resource "awx_credential_gpg_public_key" "example" {
  name           = "GPG Public Key"
  organization   = awx_organization.typed.id
  gpg_public_key = var.gpg_public_key
}
