package awx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	p "path"
	"strconv"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/mitchellh/mapstructure"

	"github.com/hashicorp/terraform-plugin-framework-validators/schemavalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// workflowJobTemplateTerraformModel maps the schema for WorkflowJobTemplate when using Data Source
type workflowJobTemplateTerraformModel struct {
	// AllowSimultaneous ""
	AllowSimultaneous types.Bool `tfsdk:"allow_simultaneous" json:"allow_simultaneous"`
	// AskInventoryOnLaunch ""
	AskInventoryOnLaunch types.Bool `tfsdk:"ask_inventory_on_launch" json:"ask_inventory_on_launch"`
	// AskLabelsOnLaunch ""
	AskLabelsOnLaunch types.Bool `tfsdk:"ask_labels_on_launch" json:"ask_labels_on_launch"`
	// AskLimitOnLaunch ""
	AskLimitOnLaunch types.Bool `tfsdk:"ask_limit_on_launch" json:"ask_limit_on_launch"`
	// AskScmBranchOnLaunch ""
	AskScmBranchOnLaunch types.Bool `tfsdk:"ask_scm_branch_on_launch" json:"ask_scm_branch_on_launch"`
	// AskSkipTagsOnLaunch ""
	AskSkipTagsOnLaunch types.Bool `tfsdk:"ask_skip_tags_on_launch" json:"ask_skip_tags_on_launch"`
	// AskTagsOnLaunch ""
	AskTagsOnLaunch types.Bool `tfsdk:"ask_tags_on_launch" json:"ask_tags_on_launch"`
	// AskVariablesOnLaunch ""
	AskVariablesOnLaunch types.Bool `tfsdk:"ask_variables_on_launch" json:"ask_variables_on_launch"`
	// Description "Optional description of this workflow job template."
	Description types.String `tfsdk:"description" json:"description"`
	// ExtraVars ""
	ExtraVars types.String `tfsdk:"extra_vars" json:"extra_vars"`
	// ID "Database ID for this workflow job template."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// Inventory "Inventory applied as a prompt, assuming job template prompts for inventory"
	Inventory types.Int64 `tfsdk:"inventory" json:"inventory"`
	// JobTags ""
	JobTags types.String `tfsdk:"job_tags" json:"job_tags"`
	// Limit ""
	Limit types.String `tfsdk:"limit" json:"limit"`
	// Name "Name of this workflow job template."
	Name types.String `tfsdk:"name" json:"name"`
	// Organization "The organization used to determine access to this template."
	Organization types.Int64 `tfsdk:"organization" json:"organization"`
	// ScmBranch ""
	ScmBranch types.String `tfsdk:"scm_branch" json:"scm_branch"`
	// SkipTags ""
	SkipTags types.String `tfsdk:"skip_tags" json:"skip_tags"`
	// SurveyEnabled ""
	SurveyEnabled types.Bool `tfsdk:"survey_enabled" json:"survey_enabled"`
	// WebhookCredential "Personal Access Token for posting back the status to the service API"
	WebhookCredential types.Int64 `tfsdk:"webhook_credential" json:"webhook_credential"`
	// WebhookService "Service that webhook requests will be accepted from"
	WebhookService types.String `tfsdk:"webhook_service" json:"webhook_service"`
}

