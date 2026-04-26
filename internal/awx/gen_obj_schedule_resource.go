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

type scheduleResource = framework.GenericResource[scheduleTerraformModel, scheduleBodyRequestModel, *scheduleTerraformModel]

// NewScheduleResource is a helper function to simplify the provider implementation.
func NewScheduleResource() resource.Resource {
	return &scheduleResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "schedule", Endpoint: "/api/v2/schedules/"}},
		Cfg: framework.ResourceCfg[scheduleTerraformModel, scheduleBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description: "Optional description of this schedule.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
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
					"enabled": schema.BoolAttribute{
						Description: "Enables processing of this schedule.",
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
					"extra_data": schema.StringAttribute{
						Description: "Extra data",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"forks": schema.Int64Attribute{
						Description: "Forks",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"inventory": schema.Int64Attribute{
						Description: "Inventory applied as a prompt, assuming job template prompts for inventory",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"job_slice_count": schema.Int64Attribute{
						Description: "Job slice count",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"job_tags": schema.StringAttribute{
						Description: "Job tags",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"job_type": schema.StringAttribute{
						Description: "Job type",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"",
								"run",
								"check",
							),
						},
					},
					"limit": schema.StringAttribute{
						Description: "Limit",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this schedule.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"rrule": schema.StringAttribute{
						Description: "A value representing the schedules iCal recurrence rule.",
						Required:    true,
					},
					"scm_branch": schema.StringAttribute{
						Description: "Scm branch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"skip_tags": schema.StringAttribute{
						Description: "Skip tags",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"timeout": schema.Int64Attribute{
						Description: "Timeout",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"unified_job_template": schema.Int64Attribute{
						Description: "Unified job template",
						Required:    true,
					},
					"verbosity": schema.StringAttribute{
						Description: "Verbosity",
						Optional:    true,
						Computed:    true,
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
					"dtend": schema.StringAttribute{
						Description: "The last occurrence of the schedule occurs before this time, aftewards the schedule expires.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"dtstart": schema.StringAttribute{
						Description: "The first occurrence of the schedule occurs on or after this time.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this schedule.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"next_run": schema.StringAttribute{
						Description: "The next time that the scheduled action will run.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"timezone": schema.StringAttribute{
						Description: "The timezone this schedule runs in. This field is extracted from the RRULE. If the timezone in the RRULE is a link to another timezone, the link will be reflected in this field.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"until": schema.StringAttribute{
						Description: "The date this schedule will end. This field is computed from the RRULE. If the schedule does not end an empty string will be returned",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *scheduleTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			ApiVersion:   ApiVersion,
			ResourceName: "Schedule",
		},
	}
}
