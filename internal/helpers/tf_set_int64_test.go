package helpers_test

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAttrValueSetInt64(t *testing.T) {
	type model struct {
		Value types.Int64 `tfsdk:"value"`
	}

	t.Run("set value correctly as int", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetInt64(&state.Value, 10)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.EqualValues(t, 10, state.Value.ValueInt64())
	})

	t.Run("set value correctly as int64", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetInt64(&state.Value, int64(10))
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.EqualValues(t, 10, state.Value.ValueInt64())
	})

	t.Run("set value correctly as json.Number", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetInt64(&state.Value, json.Number("10"))
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.EqualValues(t, 10, state.Value.ValueInt64())
	})

	t.Run("invalid value with json.Number", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetInt64(&state.Value, json.Number("f"))
		require.Error(t, err)
		require.True(t, d.HasError())
	})

	t.Run("null value", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetInt64(&state.Value, nil)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.True(t, state.Value.IsNull())
	})

	t.Run("obj is nil error", func(t *testing.T) {
		var d, err = helpers.AttrValueSetInt64(nil, "test")
		require.Error(t, err)
		require.True(t, d.HasError())
	})

	t.Run("value is wrong data type", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetInt64(&state.Value, "false")
		require.Error(t, err)
		require.True(t, d.HasError())
	})

}
