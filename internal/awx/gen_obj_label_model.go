package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// labelTerraformModel maps the schema for Label when using Data Source
type labelTerraformModel struct {
	// ID "Database ID for this label."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// Name "Name of this label."
	Name types.String `tfsdk:"name" json:"name"`
	// Organization "Organization this label belongs to."
	Organization types.Int64 `tfsdk:"organization" json:"organization"`
}

// Clone the object
func (o *labelTerraformModel) Clone() labelTerraformModel {
	return labelTerraformModel{
		ID:           o.ID,
		Name:         o.Name,
		Organization: o.Organization,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Label
func (o *labelTerraformModel) BodyRequest() (req labelBodyRequestModel) {
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	return
}

func (o *labelTerraformModel) setID(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *labelTerraformModel) setName(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *labelTerraformModel) setOrganization(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Organization, data)
}

func (o *labelTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
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

// labelBodyRequestModel maps the schema for Label for creating and updating the data
type labelBodyRequestModel struct {
	// Name "Name of this label."
	Name string `json:"name"`
	// Organization "Organization this label belongs to."
	Organization int64 `json:"organization"`
}
