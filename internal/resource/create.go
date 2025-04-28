package resource

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	p "path"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/ilijamt/terraform-provider-awx/internal/client"
)

type createResource interface {
	Updater
	RequestBody
}

func Create(ctx context.Context, client client.Client, rci CallInfo, data createResource) (d diag.Diagnostics, err error) {
	d = make(diag.Diagnostics, 0)

	if client == nil {
		err = errors.Join(err, fmt.Errorf("client is nil"))
	}
	if data == nil {
		err = errors.Join(err, fmt.Errorf("data is nil"))
	}

	if err != nil {
		d.AddError("unable to create resource", err.Error())
		return d, err
	}

	var r *http.Request
	var endpoint = p.Clean(rci.Endpoint) + "/"
	var buf bytes.Buffer
	payload, _ := data.RequestBody()
	buf.Write(payload)

	tflog.Debug(ctx, "Preparing a request to create a resource", map[string]any{
		"data":     data,
		"method":   http.MethodPost,
		"rci":      rci,
		"endpoint": endpoint,
	})
	if r, err = client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		d.AddError(
			fmt.Sprintf("Unable to create a new request for %s on %s for create", rci.Name, endpoint),
			err.Error(),
		)
		return
	}

	var rData map[string]any
	if rData, err = client.Do(ctx, r); err != nil {
		d.AddError(
			fmt.Sprintf("Unable to create resource for %s on %s", rci.Name, endpoint),
			err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "Resource created", map[string]any{
		"rci":  rci,
		"data": rData,
	})

	var dState diag.Diagnostics
	dState, err = data.UpdateWithApiData(rci.Callee, rci.Source, rData)
	d.Append(dState...)
	return d, err
}
