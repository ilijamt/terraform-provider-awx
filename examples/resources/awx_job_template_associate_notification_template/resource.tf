resource "awx_job_template_associate_notification_template" "demo_job_template" {
  notification_template_id = awx_notification_template.demo_webhook_notification.id
  option                   = "started"
  job_template_id          = awx_job_template.demo_job_template.id
}
