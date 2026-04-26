package helpers_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/stretchr/testify/require"
)

func TestListAsStringSlice(t *testing.T) {
	t.Run("null list returns empty slice", func(t *testing.T) {
		out := helpers.ListAsStringSlice(types.ListNull(types.StringType), false)
		require.Empty(t, out)
		require.NotNil(t, out)
	})

	t.Run("empty list returns empty slice", func(t *testing.T) {
		out := helpers.ListAsStringSlice(types.ListValueMust(types.StringType, []attr.Value{}), false)
		require.Empty(t, out)
		require.NotNil(t, out)
	})

	t.Run("list of strings", func(t *testing.T) {
		out := helpers.ListAsStringSlice(types.ListValueMust(types.StringType, []attr.Value{
			types.StringValue("a"),
			types.StringValue("b"),
		}), false)
		require.Equal(t, []string{"a", "b"}, out)
	})

	t.Run("list of strings with trim", func(t *testing.T) {
		out := helpers.ListAsStringSlice(types.ListValueMust(types.StringType, []attr.Value{
			types.StringValue(" a\n"),
			types.StringValue(" b "),
		}), true)
		require.Equal(t, []string{"a", "b"}, out)
	})

	t.Run("non-string element falls back to String()", func(t *testing.T) {
		out := helpers.ListAsStringSlice(types.ListValueMust(types.Int64Type, []attr.Value{
			types.Int64Value(7),
		}), false)
		require.Equal(t, []string{"7"}, out)
	})
}
