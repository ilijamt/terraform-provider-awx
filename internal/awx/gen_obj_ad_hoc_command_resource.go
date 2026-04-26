package awx

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type adHocCommandResource = framework.GenericResource[adHocCommandTerraformModel, adHocCommandBodyRequestModel, *adHocCommandTerraformModel]

// NewAdHocCommandResource is a helper function to simplify the provider implementation.
func NewAdHocCommandResource() resource.Resource {
	return &adHocCommandResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "ad_hoc_command", Endpoint: "/api/v2/ad_hoc_commands/"}},
		Cfg: framework.ResourceCfg[adHocCommandTerraformModel, adHocCommandBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"become_enabled": schema.BoolAttribute{
						Description: "Become enabled",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"credential": schema.Int64Attribute{
						Description: "Credential",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"diff_mode": schema.BoolAttribute{
						Description: "Diff mode",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"execution_environment": schema.Int64Attribute{
						Description: "The container image to be used for execution.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"extra_vars": schema.StringAttribute{
						Description: "Extra vars",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"forks": schema.Int64Attribute{
						Description: "Forks",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{
							int64validator.Between(0, 2147483647),
						},
					},
					"inventory": schema.Int64Attribute{
						Description: "Inventory",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"job_type": schema.StringAttribute{
						Description: "Job type",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`run`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"run",
								"check",
							),
						},
					},
					"limit": schema.StringAttribute{
						Description: "Limit",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"module_args": schema.StringAttribute{
						Description: "Module args",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"module_name": schema.StringAttribute{
						Description: "Module name",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`command`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"command",
								"shell",
								"yum",
								"apt",
								"apt_key",
								"apt_repository",
								"apt_rpm",
								"service",
								"group",
								"user",
								"mount",
								"ping",
								"selinux",
								"setup",
								"win_ping",
								"win_service",
								"win_updates",
								"win_group",
								"win_user",
							),
						},
					},
					"verbosity": schema.StringAttribute{
						Description: "Verbosity",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`0`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"0",
								"1",
								"2",
								"3",
								"4",
								"5",
							),
						},
					},
					"canceled_on": schema.StringAttribute{
						Description: "The date and time when the cancel request was sent.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"controller_node": schema.StringAttribute{
						Description: "The instance that managed the execution environment.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"elapsed": schema.Float64Attribute{
						Description: "Elapsed time in seconds that the job ran.",
						Computed:    true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
					"execution_node": schema.StringAttribute{
						Description: "The node the job executed on.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"failed": schema.BoolAttribute{
						Description: "Failed",
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"finished": schema.StringAttribute{
						Description: "The date and time the job finished execution.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this ad hoc command.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"job_explanation": schema.StringAttribute{
						Description: "A status field to indicate the state of the job if it wasn't able to run and capture stdout",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"launch_type": schema.StringAttribute{
						Description: "Launch type",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"manual",
								"relaunch",
								"callback",
								"scheduled",
								"dependency",
								"workflow",
								"webhook",
								"sync",
								"scm",
							),
						},
					},
					"launched_by": schema.Int64Attribute{
						Description: "Launched by",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this ad hoc command.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"started": schema.StringAttribute{
						Description: "The date and time the job was queued for starting.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"status": schema.StringAttribute{
						Description: "Status",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"new",
								"pending",
								"waiting",
								"running",
								"successful",
								"failed",
								"error",
								"canceled",
							),
						},
					},
					"work_unit_id": schema.StringAttribute{
						Description: "The Receptor work unit ID associated with this job.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor: func(m *adHocCommandTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:      "id",
			Hook: func(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *adHocCommandTerraformModel) error {
				return hooks.RequireResourceStateOrOrig(ctx, apiVersion, source, callee, orig, state)
			},
			ApiVersion:   ApiVersion,
			ResourceName: "AdHocCommand",
		},
	}
}
