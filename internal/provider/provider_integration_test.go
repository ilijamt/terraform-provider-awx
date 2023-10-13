package provider_test

import (
	d "github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	r "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/ilijamt/terraform-provider-awx/internal/provider"
	"github.com/ilijamt/terraform-provider-awx/version"
	"testing"
)

// testAccProtoV6ProviderFactoriesUnique is used to ensure that the provider instance used for
// each acceptance test is unique.
// This is necessary because this provider make use of state stored in the provider instance.
func testAccProtoV6ProviderFactoriesUnique() map[string]func() (tfprotov6.ProviderServer, error) {
	var resources = func() []func() r.Resource {
		return []func() r.Resource{}
	}
	var dataSources = func() []func() d.DataSource {
		return []func() d.DataSource{}
	}

	return map[string]func() (tfprotov6.ProviderServer, error){
		"awx": providerserver.NewProtocol6WithError(provider.NewFuncProvider(version.Version, resources(), dataSources())()),
	}
}

// testAccPreCheck ensures that the environment is properly configured for acceptance testing.
func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
}

const (
	providerConfig = `
provider "awx" {
  username             = "admin"
  password             = "admin"
  host                 = "http://localhost"
  insecure_skip_verify = "true"
}
`
)

func TestProviderIntegration(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactoriesUnique(),
		ExternalProviders:        map[string]resource.ExternalProvider{},
		Steps: []resource.TestStep{
			{
				Config: providerConfig,
				Check: func(state *terraform.State) error {
					return nil
				},
			},
		},
	})
}
