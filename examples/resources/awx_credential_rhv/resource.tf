resource "awx_credential_rhv" "example" {
  name         = "Red Hat Virtualization"
  organization = awx_organization.typed.id
  host         = "https://service.example.com"
  username     = "service-account"
  password     = var.secrets["rhv_password"]
}
