package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/ilijamt/envwrap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProviderConfigureFromEnvironment(t *testing.T) {
	var defaultEnvs = []string{"AWX_HOST", "AWX_USERNAME", "AWX_PASSWORD", "TOWER_HOST", "TOWER_PASSWORD", "TOWER_USERNAME", "TOWER_AUTH_TOKEN", "AWX_AUTH_TOKEN"}
	var tests = []struct {
		in   map[string]string
		null []string
		out  Model
	}{
		{
			in:   map[string]string{"AWX_HOST": "host"},
			out:  Model{Hostname: types.StringValue("host")},
			null: []string{"username", "password", "insecure_skip_verify", "token"},
		},
		{
			in:   map[string]string{"TOWER_HOST": "tower-host", "AWX_HOST": "awx-host"},
			out:  Model{Hostname: types.StringValue("tower-host")},
			null: []string{"username", "password", "insecure_skip_verify", "token"},
		},
		{
			in:   map[string]string{"TOWER_USERNAME": "tower-username"},
			out:  Model{Username: types.StringValue("tower-username")},
			null: []string{"hostname", "password", "insecure_skip_verify", "token"},
		},
		{
			in:   map[string]string{"TOWER_HOST": "tower-host", "AWX_HOST": "awx-host"},
			out:  Model{Hostname: types.StringValue("tower-host")},
			null: []string{"username", "password", "insecure_skip_verify", "token"},
		},
		{
			in:   map[string]string{"AWX_PASSWORD": "password"},
			out:  Model{Password: types.StringValue("password")},
			null: []string{"hostname", "username", "insecure_skip_verify", "token"},
		},
		{
			in:   map[string]string{"TOWER_VERIFY_SSL": "true", "AWX_VERIFY_SSL": "false"},
			out:  Model{VerifySSL: types.BoolValue(true)},
			null: []string{"hostname", "username", "password", "token"},
		},
		{
			in:   map[string]string{"AWX_AUTH_TOKEN": "awx-auth-token"},
			out:  Model{Token: types.StringValue("awx-auth-token")},
			null: []string{"hostname", "username", "password"},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.in), func(t *testing.T) {
			env := envwrap.NewStorage()
			for _, v := range defaultEnvs {
				if _, ok := test.in[v]; !ok {
					_ = env.Store(v, "")
				}
			}
			for k, v := range test.in {
				_ = env.Store(k, v)
			}

			var config = new(Model)
			configureFromEnvironment(context.Background(), config)
			assert.EqualValues(t, test.out, *config)

			for _, n := range test.null {
				switch n {
				case "hostname":
					assert.True(t, config.Hostname.IsNull())
				case "username":
					assert.True(t, config.Username.IsNull())
				case "password":
					assert.True(t, config.Password.IsNull())
				case "token":
					assert.True(t, config.Token.IsNull())
				case "insecure_skip_verify":
					assert.True(t, config.VerifySSL.IsNull())
				}
			}

			_ = env.ReleaseAll()
		})
	}
}

func TestConfigureDefaults(t *testing.T) {
	t.Run("value already set should not override it", func(t *testing.T) {
		var config = &Model{VerifySSL: types.BoolValue(false)}
		require.False(t, config.VerifySSL.IsNull())
		configureDefaults(context.Background(), config)
		require.False(t, config.VerifySSL.ValueBool())
	})

	t.Run("value is set should not override it", func(t *testing.T) {
		var config = &Model{}
		require.True(t, config.VerifySSL.IsNull())
		configureDefaults(context.Background(), config)
		require.True(t, config.VerifySSL.ValueBool())
	})
}
