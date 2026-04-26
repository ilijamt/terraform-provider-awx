package framework

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// LookupCredentialTypeIDByNamespace resolves the AWX-side credential_type ID
// for a managed credential type identified by its namespace (e.g. "aws",
// "ssh", "vault"). The namespace is the canonical key AWX uses internally —
// names like "Amazon Web Services" can vary by language pack but namespace is
// stable across instances.
//
// Returns 0 + an error diagnostic if no matching credential type is found.
// Used by the OnConfigure hook on each typed credential resource so the
// resource picks up the correct local ID instead of a hardcoded one.
func LookupCredentialTypeIDByNamespace(ctx context.Context, client Requester, namespace string) (int64, diag.Diagnostics) {
	diags := diag.Diagnostics{}
	if namespace == "" {
		diags.AddError("Empty credential type namespace", "namespace must not be empty")
		return 0, diags
	}

	endpoint := fmt.Sprintf("/api/v2/credential_types/?managed=true&namespace=%s", url.QueryEscape(namespace))
	data, d := ReadRequest(ctx, client, endpoint, fmt.Sprintf("CredentialType[%s]", namespace))
	diags.Append(d...)
	if d.HasError() {
		return 0, diags
	}

	results, ok := data["results"].([]any)
	if !ok || len(results) == 0 {
		diags.AddError(
			fmt.Sprintf("No managed credential_type found for namespace %q", namespace),
			fmt.Sprintf("AWX returned 0 results for %s. Confirm the AWX instance has the managed credential type for this namespace.", endpoint),
		)
		return 0, diags
	}

	first, ok := results[0].(map[string]any)
	if !ok {
		diags.AddError(
			fmt.Sprintf("Unexpected response shape for credential_type lookup (%s)", namespace),
			"results[0] is not an object",
		)
		return 0, diags
	}

	switch id := first["id"].(type) {
	case json.Number:
		v, err := id.Int64()
		if err != nil {
			diags.AddError(
				fmt.Sprintf("Unparsable id for credential_type %q", namespace),
				err.Error(),
			)
			return 0, diags
		}
		return v, diags
	case float64:
		return int64(id), diags
	case int64:
		return id, diags
	case int:
		return int64(id), diags
	default:
		diags.AddError(
			fmt.Sprintf("Unexpected id type for credential_type %q", namespace),
			fmt.Sprintf("expected number, got %T", first["id"]),
		)
		return 0, diags
	}
}
