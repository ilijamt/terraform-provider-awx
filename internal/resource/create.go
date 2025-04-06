package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/ilijamt/terraform-provider-awx/internal/client"
)

func Create(ctx context.Context, client client.Client, name, endpoint string) (d diag.Diagnostics, err error) {
	return d, err
}
