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

// TestIntegration_Schedule covers issue #173: an unset diff_mode must stay out
// of the create payload, because AWX answers "Field is not allowed on launch"
// for unified job templates that do not prompt for it.
func TestIntegration_Schedule(t *testing.T) {
	httpClient := NewVCRClient(t, "schedule")
	cfg := ReadFixture(t, filepath.Join("schedule", "main.tf"))
	updated := ReadFixture(t, filepath.Join("schedule", "update.tf"))

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
					resource.TestCheckResourceAttr("awx_schedule.inventory_refresh", "name", "Refresh example inventory every 6 hours"),
					resource.TestCheckResourceAttr("awx_schedule.inventory_refresh", "enabled", "true"),
					resource.TestCheckResourceAttrPair("awx_schedule.inventory_refresh", "unified_job_template", "awx_inventory_source.schedule", "id"),
					resource.TestCheckNoResourceAttr("awx_schedule.inventory_refresh", "diff_mode"),
					resource.TestCheckResourceAttrSet("awx_schedule.inventory_refresh", "next_run"),
				),
			},
			{
				Config: providerHeader(t) + updated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("awx_schedule.inventory_refresh", "name", "Refresh example inventory every 12 hours"),
					resource.TestCheckResourceAttr("awx_schedule.inventory_refresh", "enabled", "false"),
					resource.TestCheckNoResourceAttr("awx_schedule.inventory_refresh", "next_run"),
				),
			},
			{
				ResourceName:      "awx_schedule.inventory_refresh",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
