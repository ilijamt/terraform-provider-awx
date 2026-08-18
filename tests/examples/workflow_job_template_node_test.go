//go:build integration

package examples

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/ilijamt/terraform-provider-awx/internal/awx"
	"github.com/ilijamt/terraform-provider-awx/internal/provider"
	"github.com/ilijamt/terraform-provider-awx/version"
)

// TestIntegration_WorkflowJobTemplateNode covers workflow nodes, the
// success/failure/always link resources, and the approval gate. The fixture
// points nodes at a second workflow so no project sync lands in the cassette.
func TestIntegration_WorkflowJobTemplateNode(t *testing.T) {
	httpClient := NewVCRClient(t, "workflow_job_template_node")
	cfg := ReadFixture(t, filepath.Join("workflow_job_template_node", "main.tf"))
	updated := ReadFixture(t, filepath.Join("workflow_job_template_node", "update.tf"))

	factories := map[string]func() (tfprotov6.ProviderServer, error){
		"awx": providerserver.NewProtocol6WithError(
			provider.NewFuncProvider(version.Version, httpClient, awx.Resources(), awx.DataSources())(),
		),
	}

	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: factories,
		Steps: []resource.TestStep{
			{
				Config: providerHeader(t) + cfg,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("awx_workflow_job_template_node.build", "identifier", "build"),
					resource.TestCheckResourceAttrPair(
						"awx_workflow_job_template_node.build", "workflow_job_template",
						"awx_workflow_job_template.deploy", "id"),
					// No identifier in config, so AWX generates a unique one.
					// A pinned default would collide on the second node.
					resource.TestCheckResourceAttrSet("awx_workflow_job_template_node.notify", "identifier"),
					resource.TestCheckResourceAttr("awx_workflow_job_template_node_approval.gate", "name", "Approve deploy"),
					resource.TestCheckResourceAttr("awx_workflow_job_template_node_approval.gate", "timeout", "3600"),
					resource.TestCheckResourceAttrSet("awx_workflow_job_template_node_approval.gate", "id"),
					resource.TestCheckResourceAttr("awx_workflow_job_template_node.notify", "all_parents_must_converge", "false"),
					resource.TestCheckResourceAttrPair(
						"awx_workflow_job_template_node_associate_success_node.build_to_test", "success_node_id",
						"awx_workflow_job_template_node.test", "id"),
					resource.TestCheckResourceAttrPair(
						"data.awx_workflow_job_template_node.build", "identifier",
						"awx_workflow_job_template_node.build", "identifier"),
				),
			},
			{
				Config: providerHeader(t) + updated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("awx_workflow_job_template_node.build", "identifier", "build-updated"),
					resource.TestCheckResourceAttr("awx_workflow_job_template_node.notify", "all_parents_must_converge", "true"),
					resource.TestCheckResourceAttrPair(
						"awx_workflow_job_template_node_associate_always_node.test_to_notify", "always_node_id",
						"awx_workflow_job_template_node.notify", "id"),
					resource.TestCheckResourceAttr("awx_workflow_job_template_node_approval.gate", "name", "Approve deploy (updated)"),
					resource.TestCheckResourceAttr("awx_workflow_job_template_node_approval.gate", "timeout", "7200"),
					// Step 1 creates the node before the association and never
					// re-reads it, so links only show from step 2 on.
					resource.TestCheckResourceAttr("awx_workflow_job_template_node.build", "success_nodes.#", "1"),
					resource.TestCheckResourceAttrPair(
						"awx_workflow_job_template_node.build", "success_nodes.0",
						"awx_workflow_job_template_node.test", "id"),
					resource.TestCheckResourceAttrPair(
						"awx_workflow_job_template_node.build", "failure_nodes.0",
						"awx_workflow_job_template_node.notify", "id"),
					resource.TestCheckResourceAttrPair(
						"data.awx_workflow_job_template_node.build", "success_nodes.0",
						"awx_workflow_job_template_node.test", "id"),
				),
			},
			{
				ResourceName:      "awx_workflow_job_template_node.build",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// The template holds no reference back to its node, so the node
				// id has to come in on the import string.
				ResourceName:      "awx_workflow_job_template_node_approval.gate",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: approvalImportID,
			},
		},
	})
}

func approvalImportID(s *terraform.State) (string, error) {
	node, ok := s.RootModule().Resources["awx_workflow_job_template_node.gate"]
	if !ok {
		return "", fmt.Errorf("awx_workflow_job_template_node.gate not found in state")
	}
	approval, ok := s.RootModule().Resources["awx_workflow_job_template_node_approval.gate"]
	if !ok {
		return "", fmt.Errorf("awx_workflow_job_template_node_approval.gate not found in state")
	}
	return fmt.Sprintf("%s/%s", node.Primary.Attributes["id"], approval.Primary.ID), nil
}
