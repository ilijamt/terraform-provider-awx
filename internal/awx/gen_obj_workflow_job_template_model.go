package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
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
func (o *workflowJobTemplateTerraformModel) Clone() workflowJobTemplateTerraformModel {
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for WorkflowJobTemplate
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
	{
		dg, _ := helpers.AttrValueSetBool(&o.AllowSimultaneous, data["allow_simultaneous"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.AskInventoryOnLaunch, data["ask_inventory_on_launch"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.AskLabelsOnLaunch, data["ask_labels_on_launch"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.AskLimitOnLaunch, data["ask_limit_on_launch"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.AskScmBranchOnLaunch, data["ask_scm_branch_on_launch"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.AskSkipTagsOnLaunch, data["ask_skip_tags_on_launch"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.AskTagsOnLaunch, data["ask_tags_on_launch"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.AskVariablesOnLaunch, data["ask_variables_on_launch"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Description, data["description"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.ExtraVars, data["extra_vars"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.ID, data["id"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.Inventory, data["inventory"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.JobTags, data["job_tags"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Limit, data["limit"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Name, data["name"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.Organization, data["organization"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.ScmBranch, data["scm_branch"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.SkipTags, data["skip_tags"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.SurveyEnabled, data["survey_enabled"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.WebhookCredential, data["webhook_credential"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.WebhookService, data["webhook_service"], false)
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
