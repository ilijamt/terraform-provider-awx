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
  description  = "build -> test on success, notify on failure and after test"
  organization = awx_organization.workflow.id
}

resource "awx_workflow_job_template_node" "build" {
  workflow_job_template = awx_workflow_job_template.deploy.id
  unified_job_template  = awx_workflow_job_template.child.id
  identifier            = "build-updated"
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
  all_parents_must_converge = true
}

resource "awx_workflow_job_template_node_associate_success_node" "build_to_test" {
  workflow_job_template_node_id = awx_workflow_job_template_node.build.id
  success_node_id               = awx_workflow_job_template_node.test.id
}

resource "awx_workflow_job_template_node_associate_failure_node" "build_to_notify" {
  workflow_job_template_node_id = awx_workflow_job_template_node.build.id
  failure_node_id               = awx_workflow_job_template_node.notify.id
}

# Added by the update step. The parent is test rather than build, so it does not
# trip AWX's "Relationship not allowed" rule for one child on one parent.
resource "awx_workflow_job_template_node_associate_always_node" "test_to_notify" {
  workflow_job_template_node_id = awx_workflow_job_template_node.test.id
  always_node_id                = awx_workflow_job_template_node.notify.id
}

resource "awx_workflow_job_template_node" "gate" {
  workflow_job_template = awx_workflow_job_template.deploy.id
  identifier            = "gate"
}

# AWX has no collection for approval templates, so this posts to a sub-endpoint
# of the node and manages the result on /workflow_approval_templates/.
resource "awx_workflow_job_template_node_approval" "gate" {
  workflow_job_template_node_id = awx_workflow_job_template_node.gate.id
  name                          = "Approve deploy (updated)"
  description                   = "Pauses until someone approves"
  timeout                       = 7200
}

resource "awx_workflow_job_template_node_associate_success_node" "gate_to_build" {
  workflow_job_template_node_id = awx_workflow_job_template_node.gate.id
  success_node_id               = awx_workflow_job_template_node.build.id
}

data "awx_workflow_job_template_node" "build" {
  id = awx_workflow_job_template_node.build.id
}
