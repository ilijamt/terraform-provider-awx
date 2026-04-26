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
					"description": schema.StringAttribute{
						Description: "Optional description of this schedule.",
						Computed:    true,
					},
					"diff_mode": schema.BoolAttribute{
						Description: "Diff mode",
						Computed:    true,
					},
					"dtend": schema.StringAttribute{
						Description: "The last occurrence of the schedule occurs before this time, aftewards the schedule expires.",
						Computed:    true,
					},
					"dtstart": schema.StringAttribute{
						Description: "The first occurrence of the schedule occurs on or after this time.",
						Computed:    true,
					},
					"enabled": schema.BoolAttribute{
						Description: "Enables processing of this schedule.",
						Computed:    true,
					},
					"execution_environment": schema.Int64Attribute{
						Description: "The container image to be used for execution.",
						Computed:    true,
					},
					"extra_data": schema.StringAttribute{
						Description: "Extra data",
						Computed:    true,
					},
					"forks": schema.Int64Attribute{
						Description: "Forks",
						Computed:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this schedule.",
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
						Computed:    true,
					},
					"job_slice_count": schema.Int64Attribute{
						Description: "Job slice count",
						Computed:    true,
					},
					"job_tags": schema.StringAttribute{
						Description: "Job tags",
						Computed:    true,
					},
					"job_type": schema.StringAttribute{
						Description: "Job type",
						Computed:    true,
					},
					"limit": schema.StringAttribute{
						Description: "Limit",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of this schedule.",
						Computed:    true,
					},
					"next_run": schema.StringAttribute{
						Description: "The next time that the scheduled action will run.",
						Computed:    true,
					},
					"rrule": schema.StringAttribute{
						Description: "A value representing the schedules iCal recurrence rule.",
						Computed:    true,
					},
					"scm_branch": schema.StringAttribute{
						Description: "Scm branch",
						Computed:    true,
					},
					"skip_tags": schema.StringAttribute{
						Description: "Skip tags",
						Computed:    true,
					},
					"timeout": schema.Int64Attribute{
						Description: "Timeout",
						Computed:    true,
					},
					"timezone": schema.StringAttribute{
						Description: "The timezone this schedule runs in. This field is extracted from the RRULE. If the timezone in the RRULE is a link to another timezone, the link will be reflected in this field.",
						Computed:    true,
					},
					"unified_job_template": schema.Int64Attribute{
						Description: "Unified job template",
						Computed:    true,
					},
					"until": schema.StringAttribute{
						Description: "The date this schedule will end. This field is computed from the RRULE. If the schedule does not end an empty string will be returned",
						Computed:    true,
					},
					"verbosity": schema.StringAttribute{
						Description: "Verbosity",
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
