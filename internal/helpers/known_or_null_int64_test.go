package helpers_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/require"

	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

func TestKnownOrNullInt64(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   types.Int64
		out  types.Int64
	}{
		{"unknown becomes null", types.Int64Unknown(), types.Int64Null()},
		{"null stays null", types.Int64Null(), types.Int64Null()},
		{"value passes through", types.Int64Value(7), types.Int64Value(7)},
		{"zero is a value", types.Int64Value(0), types.Int64Value(0)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.out, helpers.KnownOrNullInt64(tc.in))
		})
	}
}
