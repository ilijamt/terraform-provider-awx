package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// roleDefinitionTerraformModel maps the schema for RoleDefinition when using Data Source
type roleDefinitionTerraformModel struct {
	// ContentType "The type of resource this applies to"
	ContentType types.String `tfsdk:"content_type" json:"content_type"`
	// CreatedBy "The user who created this resource"
	CreatedBy types.Int64 `tfsdk:"created_by" json:"created_by"`
	// Description "Optional description of this role definition."
	Description types.String `tfsdk:"description" json:"description"`
	// ID "Database ID for this role definition."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// Managed ""
	Managed types.Bool `tfsdk:"managed" json:"managed"`
	// ModifiedBy "The user who last modified this resource"
	ModifiedBy types.Int64 `tfsdk:"modified_by" json:"modified_by"`
	// Name "Name of this role definition."
	Name types.String `tfsdk:"name" json:"name"`
	// Permissions ""
	Permissions types.List `tfsdk:"permissions" json:"permissions"`
}

// Clone the object
func (o *roleDefinitionTerraformModel) Clone() roleDefinitionTerraformModel {
	return roleDefinitionTerraformModel{
		ContentType: o.ContentType,
		CreatedBy:   o.CreatedBy,
		Description: o.Description,
		ID:          o.ID,
		Managed:     o.Managed,
		ModifiedBy:  o.ModifiedBy,
		Name:        o.Name,
		Permissions: o.Permissions,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for RoleDefinition
func (o *roleDefinitionTerraformModel) BodyRequest() (req roleDefinitionBodyRequestModel) {
	req.ContentType = o.ContentType.ValueString()
	req.Description = o.Description.ValueString()
	req.Name = o.Name.ValueString()
	req.Permissions = []string{}
	for _, val := range o.Permissions.Elements() {
		if _, ok := val.(types.String); ok {
			req.Permissions = append(req.Permissions, val.(types.String).ValueString())
		} else {
			req.Permissions = append(req.Permissions, val.String())
		}
	}
	return
}

func (o *roleDefinitionTerraformModel) setContentType(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.ContentType, data, false)
}

func (o *roleDefinitionTerraformModel) setCreatedBy(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.CreatedBy, data)
}

func (o *roleDefinitionTerraformModel) setDescription(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *roleDefinitionTerraformModel) setID(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *roleDefinitionTerraformModel) setManaged(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.Managed, data)
}

func (o *roleDefinitionTerraformModel) setModifiedBy(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ModifiedBy, data)
}

func (o *roleDefinitionTerraformModel) setName(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *roleDefinitionTerraformModel) setPermissions(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetListString(&o.Permissions, data, false)
}

func (o *roleDefinitionTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setContentType(data["content_type"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setCreatedBy(data["created_by"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setManaged(data["managed"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setModifiedBy(data["modified_by"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setName(data["name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setPermissions(data["permissions"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// roleDefinitionBodyRequestModel maps the schema for RoleDefinition for creating and updating the data
type roleDefinitionBodyRequestModel struct {
	// ContentType "The type of resource this applies to"
	ContentType string `json:"content_type,omitempty"`
	// Description "Optional description of this role definition."
	Description string `json:"description,omitempty"`
	// Name "Name of this role definition."
	Name string `json:"name"`
	// Permissions ""
	Permissions []string `json:"permissions"`
}
