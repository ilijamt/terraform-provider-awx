package helpers_test

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAttrValueSetJsonString(t *testing.T) {
	type model struct {
		Value types.String `tfsdk:"value"`
	}

	t.Run("obj is nil error", func(t *testing.T) {
		var d, err = helpers.AttrValueSetJsonString(nil, "test", false)
		require.Error(t, err)
		require.True(t, d.HasError())
	})

	t.Run("value is null", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetJsonString(&state.Value, nil, false)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.True(t, state.Value.IsNull())
	})

	t.Run("value is a string", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetJsonString(&state.Value, "any", false)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.EqualValues(t, `any`, state.Value.ValueString(), false)
	})

	t.Run("value is a number", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetJsonString(&state.Value, 10, false)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.EqualValues(t, "10", state.Value.ValueString(), false)
	})

	t.Run("value is a json.Number", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetJsonString(&state.Value, json.Number("10"), false)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.EqualValues(t, "10", state.Value.ValueString())
	})

	t.Run("value is map[string]any", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetJsonString(&state.Value, map[string]any{"test": "value"}, false)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.EqualValues(t, `{"test":"value"}`, state.Value.ValueString())
	})

	t.Run("value is []any", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetJsonString(&state.Value, []any{"value"}, false)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.EqualValues(t, `["value"]`, state.Value.ValueString())
	})

	t.Run("value is \"\"", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetJsonString(&state.Value, "\"\"", true)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.EqualValues(t, `""`, state.Value.ValueString())
	})

}
