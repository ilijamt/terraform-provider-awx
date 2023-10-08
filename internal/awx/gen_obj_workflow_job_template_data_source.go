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
	_ datasource.DataSource              = &workflowJobTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &workflowJobTemplateDataSource{}
)

// NewWorkflowJobTemplateDataSource is a helper function to instantiate the WorkflowJobTemplate data source.
func NewWorkflowJobTemplateDataSource() datasource.DataSource {
	return &workflowJobTemplateDataSource{}
}

// workflowJobTemplateDataSource is the data source implementation.
type workflowJobTemplateDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *workflowJobTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/workflow_job_templates/"
}

// Metadata returns the data source type name.
func (o *workflowJobTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workflow_job_template"
}

// Schema defines the schema for the data source.
func (o *workflowJobTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"allow_simultaneous": schema.BoolAttribute{
				Description: "Allow simultaneous",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_inventory_on_launch": schema.BoolAttribute{
				Description: "Ask inventory on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_labels_on_launch": schema.BoolAttribute{
				Description: "Ask labels on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_limit_on_launch": schema.BoolAttribute{
				Description: "Ask limit on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_scm_branch_on_launch": schema.BoolAttribute{
				Description: "Ask scm branch on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_skip_tags_on_launch": schema.BoolAttribute{
				Description: "Ask skip tags on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_tags_on_launch": schema.BoolAttribute{
				Description: "Ask tags on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"ask_variables_on_launch": schema.BoolAttribute{
				Description: "Ask variables on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"description": schema.StringAttribute{
				Description: "Optional description of this workflow job template.",
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
			"id": schema.Int64Attribute{
				Description: "Database ID for this workflow job template.",
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
				Description: "Inventory applied as a prompt, assuming job template prompts for inventory",
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
			"limit": schema.StringAttribute{
				Description: "Limit",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"name": schema.StringAttribute{
				Description: "Name of this workflow job template.",
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
			"organization": schema.Int64Attribute{
				Description: "The organization used to determine access to this template.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
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
			"survey_enabled": schema.BoolAttribute{
				Description: "Survey enabled",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"webhook_credential": schema.Int64Attribute{
				Description: "Personal Access Token for posting back the status to the service API",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"webhook_service": schema.StringAttribute{
				Description: "Service that webhook requests will be accepted from",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"", "github", "gitlab"}...),
				},
			},
			// Write only elements
		},
	}
}

func (o *workflowJobTemplateDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// Read refreshes the Terraform state with the latest data.
func (o *workflowJobTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state workflowJobTemplateTerraformModel
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

	// Creates a new request for WorkflowJobTemplate
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for WorkflowJobTemplate on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for WorkflowJobTemplate
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for WorkflowJobTemplate on %s", o.endpoint),
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
			"Unable to process custom hook for the state on WorkflowJobTemplate",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
