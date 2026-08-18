package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateOverrides(t *testing.T) {
	objmap := map[string]any{"actions": map[string]any{
		"GET": map[string]any{
			"id":      map[string]any{"type": "integer"},
			"status":  map[string]any{"type": "string"},
			"timeout": map[string]any{"type": "integer"},
		},
		"POST": map[string]any{
			"name":     map[string]any{"type": "string", "required": true},
			"enabled":  map[string]any{"type": "boolean"},
			"timeout":  map[string]any{"type": "integer"},
			"defaults": map[string]any{"type": "integer", "required": true, "default": 5},
		},
	}}
	base := func(item Item) Item {
		item.ApiPropertyDataKey, item.ApiPropertyResourceKey = "GET", "POST"
		return item
	}

	for _, tc := range []struct {
		name string
		item Item
		want string
	}{
		{name: "no overrides", item: Item{}},
		{
			name: "override names a field neither action carries",
			item: Item{PropertyOverrides: map[string]PropertyOverride{"nope": {Trim: true}}},
			want: `property_overrides "nope" names a field absent from both actions`,
		},
		{
			name: "nullable outside a bool",
			item: Item{PropertyOverrides: map[string]PropertyOverride{"timeout": {Nullable: boolPtr(true)}}},
			want: `property_overrides "timeout" sets nullable on a integer`,
		},
		{name: "nullable on a bool", item: Item{PropertyOverrides: map[string]PropertyOverride{"enabled": {Nullable: boolPtr(true)}}}},
		{
			name: "omit_empty on a required field",
			item: Item{PropertyOverrides: map[string]PropertyOverride{"name": {OmitEmpty: boolPtr(false)}}},
			want: `property_overrides "name" sets omit_empty on a field that never carries it`,
		},
		{
			// setDefaultValue clears required, so this one is load-bearing.
			name: "omit_empty on a required field that has a default",
			item: Item{PropertyOverrides: map[string]PropertyOverride{"defaults": {OmitEmpty: boolPtr(false)}}},
		},
		{
			name: "use_state_for_unknown off on a required field",
			item: Item{PropertyOverrides: map[string]PropertyOverride{"name": {UseStateForUnknown: boolPtr(false)}}},
			want: `property_overrides "name" turns off use_state_for_unknown on a required field, which never has it`,
		},
		{
			name: "removal names a read-only field",
			item: Item{RemoveFieldsResource: []string{"status"}},
			want: `remove_fields_resource "status" is not in the POST action`,
		},
		{name: "removal of a real write field", item: Item{RemoveFieldsResource: []string{"timeout"}}},
		{
			name: "an injected field is not absent",
			item: Item{
				PropertyOverrides:       map[string]PropertyOverride{"parent_id": {RequiresReplace: true}},
				ApiDataOverrideResource: map[string]map[string]any{"parent_id": {"type": "integer"}},
			},
		},
		{
			// The field is pruned before the templates see it, so it is read-only there.
			name: "removed write field is judged read-only",
			item: Item{
				RemoveFieldsResource: []string{"timeout"},
				PropertyOverrides:    map[string]PropertyOverride{"timeout": {UseStateForUnknown: boolPtr(false)}},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := ValidateOverrides(base(tc.item), objmap)
			if tc.want == "" {
				assert.Empty(t, got)
				return
			}
			require.Len(t, got, 1)
			assert.Contains(t, got[0], tc.want)
		})
	}
}
