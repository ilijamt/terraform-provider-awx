package framework_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
)

func TestLookupCredentialTypeIDByNamespace(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		namespace   string
		requester   framework.Requester
		wantID      int64
		wantError   bool
		wantInError string
	}{
		{
			name:      "happy path",
			namespace: "aws",
			requester: namespaceLookupRequester(t, "aws", []map[string]any{{"id": float64(5), "namespace": "aws"}}),
			wantID:    5,
		},
		{
			name:        "empty namespace",
			namespace:   "",
			requester:   nil,
			wantError:   true,
			wantInError: "namespace must not be empty",
		},
		{
			name:        "no results",
			namespace:   "ssh",
			requester:   namespaceLookupRequester(t, "ssh", []map[string]any{}),
			wantError:   true,
			wantInError: "No managed credential_type found",
		},
		{
			name:        "request fails",
			namespace:   "aws",
			requester:   failDo(),
			wantError:   true,
			wantInError: "do error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, diags := framework.LookupCredentialTypeIDByNamespace(ctx, tt.requester, tt.namespace)
			if tt.wantError {
				require.True(t, diags.HasError(), "expected diagnostic error")
				if tt.wantInError != "" {
					var combined string
					for _, d := range diags.Errors() {
						combined += d.Summary() + "|" + d.Detail() + "\n"
					}
					assert.Contains(t, combined, tt.wantInError)
				}
				return
			}
			require.False(t, diags.HasError())
			assert.Equal(t, tt.wantID, id)
		})
	}
}

// namespaceLookupRequester returns a Requester that asserts the URL contains
// the expected namespace query and returns a paginated-style response with the
// supplied results array.
func namespaceLookupRequester(t *testing.T, expectNamespace string, results []map[string]any) *mockRequester {
	t.Helper()
	return &mockRequester{
		newRequestFunc: func(_ context.Context, _, endpoint string, _ io.Reader) (*http.Request, error) {
			if !strings.Contains(endpoint, "namespace="+expectNamespace) {
				return nil, fmt.Errorf("expected namespace %q in endpoint, got %q", expectNamespace, endpoint)
			}
			if !strings.Contains(endpoint, "managed=true") {
				return nil, fmt.Errorf("expected managed=true in endpoint, got %q", endpoint)
			}
			return &http.Request{}, nil
		},
		doFunc: func(context.Context, *http.Request) (map[string]any, error) {
			any_ := make([]any, 0, len(results))
			for _, r := range results {
				any_ = append(any_, r)
			}
			return map[string]any{"results": any_}, nil
		},
	}
}
