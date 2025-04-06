package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/ilijamt/terraform-provider-awx/internal/client"
)

func GetCredentialTypeID(ctx context.Context, client client.Client, rci CallInfo, name string) (d diag.Diagnostics, err error) {
	d = make(diag.Diagnostics, 0)
	return d, err
}
