package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// roleUserAssignmentTerraformModel maps the schema for RoleUserAssignment when using Data Source
type roleUserAssignmentTerraformModel struct {
	// ContentType "The type of resource this applies to"
	ContentType types.Int64 `tfsdk:"content_type" json:"content_type"`
	// CreatedBy "The user who created this resource"
	CreatedBy types.Int64 `tfsdk:"created_by" json:"created_by"`
	// ID "Database ID for this role user assignment."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// ObjectAnsibleId "Resource id of the object this role applies to. Alternative to the object_id field."
	ObjectAnsibleId types.String `tfsdk:"object_ansible_id" json:"object_ansible_id"`
	// ObjectId "Primary key of the object this assignment applies to, null value indicates system-wide assignment"
	ObjectId types.String `tfsdk:"object_id" json:"object_id"`
	// RoleDefinition "The role definition which defines permissions conveyed by this assignment"
	RoleDefinition types.Int64 `tfsdk:"role_definition" json:"role_definition"`
	// User ""
	User types.Int64 `tfsdk:"user" json:"user"`
	// UserAnsibleId "Resource id of the user who will receive permissions from this assignment. Alternative to user field."
	UserAnsibleId types.String `tfsdk:"user_ansible_id" json:"user_ansible_id"`
}

// Clone the object
func (o *roleUserAssignmentTerraformModel) Clone() roleUserAssignmentTerraformModel {
	return roleUserAssignmentTerraformModel{
		ContentType:     o.ContentType,
		CreatedBy:       o.CreatedBy,
		ID:              o.ID,
		ObjectAnsibleId: o.ObjectAnsibleId,
		ObjectId:        o.ObjectId,
		RoleDefinition:  o.RoleDefinition,
		User:            o.User,
		UserAnsibleId:   o.UserAnsibleId,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for RoleUserAssignment
func (o *roleUserAssignmentTerraformModel) BodyRequest() (req roleUserAssignmentBodyRequestModel) {
	req.ObjectAnsibleId = o.ObjectAnsibleId.ValueString()
	req.ObjectId = o.ObjectId.ValueString()
	req.RoleDefinition = o.RoleDefinition.ValueInt64()
	req.User = o.User.ValueInt64()
	req.UserAnsibleId = o.UserAnsibleId.ValueString()
	return
}

func (o *roleUserAssignmentTerraformModel) setContentType(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ContentType, data)
}

func (o *roleUserAssignmentTerraformModel) setCreatedBy(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.CreatedBy, data)
}

func (o *roleUserAssignmentTerraformModel) setID(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *roleUserAssignmentTerraformModel) setObjectAnsibleId(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.ObjectAnsibleId, data, false)
}

func (o *roleUserAssignmentTerraformModel) setObjectId(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.ObjectId, data, false)
}

func (o *roleUserAssignmentTerraformModel) setRoleDefinition(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.RoleDefinition, data)
}

func (o *roleUserAssignmentTerraformModel) setUser(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.User, data)
}

func (o *roleUserAssignmentTerraformModel) setUserAnsibleId(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.UserAnsibleId, data, false)
}

func (o *roleUserAssignmentTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
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
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setObjectAnsibleId(data["object_ansible_id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setObjectId(data["object_id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setRoleDefinition(data["role_definition"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setUser(data["user"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setUserAnsibleId(data["user_ansible_id"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// roleUserAssignmentBodyRequestModel maps the schema for RoleUserAssignment for creating and updating the data
type roleUserAssignmentBodyRequestModel struct {
	// ObjectAnsibleId "Resource id of the object this role applies to. Alternative to the object_id field."
	ObjectAnsibleId string `json:"object_ansible_id,omitempty"`
	// ObjectId "Primary key of the object this assignment applies to, null value indicates system-wide assignment"
	ObjectId string `json:"object_id,omitempty"`
	// RoleDefinition "The role definition which defines permissions conveyed by this assignment"
	RoleDefinition int64 `json:"role_definition"`
	// User ""
	User int64 `json:"user,omitempty"`
	// UserAnsibleId "Resource id of the user who will receive permissions from this assignment. Alternative to user field."
	UserAnsibleId string `json:"user_ansible_id,omitempty"`
}
