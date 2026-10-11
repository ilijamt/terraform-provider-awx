resource "awx_notification_template_irc" "example" {
  name         = "IRC"
  organization = awx_organization.typed.id
  server       = "irc.example.com"
  port         = 6697
  nickname     = "awx"
  password     = var.secrets["irc_password"]
  use_ssl      = true
  targets      = ["#ops"]
}
