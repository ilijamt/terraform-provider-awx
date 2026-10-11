resource "awx_credential_insights" "example" {
  name         = "Insights"
  organization = awx_organization.typed.id
  username     = "service-account"
  password     = var.secrets["insights_password"]
}
