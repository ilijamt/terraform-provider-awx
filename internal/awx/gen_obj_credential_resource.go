package awx

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

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
