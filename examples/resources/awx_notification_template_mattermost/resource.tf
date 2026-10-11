resource "awx_notification_template_mattermost" "example" {
  name                     = "Mattermost"
  organization             = awx_organization.typed.id
  mattermost_url           = "https://mattermost.example.com/hooks/example"
  mattermost_no_verify_ssl = false
}
