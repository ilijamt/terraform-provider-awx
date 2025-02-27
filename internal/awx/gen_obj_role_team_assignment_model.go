package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// roleTeamAssignmentTerraformModel maps the schema for RoleTeamAssignment when using Data Source
type roleTeamAssignmentTerraformModel struct {
	// ContentType "The type of resource this applies to"
	ContentType types.Int64 `tfsdk:"content_type" json:"content_type"`
	// CreatedBy "The user who created this resource"
	CreatedBy types.Int64 `tfsdk:"created_by" json:"created_by"`
	// ID "Database ID for this role team assignment."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// ObjectAnsibleId "Resource id of the object this role applies to. Alternative to the object_id field."
	ObjectAnsibleId types.String `tfsdk:"object_ansible_id" json:"object_ansible_id"`
	// ObjectId "Primary key of the object this assignment applies to, null value indicates system-wide assignment"
	ObjectId types.String `tfsdk:"object_id" json:"object_id"`
	// RoleDefinition "The role definition which defines permissions conveyed by this assignment"
	RoleDefinition types.Int64 `tfsdk:"role_definition" json:"role_definition"`
	// Team ""
	Team types.Int64 `tfsdk:"team" json:"team"`
	// TeamAnsibleId "Resource id of the team who will receive permissions from this assignment. Alternative to team field."
	TeamAnsibleId types.String `tfsdk:"team_ansible_id" json:"team_ansible_id"`
}

// Clone the object
func (o *roleTeamAssignmentTerraformModel) Clone() roleTeamAssignmentTerraformModel {
	return roleTeamAssignmentTerraformModel{
		ContentType:     o.ContentType,
		CreatedBy:       o.CreatedBy,
		ID:              o.ID,
		ObjectAnsibleId: o.ObjectAnsibleId,
		ObjectId:        o.ObjectId,
		RoleDefinition:  o.RoleDefinition,
		Team:            o.Team,
		TeamAnsibleId:   o.TeamAnsibleId,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for RoleTeamAssignment
func (o *roleTeamAssignmentTerraformModel) BodyRequest() (req roleTeamAssignmentBodyRequestModel) {
	req.ObjectAnsibleId = o.ObjectAnsibleId.ValueString()
	req.ObjectId = o.ObjectId.ValueString()
	req.RoleDefinition = o.RoleDefinition.ValueInt64()
	req.Team = o.Team.ValueInt64()
	req.TeamAnsibleId = o.TeamAnsibleId.ValueString()
	return
}

func (o *roleTeamAssignmentTerraformModel) setContentType(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ContentType, data)
}

func (o *roleTeamAssignmentTerraformModel) setCreatedBy(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.CreatedBy, data)
}

func (o *roleTeamAssignmentTerraformModel) setID(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *roleTeamAssignmentTerraformModel) setObjectAnsibleId(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.ObjectAnsibleId, data, false)
}

func (o *roleTeamAssignmentTerraformModel) setObjectId(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.ObjectId, data, false)
}

func (o *roleTeamAssignmentTerraformModel) setRoleDefinition(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.RoleDefinition, data)
}

func (o *roleTeamAssignmentTerraformModel) setTeam(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Team, data)
}

func (o *roleTeamAssignmentTerraformModel) setTeamAnsibleId(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.TeamAnsibleId, data, false)
}

func (o *roleTeamAssignmentTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
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
	if dg, _ := o.setTeam(data["team"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setTeamAnsibleId(data["team_ansible_id"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// roleTeamAssignmentBodyRequestModel maps the schema for RoleTeamAssignment for creating and updating the data
type roleTeamAssignmentBodyRequestModel struct {
	// ObjectAnsibleId "Resource id of the object this role applies to. Alternative to the object_id field."
	ObjectAnsibleId string `json:"object_ansible_id,omitempty"`
	// ObjectId "Primary key of the object this assignment applies to, null value indicates system-wide assignment"
	ObjectId string `json:"object_id,omitempty"`
	// RoleDefinition "The role definition which defines permissions conveyed by this assignment"
	RoleDefinition int64 `json:"role_definition"`
	// Team ""
	Team int64 `json:"team,omitempty"`
	// TeamAnsibleId "Resource id of the team who will receive permissions from this assignment. Alternative to team field."
	TeamAnsibleId string `json:"team_ansible_id,omitempty"`
}
