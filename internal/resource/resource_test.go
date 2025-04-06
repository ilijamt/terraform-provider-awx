package resource_test

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/ilijamt/terraform-provider-awx/internal/resource"
)

type resourceUpdaterTest struct {
	diags diag.Diagnostics
	err   error
}

func (r resourceUpdaterTest) UpdateWithApiData(data map[string]any) (_ diag.Diagnostics, _ error) {
	return r.diags, r.err
}

var _ resource.Updater = (*resourceUpdaterTest)(nil)
