terraform {
  required_providers {
    awx = {
      source = "registry.terraform.io/ilijamt/awx"
    }
  }
}

provider "awx" {}

data "awx_credential_type" "ansible_galaxy" {
  name = "Ansible Galaxy/Automation Hub API Token"
}

data "awx_credential_type" "machine" {
  name = "Machine"
}

data "awx_credential" "ansible_galaxy" {
  name = "Ansible Galaxy"
}

resource "awx_credential" "demo_credential" {
  name            = "Demo Credential"
  credential_type = data.awx_credential_type.machine.id
  organization    = awx_organization.demo_organization.id
  inputs = jsonencode({
    username = "admin"
  })
}

data "awx_credential_object_roles" "demo_credential" {
  id = awx_credential.demo_credential.id
}

resource "awx_organization" "demo_organization" {
  name = "Demo Organization"
}

resource "awx_execution_environment" "demo_ee" {
  name  = "Demo EE"
  image = "quay.io/ansible/awx-ee:latest"
}

resource "awx_user" "demo_user" {
  username = "demo_user"
  email    = "demouser@example.com"
  password = "test"
}

resource "awx_team" "demo_team" {
  name         = "Demo Team"
  organization = awx_organization.demo_organization.id
}

data "awx_team_object_roles" "demo_team" {
  id = awx_team.demo_team.id
}

resource "awx_user_associate_role" "user_admin_of_org" {
  user_id = awx_user.demo_user.id
  role_id = data.awx_team_object_roles.demo_team.roles["Admin"]
}

resource "awx_team_associate_role" "team_access_to_credential" {
  team_id = awx_team.demo_team.id
  role_id = data.awx_credential_object_roles.demo_credential.roles["Use"]
}

resource "awx_project" "demo_project" {
  name         = "Demo Project"
  organization = awx_organization.demo_organization.id
  scm_url      = "https://github.com/ansible/ansible-tower-samples"
  scm_type     = "git"
  scm_clean    = false
}

resource "awx_inventory" "demo_inventory" {
  name         = "Demo Inventory"
  organization = awx_organization.demo_organization.id
}

data "awx_inventory_object_roles" "demo_inventory" {
  id = awx_inventory.demo_inventory.id
}

resource "awx_team_associate_role" "team_access_to_inventory" {
  team_id = awx_team.demo_team.id
  role_id = data.awx_inventory_object_roles.demo_inventory.roles["Admin"]
}

resource "awx_host" "localhost" {
  name      = "localhost"
  inventory = awx_inventory.demo_inventory.id
  enabled   = true
  variables = jsonencode({
    "ansible_connection"         = "local"
    "ansible_python_interpreter" = "{{ ansible_playbook_python }}"
    "ansible_debug"              = true
  })
}

resource "awx_group" "localhost" {
  name      = "local"
  inventory = awx_inventory.demo_inventory.id
}

resource "awx_host_associate_group" "localhost_to_group_local" {
  host_id  = awx_host.localhost.id
  group_id = awx_group.localhost.id
}

resource "awx_job_template" "demo_job_template" {
  name               = "Demo Job Template"
  inventory          = awx_inventory.demo_inventory.id
  verbosity          = 0
  job_type           = "run"
  playbook           = "hello_world.yml"
  project            = awx_project.demo_project.id
  scm_branch         = ""
  job_slice_count    = 2
  allow_simultaneous = true
}

# resource "awx_job_template_survey_spec" "demo_job_template" {
#   job_template_id = awx_job_template.demo_job_template.id
#   spec = jsonencode([{
#     type                 = "float"
#     max                  = 1024
#     min                  = 0
#     choices              = ""
#     default              = ""
#     question_description = ""
#     new_question         = true
#     question_name        = "What is the percentage of failure?"
#     required             = true
#     variable             = "pct_failure"
#   }])
# }

resource "awx_job_template_associate_notification_template" "demo_job_template" {
  notification_template_id = awx_notification_template.demo_webhook_notification.id
  option                   = "started"
  job_template_id          = awx_job_template.demo_job_template.id
}

resource "awx_schedule" "demo_job" {
  enabled              = true
  name                 = "Run Demo Job every month"
  rrule                = "DTSTART;TZID=Europe/Amsterdam:20221111T103000 RRULE:INTERVAL=1;FREQ=MONTHLY;BYMONTHDAY=1"
  unified_job_template = awx_job_template.demo_job_template.id
}

resource "awx_job_template_associate_credential" "demo_machine" {
  credential_id   = awx_credential.demo_credential.id
  job_template_id = awx_job_template.demo_job_template.id
}

resource "awx_workflow_job_template" "demo_workflow_template" {
  name         = "Demo Workflow Job"
  organization = awx_organization.demo_organization.id
}

resource "awx_workflow_job_template_associate_notification_template" "demo_workflow_template" {
  notification_template_id = awx_notification_template.demo_webhook_notification.id
  option                   = "started"
  workflow_job_template_id = awx_workflow_job_template.demo_workflow_template.id
}


resource "awx_schedule" "demo_workflow_template" {
  enabled              = true
  name                 = "Run Demo Workflow Job every month"
  rrule                = "DTSTART;TZID=Europe/Amsterdam:20221111T103000 RRULE:INTERVAL=1;FREQ=MONTHLY;BYMONTHDAY=1"
  unified_job_template = awx_workflow_job_template.demo_workflow_template.id
}

data "awx_workflow_job_template_object_roles" "demo_workflow_template" {
  id = awx_workflow_job_template.demo_workflow_template.id
}

resource "awx_team_associate_role" "team_access_to_demo_workflow_template" {
  team_id = awx_team.demo_team.id
  role_id = data.awx_workflow_job_template_object_roles.demo_workflow_template.roles["Admin"]
}

resource "awx_notification_template" "demo_webhook_notification" {
  name              = "Demo Webhook Notification"
  notification_type = "webhook"
  organization      = awx_organization.demo_organization.id
  notification_configuration = jsonencode({
    "url"                      = "http://example.com"
    "http_method"              = "POST"
    "disable_ssl_verification" = true
    "username"                 = ""
    "password"                 = ""
    "headers"                  = {}
  })
}

