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

// TestIntegration_Application exercises examples/application:
//
//   - Create an organization-scoped OAuth2 awx_application (confidential,
//     authorization-code).
//   - Resolve it through both data source search groups: by_id and
//     by_name_organization. The lookups are mutually exclusive in schema
//     (id conflicts with name+organization), so both code paths matter.
//   - Update name, description, and redirect_uris in place.
//   - Import round-trip — `client_secret` is excluded because AWX returns
//     `$encrypted$` for it after creation.
func TestIntegration_Application(t *testing.T) {
	httpClient := NewVCRClient(t, "application")
	cfg := ReadFixture(t, filepath.Join("application", "main.tf"))
	updated := ReadFixture(t, filepath.Join("application", "update.tf"))

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
					resource.TestCheckResourceAttr("awx_organization.application", "name", "Application"),

					resource.TestCheckResourceAttr("awx_application.application", "name", "Application"),
					resource.TestCheckResourceAttr("awx_application.application", "client_type", "confidential"),
					resource.TestCheckResourceAttr("awx_application.application", "authorization_grant_type", "authorization-code"),
					resource.TestCheckResourceAttr("awx_application.application", "redirect_uris", "https://localhost"),
					resource.TestCheckResourceAttrPair("awx_application.application", "organization", "awx_organization.application", "id"),
					resource.TestCheckResourceAttrSet("awx_application.application", "client_id"),

					resource.TestCheckResourceAttrPair("data.awx_application.application_by_id", "id", "awx_application.application", "id"),
					resource.TestCheckResourceAttrPair("data.awx_application.application_by_id", "name", "awx_application.application", "name"),

					resource.TestCheckResourceAttrPair("data.awx_application.application_by_name_org", "id", "awx_application.application", "id"),
					resource.TestCheckResourceAttrPair("data.awx_application.application_by_name_org", "organization", "awx_organization.application", "id"),
				),
			},
			{
				Config: providerHeader(t) + updated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("awx_application.application", "name", "Application (updated)"),
					resource.TestCheckResourceAttr("awx_application.application", "description", "OAuth2 application created by TestIntegration_Application (updated)"),
					resource.TestCheckResourceAttr("awx_application.application", "redirect_uris", "https://localhost https://localhost/callback"),
					resource.TestCheckResourceAttrPair("data.awx_application.application_by_id", "name", "awx_application.application", "name"),
				),
			},
			{
				ResourceName:            "awx_application.application",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"client_secret"},
			},
		},
	})
}
