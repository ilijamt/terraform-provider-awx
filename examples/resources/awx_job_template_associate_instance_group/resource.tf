resource "awx_job_template_associate_instance_group" "ig" {
  instance_group_id = awx_instance_group.ig.id
  job_template_id   = awx_job_template.instance_group.id
}
