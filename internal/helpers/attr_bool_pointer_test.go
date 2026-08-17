package helpers_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/require"

	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

func TestAttrBoolPointer(t *testing.T) {
	for _, tc := range []struct {
		name     string
		val      types.Bool
		expected *bool
	}{
		{"null is unset", types.BoolNull(), nil},
		{"unknown is unset", types.BoolUnknown(), nil},
		{"explicit false is kept", types.BoolValue(false), boolPtr(false)},
		{"explicit true is kept", types.BoolValue(true), boolPtr(true)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, helpers.AttrBoolPointer(tc.val))
		})
	}
}

func boolPtr(v bool) *bool { return &v }
