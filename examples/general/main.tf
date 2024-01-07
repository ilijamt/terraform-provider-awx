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

resource "awx_instance_group" "ig" {
  name               = "Demo Instance Group"
  is_container_group = true
  pod_spec_override = jsonencode({
    "apiVersion" : "v1",
    "kind" : "Pod",
    "metadata" : {
      "namespace" : "awx"
    },
    "spec" : {
      "serviceAccountName" : "default",
      "automountServiceAccountToken" : false,
      "containers" : [
        {
          "image" : "quay.io/ansible/awx-ee:latest",
          "name" : "worker",
          "args" : [
            "ansible-runner",
            "worker",
            "--private-data-dir=/runner"
          ],
          "resources" : {
            "requests" : {
              "cpu" : "250m",
              "memory" : "100Mi"
            }
          }
        }
      ]
    }
  })
}

resource "awx_project" "demo_project" {
  name         = "Job Demo Project"
  organization = awx_organization.test.id
  scm_url      = "https://github.com/ansible/ansible-tower-samples"
  scm_type     = "git"
  scm_clean    = false
}

resource "awx_inventory" "demo_inventory" {
  name         = "Job Demo Inventory"
  organization = awx_organization.test.id
}


resource "awx_job_template" "instance_group" {
  name               = "Job Demo Instance Group"
  inventory          = awx_inventory.demo_inventory.id
  verbosity          = 0
  job_type           = "run"
  playbook           = "hello_world.yml"
  project            = awx_project.demo_project.id
  scm_branch         = ""
  job_slice_count    = 2
  allow_simultaneous = true
}

resource "awx_job_template_associate_instance_group" "ig" {
  instance_group_id = awx_instance_group.ig.id
  job_template_id   = awx_job_template.instance_group.id
}