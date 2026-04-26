package awx

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type teamResource = framework.GenericResource[teamTerraformModel, teamBodyRequestModel, *teamTerraformModel]

// NewTeamResource is a helper function to simplify the provider implementation.
func NewTeamResource() resource.Resource {
	return &teamResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "team", Endpoint: "/api/v2/teams/"}},
		Cfg: framework.ResourceCfg[teamTerraformModel, teamBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description: "Optional description of this team.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this team.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"organization": schema.Int64Attribute{
						Description: "Organization",
						Required:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this team.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *teamTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			ApiVersion:   ApiVersion,
			ResourceName: "Team",
		},
	}
}
