package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type roleTeamAssignmentTerraformModel struct {
	ContentType     types.String `tfsdk:"content_type" json:"content_type"`
	CreatedBy       types.Int64  `tfsdk:"created_by" json:"created_by"`
	ID              types.Int64  `tfsdk:"id" json:"id"`
	ObjectAnsibleId types.String `tfsdk:"object_ansible_id" json:"object_ansible_id"`
	ObjectId        types.String `tfsdk:"object_id" json:"object_id"`
	RoleDefinition  types.Int64  `tfsdk:"role_definition" json:"role_definition"`
	Team            types.Int64  `tfsdk:"team" json:"team"`
	TeamAnsibleId   types.String `tfsdk:"team_ansible_id" json:"team_ansible_id"`
}

func (o *roleTeamAssignmentTerraformModel) Clone() roleTeamAssignmentTerraformModel {
	return *o
}

func (o *roleTeamAssignmentTerraformModel) BodyRequest() *roleTeamAssignmentBodyRequestModel {
	var req roleTeamAssignmentBodyRequestModel
	req.ObjectAnsibleId = o.ObjectAnsibleId.ValueString()
	req.ObjectId = o.ObjectId.ValueString()
	req.RoleDefinition = o.RoleDefinition.ValueInt64()
	req.Team = o.Team.ValueInt64()
	req.TeamAnsibleId = o.TeamAnsibleId.ValueString()
	return &req
}

func (o *roleTeamAssignmentTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.ContentType, data["content_type"], false))
	collect(helpers.AttrValueSetInt64(&o.CreatedBy, data["created_by"]))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.ObjectAnsibleId, data["object_ansible_id"], false))
	collect(helpers.AttrValueSetString(&o.ObjectId, data["object_id"], false))
	collect(helpers.AttrValueSetInt64(&o.RoleDefinition, data["role_definition"]))
	collect(helpers.AttrValueSetInt64(&o.Team, data["team"]))
	collect(helpers.AttrValueSetString(&o.TeamAnsibleId, data["team_ansible_id"], false))
	return diags, nil
}

type roleTeamAssignmentBodyRequestModel struct {
	ObjectAnsibleId string `json:"object_ansible_id,omitempty"`
	ObjectId        string `json:"object_id,omitempty"`
	RoleDefinition  int64  `json:"role_definition"`
	Team            int64  `json:"team,omitempty"`
	TeamAnsibleId   string `json:"team_ansible_id,omitempty"`
}

type roleTeamAssignmentResource = framework.GenericResource[roleTeamAssignmentTerraformModel, roleTeamAssignmentBodyRequestModel, *roleTeamAssignmentTerraformModel]

// NewRoleTeamAssignmentResource is a helper function to simplify the provider implementation.
func NewRoleTeamAssignmentResource() resource.Resource {
	return &roleTeamAssignmentResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "role_team_assignment", Endpoint: "/api/v2/role_team_assignments/"}},
		Cfg: framework.ResourceCfg[roleTeamAssignmentTerraformModel, roleTeamAssignmentBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"object_ansible_id": schema.StringAttribute{
						Description: "Resource id of the object this role applies to. Alternative to the object_id field.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
							stringplanmodifier.RequiresReplace(),
						},
					},
					"object_id": schema.StringAttribute{
						Description: "Primary key of the object this assignment applies to, null value indicates system-wide assignment",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
							stringplanmodifier.RequiresReplace(),
						},
					},
					"role_definition": schema.Int64Attribute{
						Description: "The role definition which defines permissions conveyed by this assignment",
						Required:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplace(),
						},
					},
					"team": schema.Int64Attribute{
						Description: "Team",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
							int64planmodifier.RequiresReplace(),
						},
					},
					"team_ansible_id": schema.StringAttribute{
						Description: "Resource id of the team who will receive permissions from this assignment. Alternative to team field.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
							stringplanmodifier.RequiresReplace(),
						},
					},
					"content_type": schema.StringAttribute{
						Description: "The type of resource this applies to",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"awx.credential",
								"awx.executionenvironment",
								"awx.instancegroup",
								"awx.inventory",
								"awx.jobtemplate",
								"awx.notificationtemplate",
								"awx.project",
								"awx.workflowjobtemplate",
								"shared.organization",
								"shared.team",
							),
						},
					},
					"created_by": schema.Int64Attribute{
						Description: "The user who created this resource",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this role team assignment.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *roleTeamAssignmentTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			ApiVersion:   ApiVersion,
			ResourceName: "RoleTeamAssignment",
		},
	}
}

type roleTeamAssignmentDataSource = framework.GenericDataSource[roleTeamAssignmentTerraformModel, *roleTeamAssignmentTerraformModel]

// NewRoleTeamAssignmentDataSource is a helper function to instantiate the RoleTeamAssignment data source.
func NewRoleTeamAssignmentDataSource() datasource.DataSource {
	return &roleTeamAssignmentDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "role_team_assignment", Endpoint: "/api/v2/role_team_assignments/"}},
		Cfg: framework.DataSourceCfg[roleTeamAssignmentTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"content_type": dschema.StringAttribute{
						Description: "The type of resource this applies to",
						Computed:    true,
					},
					"created_by": dschema.Int64Attribute{
						Description: "The user who created this resource",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this role team assignment.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
							),
						},
					},
					"object_ansible_id": dschema.StringAttribute{
						Description: "Resource id of the object this role applies to. Alternative to the object_id field.",
						Computed:    true,
					},
					"object_id": dschema.StringAttribute{
						Description: "Primary key of the object this assignment applies to, null value indicates system-wide assignment",
						Computed:    true,
					},
					"role_definition": dschema.Int64Attribute{
						Description: "The role definition which defines permissions conveyed by this assignment",
						Computed:    true,
					},
					"team": dschema.Int64Attribute{
						Description: "Team",
						Computed:    true,
					},
					"team_ansible_id": dschema.StringAttribute{
						Description: "Resource id of the team who will receive permissions from this assignment. Alternative to team field.",
						Computed:    true,
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "RoleTeamAssignment",
		},
	}
}
