package provider

import (
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"

	"github.com/ilijamt/terraform-provider-awx/internal/client"
)

func TestAccPreCheck(t *testing.T) {}

func TestFactories(t *testing.T, name string, httpClient *http.Client, awxClient client.Client, version string, fnResources []func() resource.Resource, fnDataSources []func() datasource.DataSource) (factory map[string]func() (tfprotov6.ProviderServer, error)) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		name: providerserver.NewProtocol6WithError(
			NewFuncProvider(version, httpClient, awxClient, fnResources, fnDataSources)(),
		),
	}
}

func TestEmptyResources(t *testing.T) []func() resource.Resource {
	return []func() resource.Resource{}
}

func TestEmptyDataSources(t *testing.T) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}
