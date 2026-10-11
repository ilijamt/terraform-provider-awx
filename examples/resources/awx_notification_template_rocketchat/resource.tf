resource "awx_notification_template_rocketchat" "example" {
  name                     = "Rocket.Chat"
  organization             = awx_organization.typed.id
  rocketchat_url           = "https://rocketchat.example.com/hooks/example"
  rocketchat_no_verify_ssl = false
}
