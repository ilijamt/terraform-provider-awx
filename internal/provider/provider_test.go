package provider_test

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
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
	awxProvider := providerserver.NewProtocol6WithError(provider.NewFuncProvider("test", resources(), dataSources())())
	require.NotNil(t, awxProvider)
	frameworkServer, err := awxProvider()
	require.NoError(t, err)
	require.NotNil(t, frameworkServer)
	schema, err := frameworkServer.GetProviderSchema(context.Background(), &tfprotov6.GetProviderSchemaRequest{})
	require.NoError(t, err)
	require.NotNil(t, schema)
	_, err = frameworkServer.ValidateProviderConfig(context.Background(), &tfprotov6.ValidateProviderConfigRequest{})
	require.NoError(t, err)
}

func TestProviderConfiguration(t *testing.T) {
	var resources = func() []func() resource.Resource {
		return []func() resource.Resource{}
	}
	var dataSources = func() []func() datasource.DataSource {
		return []func() datasource.DataSource{}
	}

	awxProvider := providerserver.NewProtocol6WithError(provider.NewFuncProvider("test", resources(), dataSources())())
	require.NotNil(t, awxProvider)
	frameworkServer, err := awxProvider()
	require.NoError(t, err)
	require.NotNil(t, frameworkServer)

	t.Run("valid configuration", func(t *testing.T) {
		var ConfigDataType = tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"hostname":             tftypes.String,
				"username":             tftypes.String,
				"password":             tftypes.String,
				"insecure_skip_verify": tftypes.Bool,
			},
		}
		config, err := tfprotov6.NewDynamicValue(ConfigDataType, tftypes.NewValue(ConfigDataType, map[string]tftypes.Value{
			"hostname":             tftypes.NewValue(tftypes.String, "host"),
			"username":             tftypes.NewValue(tftypes.String, "username"),
			"password":             tftypes.NewValue(tftypes.String, "password"),
			"insecure_skip_verify": tftypes.NewValue(tftypes.Bool, true),
		}))
		response, err := frameworkServer.ConfigureProvider(context.Background(), &tfprotov6.ConfigureProviderRequest{
			Config: &config,
		})
		require.NoError(t, err)
		require.NotNil(t, response)
		require.Empty(t, response.Diagnostics)
	})
}
