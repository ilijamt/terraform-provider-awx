resource "awx_workflow_job_template_associate_notification_template" "demo_workflow_template" {
  notification_template_id = awx_notification_template.demo_webhook_notification.id
  option                   = "started"
  workflow_job_template_id = awx_workflow_job_template.demo_workflow_template.id
}
