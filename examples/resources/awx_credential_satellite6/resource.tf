resource "awx_credential_satellite6" "example" {
  name         = "Red Hat Satellite 6"
  organization = awx_organization.typed.id
  host         = "https://service.example.com"
  username     = "service-account"
  password     = var.secrets["satellite6_password"]
}
