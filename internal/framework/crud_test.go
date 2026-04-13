package framework_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

func TestCreateUpdateRequest(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name          string
		requester     framework.Requester
		method        string
		body          any
		expectError   bool
		expectDataKey string
		expectDataVal any
	}{
		{
			name:          "success POST",
			requester:     successRequester(map[string]any{"id": 1}),
			method:        http.MethodPost,
			body:          map[string]string{"name": "test"},
			expectError:   false,
			expectDataKey: "id",
			expectDataVal: 1,
		},
		{
			name:          "success PATCH",
			requester:     successRequester(map[string]any{"id": 2}),
			method:        http.MethodPatch,
			body:          map[string]string{"name": "test"},
			expectError:   false,
			expectDataKey: "id",
			expectDataVal: 2,
		},
		{
			name:        "NewRequest fails",
			requester:   failNewRequest(),
			method:      http.MethodPost,
			body:        nil,
			expectError: true,
		},
		{
			name:        "Do fails",
			requester:   failDo(),
			method:      http.MethodPost,
			body:        nil,
			expectError: true,
		},
		{
			name:        "nil body succeeds",
			requester:   successRequester(map[string]any{}),
			method:      http.MethodPost,
			body:        nil,
			expectError: false,
		},
		{
			// Documents current behavior: encoding errors are silently
			// ignored (crud.go:51). An unencodable body results in an
			// empty JSON payload being sent.
			name:        "unencodable body silent fail",
			requester:   successRequester(map[string]any{}),
			method:      http.MethodPost,
			body:        make(chan int),
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, diags := framework.CreateUpdateRequest(ctx, tt.requester, tt.method, "/api/v2/test/", tt.body, "TestResource", "create")
			if tt.expectError {
				assert.True(t, diags.HasError())
				return
			}
			require.False(t, diags.HasError())
			if tt.expectDataKey != "" {
				assert.Equal(t, tt.expectDataVal, data[tt.expectDataKey])
			}
		})
	}
}

func TestReadRequest(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name          string
		requester     framework.Requester
		expectError   bool
		expectDataKey string
		expectDataVal any
	}{
		{
			name:          "success",
			requester:     successRequester(map[string]any{"name": "test"}),
			expectError:   false,
			expectDataKey: "name",
			expectDataVal: "test",
		},
		{
			name:        "NewRequest fails",
			requester:   failNewRequest(),
			expectError: true,
		},
		{
			name:        "Do fails",
			requester:   failDo(),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, diags := framework.ReadRequest(ctx, tt.requester, "/api/v2/test/1/", "TestResource")
			if tt.expectError {
				assert.True(t, diags.HasError())
				return
			}
			require.False(t, diags.HasError())
			if tt.expectDataKey != "" {
				assert.Equal(t, tt.expectDataVal, data[tt.expectDataKey])
			}
		})
	}
}

func TestDeleteRequest(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		requester   framework.Requester
		expectError bool
	}{
		{
			name:        "success",
			requester:   successRequester(nil),
			expectError: false,
		},
		{
			name:        "NewRequest fails",
			requester:   failNewRequest(),
			expectError: true,
		},
		{
			name:        "Do fails",
			requester:   failDo(),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := framework.DeleteRequest(ctx, tt.requester, "/api/v2/test/1/", "TestResource")
			assert.Equal(t, tt.expectError, diags.HasError())
		})
	}
}
