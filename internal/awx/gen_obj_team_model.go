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
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Team
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
	{
		dg, _ := helpers.AttrValueSetString(&o.Description, data["description"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.ID, data["id"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Name, data["name"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.Organization, data["organization"])
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
