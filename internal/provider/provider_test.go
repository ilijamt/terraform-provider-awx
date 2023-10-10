package provider_test

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/ilijamt/terraform-provider-awx/internal/provider"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func TestProvider(t *testing.T) {
	var resources = func() []func() resource.Resource {
		return []func() resource.Resource{}
	}
	var dataSources = func() []func() datasource.DataSource {
		return []func() datasource.DataSource{}
	}
	awxProvider := providerserver.NewProtocol6WithError(provider.New("test", resources(), dataSources())())
	require.NotNil(t, awxProvider)
	frameworkServer, err := awxProvider()
	require.NoError(t, err)
	require.NotNil(t, frameworkServer)
	schema, err := frameworkServer.GetProviderSchema(context.Background(), &tfprotov6.GetProviderSchemaRequest{})
	require.NoError(t, err)
	require.NotNil(t, schema)
}
