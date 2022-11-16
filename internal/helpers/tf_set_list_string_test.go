package helpers_test

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAttrValueSetListString(t *testing.T) {
	type model struct {
		Value types.List `tfsdk:"value"`
	}

	t.Run("obj is nil error", func(t *testing.T) {
		var d, err = helpers.AttrValueSetListString(nil, "test", false)
		require.Error(t, err)
		require.True(t, d.HasError())
	})

	t.Run("value is null should return empty list", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetListString(&state.Value, nil, false)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.Empty(t, state.Value.Elements())
	})

	t.Run("value is a types.List", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetListString(&state.Value,
			types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("test"),
			}), false)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.NotEmpty(t, state.Value.Elements())
		require.Len(t, state.Value.Elements(), 1)
	})

	t.Run("value is a []any", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetListString(&state.Value,
			[]any{"test"}, false)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.NotEmpty(t, state.Value.Elements())
		require.Len(t, state.Value.Elements(), 1)
	})

	t.Run("value is a []string", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetListString(&state.Value,
			[]string{"test"}, false)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.NotEmpty(t, state.Value.Elements())
		require.Len(t, state.Value.Elements(), 1)
	})

	t.Run("value is a []any with trim", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetListString(&state.Value,
			[]any{" test "}, true)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.NotEmpty(t, state.Value.Elements())
		require.Len(t, state.Value.Elements(), 1)
		require.EqualValues(t, "test", state.Value.Elements()[0].(types.String).ValueString())
	})

	t.Run("value is a []string with trim", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetListString(&state.Value,
			[]string{" test "}, true)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.NotEmpty(t, state.Value.Elements())
		require.Len(t, state.Value.Elements(), 1)
		require.EqualValues(t, "test", state.Value.Elements()[0].(types.String).ValueString())
	})

	t.Run("value is wrong data type", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetListString(&state.Value, 10, false)
		require.Error(t, err)
		require.True(t, d.HasError())
	})

}
