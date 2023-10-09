package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// teamTerraformModel maps the schema for Team when using Data Source
type teamTerraformModel struct {
	// Description "Optional description of this team."
	Description types.String `tfsdk:"description" json:"description"`
	// ID "Database ID for this team."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// Name "Name of this team."
	Name types.String `tfsdk:"name" json:"name"`
	// Organization ""
	Organization types.Int64 `tfsdk:"organization" json:"organization"`
}

// Clone the object
func (o *teamTerraformModel) Clone() teamTerraformModel {
	return teamTerraformModel{
		Description:  o.Description,
		ID:           o.ID,
		Name:         o.Name,
		Organization: o.Organization,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Team
func (o *teamTerraformModel) BodyRequest() (req teamBodyRequestModel) {
	req.Description = o.Description.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	return
}

func (o *teamTerraformModel) setDescription(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *teamTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *teamTerraformModel) setName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *teamTerraformModel) setOrganization(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Organization, data)
}

func (o *teamTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setName(data["name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setOrganization(data["organization"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// teamBodyRequestModel maps the schema for Team for creating and updating the data
type teamBodyRequestModel struct {
	// Description "Optional description of this team."
	Description string `json:"description,omitempty"`
	// Name "Name of this team."
	Name string `json:"name"`
	// Organization ""
	Organization int64 `json:"organization"`
}

type teamObjectRolesModel struct {
	ID    types.Int64 `tfsdk:"id"`
	Roles types.Map   `tfsdk:"roles"`
}
