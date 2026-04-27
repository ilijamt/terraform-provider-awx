package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type credentialTerraformModel struct {
	Cloud          types.Bool   `tfsdk:"cloud" json:"cloud"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Description    types.String `tfsdk:"description" json:"description"`
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Inputs         types.String `tfsdk:"inputs" json:"inputs"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Kubernetes     types.Bool   `tfsdk:"kubernetes" json:"kubernetes"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	Name           types.String `tfsdk:"name" json:"name"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
}

func (o *credentialTerraformModel) Clone() credentialTerraformModel {
	return *o
}

func (o *credentialTerraformModel) BodyRequest() *credentialBodyRequestModel {
	var req credentialBodyRequestModel
	req.CredentialType = o.CredentialType.ValueInt64()
	req.Description = o.Description.ValueString()
	req.Inputs = json.RawMessage(o.Inputs.ValueString())
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	return &req
}

func (o *credentialTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetBool(&o.Cloud, data["cloud"]))
	collect(helpers.AttrValueSetInt64(&o.CredentialType, data["credential_type"]))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetJsonString(&o.Inputs, data["inputs"], false))
	collect(helpers.AttrValueSetString(&o.Kind, data["kind"], false))
	collect(helpers.AttrValueSetBool(&o.Kubernetes, data["kubernetes"]))
	collect(helpers.AttrValueSetBool(&o.Managed, data["managed"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	return diags, nil
}

type credentialBodyRequestModel struct {
	CredentialType int64           `json:"credential_type"`
	Description    string          `json:"description,omitempty"`
	Inputs         json.RawMessage `json:"inputs,omitempty"`
	Name           string          `json:"name"`
	Organization   int64           `json:"organization,omitempty"`
	Team           int64           `json:"team,omitempty"`
	User           int64           `json:"user,omitempty"`
}

type credentialResource = framework.GenericResource[credentialTerraformModel, credentialBodyRequestModel, *credentialTerraformModel]

// NewCredentialResource is a helper function to simplify the provider implementation.
func NewCredentialResource() resource.Resource {
	return &credentialResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.ResourceCfg[credentialTerraformModel, credentialBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"credential_type": schema.Int64Attribute{
						Description: "Specify the type of credential you want to create. Refer to the documentation for details on each type.",
						Required:    true,
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this credential.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"inputs": schema.StringAttribute{
						Description: "Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this credential.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"organization": schema.Int64Attribute{
						Description: "Inherit permissions from organization roles. If provided on creation, do not give either user or team.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{
							// exactly_one_of_org_user_team
							int64validator.ExactlyOneOf(path.MatchRoot("organization"), path.MatchRoot("team"), path.MatchRoot("user")),
						},
					},
					"team": schema.Int64Attribute{
						Description: "Write-only field used to add team to owner role. If provided, do not give either user or organization. Only valid for creation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{
							// exactly_one_of_org_user_team
							int64validator.ExactlyOneOf(path.MatchRoot("organization"), path.MatchRoot("team"), path.MatchRoot("user")),
						},
					},
					"user": schema.Int64Attribute{
						Description: "Write-only field used to add user to owner role. If provided, do not give either team or organization. Only valid for creation.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{
							// exactly_one_of_org_user_team
							int64validator.ExactlyOneOf(path.MatchRoot("organization"), path.MatchRoot("team"), path.MatchRoot("user")),
						},
					},
					"cloud": schema.BoolAttribute{
						Description: "Cloud",
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this credential.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"kind": schema.StringAttribute{
						Description: "Kind",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"kubernetes": schema.BoolAttribute{
						Description: "Kubernetes",
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"managed": schema.BoolAttribute{
						Description: "Managed",
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor: func(m *credentialTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:      "id",
			Hook:       hookCredential,
			WriteOnlyPlanToBody: func(plan *credentialTerraformModel, body *credentialBodyRequestModel) {
				body.Team = plan.Team.ValueInt64()
				body.User = plan.User.ValueInt64()
			},
			WriteOnlyPlanToState: func(plan, state *credentialTerraformModel) {
				state.Team = types.Int64Value(plan.Team.ValueInt64())
				state.User = types.Int64Value(plan.User.ValueInt64())
			},
			ApiVersion:   ApiVersion,
			ResourceName: "Credential",
		},
	}
}

type credentialDataSource = framework.GenericDataSource[credentialTerraformModel, *credentialTerraformModel]

// NewCredentialDataSource is a helper function to instantiate the Credential data source.
func NewCredentialDataSource() datasource.DataSource {
	return &credentialDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"cloud": dschema.BoolAttribute{
						Description: "Cloud",
						Computed:    true,
					},
					"credential_type": dschema.Int64Attribute{
						Description: "Specify the type of credential you want to create. Refer to the documentation for details on each type.",
						Computed:    true,
					},
					"description": dschema.StringAttribute{
						Description: "Optional description of this credential.",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this credential.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"inputs": dschema.StringAttribute{
						Description: "Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax.",
						Computed:    true,
					},
					"kind": dschema.StringAttribute{
						Description: "Kind",
						Computed:    true,
					},
					"kubernetes": dschema.BoolAttribute{
						Description: "Kubernetes",
						Computed:    true,
					},
					"managed": dschema.BoolAttribute{
						Description: "Managed",
						Computed:    true,
					},
					"name": dschema.StringAttribute{
						Description: "Name of this credential.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"organization": dschema.Int64Attribute{
						Description: "Inherit permissions from organization roles. If provided on creation, do not give either user or team.",
						Computed:    true,
					},
					"team": dschema.Int64Attribute{
						Description: "Write-only field used to add team to owner role. If provided, do not give either user or organization. Only valid for creation.",
						Optional:    true,
						Computed:    true,
					},
					"user": dschema.Int64Attribute{
						Description: "Write-only field used to add user to owner role. If provided, do not give either team or organization. Only valid for creation.",
						Optional:    true,
						Computed:    true,
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
				{Name: "by_name", URLSuffix: "/?name__exact=%s", Fields: []framework.SearchField{
					{Name: "name", Type: "string", URLEscape: true},
				}},
			},
			Hook:         hookCredential,
			ApiVersion:   ApiVersion,
			ResourceName: "Credential",
		},
	}
}
