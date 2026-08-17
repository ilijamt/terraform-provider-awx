package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNormalizeResourcePayload(t *testing.T) {
	for _, tt := range []struct {
		name string
		in   map[string]any
		want map[string]any
	}{
		{
			name: "sorts related_search_fields",
			in:   map[string]any{"related_search_fields": []any{"modified_by__search", "created_by__search", "labels__search"}},
			want: map[string]any{"related_search_fields": []any{"created_by__search", "labels__search", "modified_by__search"}},
		},
		{
			name: "leaves ordered lists alone",
			in:   map[string]any{"renders": []any{"application/json", "text/html"}},
			want: map[string]any{"renders": []any{"application/json", "text/html"}},
		},
		{
			name: "tolerates a missing key",
			in:   map[string]any{"actions": map[string]any{}},
			want: map[string]any{"actions": map[string]any{}},
		},
		{
			name: "leaves a mixed list untouched",
			in:   map[string]any{"related_search_fields": []any{"b", 1, "a"}},
			want: map[string]any{"related_search_fields": []any{"b", 1, "a"}},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, NormalizeResourcePayload(tt.in))
		})
	}
}

func TestNormalizeCredentialTypePayload(t *testing.T) {
	in := map[string]any{
		"name":     "Machine",
		"kind":     "ssh",
		"created":  "2024-07-12T20:38:21.618192Z",
		"modified": "2024-07-12T20:38:21.618192Z",
		"inputs": map[string]any{"fields": []any{
			map[string]any{"id": "cloud_name", "choices": []any{"AzureGermanCloud", "AzureChinaCloud", "AzureCloud"}},
			map[string]any{"id": "secret"},
		}},
	}

	require.Equal(t, map[string]any{
		"name": "Machine",
		"kind": "ssh",
		"inputs": map[string]any{"fields": []any{
			map[string]any{"id": "cloud_name", "choices": []any{"AzureChinaCloud", "AzureCloud", "AzureGermanCloud"}},
			map[string]any{"id": "secret"},
		}},
	}, NormalizeCredentialTypePayload(in))
}
