resource "awx_credential_thycotic_tss" "example" {
  name         = "Thycotic Secret Server"
  organization = awx_organization.typed.id
  server_url   = "https://tss.example.com"
  username     = "service-account"
  password     = var.secrets["thycotic_tss_password"]
}
