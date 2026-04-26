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

// TestIntegration_ObjectRoles consolidates examples/objectroles and
// examples/roles. It exercises:
//
//   - object_roles data sources for organization, team, credential, and
//     the default instance_group
//   - team_associate_role / user_associate_role wired off those data sources
//   - update step drops one team_associate_role to exercise dissociation
//
// The associate_role resources expose `team_id` / `user_id` / `role_id`
// (no synthetic `id`), so the import step targets awx_organization.org
// instead.
func TestIntegration_ObjectRoles(t *testing.T) {
	httpClient := NewVCRClient(t, "object_roles")
	cfg := ReadFixture(t, filepath.Join("object_roles", "main.tf"))
	updated := ReadFixture(t, filepath.Join("object_roles", "update.tf"))

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
					resource.TestCheckResourceAttr("awx_organization.org", "name", "Object Roles"),
					resource.TestCheckResourceAttr("awx_team.team", "name", "Team"),
					resource.TestCheckResourceAttr("awx_user.user", "username", "object_roles_user"),

					resource.TestCheckResourceAttrPair("data.awx_organization_object_roles.org", "id", "awx_organization.org", "id"),
					resource.TestCheckResourceAttrSet("data.awx_organization_object_roles.org", "roles.%"),
					resource.TestCheckResourceAttrSet("data.awx_team_object_roles.team", "roles.%"),
					resource.TestCheckResourceAttrSet("data.awx_credential_object_roles.machine", "roles.%"),
					resource.TestCheckResourceAttrSet("data.awx_instance_group_object_roles.default", "roles.%"),

					resource.TestCheckResourceAttrPair("awx_team_associate_role.team_org_execute", "team_id", "awx_team.team", "id"),
					resource.TestCheckResourceAttrSet("awx_team_associate_role.team_org_execute", "role_id"),
					resource.TestCheckResourceAttrSet("awx_team_associate_role.team_org_inventory_admin", "role_id"),
					resource.TestCheckResourceAttrSet("awx_team_associate_role.team_credential_use", "role_id"),
					resource.TestCheckResourceAttrSet("awx_team_associate_role.team_instance_group_use", "role_id"),
					resource.TestCheckResourceAttrPair("awx_user_associate_role.user_team_admin", "user_id", "awx_user.user", "id"),
					resource.TestCheckResourceAttrSet("awx_user_associate_role.user_team_admin", "role_id"),
				),
			},
			{
				Config: providerHeader(t) + updated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("awx_credential.machine", "description", "Updated"),
					resource.TestCheckResourceAttrPair("awx_team_associate_role.team_org_execute", "team_id", "awx_team.team", "id"),
					resource.TestCheckResourceAttrPair("awx_team_associate_role.team_credential_use", "team_id", "awx_team.team", "id"),
					resource.TestCheckResourceAttrPair("awx_team_associate_role.team_instance_group_use", "team_id", "awx_team.team", "id"),
					resource.TestCheckResourceAttrPair("awx_user_associate_role.user_team_admin", "user_id", "awx_user.user", "id"),
				),
			},
			{
				ResourceName:      "awx_organization.org",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
