//go:build integration

package examples

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/ilijamt/terraform-provider-awx/internal/awx"
	"github.com/ilijamt/terraform-provider-awx/internal/provider"
	"github.com/ilijamt/terraform-provider-awx/version"
)

// TestIntegration_Inventory exercises examples/inventory against a recorded
// cassette. Run with AWX_VCR_RECORD=1 against a freshly bootstrapped local
// AWX to refresh the cassette; otherwise it replays without touching the
// network.
func TestIntegration_Inventory(t *testing.T) {
	httpClient := NewVCRClient(t, "inventory")
	cfg := ReadFixture(t, filepath.Join("inventory", "main.tf"))
	// Updated config lives in a sibling .tf so you can diff main.tf vs
	// update.tf to see what the update step is exercising. Re-record the
	// cassette (`make test-integration-record`) after changing update.tf.
	updated := ReadFixture(t, filepath.Join("inventory", "update.tf"))

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
					resource.TestCheckResourceAttr("awx_organization.inventory", "name", "Inventory"),
					resource.TestCheckResourceAttr("awx_inventory.inventory", "name", "Example inventory"),
					resource.TestCheckResourceAttrSet("awx_inventory.inventory", "id"),
					resource.TestCheckResourceAttr("awx_host.example_host_vars", "name", "vars.example.com"),
					resource.TestCheckResourceAttrPair("data.awx_inventory_object_roles.inventory", "id", "awx_inventory.inventory", "id"),
					resource.TestCheckResourceAttrSet("data.awx_inventory_object_roles.inventory", "roles.%"),
				),
			},
			{
				Config: providerHeader(t) + updated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("awx_inventory.inventory", "name", "Example inventory (updated)"),
					resource.TestCheckResourceAttr("awx_host.example_host_vars", "name", "vars-updated.example.com"),
					resource.TestCheckResourceAttrPair("data.awx_inventory_object_roles.inventory", "id", "awx_inventory.inventory", "id"),
					resource.TestCheckResourceAttrSet("data.awx_inventory_object_roles.inventory", "roles.%"),
				),
			},
			{
				// Import the inventory by its numeric ID and verify the
				// imported state round-trips against the prior step's state.
				ResourceName:      "awx_inventory.inventory",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// providerHeader emits a provider block pointing at the stable test host.
// In replay mode the token is a placeholder. In record mode the
// bootstrap-generated token is loaded from disk.
func providerHeader(t *testing.T) string {
	token := FakeToken
	if IsRecording() {
		token = LoadBootstrapToken(t)
	}
	return fmt.Sprintf(`
provider "awx" {
  hostname   = %q
  token      = %q
  verify_ssl = false
}
`, StableHost, token)
}
