resource "awx_credential_bitbucket_dc_token" "example" {
  name         = "Bitbucket Data Center HTTP Access Token"
  organization = awx_organization.typed.id
  token        = var.secrets["bitbucket_dc_token_token"]
}
