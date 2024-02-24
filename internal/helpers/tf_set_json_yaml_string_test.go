package helpers_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

func TestAttrValueSetJsonYamlString(t *testing.T) {
	type model struct {
		Value types.String `tfsdk:"value"`
	}

	t.Run("obj is nil error", func(t *testing.T) {
		var d, err = helpers.AttrValueSetJsonYamlString(nil, "test", false)
		require.Error(t, err)
		require.True(t, d.HasError())
	})

	t.Run("value is nil", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetJsonYamlString(&state.Value, nil, false)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.True(t, state.Value.IsNull())
	})

	t.Run("value is a json.Number", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetJsonYamlString(&state.Value, json.Number("10"), false)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.EqualValues(t, "10", state.Value.ValueString())
	})

	t.Run("value is map[string]any", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetJsonYamlString(&state.Value, map[string]any{"test": "value"}, false)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.EqualValues(t, `{"test":"value"}`, state.Value.ValueString())
	})

	t.Run("value is []any", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetJsonYamlString(&state.Value, []any{"value"}, false)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.EqualValues(t, `["value"]`, state.Value.ValueString())
	})

	t.Run("value is not a json or yaml", func(t *testing.T) {
		tests := []struct {
			in   any
			trim bool
		}{
			{"\"\"", true},
			{"w", true},
			{"{ what", true},
		}
		for _, test := range tests {
			t.Run(fmt.Sprintf("trim=%t,in=%v", test.trim, test.in), func(t *testing.T) {
				var state model
				var d, err = helpers.AttrValueSetJsonYamlString(&state.Value, test.in, test.trim)
				assert.Error(t, err)
				assert.True(t, d.HasError())
			})
		}

	})

	t.Run("input is yaml should auto convert to json", func(t *testing.T) {
		for _, trim := range []bool{true, false} {
			t.Run(fmt.Sprintf("trim:%t", trim), func(t *testing.T) {
				var state model
				var d, err = helpers.AttrValueSetJsonYamlString(&state.Value, "a: 1\nb: 2\n", trim)
				require.NoError(t, err)
				require.False(t, d.HasError())
				require.False(t, state.Value.IsNull())
				require.EqualValues(t, `{"a":1,"b":2}`, state.Value.ValueString())
			})
		}
	})

	t.Run("input is json should remain as json", func(t *testing.T) {
		for _, trim := range []bool{true, false} {
			t.Run(fmt.Sprintf("trim:%t", trim), func(t *testing.T) {
				var state model
				var d, err = helpers.AttrValueSetJsonYamlString(&state.Value, `{"a":1,"b":2}`, trim)
				require.NoError(t, err)
				require.False(t, d.HasError())
				require.False(t, state.Value.IsNull())
				require.EqualValues(t, `{"a":1,"b":2}`, state.Value.ValueString())
			})
		}
	})
}
