package resource

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/ilijamt/terraform-provider-awx/internal/client"
)

func GetCredentialTypeID(ctx context.Context, client client.Client, rci CallInfo, name string) (d diag.Diagnostics, err error) {
	d = make(diag.Diagnostics, 0)
	if client == nil {
		err = errors.Join(err, fmt.Errorf("client is nil"))
	}
	name = strings.TrimSpace(name)
	if name == "" {
		err = errors.Join(err, fmt.Errorf("name is empty"))
	}

	if err != nil {
		d.AddError("unable to fetch credential type id", err.Error())
		return d, err
	}

	return d, err
}
