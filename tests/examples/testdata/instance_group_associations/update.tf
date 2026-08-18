resource "awx_organization" "org" {
  name = "Instance Group Association Organization"
}

resource "awx_instance_group" "first" {
  name = "Association Group One"
}

resource "awx_instance_group" "second" {
  name = "Association Group Two"
}

resource "awx_project" "project" {
  name         = "Instance Group Association Project"
  organization = awx_organization.org.id
  scm_url      = "https://github.com/ansible/ansible-tower-samples"
  scm_type     = "git"
  scm_clean    = false

  wait_for_sync = true

  timeouts {
    create = "10m"
    update = "5m"
  }
}

resource "awx_inventory" "inventory" {
  name         = "Instance Group Association Inventory"
  organization = awx_organization.org.id
}

resource "awx_job_template" "job_template" {
  name      = "Instance Group Association Job Template"
  inventory = awx_inventory.inventory.id
  job_type  = "run"
  playbook  = "hello_world.yml"
  project   = awx_project.project.id
}

resource "awx_organization_associate_instance_group" "org_link" {
  organization_id   = awx_organization.org.id
  instance_group_id = awx_instance_group.second.id
}

resource "awx_job_template_associate_instance_group" "jt_link" {
  job_template_id   = awx_job_template.job_template.id
  instance_group_id = awx_instance_group.second.id
}
