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

// TestIntegration_Instance covers mesh node registration. AWX refuses DELETE on
// an instance, so destroy goes out as node_state "deprovisioning" and the
// cassette is the proof that it removes the record.
func TestIntegration_Instance(t *testing.T) {
	httpClient := NewVCRClient(t, "instance")
	cfg := ReadFixture(t, filepath.Join("instance", "main.tf"))
	updated := ReadFixture(t, filepath.Join("instance", "update.tf"))

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
					resource.TestCheckResourceAttr("awx_instance.execution", "node_type", "execution"),
					resource.TestCheckResourceAttr("awx_instance.execution", "listener_port", "27199"),
					resource.TestCheckResourceAttr("awx_instance.execution", "enabled", "true"),
					// AWX serves capacity_adjustment as the string "1.00".
					resource.TestCheckResourceAttr("awx_instance.execution", "capacity_adjustment", "1"),
					resource.TestCheckResourceAttr("awx_instance.hop", "node_type", "hop"),
					resource.TestCheckResourceAttrPair(
						"data.awx_instance.execution", "id",
						"awx_instance.execution", "id"),
				),
			},
			{
				Config: providerHeader(t) + updated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("awx_instance.execution", "enabled", "false"),
					resource.TestCheckResourceAttr("awx_instance.execution", "capacity_adjustment", "0.5"),
				),
			},
			{
				ResourceName:      "awx_instance.execution",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
