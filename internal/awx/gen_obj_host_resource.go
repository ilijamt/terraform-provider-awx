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

type hostResource = framework.GenericResource[hostTerraformModel, hostBodyRequestModel, *hostTerraformModel]

// NewHostResource is a helper function to simplify the provider implementation.
func NewHostResource() resource.Resource {
	return &hostResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "host", Endpoint: "/api/v2/hosts/"}},
		Cfg: framework.ResourceCfg[hostTerraformModel, hostBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Request elements
					"description": schema.StringAttribute{
						Description: "Optional description of this host.",
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
					"enabled": schema.BoolAttribute{
						Description: "Is this host online and available for running jobs?",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					"instance_id": schema.StringAttribute{
						Description: "The value used by the remote inventory source to uniquely identify the host",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(1024),
						},
					},
					"inventory": schema.Int64Attribute{
						Description:   "Inventory",
						Sensitive:     false,
						Required:      true,
						Optional:      false,
						Computed:      false,
						PlanModifiers: []planmodifier.Int64{},
						Validators:    []validator.Int64{},
					},
					"name": schema.StringAttribute{
						Description:   "Name of this host.",
						Sensitive:     false,
						Required:      true,
						Optional:      false,
						Computed:      false,
						PlanModifiers: []planmodifier.String{},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"variables": schema.StringAttribute{
						Description: "Host variables in JSON or YAML format.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					// Write only elements
					// Data only elements
					"id": schema.Int64Attribute{
						Description: "Database ID for this host.",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"last_job": schema.Int64Attribute{
						Description: "",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"last_job_host_summary": schema.Int64Attribute{
						Description: "",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:      func(m *hostTerraformModel) any { return m.ID.ValueInt64() },
			ImportStateFunc: hostResourceImportState,
			ApiVersion:      ApiVersion,
			ResourceName:    "Host",
		},
	}
}

func hostResourceImportState(_ context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the Host.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(context.Background(), path.Root("id"), id)...)
}
