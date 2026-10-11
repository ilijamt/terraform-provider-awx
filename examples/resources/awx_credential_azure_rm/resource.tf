resource "awx_credential_azure_rm" "example" {
  name         = "Microsoft Azure Resource Manager"
  organization = awx_organization.typed.id
  subscription = "subscription-id"
}
