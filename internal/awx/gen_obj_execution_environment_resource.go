package awx

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
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
					// Request elements
					"credential": schema.Int64Attribute{
						Description: "Credential",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{},
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this execution environment.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"image": schema.StringAttribute{
						Description:   "The full image location, including the container registry, image name, and version tag.",
						Sensitive:     false,
						Required:      true,
						Optional:      false,
						Computed:      false,
						PlanModifiers: []planmodifier.String{},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(1024),
						},
					},
					"name": schema.StringAttribute{
						Description:   "Name of this execution environment.",
						Sensitive:     false,
						Required:      true,
						Optional:      false,
						Computed:      false,
						PlanModifiers: []planmodifier.String{},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"organization": schema.Int64Attribute{
						Description: "The organization used to determine access to this execution environment.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{},
					},
					"pull": schema.StringAttribute{
						Description: "Pull image before running?",
						Sensitive:   false,
						Required:    false,
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
					// Write only elements
					// Data only elements
					"id": schema.Int64Attribute{
						Description: "Database ID for this execution environment.",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"managed": schema.BoolAttribute{
						Description: "",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:      func(m *executionEnvironmentTerraformModel) any { return m.ID.ValueInt64() },
			ImportStateFunc: executionEnvironmentResourceImportState,
			ApiVersion:      ApiVersion,
			ResourceName:    "ExecutionEnvironment",
		},
	}
}

func executionEnvironmentResourceImportState(_ context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the ExecutionEnvironment.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(context.Background(), path.Root("id"), id)...)
}
