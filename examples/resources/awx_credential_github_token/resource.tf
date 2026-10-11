resource "awx_credential_github_token" "example" {
  name         = "GitHub Personal Access Token"
  organization = awx_organization.typed.id
  token        = var.secrets["github_token_token"]
}