// Clone the object
func (o workflowJobTemplateTerraformModel) Clone() workflowJobTemplateTerraformModel {
	return workflowJobTemplateTerraformModel{
		AllowSimultaneous:    o.AllowSimultaneous,
		AskInventoryOnLaunch: o.AskInventoryOnLaunch,
		AskLabelsOnLaunch:    o.AskLabelsOnLaunch,
		AskLimitOnLaunch:     o.AskLimitOnLaunch,
		AskScmBranchOnLaunch: o.AskScmBranchOnLaunch,
		AskSkipTagsOnLaunch:  o.AskSkipTagsOnLaunch,
		AskTagsOnLaunch:      o.AskTagsOnLaunch,
		AskVariablesOnLaunch: o.AskVariablesOnLaunch,
		Description:          o.Description,
		ExtraVars:            o.ExtraVars,
		ID:                   o.ID,
		Inventory:            o.Inventory,
		JobTags:              o.JobTags,
		Limit:                o.Limit,
		Name:                 o.Name,
		Organization:         o.Organization,
		ScmBranch:            o.ScmBranch,
		SkipTags:             o.SkipTags,
		SurveyEnabled:        o.SurveyEnabled,
		WebhookCredential:    o.WebhookCredential,
		WebhookService:       o.WebhookService,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for WorkflowJobTemplate
func (o workflowJobTemplateTerraformModel) BodyRequest() (req workflowJobTemplateBodyRequestModel) {
	req.AllowSimultaneous = o.AllowSimultaneous.ValueBool()
	req.AskInventoryOnLaunch = o.AskInventoryOnLaunch.ValueBool()
	req.AskLabelsOnLaunch = o.AskLabelsOnLaunch.ValueBool()
	req.AskLimitOnLaunch = o.AskLimitOnLaunch.ValueBool()
	req.AskScmBranchOnLaunch = o.AskScmBranchOnLaunch.ValueBool()
	req.AskSkipTagsOnLaunch = o.AskSkipTagsOnLaunch.ValueBool()
	req.AskTagsOnLaunch = o.AskTagsOnLaunch.ValueBool()
	req.AskVariablesOnLaunch = o.AskVariablesOnLaunch.ValueBool()
	req.Description = o.Description.ValueString()
	req.ExtraVars = json.RawMessage(o.ExtraVars.ValueString())
	req.Inventory = o.Inventory.ValueInt64()
	req.JobTags = o.JobTags.ValueString()
	req.Limit = o.Limit.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	req.ScmBranch = o.ScmBranch.ValueString()
	req.SkipTags = o.SkipTags.ValueString()
	req.SurveyEnabled = o.SurveyEnabled.ValueBool()
	req.WebhookCredential = o.WebhookCredential.ValueInt64()
	req.WebhookService = o.WebhookService.ValueString()
	return
}

func (o *workflowJobTemplateTerraformModel) setAllowSimultaneous(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AllowSimultaneous, data)
}

func (o *workflowJobTemplateTerraformModel) setAskInventoryOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskInventoryOnLaunch, data)
}

func (o *workflowJobTemplateTerraformModel) setAskLabelsOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskLabelsOnLaunch, data)
}

func (o *workflowJobTemplateTerraformModel) setAskLimitOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskLimitOnLaunch, data)
}

func (o *workflowJobTemplateTerraformModel) setAskScmBranchOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskScmBranchOnLaunch, data)
}

func (o *workflowJobTemplateTerraformModel) setAskSkipTagsOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskSkipTagsOnLaunch, data)
}

func (o *workflowJobTemplateTerraformModel) setAskTagsOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskTagsOnLaunch, data)
}

func (o *workflowJobTemplateTerraformModel) setAskVariablesOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskVariablesOnLaunch, data)
}

func (o *workflowJobTemplateTerraformModel) setDescription(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *workflowJobTemplateTerraformModel) setExtraVars(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.ExtraVars, data, false)
}

func (o *workflowJobTemplateTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *workflowJobTemplateTerraformModel) setInventory(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Inventory, data)
}

func (o *workflowJobTemplateTerraformModel) setJobTags(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.JobTags, data, false)
}

func (o *workflowJobTemplateTerraformModel) setLimit(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Limit, data, false)
}

func (o *workflowJobTemplateTerraformModel) setName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *workflowJobTemplateTerraformModel) setOrganization(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Organization, data)
}

func (o *workflowJobTemplateTerraformModel) setScmBranch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.ScmBranch, data, false)
}

