package provider_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/require"

	"github.com/ilijamt/terraform-provider-awx/internal/provider"

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
	awxProvider := providerserver.NewProtocol6WithError(provider.NewFuncProvider("test", nil, resources(), dataSources())())
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

	awxProvider := providerserver.NewProtocol6WithError(provider.NewFuncProvider("test", nil, resources(), dataSources())())
	require.NotNil(t, awxProvider)
	frameworkServer, err := awxProvider()
	require.NoError(t, err)
	require.NotNil(t, frameworkServer)

	var ConfigDataType = tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"hostname":   tftypes.String,
			"username":   tftypes.String,
			"password":   tftypes.String,
			"verify_ssl": tftypes.Bool,
			"token":      tftypes.String,
		},
	}

	t.Run("valid configuration", func(t *testing.T) {
		config, err := tfprotov6.NewDynamicValue(ConfigDataType, tftypes.NewValue(ConfigDataType, map[string]tftypes.Value{
			"hostname":   tftypes.NewValue(tftypes.String, "host"),
			"username":   tftypes.NewValue(tftypes.String, "username"),
			"password":   tftypes.NewValue(tftypes.String, "password"),
			"verify_ssl": tftypes.NewValue(tftypes.Bool, true),
			"token":      tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
		}))
		require.NoError(t, err)
		response, err := frameworkServer.ConfigureProvider(context.Background(), &tfprotov6.ConfigureProviderRequest{
			Config: &config,
		})
		require.NoError(t, err)
		require.NotNil(t, response)
		require.Empty(t, response.Diagnostics)
	})

	t.Run("unknown values for configuration", func(t *testing.T) {
		config, err := tfprotov6.NewDynamicValue(ConfigDataType, tftypes.NewValue(ConfigDataType, map[string]tftypes.Value{
			"hostname":   tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			"username":   tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			"password":   tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			"verify_ssl": tftypes.NewValue(tftypes.Bool, true),
			"token":      tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
		}))
		require.NoError(t, err)
		response, err := frameworkServer.ConfigureProvider(context.Background(), &tfprotov6.ConfigureProviderRequest{
			Config: &config,
		})
		require.NoError(t, err)
		require.NotNil(t, response)
		require.Len(t, response.Diagnostics, 2)
	})

	t.Run("empty values for configuration", func(t *testing.T) {
		config, err := tfprotov6.NewDynamicValue(ConfigDataType, tftypes.NewValue(ConfigDataType, map[string]tftypes.Value{
			"hostname":   tftypes.NewValue(tftypes.String, ""),
			"username":   tftypes.NewValue(tftypes.String, ""),
			"password":   tftypes.NewValue(tftypes.String, ""),
			"verify_ssl": tftypes.NewValue(tftypes.Bool, true),
			"token":      tftypes.NewValue(tftypes.String, ""),
		}))
		require.NoError(t, err)
		response, err := frameworkServer.ConfigureProvider(context.Background(), &tfprotov6.ConfigureProviderRequest{
			Config: &config,
		})
		require.NoError(t, err)
		require.NotNil(t, response)
		require.Len(t, response.Diagnostics, 2)
	})

	t.Run("configuration", func(t *testing.T) {
		var tests = []struct {
			in         map[string]tftypes.Value
			errLen     int
			errSummary []string
		}{
			{
				in: map[string]tftypes.Value{
					"hostname":   tftypes.NewValue(tftypes.String, ""),
					"username":   tftypes.NewValue(tftypes.String, "username"),
					"password":   tftypes.NewValue(tftypes.String, "password"),
					"verify_ssl": tftypes.NewValue(tftypes.Bool, true),
					"token":      tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
				},
				errLen:     1,
				errSummary: []string{"Unknown AWX API Host"},
			},
			{
				in: map[string]tftypes.Value{
					"hostname":   tftypes.NewValue(tftypes.String, "hostname"),
					"username":   tftypes.NewValue(tftypes.String, ""),
					"password":   tftypes.NewValue(tftypes.String, "password"),
					"verify_ssl": tftypes.NewValue(tftypes.Bool, true),
					"token":      tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
				},
				errLen:     1,
				errSummary: []string{"Unknown AWX API Username"},
			},
			{
				in: map[string]tftypes.Value{
					"hostname":   tftypes.NewValue(tftypes.String, "hostname"),
					"username":   tftypes.NewValue(tftypes.String, "username"),
					"password":   tftypes.NewValue(tftypes.String, ""),
					"verify_ssl": tftypes.NewValue(tftypes.Bool, true),
					"token":      tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
				},
				errLen:     1,
				errSummary: []string{"Unknown AWX API Password"},
			},
			{
				in: map[string]tftypes.Value{
					"hostname":   tftypes.NewValue(tftypes.String, "hostname"),
					"username":   tftypes.NewValue(tftypes.String, ""),
					"password":   tftypes.NewValue(tftypes.String, ""),
					"verify_ssl": tftypes.NewValue(tftypes.Bool, true),
					"token":      tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
				},
				errLen:     1,
				errSummary: []string{`must provide one of ["username", "password"] or "token".`},
			},
			{
				in: map[string]tftypes.Value{
					"hostname":   tftypes.NewValue(tftypes.String, ""),
					"username":   tftypes.NewValue(tftypes.String, "username"),
					"password":   tftypes.NewValue(tftypes.String, ""),
					"verify_ssl": tftypes.NewValue(tftypes.Bool, true),
					"token":      tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
				},
				errLen:     2,
				errSummary: []string{"Unknown AWX API Host", "Unknown AWX API Password"},
			},
			{
				in: map[string]tftypes.Value{
					"hostname":   tftypes.NewValue(tftypes.String, ""),
					"username":   tftypes.NewValue(tftypes.String, ""),
					"password":   tftypes.NewValue(tftypes.String, ""),
					"verify_ssl": tftypes.NewValue(tftypes.Bool, true),
					"token":      tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
				},
				errLen:     2,
				errSummary: []string{"Unknown AWX API Host", `must provide one of ["username", "password"] or "token".`},
			},
			{
				in: map[string]tftypes.Value{
					"hostname":   tftypes.NewValue(tftypes.String, "hostname"),
					"username":   tftypes.NewValue(tftypes.String, ""),
					"password":   tftypes.NewValue(tftypes.String, ""),
					"verify_ssl": tftypes.NewValue(tftypes.Bool, true),
					"token":      tftypes.NewValue(tftypes.String, "token"),
				},
				errLen: 0,
			},
			{
				in: map[string]tftypes.Value{
					"hostname":   tftypes.NewValue(tftypes.String, "hostname"),
					"username":   tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
					"password":   tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
					"verify_ssl": tftypes.NewValue(tftypes.Bool, true),
					"token":      tftypes.NewValue(tftypes.String, "token"),
				},
				errLen: 0,
			},
			{
				in: map[string]tftypes.Value{
					"hostname":   tftypes.NewValue(tftypes.String, "hostname"),
					"username":   tftypes.NewValue(tftypes.String, "username"),
					"password":   tftypes.NewValue(tftypes.String, "password"),
					"verify_ssl": tftypes.NewValue(tftypes.Bool, true),
					"token":      tftypes.NewValue(tftypes.String, ""),
				},
				errLen: 0,
			},
			{
				in: map[string]tftypes.Value{
					"hostname":   tftypes.NewValue(tftypes.String, "hostname"),
					"username":   tftypes.NewValue(tftypes.String, "username"),
					"password":   tftypes.NewValue(tftypes.String, "password"),
					"verify_ssl": tftypes.NewValue(tftypes.Bool, true),
					"token":      tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
				},
				errLen: 0,
			},
			{
				in: map[string]tftypes.Value{
					"hostname":   tftypes.NewValue(tftypes.String, "hostname"),
					"username":   tftypes.NewValue(tftypes.String, "username"),
					"password":   tftypes.NewValue(tftypes.String, "password"),
					"verify_ssl": tftypes.NewValue(tftypes.Bool, true),
					"token":      tftypes.NewValue(tftypes.String, "token"),
				},
				errLen:     1,
				errSummary: []string{`must provide one of ["username", "password"] or "token".`},
			},
		}

		for _, test := range tests {
			t.Run(fmt.Sprintf("%s", test.in), func(t *testing.T) {
				config, err := tfprotov6.NewDynamicValue(ConfigDataType,
					tftypes.NewValue(ConfigDataType, test.in))
				require.NoError(t, err)
				response, err := frameworkServer.ConfigureProvider(context.Background(), &tfprotov6.ConfigureProviderRequest{
					Config: &config,
				})
				require.NoError(t, err)
				require.NotNil(t, response)
				require.Len(t, response.Diagnostics, test.errLen)
				var summary []string
				for _, d := range response.Diagnostics {
					summary = append(summary, d.Summary)
				}
				require.EqualValues(t, summary, test.errSummary)
			})
		}
	})

}
