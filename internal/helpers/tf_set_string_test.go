package helpers_test

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAttrValueSetString(t *testing.T) {
	type model struct {
		Value types.String `tfsdk:"value"`
	}

	t.Run("set value correctly", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetString(&state.Value, "test")
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.EqualValues(t, "test", state.Value.ValueString())
	})

	t.Run("obj is nil error", func(t *testing.T) {
		var d, err = helpers.AttrValueSetString(nil, "test")
		require.Error(t, err)
		require.True(t, d.HasError())
	})

	t.Run("value is null", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetString(&state.Value, nil)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.True(t, state.Value.IsNull())
	})

	t.Run("decode the number as string", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetString(&state.Value, json.Number("10"))
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.EqualValues(t, "10", state.Value.ValueString())
	})

	t.Run("value is wrong data type", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetString(&state.Value, 10)
		require.Error(t, err)
		require.True(t, d.HasError())
	})

}
