resource "awx_application" "application" {
  organization             = awx_organization.application.id
  authorization_grant_type = "authorization-code"
  name                     = "Application"
  client_type              = "confidential"
  redirect_uris            = "https://localhost"
}
