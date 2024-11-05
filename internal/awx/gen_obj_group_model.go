package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// groupTerraformModel maps the schema for Group when using Data Source
type groupTerraformModel struct {
	// Description "Optional description of this group."
	Description types.String `tfsdk:"description" json:"description"`
	// ID "Database ID for this group."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// Inventory ""
	Inventory types.Int64 `tfsdk:"inventory" json:"inventory"`
	// Name "Name of this group."
	Name types.String `tfsdk:"name" json:"name"`
	// Variables "Group variables in JSON or YAML format."
	Variables types.String `tfsdk:"variables" json:"variables"`
}

// Clone the object
func (o *groupTerraformModel) Clone() groupTerraformModel {
	return groupTerraformModel{
		Description: o.Description,
		ID:          o.ID,
		Inventory:   o.Inventory,
		Name:        o.Name,
		Variables:   o.Variables,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Group
func (o *groupTerraformModel) BodyRequest() (req groupBodyRequestModel) {
	req.Description = o.Description.ValueString()
	req.Inventory = o.Inventory.ValueInt64()
	req.Name = o.Name.ValueString()
	req.Variables = json.RawMessage(o.Variables.ValueString())
	return
}

func (o *groupTerraformModel) setDescription(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *groupTerraformModel) setID(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *groupTerraformModel) setInventory(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Inventory, data)
}

func (o *groupTerraformModel) setName(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *groupTerraformModel) setVariables(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetJsonString(&o.Variables, data, false)
}

func (o *groupTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setInventory(data["inventory"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setName(data["name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setVariables(data["variables"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// groupBodyRequestModel maps the schema for Group for creating and updating the data
type groupBodyRequestModel struct {
	// Description "Optional description of this group."
	Description string `json:"description,omitempty"`
	// Inventory ""
	Inventory int64 `json:"inventory"`
	// Name "Name of this group."
	Name string `json:"name"`
	// Variables "Group variables in JSON or YAML format."
	Variables json.RawMessage `json:"variables,omitempty"`
}
