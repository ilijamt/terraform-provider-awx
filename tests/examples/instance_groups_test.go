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

// TestIntegration_InstanceGroups exercises examples/instance_groups:
//
//   - Create a container instance group with a JSON-encoded pod_spec_override.
//   - Resolve it through the awx_instance_group data source by name and the
//     awx_instance_group_object_roles data source by id.
//   - Update both name and pod_spec_override (the JSON-string round-trip is
//     the interesting bit — proves the resource doesn't churn on equivalent
//     JSON or drop nested fields).
//   - Import round-trip on the instance group.
func TestIntegration_InstanceGroups(t *testing.T) {
	httpClient := NewVCRClient(t, "instance_groups")
	cfg := ReadFixture(t, filepath.Join("instance_groups", "main.tf"))
	updated := ReadFixture(t, filepath.Join("instance_groups", "update.tf"))

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
					resource.TestCheckResourceAttr("awx_instance_group.ig", "name", "Demo Instance Group"),
					resource.TestCheckResourceAttr("awx_instance_group.ig", "is_container_group", "true"),
					resource.TestCheckResourceAttrSet("awx_instance_group.ig", "id"),
					resource.TestCheckResourceAttrSet("awx_instance_group.ig", "pod_spec_override"),

					resource.TestCheckResourceAttrPair("data.awx_instance_group.ig", "id", "awx_instance_group.ig", "id"),
					resource.TestCheckResourceAttrPair("data.awx_instance_group_object_roles.ig", "id", "awx_instance_group.ig", "id"),
					resource.TestCheckResourceAttrSet("data.awx_instance_group_object_roles.ig", "roles.%"),
				),
			},
			{
				Config: providerHeader(t) + updated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("awx_instance_group.ig", "name", "Demo Instance Group (updated)"),
					resource.TestCheckResourceAttr("awx_instance_group.ig", "is_container_group", "true"),
					resource.TestCheckResourceAttrPair("data.awx_instance_group.ig", "id", "awx_instance_group.ig", "id"),
				),
			},
			{
				ResourceName:      "awx_instance_group.ig",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
