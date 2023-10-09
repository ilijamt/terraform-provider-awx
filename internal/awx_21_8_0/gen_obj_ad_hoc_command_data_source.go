package awx

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	p "path"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"

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
	_ datasource.DataSource              = &adHocCommandDataSource{}
	_ datasource.DataSourceWithConfigure = &adHocCommandDataSource{}
)

// NewAdHocCommandDataSource is a helper function to instantiate the AdHocCommand data source.
func NewAdHocCommandDataSource() datasource.DataSource {
	return &adHocCommandDataSource{}
}

// adHocCommandDataSource is the data source implementation.
type adHocCommandDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *adHocCommandDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/ad_hoc_commands/"
}

// Metadata returns the data source type name.
func (o *adHocCommandDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ad_hoc_command"
}

// Schema defines the schema for the data source.
func (o *adHocCommandDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"become_enabled": schema.BoolAttribute{
				Description: "Become enabled",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"canceled_on": schema.StringAttribute{
				Description: "The date and time when the cancel request was sent.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"controller_node": schema.StringAttribute{
				Description: "The instance that managed the execution environment.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"credential": schema.Int64Attribute{
				Description: "Credential",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"diff_mode": schema.BoolAttribute{
				Description: "Diff mode",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"elapsed": schema.Float64Attribute{
				Description: "Elapsed time in seconds that the job ran.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Float64{},
			},
			"execution_environment": schema.Int64Attribute{
				Description: "The container image to be used for execution.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"execution_node": schema.StringAttribute{
				Description: "The node the job executed on.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"extra_vars": schema.StringAttribute{
				Description: "Extra vars",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"failed": schema.BoolAttribute{
				Description: "Failed",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"finished": schema.StringAttribute{
				Description: "The date and time the job finished execution.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"forks": schema.Int64Attribute{
				Description: "Forks",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"id": schema.Int64Attribute{
				Description: "Database ID for this ad hoc command.",
				Sensitive:   false,
				Optional:    true,
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.ExactlyOneOf(
						path.MatchRoot("id"),
						path.MatchRoot("name"),
					),
				},
			},
			"inventory": schema.Int64Attribute{
				Description: "Inventory",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"job_explanation": schema.StringAttribute{
				Description: "A status field to indicate the state of the job if it wasn't able to run and capture stdout",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"job_type": schema.StringAttribute{
				Description: "Job type",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"run", "check"}...),
				},
			},
			"launch_type": schema.StringAttribute{
				Description: "Launch type",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"manual", "relaunch", "callback", "scheduled", "dependency", "workflow", "webhook", "sync", "scm"}...),
				},
			},
			"launched_by": schema.Int64Attribute{
				Description: "Launched by",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"limit": schema.StringAttribute{
				Description: "Limit",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"module_args": schema.StringAttribute{
				Description: "Module args",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"module_name": schema.StringAttribute{
				Description: "Module name",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"command", "shell", "yum", "apt", "apt_key", "apt_repository", "apt_rpm", "service", "group", "user", "mount", "ping", "selinux", "setup", "win_ping", "win_service", "win_updates", "win_group", "win_user"}...),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of this ad hoc command.",
				Sensitive:   false,
				Optional:    true,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.ExactlyOneOf(
						path.MatchRoot("id"),
						path.MatchRoot("name"),
					),
				},
			},
			"started": schema.StringAttribute{
				Description: "The date and time the job was queued for starting.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"status": schema.StringAttribute{
				Description: "Status",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"new", "pending", "waiting", "running", "successful", "failed", "error", "canceled"}...),
				},
			},
			"verbosity": schema.StringAttribute{
				Description: "Verbosity",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"0", "1", "2", "3", "4", "5"}...),
				},
			},
			"work_unit_id": schema.StringAttribute{
				Description: "The Receptor work unit ID associated with this job.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			// Write only elements
		},
	}
}

func (o *adHocCommandDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// Read refreshes the Terraform state with the latest data.
func (o *adHocCommandDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state adHocCommandTerraformModel
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

	// Evaluate group 'by_name' based on the schema definition
	var groupByNameExists = func() bool {
		var groupByNameExists = true
		var paramsByName = []any{o.endpoint}
		var attrName types.String
		req.Config.GetAttribute(ctx, path.Root("name"), &attrName)
		groupByNameExists = groupByNameExists && (!attrName.IsNull() && !attrName.IsUnknown())
		paramsByName = append(paramsByName, url.PathEscape(attrName.ValueString()))
		if groupByNameExists {
			endpoint = p.Clean(fmt.Sprintf("%s/?name__exact=%s", paramsByName...))
		}
		return groupByNameExists
	}()
	searchDefined = searchDefined || groupByNameExists

	if !searchDefined {
		var detailMessage string
		resp.Diagnostics.AddError(
			"missing configuration for one of the predefined search groups",
			detailMessage,
		)
		return
	}

	// Creates a new request for AdHocCommand
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for AdHocCommand on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for AdHocCommand
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for AdHocCommand on %s", endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics

	if data, d, err = helpers.ExtractDataIfSearchResult(data); err != nil {
		resp.Diagnostics.Append(d...)
		return
	}

	if d, err = state.updateFromApiData(data); err != nil {
		resp.Diagnostics.Append(d...)
		return
	}

	// Set state
	if err = hooks.RequireResourceStateOrOrig(ctx, ApiVersion, hooks.SourceData, hooks.CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			"Unable to process custom hook for the state on AdHocCommand",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
