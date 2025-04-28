package resource

import (
	"context"
	"fmt"
	"net/http"
	p "path"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/ilijamt/terraform-provider-awx/internal/client"
)

type deleteResource interface {
	Id
}

// Delete performs a DELETE request to remove a resource from an API.
//
// It constructs the appropriate endpoint URL by combining the base endpoint from CallInfo
// with the resource ID and sends a DELETE request to that endpoint.
func Delete(ctx context.Context, client client.Client, rci CallInfo, data deleteResource) (d diag.Diagnostics, err error) {
	d = make(diag.Diagnostics, 0)
	if client == nil {
		err = fmt.Errorf("client is nil")
		d.AddError("unable to delete resource", err.Error())
		return d, err
	}

	var id string
	if id, err = data.GetId(); err != nil {
		d.AddError("unable to get id for resource", err.Error())
		return d, err
	}

	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s/%s", rci.Endpoint, id)) + "/"
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
