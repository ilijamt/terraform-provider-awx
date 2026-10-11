resource "awx_credential_terraform" "example" {
  name          = "Terraform backend configuration"
  organization  = awx_organization.typed.id
  configuration = var.secrets["terraform_configuration"]
}
