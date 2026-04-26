package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type executionEnvironmentTerraformModel struct {
	Credential   types.Int64  `tfsdk:"credential" json:"credential"`
	Description  types.String `tfsdk:"description" json:"description"`
	ID           types.Int64  `tfsdk:"id" json:"id"`
	Image        types.String `tfsdk:"image" json:"image"`
	Managed      types.Bool   `tfsdk:"managed" json:"managed"`
	Name         types.String `tfsdk:"name" json:"name"`
	Organization types.Int64  `tfsdk:"organization" json:"organization"`
	Pull         types.String `tfsdk:"pull" json:"pull"`
}

func (o *executionEnvironmentTerraformModel) Clone() executionEnvironmentTerraformModel {
	return *o
}

func (o *executionEnvironmentTerraformModel) BodyRequest() *executionEnvironmentBodyRequestModel {
	var req executionEnvironmentBodyRequestModel
	req.Credential = o.Credential.ValueInt64()
	req.Description = o.Description.ValueString()
	req.Image = o.Image.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	req.Pull = o.Pull.ValueString()
	return &req
}

func (o *executionEnvironmentTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.Credential, data["credential"]))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.Image, data["image"], false))
	collect(helpers.AttrValueSetBool(&o.Managed, data["managed"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	collect(helpers.AttrValueSetString(&o.Pull, data["pull"], false))
	return diags, nil
}

type executionEnvironmentBodyRequestModel struct {
	Credential   int64  `json:"credential,omitempty"`
	Description  string `json:"description,omitempty"`
	Image        string `json:"image"`
	Name         string `json:"name"`
	Organization int64  `json:"organization,omitempty"`
	Pull         string `json:"pull,omitempty"`
}
