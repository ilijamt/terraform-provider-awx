resource "awx_organization_associate_galaxy_credential" "default_demo_organization" {
  galaxy_credential_id = data.awx_credential.ansible_galaxy.id
  organization_id      = awx_organization.demo_organization.id
}
