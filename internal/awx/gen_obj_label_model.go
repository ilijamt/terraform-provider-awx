package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type labelTerraformModel struct {
	ID           types.Int64  `tfsdk:"id" json:"id"`
	Name         types.String `tfsdk:"name" json:"name"`
	Organization types.Int64  `tfsdk:"organization" json:"organization"`
}

func (o *labelTerraformModel) Clone() labelTerraformModel {
	return *o
}

func (o *labelTerraformModel) BodyRequest() *labelBodyRequestModel {
	var req labelBodyRequestModel
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	return &req
}

func (o *labelTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	return diags, nil
}

type labelBodyRequestModel struct {
	Name         string `json:"name"`
	Organization int64  `json:"organization"`
}
