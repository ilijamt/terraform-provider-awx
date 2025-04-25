package resource

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/ilijamt/terraform-provider-awx/internal/client"
)

func Create(ctx context.Context, client client.Client, rci CallInfo, data any) (d diag.Diagnostics, err error) {
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

	return d, err
}
