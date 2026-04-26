package framework_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

func TestCleanEndpoint(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "already clean with trailing slash",
			input:    "/api/v2/inventories/",
			expected: "/api/v2/inventories/",
		},
		{
			name:     "missing trailing slash",
			input:    "/api/v2/inventories",
			expected: "/api/v2/inventories/",
		},
		{
			name:     "double slashes collapsed",
			input:    "/api/v2//inventories/",
			expected: "/api/v2/inventories/",
		},
		{
			name:     "root path",
			input:    "/",
			expected: "//",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "./",
		},
		{
			name:     "dot-dot segments resolved",
			input:    "/api/v2/../v3/inventories",
			expected: "/api/v3/inventories/",
		},
		{
			name:     "multiple trailing slashes",
			input:    "/api/v2/inventories///",
			expected: "/api/v2/inventories/",
		},
		{
			name:     "single segment",
			input:    "/inventories",
			expected: "/inventories/",
		},
		{
			name:     "deeply nested",
			input:    "/a/b/c/d/e",
			expected: "/a/b/c/d/e/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, framework.CleanEndpoint(tt.input))
		})
	}
}

func TestEndpointWithID(t *testing.T) {
	tests := []struct {
		name     string
		base     string
		id       any
		expected string
	}{
		{
			name:     "int64 id",
			base:     "/api/v2/inventories",
			id:       int64(42),
			expected: "/api/v2/inventories/42/",
		},
		{
			name:     "string id",
			base:     "/api/v2/inventories",
			id:       "42",
			expected: "/api/v2/inventories/42/",
		},
		{
			name:     "base with trailing slash",
			base:     "/api/v2/inventories/",
			id:       int64(1),
			expected: "/api/v2/inventories/1/",
		},
		{
			name:     "int id",
			base:     "/api/v2/inventories",
			id:       int(7),
			expected: "/api/v2/inventories/7/",
		},
		{
			name:     "zero id",
			base:     "/api/v2/inventories",
			id:       int64(0),
			expected: "/api/v2/inventories/0/",
		},
		{
			name:     "negative id",
			base:     "/api/v2/inventories",
			id:       int64(-1),
			expected: "/api/v2/inventories/-1/",
		},
		{
			name:     "empty base",
			base:     "",
			id:       int64(5),
			expected: "/5/",
		},
		{
			name:     "float id uses fmt %v",
			base:     "/api/v2/inventories",
			id:       3.14,
			expected: "/api/v2/inventories/3.14/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, framework.EndpointWithID(tt.base, tt.id))
		})
	}
}
