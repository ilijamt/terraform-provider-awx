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

func TestPropertyUseStateForUnknown(t *testing.T) {
	require.True(t, updateProperty(t, "next_run", "datetime", PropertyOverride{}).UseStateForUnknown)
	require.False(t, updateProperty(t, "next_run", "datetime", PropertyOverride{
		UseStateForUnknown: boolPtr(false),
	}).UseStateForUnknown)
}

// AWX reports bool defaults (host enabled, schedule enabled) and they used to be
// dropped, so an unset attribute went out as an explicit false and flipped the
// server's own default.
func TestPropertyBoolDefault(t *testing.T) {
	for _, tc := range []struct {
		name     string
		awxValue any
		expected string
	}{
		{name: "true reaches the schema", awxValue: true, expected: "booldefault.StaticBool(true)"},
		{name: "false reaches the schema", awxValue: false, expected: "booldefault.StaticBool(false)"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var p Property
			p.Name = "enabled"
			values := map[string]any{"type": "boolean", "label": "Enabled", "default": tc.awxValue}
			require.NoError(t, p.Update(TypeWrite, PropertyOverride{}, values, Item{Name: "Host"}))
			require.Equal(t, tc.expected, p.DefaultValue)
		})
	}
}
