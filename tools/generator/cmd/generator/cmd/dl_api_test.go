package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ilijamt/terraform-provider-awx/tools/generator/internal"
)

func TestMissingActions(t *testing.T) {
	actions := func(keys ...string) map[string]any {
		out := make(map[string]any, len(keys))
		for _, k := range keys {
			out[k] = map[string]any{}
		}
		return map[string]any{"actions": out}
	}

	for _, tt := range []struct {
		name    string
		item    internal.Item
		payload map[string]any
		want    []string
	}{
		{
			name:    "a complete payload is accepted",
			item:    internal.Item{ApiPropertyResourceKey: "POST", ApiPropertyDataKey: "GET"},
			payload: actions("GET", "POST"),
		},
		{
			name:    "AWX hiding POST on an unseeded instance is caught",
			item:    internal.Item{ApiPropertyResourceKey: "POST", ApiPropertyDataKey: "GET"},
			payload: actions("GET"),
			want:    []string{"POST"},
		},
		{
			name:    "settings need their PUT block",
			item:    internal.Item{ApiPropertyResourceKey: "PUT", ApiPropertyDataKey: "GET"},
			payload: actions("GET"),
			want:    []string{"PUT"},
		},
		{
			name:    "a shared key is reported once",
			item:    internal.Item{ApiPropertyResourceKey: "GET", ApiPropertyDataKey: "GET"},
			payload: actions(),
			want:    []string{"GET"},
		},
		{
			name:    "a payload without actions reports everything",
			item:    internal.Item{ApiPropertyResourceKey: "POST", ApiPropertyDataKey: "GET"},
			payload: map[string]any{},
			want:    []string{"POST", "GET"},
		},
		{
			name: "a data-source-only item never needs the write block",
			item: internal.Item{
				ApiPropertyResourceKey: "POST",
				ApiPropertyDataKey:     "GET",
				NoTerraformResource:    true,
			},
			payload: actions("GET"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, missingActions(tt.item, tt.payload))
		})
	}
}
