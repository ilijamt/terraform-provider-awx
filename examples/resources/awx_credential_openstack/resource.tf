resource "awx_credential_openstack" "example" {
  name         = "OpenStack"
  organization = awx_organization.typed.id
  username     = "service-account"
  password     = var.secrets["openstack_password"]
  host         = "https://service.example.com"
  project      = "my-project"
  verify_ssl   = true
}
