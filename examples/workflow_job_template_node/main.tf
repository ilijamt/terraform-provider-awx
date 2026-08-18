terraform {
  required_providers {
    awx = {
      source = "ilijamt/awx"
    }
  }
}

provider "awx" {}

resource "awx_organization" "workflow" {
  name = "Workflow nodes"
}

resource "awx_inventory" "workflow" {
  name         = "Example workflow inventory"
  organization = awx_organization.workflow.id
}

resource "awx_project" "workflow" {
  name         = "Example workflow project"
  organization = awx_organization.workflow.id
  scm_type     = "git"
  scm_url      = "https://github.com/ansible/ansible-tower-samples"

  wait_for_sync = true

  timeouts {
    create = "10m"
    update = "5m"
  }
}

resource "awx_job_template" "build" {
  name      = "Example build"
  job_type  = "run"
  inventory = awx_inventory.workflow.id
  project   = awx_project.workflow.id
  playbook  = "hello_world.yml"
}

# A node may only override a launch-time field the job template prompts for.
# Without these three, setting limit/job_tags/verbosity on the node below fails
# with "Field is not configured to prompt on launch."
resource "awx_job_template" "test" {
  name                    = "Example test"
  job_type                = "run"
  inventory               = awx_inventory.workflow.id
  project                 = awx_project.workflow.id
  playbook                = "hello_world.yml"
  ask_limit_on_launch     = true
  ask_tags_on_launch      = true
  ask_verbosity_on_launch = true
}

resource "awx_job_template" "notify" {
  name      = "Example notify"
  job_type  = "run"
  inventory = awx_inventory.workflow.id
  project   = awx_project.workflow.id
  playbook  = "hello_world.yml"
}

resource "awx_workflow_job_template" "deploy" {
  name         = "Example deploy workflow"
  description  = "build -> test on success, notify on failure"
  organization = awx_organization.workflow.id
}

# Leave identifier unset and AWX generates a uuid. Set it and the value must be
# unique within the workflow. Changing workflow_job_template forces replacement:
# AWX turns the field read-only once the node exists.
resource "awx_workflow_job_template_node" "build" {
  workflow_job_template = awx_workflow_job_template.deploy.id
  unified_job_template  = awx_job_template.build.id
  identifier            = "build"
}

resource "awx_workflow_job_template_node" "test" {
  workflow_job_template = awx_workflow_job_template.deploy.id
  unified_job_template  = awx_job_template.test.id
  identifier            = "test"

  # Accepted only because awx_job_template.test sets the matching ask_*_on_launch
  # flags. verbosity is a string because AWX reports it as a choice.
  limit     = "web"
  job_tags  = "deploy"
  verbosity = "1"
}

resource "awx_workflow_job_template_node" "notify" {
  workflow_job_template = awx_workflow_job_template.deploy.id
  unified_job_template  = awx_job_template.notify.id
  identifier            = "notify"
}

# Referencing both node ids gives Terraform the dependency edges, so children are
# created before the edge and torn down after it.
resource "awx_workflow_job_template_node_associate_success_node" "build_to_test" {
  workflow_job_template_node_id = awx_workflow_job_template_node.build.id
  success_node_id               = awx_workflow_job_template_node.test.id
}

resource "awx_workflow_job_template_node_associate_failure_node" "build_to_notify" {
  workflow_job_template_node_id = awx_workflow_job_template_node.build.id
  failure_node_id               = awx_workflow_job_template_node.notify.id
}

# always_nodes run whatever the parent's outcome. AWX rejects a child already
# linked to the same parent through another relationship, and rejects cycles.
resource "awx_workflow_job_template_node_associate_always_node" "test_to_notify" {
  workflow_job_template_node_id = awx_workflow_job_template_node.test.id
  always_node_id                = awx_workflow_job_template_node.notify.id
}

# An approval node is a plain node whose unified_job_template is left unset,
# then converted. AWX has no collection for approval templates, so this resource
# posts to a sub-endpoint of the node and manages the result separately.
resource "awx_workflow_job_template_node" "gate" {
  workflow_job_template = awx_workflow_job_template.deploy.id
  identifier            = "gate"
}

resource "awx_workflow_job_template_node_approval" "gate" {
  workflow_job_template_node_id = awx_workflow_job_template_node.gate.id
  name                          = "Approve deploy"
  description                   = "Pauses the workflow until someone approves"
  timeout                       = 3600
}

# Approval nodes take the same link resources as any other node.
resource "awx_workflow_job_template_node_associate_success_node" "gate_to_build" {
  workflow_job_template_node_id = awx_workflow_job_template_node.gate.id
  success_node_id               = awx_workflow_job_template_node.build.id
}

data "awx_workflow_job_template_node" "build" {
  id = awx_workflow_job_template_node.build.id

  depends_on = [
    awx_workflow_job_template_node_associate_success_node.build_to_test,
    awx_workflow_job_template_node_associate_failure_node.build_to_notify,
  ]
}

output "build_node_identifier" {
  value = data.awx_workflow_job_template_node.build.identifier
}

# Read-only on the node. Without depends_on the data source can read it before
# the link resources have run.
output "build_node_links" {
  value = {
    success = data.awx_workflow_job_template_node.build.success_nodes
    failure = data.awx_workflow_job_template_node.build.failure_nodes
    always  = data.awx_workflow_job_template_node.build.always_nodes
  }
}
