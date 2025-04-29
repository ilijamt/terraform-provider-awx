package provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/ilijamt/terraform-provider-awx/internal/provider"
	"github.com/ilijamt/terraform-provider-awx/version"
)

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
	factories := provider.TestFactories(t, "awx", nil, version.Version, provider.TestEmptyResources(t), provider.TestEmptyDataSources(t))
	resource.UnitTest(t, resource.TestCase{
		PreCheck:                 func() { provider.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: factories,
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
