package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func boolPtr(v bool) *bool { return &v }

func updateProperty(t *testing.T, name, awxType string, override PropertyOverride) Property {
	t.Helper()
	var p Property
	p.Name = name
	values := map[string]any{"type": awxType, "label": name}
	require.NoError(t, p.Update(TypeWrite, override, values, Item{Name: "Schedule", TypeName: "schedule"}))
	return p
}

func TestPropertyNullable(t *testing.T) {
	for _, tc := range []struct {
		name          string
		awxType       string
		override      PropertyOverride
		expectedType  string
		expectedValue string
	}{
		{
			name:          "bool defaults to a plain value so an explicit false is still sent",
			awxType:       "boolean",
			expectedType:  "bool",
			expectedValue: "o.DiffMode.ValueBool()",
		},
		{
			name:          "nullable bool becomes a pointer",
			awxType:       "boolean",
			override:      PropertyOverride{Nullable: boolPtr(true)},
			expectedType:  "*bool",
			expectedValue: "helpers.AttrBoolPointer(o.DiffMode)",
		},
		{
			// Only bool is forced to always-send, so nothing else needs it.
			name:          "nullable is ignored for non-bool types",
			awxType:       "string",
			override:      PropertyOverride{Nullable: boolPtr(true)},
			expectedType:  "string",
			expectedValue: "o.DiffMode.ValueString()",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			p := updateProperty(t, "diff_mode", tc.awxType, tc.override)
			require.Equal(t, tc.expectedType, p.Generated.BodyRequestModelType)
			require.Equal(t, tc.expectedValue, p.Generated.ModelBodyRequestValue)
		})
	}
}

func TestPropertyUseStateForUnknown(t *testing.T) {
	require.True(t, updateProperty(t, "next_run", "datetime", PropertyOverride{}).UseStateForUnknown)
	require.False(t, updateProperty(t, "next_run", "datetime", PropertyOverride{
		UseStateForUnknown: boolPtr(false),
	}).UseStateForUnknown)
}