func (o *workflowJobTemplateTerraformModel) setSkipTags(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SkipTags, data, false)
}

func (o *workflowJobTemplateTerraformModel) setSurveyEnabled(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.SurveyEnabled, data)
}

func (o *workflowJobTemplateTerraformModel) setWebhookCredential(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.WebhookCredential, data)
}

func (o *workflowJobTemplateTerraformModel) setWebhookService(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.WebhookService, data, false)
}

func (o *workflowJobTemplateTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setAllowSimultaneous(data["allow_simultaneous"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAskInventoryOnLaunch(data["ask_inventory_on_launch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAskLabelsOnLaunch(data["ask_labels_on_launch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAskLimitOnLaunch(data["ask_limit_on_launch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAskScmBranchOnLaunch(data["ask_scm_branch_on_launch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAskSkipTagsOnLaunch(data["ask_skip_tags_on_launch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAskTagsOnLaunch(data["ask_tags_on_launch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAskVariablesOnLaunch(data["ask_variables_on_launch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setExtraVars(data["extra_vars"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setInventory(data["inventory"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setJobTags(data["job_tags"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setLimit(data["limit"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setName(data["name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setOrganization(data["organization"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setScmBranch(data["scm_branch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSkipTags(data["skip_tags"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSurveyEnabled(data["survey_enabled"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setWebhookCredential(data["webhook_credential"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setWebhookService(data["webhook_service"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// workflowJobTemplateBodyRequestModel maps the schema for WorkflowJobTemplate for creating and updating the data
type workflowJobTemplateBodyRequestModel struct {
	// AllowSimultaneous ""
	AllowSimultaneous bool `json:"allow_simultaneous"`
	// AskInventoryOnLaunch ""
	AskInventoryOnLaunch bool `json:"ask_inventory_on_launch"`
	// AskLabelsOnLaunch ""
	AskLabelsOnLaunch bool `json:"ask_labels_on_launch"`
	// AskLimitOnLaunch ""
	AskLimitOnLaunch bool `json:"ask_limit_on_launch"`
	// AskScmBranchOnLaunch ""
	AskScmBranchOnLaunch bool `json:"ask_scm_branch_on_launch"`
	// AskSkipTagsOnLaunch ""
	AskSkipTagsOnLaunch bool `json:"ask_skip_tags_on_launch"`
	// AskTagsOnLaunch ""
	AskTagsOnLaunch bool `json:"ask_tags_on_launch"`
	// AskVariablesOnLaunch ""
	AskVariablesOnLaunch bool `json:"ask_variables_on_launch"`
	// Description "Optional description of this workflow job template."
	Description string `json:"description,omitempty"`
	// ExtraVars ""
	ExtraVars json.RawMessage `json:"extra_vars,omitempty"`
	// Inventory "Inventory applied as a prompt, assuming job template prompts for inventory"
	Inventory int64 `json:"inventory,omitempty"`
	// JobTags ""
	JobTags string `json:"job_tags,omitempty"`
	// Limit ""
	Limit string `json:"limit,omitempty"`
	// Name "Name of this workflow job template."
	Name string `json:"name"`
	// Organization "The organization used to determine access to this template."
	Organization int64 `json:"organization,omitempty"`
	// ScmBranch ""
	ScmBranch string `json:"scm_branch,omitempty"`
	// SkipTags ""
	SkipTags string `json:"skip_tags,omitempty"`
	// SurveyEnabled ""
	SurveyEnabled bool `json:"survey_enabled"`
	// WebhookCredential "Personal Access Token for posting back the status to the service API"
	WebhookCredential int64 `json:"webhook_credential,omitempty"`
	// WebhookService "Service that webhook requests will be accepted from"
	WebhookService string `json:"webhook_service,omitempty"`
}

type workflowJobTemplateObjectRolesModel struct {
	ID    types.Int64 `tfsdk:"id"`
	Roles types.Map   `tfsdk:"roles"`
}

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

// GetSchema defines the schema for the data source.
func (o *workflowJobTemplateDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"WorkflowJobTemplate",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"allow_simultaneous": {
					Description: "Allow simultaneous",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"ask_inventory_on_launch": {
					Description: "Ask inventory on launch",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"ask_labels_on_launch": {
					Description: "Ask labels on launch",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"ask_limit_on_launch": {
					Description: "Ask limit on launch",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"ask_scm_branch_on_launch": {
					Description: "Ask scm branch on launch",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"ask_skip_tags_on_launch": {
					Description: "Ask skip tags on launch",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"ask_tags_on_launch": {
					Description: "Ask tags on launch",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"ask_variables_on_launch": {
					Description: "Ask variables on launch",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"description": {
					Description: "Optional description of this workflow job template.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"extra_vars": {
					Description: "Extra vars",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"id": {
					Description: "Database ID for this workflow job template.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.ExactlyOneOf(
							path.MatchRoot("id"),
							path.MatchRoot("name"),
						),
					},
				},
				"inventory": {
					Description: "Inventory applied as a prompt, assuming job template prompts for inventory",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"job_tags": {
					Description: "Job tags",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"limit": {
					Description: "Limit",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"name": {
					Description: "Name of this workflow job template.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.ExactlyOneOf(
							path.MatchRoot("id"),
							path.MatchRoot("name"),
						),
					},
				},
				"organization": {
					Description: "The organization used to determine access to this template.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"scm_branch": {
					Description: "Scm branch",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"skip_tags": {
					Description: "Skip tags",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"survey_enabled": {
					Description: "Survey enabled",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"webhook_credential": {
					Description: "Personal Access Token for posting back the status to the service API",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"webhook_service": {
					Description: "Service that webhook requests will be accepted from",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"", "github", "gitlab"}...),
					},
				},
				// Write only elements
			},
		}), nil
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

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &workflowJobTemplateResource{}
	_ resource.ResourceWithConfigure   = &workflowJobTemplateResource{}
	_ resource.ResourceWithImportState = &workflowJobTemplateResource{}
)

// NewWorkflowJobTemplateResource is a helper function to simplify the provider implementation.
func NewWorkflowJobTemplateResource() resource.Resource {
	return &workflowJobTemplateResource{}
}

type workflowJobTemplateResource struct {
	client   c.Client
	endpoint string
}

func (o *workflowJobTemplateResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/workflow_job_templates/"
}

func (o workflowJobTemplateResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_workflow_job_template"
}

func (o workflowJobTemplateResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"WorkflowJobTemplate",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"allow_simultaneous": {
					Description: "Allow simultaneous",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"ask_inventory_on_launch": {
					Description: "Ask inventory on launch",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"ask_labels_on_launch": {
					Description: "Ask labels on launch",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"ask_limit_on_launch": {
					Description: "Ask limit on launch",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"ask_scm_branch_on_launch": {
					Description: "Ask scm branch on launch",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"ask_skip_tags_on_launch": {
					Description: "Ask skip tags on launch",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"ask_tags_on_launch": {
					Description: "Ask tags on launch",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"ask_variables_on_launch": {
					Description: "Ask variables on launch",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"description": {
					Description: "Optional description of this workflow job template.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"extra_vars": {
					Description: "Extra vars",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"inventory": {
					Description: "Inventory applied as a prompt, assuming job template prompts for inventory",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"job_tags": {
					Description: "Job tags",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"limit": {
					Description: "Limit",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"name": {
					Description:   "Name of this workflow job template.",
					Type:          types.StringType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(512),
					},
				},
				"organization": {
					Description: "The organization used to determine access to this template.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"scm_branch": {
					Description: "Scm branch",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"skip_tags": {
					Description: "Skip tags",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"survey_enabled": {
					Description: "Survey enabled",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"webhook_credential": {
					Description: "Personal Access Token for posting back the status to the service API",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"webhook_service": {
					Description: "Service that webhook requests will be accepted from",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"", "github", "gitlab"}...),
					},
				},
				// Write only elements
				// Data only elements
				"id": {
					Description: "Database ID for this workflow job template.",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
			},
		}), nil
}

func (o *workflowJobTemplateResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the WorkflowJobTemplate.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (o *workflowJobTemplateResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state workflowJobTemplateTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for WorkflowJobTemplate
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for WorkflowJobTemplate on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new WorkflowJobTemplate resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for WorkflowJobTemplate on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *workflowJobTemplateResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state workflowJobTemplateTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for WorkflowJobTemplate
	var r *http.Request
	var id = state.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for WorkflowJobTemplate on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for WorkflowJobTemplate from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for WorkflowJobTemplate on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *workflowJobTemplateResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state workflowJobTemplateTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for WorkflowJobTemplate
	var r *http.Request
	var id = plan.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for WorkflowJobTemplate on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new WorkflowJobTemplate resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for WorkflowJobTemplate on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *workflowJobTemplateResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state workflowJobTemplateTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for WorkflowJobTemplate
	var r *http.Request
	var id = state.ID
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for WorkflowJobTemplate on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	// Delete existing WorkflowJobTemplate
	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to delete resource for WorkflowJobTemplate on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}

var (
	_ datasource.DataSource              = &workflowJobTemplateObjectRolesDataSource{}
	_ datasource.DataSourceWithConfigure = &workflowJobTemplateObjectRolesDataSource{}
)

// NewWorkflowJobTemplateObjectRolesDataSource is a helper function to instantiate the WorkflowJobTemplate data source.
func NewWorkflowJobTemplateObjectRolesDataSource() datasource.DataSource {
	return &workflowJobTemplateObjectRolesDataSource{}
}

// workflowJobTemplateObjectRolesDataSource is the data source implementation.
type workflowJobTemplateObjectRolesDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *workflowJobTemplateObjectRolesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/workflow_job_templates/%d/object_roles/"
}

// Metadata returns the data source type name.
func (o *workflowJobTemplateObjectRolesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workflow_job_template_object_roles"
}

// GetSchema defines the schema for the data source.
func (o *workflowJobTemplateObjectRolesDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Version: helpers.SchemaVersion,
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description: "WorkflowJobTemplate ID",
				Type:        types.Int64Type,
				Required:    true,
			},
			"roles": {
				Description: "Roles for workflowjobtemplate",
				Type:        types.MapType{ElemType: types.Int64Type},
				Computed:    true,
			},
		},
	}, nil
}

// Read refreshes the Terraform state with the latest data.
func (o *workflowJobTemplateObjectRolesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state workflowJobTemplateObjectRolesModel
	var err error
	var id types.Int64

	if d := req.Config.GetAttribute(ctx, path.Root("id"), &id); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	state.ID = types.Int64Value(id.ValueInt64())

	// Creates a new request for Credential
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf(o.endpoint, id.ValueInt64()), nil); err != nil {
		resp.Diagnostics.AddError(
			"Unable to create a new request for workflowJobTemplate",
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Credential
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			"Unable to fetch the request for workflowjobtemplate object roles",
			err.Error(),
		)
		return
	}

	var sr searchResultObjectRole
	if err = mapstructure.Decode(data, &sr); err != nil {
		resp.Diagnostics.AddError(
			"Unable to decode the search result data for workflowjobtemplate",
			err.Error(),
		)
		return
	}

	var in = make(map[string]attr.Value, sr.Count)
	for _, role := range sr.Results {
		in[role.Name] = types.Int64Value(role.ID)
	}

	var d diag.Diagnostics
	if state.Roles, d = types.MapValue(types.Int64Type, in); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}

	// Set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

var (
	_ resource.Resource              = &workflowJobTemplateAssociateDisassociateNotificationTemplate{}
	_ resource.ResourceWithConfigure = &workflowJobTemplateAssociateDisassociateNotificationTemplate{}
)

type workflowJobTemplateAssociateDisassociateNotificationTemplateTerraformModel struct {
	WorkflowJobTemplateID  types.Int64  `tfsdk:"workflow_job_template_id"`
	NotificationTemplateID types.Int64  `tfsdk:"notification_template_id"`
	Option                 types.String `tfsdk:"option"`
}

// NewWorkflowJobTemplateAssociateDisassociateNotificationTemplateResource is a helper function to simplify the provider implementation.
func NewWorkflowJobTemplateAssociateDisassociateNotificationTemplateResource() resource.Resource {
	return &workflowJobTemplateAssociateDisassociateNotificationTemplate{}
}

type workflowJobTemplateAssociateDisassociateNotificationTemplate struct {
	client   c.Client
	endpoint string
}

func (o *workflowJobTemplateAssociateDisassociateNotificationTemplate) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/workflow_job_templates/%d/notification_templates_%s/"
}

func (o workflowJobTemplateAssociateDisassociateNotificationTemplate) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_workflow_job_template_associate_notification_template"
}

func (o workflowJobTemplateAssociateDisassociateNotificationTemplate) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"WorkflowJobTemplate/Associate",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				"workflow_job_template_id": {
					Description: "Database ID for this WorkflowJobTemplate.",
					Required:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.RequiresReplace(),
					},
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.AlsoRequires(
							path.MatchRoot("option"),
							path.MatchRoot("notification_template_id"),
						),
					},
				},
				"option": {
					Description: "Notification Option",
					Required:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.RequiresReplace(),
					},
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.AlsoRequires(
							path.MatchRoot("workflow_job_template_id"),
						),
						stringvalidator.OneOf([]string{"approval", "started", "success", "error"}...),
					},
				},
				"notification_template_id": {
					Description: "Database ID of the notificationtemplate to assign.",
					Required:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.RequiresReplace(),
					},
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.AlsoRequires(
							path.MatchRoot("option"),
							path.MatchRoot("workflow_job_template_id"),
						),
					},
				},
			},
		},
	), nil
}

func (o *workflowJobTemplateAssociateDisassociateNotificationTemplate) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error

	// Retrieve values from state
	var plan, state workflowJobTemplateAssociateDisassociateNotificationTemplateTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for association of WorkflowJobTemplate
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.WorkflowJobTemplateID.ValueInt64(), plan.Option.ValueString())) + "/"
	var buf bytes.Buffer
	var bodyRequest = associateDisassociateRequestModel{ID: plan.NotificationTemplateID.ValueInt64(), Disassociate: false}
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for WorkflowJobTemplate on %s for create of type notification_job_workflow_template", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to associate for WorkflowJobTemplate on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	state.WorkflowJobTemplateID = plan.WorkflowJobTemplateID
	state.NotificationTemplateID = plan.NotificationTemplateID
	state.Option = plan.Option

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *workflowJobTemplateAssociateDisassociateNotificationTemplate) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state workflowJobTemplateAssociateDisassociateNotificationTemplateTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for disassociation of WorkflowJobTemplate
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.WorkflowJobTemplateID.ValueInt64(), state.Option.ValueString())) + "/"
	var buf bytes.Buffer
	var bodyRequest = associateDisassociateRequestModel{ID: state.NotificationTemplateID.ValueInt64(), Disassociate: true}
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for WorkflowJobTemplate on %s for delete of type notification_job_workflow_template", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to disassociate for WorkflowJobTemplate on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}

func (o *workflowJobTemplateAssociateDisassociateNotificationTemplate) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
}

func (o *workflowJobTemplateAssociateDisassociateNotificationTemplate) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
}

var (
	_ resource.Resource                = &workflowJobTemplateSurvey{}
	_ resource.ResourceWithConfigure   = &workflowJobTemplateSurvey{}
	_ resource.ResourceWithImportState = &workflowJobTemplateSurvey{}
)

type workflowJobTemplateSurveyTerraformModel struct {
	WorkflowJobTemplateID types.Int64  `tfsdk:"workflow_job_template_id"`
	Spec                  types.String `tfsdk:"spec"`
}

func (o workflowJobTemplateSurveyTerraformModel) Clone() workflowJobTemplateSurveyTerraformModel {
	return workflowJobTemplateSurveyTerraformModel{
		WorkflowJobTemplateID: types.Int64Value(o.WorkflowJobTemplateID.ValueInt64()),
		Spec:                  types.StringValue(o.Spec.ValueString()),
	}
}

func (o workflowJobTemplateSurveyTerraformModel) BodyRequest() workflowJobTemplateSurveyModel {
	return workflowJobTemplateSurveyModel{
		Spec: json.RawMessage(o.Spec.ValueString()),
	}
}

type workflowJobTemplateSurveyModel struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Spec        json.RawMessage `json:"spec"`
}

// NewWorkflowJobTemplateSurveyResource is a helper function to simplify the provider implementation.
func NewWorkflowJobTemplateSurveyResource() resource.Resource {
	return &workflowJobTemplateSurvey{}
}

type workflowJobTemplateSurvey struct {
	client   c.Client
	endpoint string
}

func (o *workflowJobTemplateSurvey) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/workflow_job_templates/%d/survey_spec/"
}

func (o workflowJobTemplateSurvey) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_workflow_job_template_survey_spec"
}

func (o workflowJobTemplateSurvey) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"WorkflowJobTemplate/Survey",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				"workflow_job_template_id": {
					Description: "Database ID for this WorkflowJobTemplate.",
					Required:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.RequiresReplace(),
					},
				},
				"spec": {
					Description: "The survey spec for this WorkflowJobTemplate.",
					Required:    true,
					Type:        types.StringType,
				},
			},
		}), nil
}

