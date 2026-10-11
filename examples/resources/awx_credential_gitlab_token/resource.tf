resource "awx_credential_gitlab_token" "example" {
  name         = "GitLab Personal Access Token"
  organization = awx_organization.typed.id
  token        = var.secrets["gitlab_token_token"]
}
