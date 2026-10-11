resource "awx_credential_aws" "example" {
  name         = "Amazon Web Services"
  organization = awx_organization.typed.id
  username     = "service-account"
  password     = var.secrets["aws_password"]
}
