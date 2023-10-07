terraform {
  required_providers {
    awx = {
      source = "registry.terraform.io/ilijamt/awx"
    }
  }
}

provider "awx" {}

data "awx_execution_environment" "latest" {
  name = "AWX EE (latest)"
}

data "awx_execution_environment" "control_plane_ee" {
  name = "Control Plane Execution Environment"
}

resource "awx_execution_environment" "test" {
  name  = "Default"
  image = "quay.io/ansible/awx-ee:latest"
}

resource "awx_organization" "test" {
  name                = "Test"
  default_environment = awx_execution_environment.test.id
}

data "awx_organization_object_roles" "test" {
  id = awx_organization.test.id
}

resource "awx_team" "test" {
  name         = "Test"
  organization = awx_organization.test.id
}

data "awx_team_object_roles" "test" {
  id = awx_team.test.id
}

data "awx_credential_type" "container_registry" {
  name = "Container Registry"
}

resource "awx_credential" "container_registry" {
  name            = "Container Registry"
  credential_type = data.awx_credential_type.container_registry.id
  organization    = awx_organization.test.id
  inputs = jsonencode({
    "host" : "quay.io",
    "verify_ssl" : true,
    "username" : "test",
  })
}

data "awx_credential_object_roles" "container_registry" {
  id = awx_credential.container_registry.id
}

data "awx_user" "admin" {
  username = "admin"
}

resource "awx_user" "ilijamt" {
  username     = "ilijamt"
  password     = "test"
  is_superuser = true
}

resource "awx_label" "test" {
  name         = "test-label"
  organization = awx_organization.test.id
}

data "awx_label" "test" {
  name         = "test-label"
  organization = awx_organization.test.id
  depends_on   = [awx_label.test]
}

data "awx_label" "test_by_id" {
  id         = awx_label.test.id
  depends_on = [awx_label.test]
}

resource "awx_user" "demo" {
  username = "demo"
  password = "test"
}

# resource "awx_job_template" "demo" {
#     allow_simultaneous                  = false
#     ask_credential_on_launch            = false
#     ask_diff_mode_on_launch             = false
#     ask_execution_environment_on_launch = false
#     ask_forks_on_launch                 = false
#     ask_instance_groups_on_launch       = false
#     ask_inventory_on_launch             = true
#     ask_job_slice_count_on_launch       = false
#     ask_job_type_on_launch              = false
#     ask_labels_on_launch                = false
#     ask_limit_on_launch                 = true
#     ask_scm_branch_on_launch            = false
#     ask_skip_tags_on_launch             = true
#     ask_tags_on_launch                  = true
#     ask_timeout_on_launch               = false
#     ask_variables_on_launch             = true
#     ask_verbosity_on_launch             = false
#     become_enabled                      = false
#     diff_mode                           = true
#     execution_environment               = awx_execution_environment.default_environment.id
#     extra_vars                          = jsonencode({})
#     force_handlers                      = false
#     forks                               = 0
#     job_slice_count                     = 1
#     job_type                            = "run"
#     name                                = "demo"
#     organization                        = awx_organization.test.id
#     playbook                            = "demo.yml"
#     prevent_instance_group_fallback     = false
#     project                             = 11
#     survey_enabled                      = false
#     timeout                             = 0
#     use_fact_cache                      = true
# }
