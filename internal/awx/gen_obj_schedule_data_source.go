package awx

import (
	"context"
	"fmt"
	"net/http"
	p "path"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &scheduleDataSource{}
	_ datasource.DataSourceWithConfigure = &scheduleDataSource{}
)

// NewScheduleDataSource is a helper function to instantiate the Schedule data source.
func NewScheduleDataSource() datasource.DataSource {
	return &scheduleDataSource{}
}

// scheduleDataSource is the data source implementation.
type scheduleDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *scheduleDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/schedules/"
}

// Metadata returns the data source type name.
func (o *scheduleDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_schedule"
}

// Schema defines the schema for the data source.
func (o *scheduleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
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
				Description: "The date this schedule will end. This field is computed from the RRULE. If the schedule does not end an emptry string will be returned",
				Computed:    true,
			},
			"verbosity": schema.StringAttribute{
				Description: "Verbosity",
				Computed:    true,
			},
			// Write only elements
		},
	}
}

func (o *scheduleDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
		),
	}
}

// Read refreshes the Terraform state with the latest data.
func (o *scheduleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state scheduleTerraformModel
	var err error
	var endpoint string
	var searchDefined bool

	// Only one group should evaluate to True, terraform should prevent from being able to set
	// the conflicting groups

	// Evaluate group 'by_id' based on the schema definition
	var groupByIdExists = func() bool {
		var groupByIdExists = true
		var paramsById = []any{o.endpoint}
		var attrID types.Int64
		req.Config.GetAttribute(ctx, path.Root("id"), &attrID)
		groupByIdExists = groupByIdExists && (!attrID.IsNull() && !attrID.IsUnknown())
		paramsById = append(paramsById, attrID.ValueInt64())
		if groupByIdExists {
			endpoint = p.Clean(fmt.Sprintf("%s/%d/", paramsById...))
		}
		return groupByIdExists
	}()
	searchDefined = searchDefined || groupByIdExists

	if !searchDefined {
		var detailMessage string
		resp.Diagnostics.AddError(
			"missing configuration for one of the predefined search groups",
			detailMessage,
		)
		return
	}

	// Creates a new request for Schedule
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Schedule on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Schedule
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Schedule on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics

	if data, d, err = extractDataIfSearchResult(data); err != nil {
		resp.Diagnostics.Append(d...)
		return
	}

	if d, err = state.updateFromApiData(data); err != nil {
		resp.Diagnostics.Append(d...)
		return
	}

	// Set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
