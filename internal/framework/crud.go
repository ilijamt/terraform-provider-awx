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

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
)

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
		tflog.Trace(ctx, fmt.Sprintf("[%s/%s] Request failed", resourceName, operation), map[string]any{
			"method":   method,
			"endpoint": endpoint,
			"response": data,
			"error":    err.Error(),
		})
		diags.AddError(
			fmt.Sprintf("Unable to %s resource for %s on %s", operation, resourceName, endpoint),
			err.Error(),
		)
		return nil, diags
	}

	tflog.Trace(ctx, fmt.Sprintf("[%s/%s] Request succeeded", resourceName, operation), map[string]any{
		"method":   method,
		"endpoint": endpoint,
		"response": data,
	})

	return data, diags
}

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

func ReadRequest(ctx context.Context, r Requester, endpoint string, resourceName string) (map[string]any, diag.Diagnostics) {
	return doRequest(ctx, r, http.MethodGet, endpoint, nil, resourceName, "read")
}

// ReadRequestAllowMissing reports a 404 as found=false rather than an error.
func ReadRequestAllowMissing(ctx context.Context, r Requester, endpoint string, resourceName string) (map[string]any, bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	req, err := r.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		diags.AddError(
			fmt.Sprintf("Unable to create a new request for %s on %s for read", resourceName, endpoint),
			err.Error(),
		)
		return nil, false, diags
	}

	data, err := r.Do(ctx, req)
	if err != nil {
		if c.IsNotFound(err) {
			tflog.Debug(ctx, fmt.Sprintf("[%s/read] gone from AWX, dropping from state", resourceName), map[string]any{
				"endpoint": endpoint,
			})
			return nil, false, diags
		}
		diags.AddError(
			fmt.Sprintf("Unable to read resource for %s on %s", resourceName, endpoint),
			err.Error(),
		)
		return nil, false, diags
	}

	return data, true, diags
}

func DeleteRequest(ctx context.Context, r Requester, endpoint string, resourceName string) diag.Diagnostics {
	_, diags := doRequest(ctx, r, http.MethodDelete, endpoint, nil, resourceName, "delete")
	return diags
}
