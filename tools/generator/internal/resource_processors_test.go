package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourceProcessorSeedsReadAction(t *testing.T) {
	for _, tt := range []struct {
		name string
		item string
		in   map[string]any
		want map[string]any
	}{
		{
			name: "config gets the read action AWX omits",
			item: "Config",
			in:   map[string]any{"description": "Configuration"},
			want: map[string]any{
				"description": "Configuration",
				"actions":     map[string]any{"GET": map[string]any{}},
			},
		},
		{
			name: "ping gets it too",
			item: "Ping",
			in:   map[string]any{},
			want: map[string]any{"actions": map[string]any{"GET": map[string]any{}}},
		},
		{
			name: "an endpoint that reports its own actions is left alone",
			item: "Config",
			in:   map[string]any{"actions": map[string]any{"GET": map[string]any{"version": map[string]any{}}}},
			want: map[string]any{"actions": map[string]any{"GET": map[string]any{"version": map[string]any{}}}},
		},
		{
			name: "an unregistered item passes through untouched",
			item: "Inventory",
			in:   map[string]any{"description": "Inventories"},
			want: map[string]any{"description": "Inventories"},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResourceProcessor(tt.item, tt.in)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
