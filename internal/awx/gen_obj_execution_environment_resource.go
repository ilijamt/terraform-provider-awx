package awx

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type executionEnvironmentResource = framework.GenericResource[executionEnvironmentTerraformModel, executionEnvironmentBodyRequestModel, *executionEnvironmentTerraformModel]

// NewExecutionEnvironmentResource is a helper function to simplify the provider implementation.
func NewExecutionEnvironmentResource() resource.Resource {
	return &executionEnvironmentResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "execution_environment", Endpoint: "/api/v2/execution_environments/"}},
		Cfg: framework.ResourceCfg[executionEnvironmentTerraformModel, executionEnvironmentBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"credential": schema.Int64Attribute{
						Description: "Credential",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this execution environment.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"image": schema.StringAttribute{
						Description: "The full image location, including the container registry, image name, and version tag.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(1024),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this execution environment.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"organization": schema.Int64Attribute{
						Description: "The organization used to determine access to this execution environment.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"pull": schema.StringAttribute{
						Description: "Pull image before running?",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"",
								"always",
								"missing",
								"never",
							),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this execution environment.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
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
			IDAccessor:   func(m *executionEnvironmentTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			ApiVersion:   ApiVersion,
			ResourceName: "ExecutionEnvironment",
		},
	}
}
