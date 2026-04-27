package awx

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
)

type workflowJobTemplateTerraformModel struct {
	AllowSimultaneous    types.Bool   `tfsdk:"allow_simultaneous" json:"allow_simultaneous"`
	AskInventoryOnLaunch types.Bool   `tfsdk:"ask_inventory_on_launch" json:"ask_inventory_on_launch"`
	AskLabelsOnLaunch    types.Bool   `tfsdk:"ask_labels_on_launch" json:"ask_labels_on_launch"`
	AskLimitOnLaunch     types.Bool   `tfsdk:"ask_limit_on_launch" json:"ask_limit_on_launch"`
	AskScmBranchOnLaunch types.Bool   `tfsdk:"ask_scm_branch_on_launch" json:"ask_scm_branch_on_launch"`
	AskSkipTagsOnLaunch  types.Bool   `tfsdk:"ask_skip_tags_on_launch" json:"ask_skip_tags_on_launch"`
	AskTagsOnLaunch      types.Bool   `tfsdk:"ask_tags_on_launch" json:"ask_tags_on_launch"`
	AskVariablesOnLaunch types.Bool   `tfsdk:"ask_variables_on_launch" json:"ask_variables_on_launch"`
	Description          types.String `tfsdk:"description" json:"description"`
	ExtraVars            types.String `tfsdk:"extra_vars" json:"extra_vars"`
	ID                   types.Int64  `tfsdk:"id" json:"id"`
	Inventory            types.Int64  `tfsdk:"inventory" json:"inventory"`
	JobTags              types.String `tfsdk:"job_tags" json:"job_tags"`
	Limit                types.String `tfsdk:"limit" json:"limit"`
	Name                 types.String `tfsdk:"name" json:"name"`
	Organization         types.Int64  `tfsdk:"organization" json:"organization"`
	ScmBranch            types.String `tfsdk:"scm_branch" json:"scm_branch"`
	SkipTags             types.String `tfsdk:"skip_tags" json:"skip_tags"`
	SurveyEnabled        types.Bool   `tfsdk:"survey_enabled" json:"survey_enabled"`
	WebhookCredential    types.Int64  `tfsdk:"webhook_credential" json:"webhook_credential"`
	WebhookService       types.String `tfsdk:"webhook_service" json:"webhook_service"`
}

func (o *workflowJobTemplateTerraformModel) Clone() workflowJobTemplateTerraformModel {
	return *o
}

func (o *workflowJobTemplateTerraformModel) BodyRequest() *workflowJobTemplateBodyRequestModel {
	var req workflowJobTemplateBodyRequestModel
	req.AllowSimultaneous = o.AllowSimultaneous.ValueBool()
	req.AskInventoryOnLaunch = o.AskInventoryOnLaunch.ValueBool()
	req.AskLabelsOnLaunch = o.AskLabelsOnLaunch.ValueBool()
	req.AskLimitOnLaunch = o.AskLimitOnLaunch.ValueBool()
	req.AskScmBranchOnLaunch = o.AskScmBranchOnLaunch.ValueBool()
	req.AskSkipTagsOnLaunch = o.AskSkipTagsOnLaunch.ValueBool()
	req.AskTagsOnLaunch = o.AskTagsOnLaunch.ValueBool()
	req.AskVariablesOnLaunch = o.AskVariablesOnLaunch.ValueBool()
	req.Description = o.Description.ValueString()
	req.ExtraVars = json.RawMessage(o.ExtraVars.String())
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
	return &req
}

