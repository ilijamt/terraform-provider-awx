resource "awx_notification_template_awssns" "example" {
  name                  = "AWS SNS"
  organization          = awx_organization.typed.id
  aws_region            = "eu-west-1"
  aws_access_key_id     = "AKIAEXAMPLE"
  aws_secret_access_key = var.secrets["awssns_secret_access_key"]
  sns_topic_arn         = "arn:aws:sns:eu-west-1:123456789012:alerts"
}
