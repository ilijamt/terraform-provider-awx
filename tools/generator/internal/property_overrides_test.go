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

func TestPropertyNoDefault(t *testing.T) {
	const awxSample = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"

	t.Run("a reported default is pinned by default", func(t *testing.T) {
		var p Property
		p.Name = "identifier"
		values := map[string]any{"type": "string", "label": "Identifier", "default": awxSample}
		require.NoError(t, p.Update(TypeWrite, PropertyOverride{}, values, Item{Name: "WorkflowJobTemplateNode"}))
		require.True(t, p.HasDefaultValue)
		require.Contains(t, p.DefaultValue, awxSample)
	})

	t.Run("no_default drops it and leaves the attribute computed", func(t *testing.T) {
		var p Property
		p.Name = "identifier"
		values := map[string]any{"type": "string", "label": "Identifier", "default": awxSample}
		require.NoError(t, p.Update(TypeWrite, PropertyOverride{NoDefault: true}, values, Item{Name: "WorkflowJobTemplateNode"}))
		require.False(t, p.HasDefaultValue)
		require.Empty(t, p.DefaultValue)
		require.False(t, p.IsRequired)
		require.True(t, p.IsComputed)
	})
}

func TestPropertyRequiresReplace(t *testing.T) {
	require.False(t, updateProperty(t, "workflow_job_template", "id", PropertyOverride{}).RequiresReplace)
	require.True(t, updateProperty(t, "workflow_job_template", "id", PropertyOverride{
		RequiresReplace: true,
	}).RequiresReplace)
}

// Defaults used to be dropped for every type but string and integer, so an unset
// attribute went out as its Go zero and overwrote the server's own value.
func TestPropertyScalarDefault(t *testing.T) {
	for _, tc := range []struct {
		name     string
		awxType  string
		awxValue any
		expected string
	}{
		{name: "bool true", awxType: "boolean", awxValue: true, expected: "booldefault.StaticBool(true)"},
		{name: "bool false", awxType: "boolean", awxValue: false, expected: "booldefault.StaticBool(false)"},
		{name: "decimal", awxType: "decimal", awxValue: 1.0, expected: "float64default.StaticFloat64(1)"},
		{name: "float", awxType: "float", awxValue: 0.5, expected: "float64default.StaticFloat64(0.5)"},
		{name: "integer", awxType: "integer", awxValue: 7, expected: "int64default.StaticInt64(7)"},
		{name: "string", awxType: "string", awxValue: "x", expected: "stringdefault.StaticString(`x`)"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var p Property
			p.Name = "field"
			values := map[string]any{"type": tc.awxType, "label": "Field", "default": tc.awxValue}
			require.NoError(t, p.Update(TypeWrite, PropertyOverride{}, values, Item{Name: "Instance"}))
			require.Equal(t, tc.expected, p.DefaultValue)
		})
	}
}

func TestPropertyUseStateForUnknown(t *testing.T) {
	off := boolPtr(false)
	for _, tc := range []struct {
		name     string
		field    string
		item     Item
		override PropertyOverride
		expected bool
	}{
		{name: "on by default", field: "capacity", item: Item{IdKey: "id"}, expected: true},
		{name: "item switches it off", field: "capacity", item: Item{IdKey: "id", UseStateForUnknown: off}},
		{
			name:     "property override wins over the item",
			field:    "capacity",
			item:     Item{IdKey: "id", UseStateForUnknown: off},
			override: PropertyOverride{UseStateForUnknown: boolPtr(true)},
			expected: true,
		},
		{
			// Update addresses the resource by this value, so it cannot go unknown.
			name:     "the id keeps it whatever the item says",
			field:    "id",
			item:     Item{IdKey: "id", UseStateForUnknown: off},
			expected: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var p Property
			p.Name = tc.field
			values := map[string]any{"type": "integer", "label": tc.field}
			require.NoError(t, p.Update(TypeRead, tc.override, values, tc.item))
			require.Equal(t, tc.expected, p.UseStateForUnknown)
		})
	}
}
