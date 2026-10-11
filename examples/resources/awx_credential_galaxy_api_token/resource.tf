resource "awx_credential_galaxy_api_token" "example" {
  name         = "Ansible Galaxy/Automation Hub API Token"
  organization = awx_organization.typed.id
  url          = "https://galaxy.ansible.com/"
}