func (o *workflowJobTemplateTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetBool(&o.AllowSimultaneous, data["allow_simultaneous"]))
	collect(helpers.AttrValueSetBool(&o.AskInventoryOnLaunch, data["ask_inventory_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskLabelsOnLaunch, data["ask_labels_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskLimitOnLaunch, data["ask_limit_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskScmBranchOnLaunch, data["ask_scm_branch_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskSkipTagsOnLaunch, data["ask_skip_tags_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskTagsOnLaunch, data["ask_tags_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskVariablesOnLaunch, data["ask_variables_on_launch"]))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetJsonString(&o.ExtraVars, data["extra_vars"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetInt64(&o.Inventory, data["inventory"]))
	collect(helpers.AttrValueSetString(&o.JobTags, data["job_tags"], false))
	collect(helpers.AttrValueSetString(&o.Limit, data["limit"], false))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	collect(helpers.AttrValueSetString(&o.ScmBranch, data["scm_branch"], false))
	collect(helpers.AttrValueSetString(&o.SkipTags, data["skip_tags"], false))
	collect(helpers.AttrValueSetBool(&o.SurveyEnabled, data["survey_enabled"]))
	collect(helpers.AttrValueSetInt64(&o.WebhookCredential, data["webhook_credential"]))
	collect(helpers.AttrValueSetString(&o.WebhookService, data["webhook_service"], false))
	return diags, nil
}

type workflowJobTemplateBodyRequestModel struct {
	AllowSimultaneous    bool            `json:"allow_simultaneous"`
	AskInventoryOnLaunch bool            `json:"ask_inventory_on_launch"`
	AskLabelsOnLaunch    bool            `json:"ask_labels_on_launch"`
	AskLimitOnLaunch     bool            `json:"ask_limit_on_launch"`
	AskScmBranchOnLaunch bool            `json:"ask_scm_branch_on_launch"`
	AskSkipTagsOnLaunch  bool            `json:"ask_skip_tags_on_launch"`
	AskTagsOnLaunch      bool            `json:"ask_tags_on_launch"`
	AskVariablesOnLaunch bool            `json:"ask_variables_on_launch"`
	Description          string          `json:"description,omitempty"`
	ExtraVars            json.RawMessage `json:"extra_vars,omitempty"`
	Inventory            int64           `json:"inventory,omitempty"`
	JobTags              string          `json:"job_tags,omitempty"`
	Limit                string          `json:"limit,omitempty"`
	Name                 string          `json:"name"`
	Organization         int64           `json:"organization,omitempty"`
	ScmBranch            string          `json:"scm_branch,omitempty"`
	SkipTags             string          `json:"skip_tags,omitempty"`
	SurveyEnabled        bool            `json:"survey_enabled"`
	WebhookCredential    int64           `json:"webhook_credential,omitempty"`
	WebhookService       string          `json:"webhook_service,omitempty"`
}

type workflowJobTemplateResource = framework.GenericResource[workflowJobTemplateTerraformModel, workflowJobTemplateBodyRequestModel, *workflowJobTemplateTerraformModel]

// NewWorkflowJobTemplateResource is a helper function to simplify the provider implementation.
func NewWorkflowJobTemplateResource() resource.Resource {
	return &workflowJobTemplateResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "workflow_job_template", Endpoint: "/api/v2/workflow_job_templates/"}},
		Cfg: framework.ResourceCfg[workflowJobTemplateTerraformModel, workflowJobTemplateBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"allow_simultaneous": schema.BoolAttribute{
						Description: "Allow simultaneous",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_inventory_on_launch": schema.BoolAttribute{
						Description: "Ask inventory on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_labels_on_launch": schema.BoolAttribute{
						Description: "Ask labels on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_limit_on_launch": schema.BoolAttribute{
						Description: "Ask limit on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_scm_branch_on_launch": schema.BoolAttribute{
						Description: "Ask scm branch on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_skip_tags_on_launch": schema.BoolAttribute{
						Description: "Ask skip tags on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_tags_on_launch": schema.BoolAttribute{
						Description: "Ask tags on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"ask_variables_on_launch": schema.BoolAttribute{
						Description: "Ask variables on launch",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this workflow job template.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
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
					"inventory": schema.Int64Attribute{
						Description: "Inventory applied as a prompt, assuming job template prompts for inventory",
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
					"limit": schema.StringAttribute{
						Description: "Limit",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this workflow job template.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"organization": schema.Int64Attribute{
						Description: "The organization used to determine access to this template.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
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
					"survey_enabled": schema.BoolAttribute{
						Description: "Survey enabled",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"webhook_credential": schema.Int64Attribute{
						Description: "Personal Access Token for posting back the status to the service API",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"webhook_service": schema.StringAttribute{
						Description: "Service that webhook requests will be accepted from",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"",
								"github",
								"gitlab",
								"bitbucket_dc",
							),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this workflow job template.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor: func(m *workflowJobTemplateTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:      "id",
			Hook: func(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *workflowJobTemplateTerraformModel) error {
				return hooks.RequireResourceStateOrOrig(ctx, apiVersion, source, callee, orig, state)
			},
			ApiVersion:   ApiVersion,
			ResourceName: "WorkflowJobTemplate",
		},
	}
}

type workflowJobTemplateDataSource = framework.GenericDataSource[workflowJobTemplateTerraformModel, *workflowJobTemplateTerraformModel]

// NewWorkflowJobTemplateDataSource is a helper function to instantiate the WorkflowJobTemplate data source.
func NewWorkflowJobTemplateDataSource() datasource.DataSource {
	return &workflowJobTemplateDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "workflow_job_template", Endpoint: "/api/v2/workflow_job_templates/"}},
		Cfg: framework.DataSourceCfg[workflowJobTemplateTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"allow_simultaneous": dschema.BoolAttribute{
						Description: "Allow simultaneous",
						Computed:    true,
					},
					"ask_inventory_on_launch": dschema.BoolAttribute{
						Description: "Ask inventory on launch",
						Computed:    true,
					},
					"ask_labels_on_launch": dschema.BoolAttribute{
						Description: "Ask labels on launch",
						Computed:    true,
					},
					"ask_limit_on_launch": dschema.BoolAttribute{
						Description: "Ask limit on launch",
						Computed:    true,
					},
					"ask_scm_branch_on_launch": dschema.BoolAttribute{
						Description: "Ask scm branch on launch",
						Computed:    true,
					},
					"ask_skip_tags_on_launch": dschema.BoolAttribute{
						Description: "Ask skip tags on launch",
						Computed:    true,
					},
					"ask_tags_on_launch": dschema.BoolAttribute{
						Description: "Ask tags on launch",
						Computed:    true,
					},
					"ask_variables_on_launch": dschema.BoolAttribute{
						Description: "Ask variables on launch",
						Computed:    true,
					},
					"description": dschema.StringAttribute{
						Description: "Optional description of this workflow job template.",
						Computed:    true,
					},
					"extra_vars": dschema.StringAttribute{
						Description: "Extra vars",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this workflow job template.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"inventory": dschema.Int64Attribute{
						Description: "Inventory applied as a prompt, assuming job template prompts for inventory",
						Computed:    true,
					},
					"job_tags": dschema.StringAttribute{
						Description: "Job tags",
						Computed:    true,
					},
					"limit": dschema.StringAttribute{
						Description: "Limit",
						Computed:    true,
					},
					"name": dschema.StringAttribute{
						Description: "Name of this workflow job template.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"organization": dschema.Int64Attribute{
						Description: "The organization used to determine access to this template.",
						Computed:    true,
					},
					"scm_branch": dschema.StringAttribute{
						Description: "Scm branch",
						Computed:    true,
					},
					"skip_tags": dschema.StringAttribute{
						Description: "Skip tags",
						Computed:    true,
					},
					"survey_enabled": dschema.BoolAttribute{
						Description: "Survey enabled",
						Computed:    true,
					},
					"webhook_credential": dschema.Int64Attribute{
						Description: "Personal Access Token for posting back the status to the service API",
						Computed:    true,
					},
					"webhook_service": dschema.StringAttribute{
						Description: "Service that webhook requests will be accepted from",
						Computed:    true,
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
				{Name: "by_name", URLSuffix: "?name__exact=%s", Fields: []framework.SearchField{
					{Name: "name", Type: "string", URLEscape: true},
				}},
			},
			Hook: func(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *workflowJobTemplateTerraformModel) error {
				return hooks.RequireResourceStateOrOrig(ctx, apiVersion, source, callee, orig, state)
			},
			ApiVersion:   ApiVersion,
			ResourceName: "WorkflowJobTemplate",
		},
	}
}
