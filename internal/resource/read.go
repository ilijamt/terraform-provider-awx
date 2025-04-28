package resource

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	p "path"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/ilijamt/terraform-provider-awx/internal/client"
)

type readResource interface {
	Id
	Updater
}

// Read performs a GET request to retrieve a resource from an API and updates the resource state.
//
// It constructs the appropriate endpoint URL by combining the base endpoint from CallInfo
// with the resource ID, sends a GET request to that endpoint, and updates the provided state
// with the response data.
func Read(ctx context.Context, client client.Client, rci CallInfo, data readResource) (d diag.Diagnostics, err error) {
	d = make(diag.Diagnostics, 0)
	if client == nil {
		err = errors.Join(err, fmt.Errorf("client is nil"))
	}

	if data == nil {
		err = errors.Join(err, fmt.Errorf("state updater is nil"))
	}

	if err != nil {
		d.AddError("unable to read resource", err.Error())
		return d, err
	}

	var id string
	if id, err = data.GetId(); err != nil {
		d.AddError("unable to get id for resource", err.Error())
		return d, err
	}

	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", rci.Endpoint, id)) + "/"
	if r, err = client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		d.AddError(
			fmt.Sprintf("Unable to create a new request for %s on %s for read", rci.Name, endpoint),
			err.Error(),
		)
		return d, err
	}

	var rData map[string]any
	if rData, err = client.Do(ctx, r); err != nil {
		d.AddError(
			fmt.Sprintf("Unable to read resource for %s on %s", rci.Name, endpoint),
			err.Error(),
		)
		return d, err
	}

	var dState diag.Diagnostics
	dState, err = data.UpdateWithApiData(rci.Callee, rci.Source, rData)
	d.Append(dState...)
	return d, err
}
