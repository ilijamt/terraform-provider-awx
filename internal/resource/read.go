package resource

import (
	"context"
	"fmt"
	"net/http"
	p "path"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/client"
)

// Read performs a GET request to retrieve a resource from an API and updates the resource state.
//
// It constructs the appropriate endpoint URL by combining the base endpoint from CallInfo
// with the resource ID, sends a GET request to that endpoint, and updates the provided state
// with the response data.
func Read(ctx context.Context, client client.Client, rci CallInfo, id types.Int64, state Updater) (d diag.Diagnostics, err error) {
	var r *http.Request
	d = make(diag.Diagnostics, 0)
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", rci.Endpoint, id.ValueInt64())) + "/"
	if r, err = client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		d.AddError(
			fmt.Sprintf("Unable to create a new request for %s on %s for read", rci.Name, endpoint),
			err.Error(),
		)
		return d, err
	}

	var data map[string]any
	if data, err = client.Do(ctx, r); err != nil {
		d.AddError(
			fmt.Sprintf("Unable to read resource for %s on %s", rci.Name, endpoint),
			err.Error(),
		)
		return d, err
	}

	var dState diag.Diagnostics
	dState, err = state.UpdateWithApiData(data)
	d.Append(dState...)
	return d, err
}
