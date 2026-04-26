package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type organizationTerraformModel struct {
	DefaultEnvironment types.Int64  `tfsdk:"default_environment" json:"default_environment"`
	Description        types.String `tfsdk:"description" json:"description"`
	ID                 types.Int64  `tfsdk:"id" json:"id"`
	MaxHosts           types.Int64  `tfsdk:"max_hosts" json:"max_hosts"`
	Name               types.String `tfsdk:"name" json:"name"`
}

func (o *organizationTerraformModel) Clone() organizationTerraformModel {
	return *o
}

func (o *organizationTerraformModel) BodyRequest() *organizationBodyRequestModel {
	var req organizationBodyRequestModel
	req.DefaultEnvironment = o.DefaultEnvironment.ValueInt64()
	req.Description = o.Description.ValueString()
	req.MaxHosts = o.MaxHosts.ValueInt64()
	req.Name = o.Name.ValueString()
	return &req
}

func (o *organizationTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.DefaultEnvironment, data["default_environment"]))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetInt64(&o.MaxHosts, data["max_hosts"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	return diags, nil
}

type organizationBodyRequestModel struct {
	DefaultEnvironment int64  `json:"default_environment,omitempty"`
	Description        string `json:"description,omitempty"`
	MaxHosts           int64  `json:"max_hosts,omitempty"`
	Name               string `json:"name"`
}
