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
func (o *workflowJobTemplateTerraformModel) BodyRequest() (req workflowJobTemplateBodyRequestModel) {
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
