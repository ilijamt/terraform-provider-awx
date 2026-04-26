package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type teamTerraformModel struct {
	Description  types.String `tfsdk:"description" json:"description"`
	ID           types.Int64  `tfsdk:"id" json:"id"`
	Name         types.String `tfsdk:"name" json:"name"`
	Organization types.Int64  `tfsdk:"organization" json:"organization"`
}

func (o *teamTerraformModel) Clone() teamTerraformModel {
	return *o
}

func (o *teamTerraformModel) BodyRequest() *teamBodyRequestModel {
	var req teamBodyRequestModel
	req.Description = o.Description.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	return &req
}

func (o *teamTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	return diags, nil
}

type teamBodyRequestModel struct {
	Description  string `json:"description,omitempty"`
	Name         string `json:"name"`
	Organization int64  `json:"organization"`
}
