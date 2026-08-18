package helpers_test

import (
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/stretchr/testify/require"
)

func TestAttrValueSetListInt64(t *testing.T) {
	type model struct {
		Value types.List `tfsdk:"value"`
	}

	t.Run("obj is nil error", func(t *testing.T) {
		var d, err = helpers.AttrValueSetListInt64(nil, []any{})
		require.Error(t, err)
		require.True(t, d.HasError())
	})

	t.Run("value is null should return empty list", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetListInt64(&state.Value, nil)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.Empty(t, state.Value.Elements())
	})

	t.Run("value is a types.List", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetListInt64(&state.Value,
			types.ListValueMust(types.Int64Type, []attr.Value{types.Int64Value(7)}))
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.Len(t, state.Value.Elements(), 1)
	})

	// The client decodes with UseNumber, so ids arrive as json.Number.
	t.Run("value is a []any of json.Number", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetListInt64(&state.Value,
			[]any{json.Number("3"), json.Number("9")})
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.Equal(t, []attr.Value{types.Int64Value(3), types.Int64Value(9)}, state.Value.Elements())
	})

	t.Run("value is a []int64", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetListInt64(&state.Value, []int64{1, 2})
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.Len(t, state.Value.Elements(), 2)
	})

	t.Run("non-numeric element errors", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetListInt64(&state.Value, []any{"nope"})
		require.Error(t, err)
		require.True(t, d.HasError())
	})

	t.Run("wrong container type errors", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetListInt64(&state.Value, 42)
		require.Error(t, err)
		require.True(t, d.HasError())
	})
}
