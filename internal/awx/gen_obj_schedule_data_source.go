package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type scheduleDataSource = framework.GenericDataSource[scheduleTerraformModel, *scheduleTerraformModel]

// NewScheduleDataSource is a helper function to instantiate the Schedule data source.
func NewScheduleDataSource() datasource.DataSource {
	return &scheduleDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "schedule", Endpoint: "/api/v2/schedules/"}},
		Cfg: framework.DataSourceCfg[scheduleTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Data only elements
					"description": schema.StringAttribute{
						Description: "Optional description of this schedule.",
						Sensitive:   false,
						Computed:    true,
					},
					"diff_mode": schema.BoolAttribute{
						Description: "Diff mode",
						Sensitive:   false,
						Computed:    true,
					},
					"dtend": schema.StringAttribute{
						Description: "The last occurrence of the schedule occurs before this time, aftewards the schedule expires.",
						Sensitive:   false,
						Computed:    true,
					},
					"dtstart": schema.StringAttribute{
						Description: "The first occurrence of the schedule occurs on or after this time.",
						Sensitive:   false,
						Computed:    true,
					},
					"enabled": schema.BoolAttribute{
						Description: "Enables processing of this schedule.",
						Sensitive:   false,
						Computed:    true,
					},
					"execution_environment": schema.Int64Attribute{
						Description: "The container image to be used for execution.",
						Sensitive:   false,
						Computed:    true,
					},
					"extra_data": schema.StringAttribute{
						Description: "Extra data",
						Sensitive:   false,
						Computed:    true,
					},
					"forks": schema.Int64Attribute{
						Description: "Forks",
						Sensitive:   false,
						Computed:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this schedule.",
						Sensitive:   false,
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
							),
						},
					},
					"inventory": schema.Int64Attribute{
						Description: "Inventory applied as a prompt, assuming job template prompts for inventory",
						Sensitive:   false,
						Computed:    true,
					},
					"job_slice_count": schema.Int64Attribute{
						Description: "Job slice count",
						Sensitive:   false,
						Computed:    true,
					},
					"job_tags": schema.StringAttribute{
						Description: "Job tags",
						Sensitive:   false,
						Computed:    true,
					},
					"job_type": schema.StringAttribute{
						Description: "Job type",
						Sensitive:   false,
						Computed:    true,
					},
					"limit": schema.StringAttribute{
						Description: "Limit",
						Sensitive:   false,
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of this schedule.",
						Sensitive:   false,
						Computed:    true,
					},
					"next_run": schema.StringAttribute{
						Description: "The next time that the scheduled action will run.",
						Sensitive:   false,
						Computed:    true,
					},
					"rrule": schema.StringAttribute{
						Description: "A value representing the schedules iCal recurrence rule.",
						Sensitive:   false,
						Computed:    true,
					},
					"scm_branch": schema.StringAttribute{
						Description: "Scm branch",
						Sensitive:   false,
						Computed:    true,
					},
					"skip_tags": schema.StringAttribute{
						Description: "Skip tags",
						Sensitive:   false,
						Computed:    true,
					},
					"timeout": schema.Int64Attribute{
						Description: "Timeout",
						Sensitive:   false,
						Computed:    true,
					},
					"timezone": schema.StringAttribute{
						Description: "The timezone this schedule runs in. This field is extracted from the RRULE. If the timezone in the RRULE is a link to another timezone, the link will be reflected in this field.",
						Sensitive:   false,
						Computed:    true,
					},
					"unified_job_template": schema.Int64Attribute{
						Description: "Unified job template",
						Sensitive:   false,
						Computed:    true,
					},
					"until": schema.StringAttribute{
						Description: "The date this schedule will end. This field is computed from the RRULE. If the schedule does not end an empty string will be returned",
						Sensitive:   false,
						Computed:    true,
					},
					"verbosity": schema.StringAttribute{
						Description: "Verbosity",
						Sensitive:   false,
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
			ResourceName: "Schedule",
		},
	}
}
