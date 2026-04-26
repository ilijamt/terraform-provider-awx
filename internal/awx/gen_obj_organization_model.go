package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// organizationTerraformModel maps the schema for Organization when using Data Source
type organizationTerraformModel struct {
	// DefaultEnvironment "The default execution environment for jobs run by this organization."
	DefaultEnvironment types.Int64 `tfsdk:"default_environment" json:"default_environment"`
	// Description "Optional description of this organization."
	Description types.String `tfsdk:"description" json:"description"`
	// ID "Database ID for this organization."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// MaxHosts "Maximum number of hosts allowed to be managed by this organization."
	MaxHosts types.Int64 `tfsdk:"max_hosts" json:"max_hosts"`
	// Name "Name of this organization."
	Name types.String `tfsdk:"name" json:"name"`
}

// Clone the object
func (o *organizationTerraformModel) Clone() organizationTerraformModel {
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Organization
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
	{
		dg, _ := helpers.AttrValueSetInt64(&o.DefaultEnvironment, data["default_environment"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Description, data["description"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.ID, data["id"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.MaxHosts, data["max_hosts"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Name, data["name"], false)
		diags.Append(dg...)
	}
	return diags, nil
}

// organizationBodyRequestModel maps the schema for Organization for creating and updating the data
type organizationBodyRequestModel struct {
	// DefaultEnvironment "The default execution environment for jobs run by this organization."
	DefaultEnvironment int64 `json:"default_environment,omitempty"`
	// Description "Optional description of this organization."
	Description string `json:"description,omitempty"`
	// MaxHosts "Maximum number of hosts allowed to be managed by this organization."
	MaxHosts int64 `json:"max_hosts,omitempty"`
	// Name "Name of this organization."
	Name string `json:"name"`
}
