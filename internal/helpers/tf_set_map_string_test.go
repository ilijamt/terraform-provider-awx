package helpers_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/stretchr/testify/require"
)

func TestAttrValueSetMapString(t *testing.T) {
	type model struct {
		Value types.Map `tfsdk:"value"`
	}

	t.Run("obj is nil error", func(t *testing.T) {
		var d, err = helpers.AttrValueSetMapString(nil, "test", false)
		require.Error(t, err)
		require.True(t, d.HasError())
	})

	t.Run("value is null should return empty map", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetMapString(&state.Value, nil, false)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.Empty(t, state.Value.Elements())
	})

	t.Run("value is a types.Map", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetMapString(&state.Value,
			types.MapValueMust(types.StringType, map[string]attr.Value{
				"X-Token": types.StringValue("abc"),
			}), false)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.Len(t, state.Value.Elements(), 1)
	})

	t.Run("value is a map[string]any", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetMapString(&state.Value,
			map[string]any{"X-Token": "abc", "X-Other": "def"}, false)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.Len(t, state.Value.Elements(), 2)
		require.Equal(t, types.StringValue("abc"), state.Value.Elements()["X-Token"])
	})

	t.Run("value trims when asked", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetMapString(&state.Value,
			map[string]any{"X-Token": " abc "}, true)
		require.NoError(t, err)
		require.False(t, d.HasError())
		require.Equal(t, types.StringValue("abc"), state.Value.Elements()["X-Token"])
	})

	t.Run("non string element is an error", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetMapString(&state.Value,
			map[string]any{"X-Token": 1}, false)
		require.Error(t, err)
		require.True(t, d.HasError())
	})

	t.Run("unsupported type is an error", func(t *testing.T) {
		var state model
		var d, err = helpers.AttrValueSetMapString(&state.Value, "test", false)
		require.Error(t, err)
		require.True(t, d.HasError())
	})
}

func TestMapAsStringMap(t *testing.T) {
	t.Run("string elements are returned", func(t *testing.T) {
		var m = types.MapValueMust(types.StringType, map[string]attr.Value{
			"a": types.StringValue("1"),
			"b": types.StringValue("2"),
		})
		require.Equal(t, map[string]string{"a": "1", "b": "2"}, helpers.MapAsStringMap(m, false))
	})

	t.Run("empty map yields empty result", func(t *testing.T) {
		var m = types.MapValueMust(types.StringType, map[string]attr.Value{})
		require.Equal(t, map[string]string{}, helpers.MapAsStringMap(m, false))
	})

	t.Run("trims when asked", func(t *testing.T) {
		var m = types.MapValueMust(types.StringType, map[string]attr.Value{
			"a": types.StringValue(" 1 "),
		})
		require.Equal(t, map[string]string{"a": "1"}, helpers.MapAsStringMap(m, true))
	})

	t.Run("non string elements are skipped", func(t *testing.T) {
		var m = types.MapValueMust(types.Int64Type, map[string]attr.Value{
			"a": types.Int64Value(1),
		})
		require.Equal(t, map[string]string{}, helpers.MapAsStringMap(m, false))
	})
}
