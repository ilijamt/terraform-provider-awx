package awx

import (
	"context"
	"fmt"
	"net/http"
	p "path"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"diff_mode": schema.BoolAttribute{
				Description: "Diff mode",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"dtend": schema.StringAttribute{
				Description: "The last occurrence of the schedule occurs before this time, aftewards the schedule expires.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"dtstart": schema.StringAttribute{
				Description: "The first occurrence of the schedule occurs on or after this time.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"enabled": schema.BoolAttribute{
				Description: "Enables processing of this schedule.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"execution_environment": schema.Int64Attribute{
				Description: "The container image to be used for execution.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"extra_data": schema.StringAttribute{
				Description: "Extra data",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"forks": schema.Int64Attribute{
				Description: "Forks",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
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
				Validators:  []validator.Int64{},
			},
			"job_slice_count": schema.Int64Attribute{
				Description: "Job slice count",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"job_tags": schema.StringAttribute{
				Description: "Job tags",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"job_type": schema.StringAttribute{
				Description: "Job type",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"", "run", "check"}...),
				},
			},
			"limit": schema.StringAttribute{
				Description: "Limit",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"name": schema.StringAttribute{
				Description: "Name of this schedule.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"next_run": schema.StringAttribute{
				Description: "The next time that the scheduled action will run.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"rrule": schema.StringAttribute{
				Description: "A value representing the schedules iCal recurrence rule.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"scm_branch": schema.StringAttribute{
				Description: "Scm branch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"skip_tags": schema.StringAttribute{
				Description: "Skip tags",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"timeout": schema.Int64Attribute{
				Description: "Timeout",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"timezone": schema.StringAttribute{
				Description: "The timezone this schedule runs in. This field is extracted from the RRULE. If the timezone in the RRULE is a link to another timezone, the link will be reflected in this field.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"unified_job_template": schema.Int64Attribute{
				Description: "Unified job template",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"until": schema.StringAttribute{
				Description: "The date this schedule will end. This field is computed from the RRULE. If the schedule does not end an emptry string will be returned",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"verbosity": schema.StringAttribute{
				Description: "Verbosity",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"0", "1", "2", "3", "4", "5"}...),
				},
			},
			// Write only elements
		},
	}
}

func (o *scheduleDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
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
