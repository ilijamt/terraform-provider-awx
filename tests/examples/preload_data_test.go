//go:build integration

package examples

import (
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/ilijamt/terraform-provider-awx/internal/awx"
	"github.com/ilijamt/terraform-provider-awx/internal/provider"
	"github.com/ilijamt/terraform-provider-awx/version"
)

// TestIntegration_PreloadData exercises examples/preload_data — the
// kitchen-sink fixture that wires an organization, team, user,
// credentials, project (with wait_for_sync), inventory, host/group,
// job/workflow templates, schedules, and a webhook notification
// template together. It proves the broad set of resources cooperate
// end-to-end against a recorded cassette.
//
// The update step flips three small fields (team name, host enabled,
// job_template verbosity) so the diff between main.tf and update.tf
// stays readable. Import targets the organization — most resources in
// this fixture are association/role types whose import surface is
// covered elsewhere.
func TestIntegration_PreloadData(t *testing.T) {
	httpClient := NewVCRClient(t, "preload_data")
	cfg := ReadFixture(t, filepath.Join("preload_data", "main.tf"))
	updated := ReadFixture(t, filepath.Join("preload_data", "update.tf"))

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
					resource.TestCheckResourceAttr("awx_organization.demo_organization", "name", "Demo Organization"),
					resource.TestCheckResourceAttr("awx_team.demo_team", "name", "Demo Team"),
					resource.TestCheckResourceAttr("awx_user.demo_user", "username", "demo_user"),

					resource.TestCheckResourceAttr("awx_credential.ansible_galaxy", "name", "Ansible Galaxy"),
					resource.TestCheckResourceAttr("awx_credential.demo_credential", "name", "Demo Credential"),
					resource.TestCheckResourceAttr("awx_execution_environment.demo_ee", "name", "Demo EE"),

					resource.TestCheckResourceAttr("awx_project.demo_project", "name", "Demo Project"),
					resource.TestCheckResourceAttr("awx_project.demo_project", "wait_for_sync", "true"),
					resource.TestCheckResourceAttr("awx_inventory.demo_inventory", "name", "Demo Inventory"),
					resource.TestCheckResourceAttr("awx_host.localhost", "name", "localhost"),
					resource.TestCheckResourceAttr("awx_host.localhost", "enabled", "true"),
					resource.TestCheckResourceAttr("awx_group.localhost", "name", "local"),

					resource.TestCheckResourceAttr("awx_job_template.demo_job_template", "name", "Demo Job Template"),
					resource.TestCheckResourceAttr("awx_job_template.demo_job_template", "verbosity", "0"),
					resource.TestCheckResourceAttr("awx_workflow_job_template.demo_workflow_template", "name", "Demo Workflow Job"),
					resource.TestCheckResourceAttr("awx_notification_template.demo_webhook_notification", "name", "Demo Webhook Notification"),
					resource.TestCheckResourceAttr("awx_schedule.demo_job", "name", "Run Demo Job every month"),
					resource.TestCheckResourceAttr("awx_schedule.demo_workflow_template", "name", "Run Demo Workflow Job every month"),

					resource.TestCheckResourceAttrPair("data.awx_inventory_object_roles.demo_inventory", "id", "awx_inventory.demo_inventory", "id"),
					resource.TestCheckResourceAttrSet("data.awx_team_object_roles.demo_team", "roles.%"),
					resource.TestCheckResourceAttrSet("data.awx_credential_object_roles.demo_credential", "roles.%"),
					resource.TestCheckResourceAttrSet("data.awx_workflow_job_template_object_roles.demo_workflow_template", "roles.%"),
				),
			},
			{
				Config: providerHeader(t) + updated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("awx_team.demo_team", "name", "Demo Team (updated)"),
					resource.TestCheckResourceAttr("awx_host.localhost", "enabled", "false"),
					resource.TestCheckResourceAttr("awx_job_template.demo_job_template", "verbosity", "1"),
				),
			},
			{
				ResourceName:      "awx_organization.demo_organization",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
