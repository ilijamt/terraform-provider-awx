package awx_21_8_0

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
	return organizationTerraformModel{
		DefaultEnvironment: o.DefaultEnvironment,
		Description:        o.Description,
		ID:                 o.ID,
		MaxHosts:           o.MaxHosts,
		Name:               o.Name,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Organization
func (o *organizationTerraformModel) BodyRequest() (req organizationBodyRequestModel) {
	req.DefaultEnvironment = o.DefaultEnvironment.ValueInt64()
	req.Description = o.Description.ValueString()
	req.MaxHosts = o.MaxHosts.ValueInt64()
	req.Name = o.Name.ValueString()
	return
}

func (o *organizationTerraformModel) setDefaultEnvironment(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.DefaultEnvironment, data)
}

func (o *organizationTerraformModel) setDescription(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *organizationTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *organizationTerraformModel) setMaxHosts(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.MaxHosts, data)
}

func (o *organizationTerraformModel) setName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *organizationTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setDefaultEnvironment(data["default_environment"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setMaxHosts(data["max_hosts"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setName(data["name"]); dg.HasError() {
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

type organizationObjectRolesModel struct {
	ID    types.Int64 `tfsdk:"id"`
	Roles types.Map   `tfsdk:"roles"`
}
