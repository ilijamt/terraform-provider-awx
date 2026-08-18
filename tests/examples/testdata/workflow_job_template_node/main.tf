resource "awx_organization" "workflow" {
  name = "Workflow nodes"
}

# A second workflow stands in as the unified_job_template so no project sync
# lands in the cassette.
resource "awx_workflow_job_template" "child" {
  name         = "Example child workflow"
  organization = awx_organization.workflow.id
}

resource "awx_workflow_job_template" "deploy" {
  name         = "Example deploy workflow"
  description  = "build -> test on success, notify on failure"
  organization = awx_organization.workflow.id
}

resource "awx_workflow_job_template_node" "build" {
  workflow_job_template = awx_workflow_job_template.deploy.id
  unified_job_template  = awx_workflow_job_template.child.id
  identifier            = "build"
}

resource "awx_workflow_job_template_node" "test" {
  workflow_job_template = awx_workflow_job_template.deploy.id
  unified_job_template  = awx_workflow_job_template.child.id
  identifier            = "test"
}

# identifier left unset so AWX generates one. A pinned default would give every
# node the same value and collide with the per-workflow constraint.
resource "awx_workflow_job_template_node" "notify" {
  workflow_job_template     = awx_workflow_job_template.deploy.id
  unified_job_template      = awx_workflow_job_template.child.id
  all_parents_must_converge = false
}

resource "awx_workflow_job_template_node_associate_success_node" "build_to_test" {
  workflow_job_template_node_id = awx_workflow_job_template_node.build.id
  success_node_id               = awx_workflow_job_template_node.test.id
}

resource "awx_workflow_job_template_node_associate_failure_node" "build_to_notify" {
  workflow_job_template_node_id = awx_workflow_job_template_node.build.id
  failure_node_id               = awx_workflow_job_template_node.notify.id
}

resource "awx_workflow_job_template_node" "gate" {
  workflow_job_template = awx_workflow_job_template.deploy.id
  identifier            = "gate"
}

# AWX has no collection for approval templates, so this posts to a sub-endpoint
# of the node and manages the result on /workflow_approval_templates/.
resource "awx_workflow_job_template_node_approval" "gate" {
  workflow_job_template_node_id = awx_workflow_job_template_node.gate.id
  name                          = "Approve deploy"
  description                   = "Pauses until someone approves"
  timeout                       = 3600
}

resource "awx_workflow_job_template_node_associate_success_node" "gate_to_build" {
  workflow_job_template_node_id = awx_workflow_job_template_node.gate.id
  success_node_id               = awx_workflow_job_template_node.build.id
}

data "awx_workflow_job_template_node" "build" {
  id = awx_workflow_job_template_node.build.id
}
