package helpers_test

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAttrValueSetBool(t *testing.T) {
	type model struct {
		Value types.Bool `tfsdk:"value"`
	}

	t.Run("set value correctly to false", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetBool(&state.Value, false)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.EqualValues(t, false, state.Value.ValueBool())
	})

	t.Run("set value correctly to true", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetBool(&state.Value, true)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.EqualValues(t, true, state.Value.ValueBool())
	})

	t.Run("obj is nil error", func(t *testing.T) {
		var d, err = helpers.AttrValueSetBool(nil, "test")
		require.Error(t, err)
		require.True(t, d.HasError())
	})

	t.Run("value is null", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetBool(&state.Value, nil)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.True(t, state.Value.IsNull())
	})

	t.Run("value is wrong data type", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetBool(&state.Value, "false")
		require.Error(t, err)
		require.True(t, d.HasError())
	})
}
