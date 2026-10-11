resource "awx_credential_controller" "example" {
  name         = "Red Hat Ansible Automation Platform"
  organization = awx_organization.typed.id
  host         = "https://service.example.com"
  verify_ssl   = false
}
