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

// AWX allows only GET and DELETE on an assignment, so the update step swaps the
// role definition and expects a replacement rather than a patch.
func TestIntegration_RoleAssignments(t *testing.T) {
	httpClient := NewVCRClient(t, "role_assignments")
	cfg := ReadFixture(t, filepath.Join("role_assignments", "main.tf"))
	updated := ReadFixture(t, filepath.Join("role_assignments", "update.tf"))

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
					resource.TestCheckResourceAttr("data.awx_role_definition.inventory", "content_type", "awx.inventory"),
					// AWX returns permissions alphabetically whatever order they
					// go out in, which only a set survives.
					resource.TestCheckResourceAttr("awx_role_definition.inventory_reader", "permissions.#", "2"),
					resource.TestCheckTypeSetElemAttr("awx_role_definition.inventory_reader", "permissions.*", "awx.adhoc_inventory"),
					resource.TestCheckTypeSetElemAttr("awx_role_definition.inventory_reader", "permissions.*", "awx.view_inventory"),
					resource.TestCheckResourceAttrPair(
						"awx_role_user_assignment.reader_inventory", "role_definition",
						"data.awx_role_definition.inventory", "id"),
					resource.TestCheckResourceAttrPair(
						"awx_role_user_assignment.reader_inventory", "user",
						"awx_user.reader", "id"),
					resource.TestCheckResourceAttrPair(
						"awx_role_team_assignment.readers_inventory", "team",
						"awx_team.readers", "id"),
					resource.TestCheckResourceAttrSet("awx_role_user_assignment.reader_inventory", "object_id"),
				),
			},
			{
				Config: providerHeader(t) + updated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"awx_role_user_assignment.reader_inventory", "role_definition",
						"data.awx_role_definition.inventory", "id"),
					resource.TestCheckResourceAttr("awx_role_definition.inventory_reader", "permissions.#", "1"),
					resource.TestCheckTypeSetElemAttr("awx_role_definition.inventory_reader", "permissions.*", "awx.view_inventory"),
				),
			},
			{
				ResourceName:      "awx_role_user_assignment.reader_inventory",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
