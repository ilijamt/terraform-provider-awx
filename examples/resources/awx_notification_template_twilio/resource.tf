resource "awx_notification_template_twilio" "example" {
  name          = "Twilio"
  organization  = awx_organization.typed.id
  account_sid   = "ACexample"
  account_token = var.secrets["twilio_account_token"]
  from_number   = "+15005550006"
  to_numbers    = ["+15005550001"]
}
