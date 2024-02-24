package helpers_test

import (
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/stretchr/testify/require"
)

func TestAttrValueSetFloat64(t *testing.T) {
	type model struct {
		Value types.Float64 `tfsdk:"value"`
	}

	t.Run("set value correctly as float64", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetFloat64(&state.Value, float64(10))
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.EqualValues(t, 10, state.Value.ValueFloat64())
	})

	t.Run("set value correctly as float32", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetFloat64(&state.Value, float32(10))
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.EqualValues(t, 10, state.Value.ValueFloat64())
	})

	t.Run("set value correctly as json.Number", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetFloat64(&state.Value, json.Number("10"))
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.EqualValues(t, 10, state.Value.ValueFloat64())
	})

	t.Run("invalid value with json.Number", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetFloat64(&state.Value, json.Number("f"))
		require.Error(t, err)
		require.True(t, d.HasError())
	})

	t.Run("null value", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetFloat64(&state.Value, nil)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.True(t, state.Value.IsNull())
	})

	t.Run("obj is nil error", func(t *testing.T) {
		var d, err = helpers.AttrValueSetFloat64(nil, "test")
		require.Error(t, err)
		require.True(t, d.HasError())
	})

	t.Run("value is wrong data type", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetFloat64(&state.Value, "false")
		require.Error(t, err)
		require.True(t, d.HasError())
	})

}
