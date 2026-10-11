resource "awx_credential_gce" "example" {
  name         = "Google Compute Engine"
  organization = awx_organization.typed.id
  username     = "service-account@my-project.iam.gserviceaccount.com"
  ssh_key_data = var.ssh_private_key
}
