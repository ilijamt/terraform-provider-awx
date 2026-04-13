package framework

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// doRequest builds and executes an HTTP request, returning the decoded JSON response.
// All public CRUD functions delegate to this.
func doRequest(ctx context.Context, r Requester, method string, endpoint string, body io.Reader, resourceName string, operation string) (map[string]any, diag.Diagnostics) {
	var diags diag.Diagnostics

	req, err := r.NewRequest(ctx, method, endpoint, body)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Unable to create a new request for %s on %s for %s", resourceName, endpoint, operation),
			err.Error(),
		)
		return nil, diags
	}

	data, err := r.Do(ctx, req)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Unable to %s resource for %s on %s", operation, resourceName, endpoint),
			err.Error(),
		)
		return nil, diags
	}

	return data, diags
}

// CreateUpdateRequest encodes body as JSON, sends a POST or PATCH request, and
// returns the decoded response.
func CreateUpdateRequest(ctx context.Context, r Requester, method string, endpoint string, body any, resourceName string, operation string) (map[string]any, diag.Diagnostics) {
	tflog.Debug(ctx, fmt.Sprintf("[%s/%s] Making a request", resourceName, operation), map[string]any{
		"payload":  body,
		"method":   method,
		"endpoint": endpoint,
	})

	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(body)

	return doRequest(ctx, r, method, endpoint, &buf, resourceName, operation)
}

// ReadRequest sends a GET request and returns the decoded response.
func ReadRequest(ctx context.Context, r Requester, endpoint string, resourceName string) (map[string]any, diag.Diagnostics) {
	return doRequest(ctx, r, http.MethodGet, endpoint, nil, resourceName, "read")
}

// DeleteRequest sends a DELETE request.
func DeleteRequest(ctx context.Context, r Requester, endpoint string, resourceName string) diag.Diagnostics {
	_, diags := doRequest(ctx, r, http.MethodDelete, endpoint, nil, resourceName, "delete")
	return diags
}
