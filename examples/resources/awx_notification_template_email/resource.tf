resource "awx_notification_template_email" "example" {
  name         = "Email"
  organization = awx_organization.typed.id
  host         = "smtp.example.com"
  port         = 587
  username     = "service-account"
  password     = var.secrets["email_password"]
  use_tls      = true
  use_ssl      = false
  sender       = "awx@example.com"
  recipients   = ["ops@example.com"]
}
