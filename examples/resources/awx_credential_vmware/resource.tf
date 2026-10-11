resource "awx_credential_vmware" "example" {
  name         = "VMware vCenter"
  organization = awx_organization.typed.id
  host         = "https://service.example.com"
  username     = "service-account"
  password     = var.secrets["vmware_password"]
}