// ImportState imports the survey spec for WorkflowJobTemplate
func (o *workflowJobTemplateSurvey) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the WorkflowJobTemplate.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("workflow_job_template_id"), id)...)
}

// Delete the survey spec for WorkflowJobTemplate
func (o *workflowJobTemplateSurvey) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	var state workflowJobTemplateSurveyTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.WorkflowJobTemplateID.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for WorkflowJobTemplate on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to delete resource for WorkflowJobTemplate/Survey on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}

// Read the survey spec for WorkflowJobTemplate
func (o *workflowJobTemplateSurvey) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error
	var state workflowJobTemplateSurveyTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.WorkflowJobTemplateID.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for WorkflowJobTemplate on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for WorkflowJobTemplate/Survey on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	if val, ok := data["spec"]; ok {
		dg, _ := helpers.AttrValueSetJsonString(&state.Spec, val, false)
		if dg.HasError() {
			response.Diagnostics.Append(dg...)
		}
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

// Create the survey spec for WorkflowJobTemplate
func (o *workflowJobTemplateSurvey) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state workflowJobTemplateSurveyTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.WorkflowJobTemplateID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for WorkflowJobTemplate on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for WorkflowJobTemplate/Survey on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	state.Spec = types.StringValue(plan.Spec.ValueString())
	state.WorkflowJobTemplateID = types.Int64Value(plan.WorkflowJobTemplateID.ValueInt64())
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

// Update the survey spec for WorkflowJobTemplate
func (o *workflowJobTemplateSurvey) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state workflowJobTemplateSurveyTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.WorkflowJobTemplateID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for WorkflowJobTemplate on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for WorkflowJobTemplate/Survey on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	state.Spec = types.StringValue(plan.Spec.ValueString())
	state.WorkflowJobTemplateID = types.Int64Value(plan.WorkflowJobTemplateID.ValueInt64())
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}
