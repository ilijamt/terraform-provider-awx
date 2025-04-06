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

// Delete performs a DELETE request to remove a resource from an API.
//
// It constructs the appropriate endpoint URL by combining the base endpoint from CallInfo
// with the resource ID and sends a DELETE request to that endpoint.
func Delete(ctx context.Context, client client.Client, rci CallInfo, id types.Int64) (d diag.Diagnostics, err error) {
	var r *http.Request
	d = make(diag.Diagnostics, 0)
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", rci.Endpoint, id.ValueInt64())) + "/"
	if r, err = client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		d.AddError(
			fmt.Sprintf("Unable to create a new request for %s on %s for delete", rci.Name, endpoint),
			err.Error(),
		)
		return d, err
	}

	if _, err = client.Do(ctx, r); err != nil {
		d.AddError(
			fmt.Sprintf("Unable to delete resource for %s on %s", rci.Name, endpoint),
			err.Error(),
		)
	}
	return d, err
}
