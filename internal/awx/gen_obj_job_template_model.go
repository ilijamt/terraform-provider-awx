package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// jobTemplateTerraformModel maps the schema for JobTemplate when using Data Source
type jobTemplateTerraformModel struct {
	// AllowSimultaneous ""
	AllowSimultaneous types.Bool `tfsdk:"allow_simultaneous" json:"allow_simultaneous"`
	// AskCredentialOnLaunch ""
	AskCredentialOnLaunch types.Bool `tfsdk:"ask_credential_on_launch" json:"ask_credential_on_launch"`
	// AskDiffModeOnLaunch ""
	AskDiffModeOnLaunch types.Bool `tfsdk:"ask_diff_mode_on_launch" json:"ask_diff_mode_on_launch"`
	// AskExecutionEnvironmentOnLaunch ""
	AskExecutionEnvironmentOnLaunch types.Bool `tfsdk:"ask_execution_environment_on_launch" json:"ask_execution_environment_on_launch"`
	// AskForksOnLaunch ""
	AskForksOnLaunch types.Bool `tfsdk:"ask_forks_on_launch" json:"ask_forks_on_launch"`
	// AskInstanceGroupsOnLaunch ""
	AskInstanceGroupsOnLaunch types.Bool `tfsdk:"ask_instance_groups_on_launch" json:"ask_instance_groups_on_launch"`
	// AskInventoryOnLaunch ""
	AskInventoryOnLaunch types.Bool `tfsdk:"ask_inventory_on_launch" json:"ask_inventory_on_launch"`
	// AskJobSliceCountOnLaunch ""
	AskJobSliceCountOnLaunch types.Bool `tfsdk:"ask_job_slice_count_on_launch" json:"ask_job_slice_count_on_launch"`
	// AskJobTypeOnLaunch ""
	AskJobTypeOnLaunch types.Bool `tfsdk:"ask_job_type_on_launch" json:"ask_job_type_on_launch"`
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
	// AskTimeoutOnLaunch ""
	AskTimeoutOnLaunch types.Bool `tfsdk:"ask_timeout_on_launch" json:"ask_timeout_on_launch"`
	// AskVariablesOnLaunch ""
	AskVariablesOnLaunch types.Bool `tfsdk:"ask_variables_on_launch" json:"ask_variables_on_launch"`
	// AskVerbosityOnLaunch ""
	AskVerbosityOnLaunch types.Bool `tfsdk:"ask_verbosity_on_launch" json:"ask_verbosity_on_launch"`
	// BecomeEnabled ""
	BecomeEnabled types.Bool `tfsdk:"become_enabled" json:"become_enabled"`
	// Description "Optional description of this job template."
	Description types.String `tfsdk:"description" json:"description"`
	// DiffMode "If enabled, textual changes made to any templated files on the host are shown in the standard output"
	DiffMode types.Bool `tfsdk:"diff_mode" json:"diff_mode"`
	// ExecutionEnvironment "The container image to be used for execution."
	ExecutionEnvironment types.Int64 `tfsdk:"execution_environment" json:"execution_environment"`
	// ExtraVars ""
	ExtraVars types.String `tfsdk:"extra_vars" json:"extra_vars"`
	// ForceHandlers ""
	ForceHandlers types.Bool `tfsdk:"force_handlers" json:"force_handlers"`
	// Forks ""
	Forks types.Int64 `tfsdk:"forks" json:"forks"`
	// HostConfigKey ""
	HostConfigKey types.String `tfsdk:"host_config_key" json:"host_config_key"`
	// ID "Database ID for this job template."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// Inventory ""
	Inventory types.Int64 `tfsdk:"inventory" json:"inventory"`
	// JobSliceCount "The number of jobs to slice into at runtime. Will cause the Job Template to launch a workflow if value is greater than 1."
	JobSliceCount types.Int64 `tfsdk:"job_slice_count" json:"job_slice_count"`
	// JobTags ""
	JobTags types.String `tfsdk:"job_tags" json:"job_tags"`
	// JobType ""
	JobType types.String `tfsdk:"job_type" json:"job_type"`
	// Limit ""
	Limit types.String `tfsdk:"limit" json:"limit"`
	// Name "Name of this job template."
	Name types.String `tfsdk:"name" json:"name"`
	// Organization "The organization used to determine access to this template."
	Organization types.Int64 `tfsdk:"organization" json:"organization"`
	// Playbook ""
	Playbook types.String `tfsdk:"playbook" json:"playbook"`
	// PreventInstanceGroupFallback "If enabled, the job template will prevent adding any inventory or organization instance groups to the list of preferred instances groups to run on.If this setting is enabled and you provided an empty list, the global instance groups will be applied."
	PreventInstanceGroupFallback types.Bool `tfsdk:"prevent_instance_group_fallback" json:"prevent_instance_group_fallback"`
	// Project ""
	Project types.Int64 `tfsdk:"project" json:"project"`
	// ScmBranch "Branch to use in job run. Project default used if blank. Only allowed if project allow_override field is set to true."
	ScmBranch types.String `tfsdk:"scm_branch" json:"scm_branch"`
	// SkipTags ""
	SkipTags types.String `tfsdk:"skip_tags" json:"skip_tags"`
	// StartAtTask ""
	StartAtTask types.String `tfsdk:"start_at_task" json:"start_at_task"`
	// SurveyEnabled ""
	SurveyEnabled types.Bool `tfsdk:"survey_enabled" json:"survey_enabled"`
	// Timeout "The amount of time (in seconds) to run before the task is canceled."
	Timeout types.Int64 `tfsdk:"timeout" json:"timeout"`
	// UseFactCache "If enabled, the service will act as an Ansible Fact Cache Plugin; persisting facts at the end of a playbook run to the database and caching facts for use by Ansible."
	UseFactCache types.Bool `tfsdk:"use_fact_cache" json:"use_fact_cache"`
	// Verbosity ""
	Verbosity types.String `tfsdk:"verbosity" json:"verbosity"`
	// WebhookCredential "Personal Access Token for posting back the status to the service API"
	WebhookCredential types.Int64 `tfsdk:"webhook_credential" json:"webhook_credential"`
	// WebhookService "Service that webhook requests will be accepted from"
	WebhookService types.String `tfsdk:"webhook_service" json:"webhook_service"`
}

// Clone the object
func (o *jobTemplateTerraformModel) Clone() jobTemplateTerraformModel {
	return jobTemplateTerraformModel{
		AllowSimultaneous:               o.AllowSimultaneous,
		AskCredentialOnLaunch:           o.AskCredentialOnLaunch,
		AskDiffModeOnLaunch:             o.AskDiffModeOnLaunch,
		AskExecutionEnvironmentOnLaunch: o.AskExecutionEnvironmentOnLaunch,
		AskForksOnLaunch:                o.AskForksOnLaunch,
		AskInstanceGroupsOnLaunch:       o.AskInstanceGroupsOnLaunch,
		AskInventoryOnLaunch:            o.AskInventoryOnLaunch,
		AskJobSliceCountOnLaunch:        o.AskJobSliceCountOnLaunch,
		AskJobTypeOnLaunch:              o.AskJobTypeOnLaunch,
		AskLabelsOnLaunch:               o.AskLabelsOnLaunch,
		AskLimitOnLaunch:                o.AskLimitOnLaunch,
		AskScmBranchOnLaunch:            o.AskScmBranchOnLaunch,
		AskSkipTagsOnLaunch:             o.AskSkipTagsOnLaunch,
		AskTagsOnLaunch:                 o.AskTagsOnLaunch,
		AskTimeoutOnLaunch:              o.AskTimeoutOnLaunch,
		AskVariablesOnLaunch:            o.AskVariablesOnLaunch,
		AskVerbosityOnLaunch:            o.AskVerbosityOnLaunch,
		BecomeEnabled:                   o.BecomeEnabled,
		Description:                     o.Description,
		DiffMode:                        o.DiffMode,
		ExecutionEnvironment:            o.ExecutionEnvironment,
		ExtraVars:                       o.ExtraVars,
		ForceHandlers:                   o.ForceHandlers,
		Forks:                           o.Forks,
		HostConfigKey:                   o.HostConfigKey,
		ID:                              o.ID,
		Inventory:                       o.Inventory,
		JobSliceCount:                   o.JobSliceCount,
		JobTags:                         o.JobTags,
		JobType:                         o.JobType,
		Limit:                           o.Limit,
		Name:                            o.Name,
		Organization:                    o.Organization,
		Playbook:                        o.Playbook,
		PreventInstanceGroupFallback:    o.PreventInstanceGroupFallback,
		Project:                         o.Project,
		ScmBranch:                       o.ScmBranch,
		SkipTags:                        o.SkipTags,
		StartAtTask:                     o.StartAtTask,
		SurveyEnabled:                   o.SurveyEnabled,
		Timeout:                         o.Timeout,
		UseFactCache:                    o.UseFactCache,
		Verbosity:                       o.Verbosity,
		WebhookCredential:               o.WebhookCredential,
		WebhookService:                  o.WebhookService,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for JobTemplate
func (o *jobTemplateTerraformModel) BodyRequest() (req jobTemplateBodyRequestModel) {
	req.AllowSimultaneous = o.AllowSimultaneous.ValueBool()
	req.AskCredentialOnLaunch = o.AskCredentialOnLaunch.ValueBool()
	req.AskDiffModeOnLaunch = o.AskDiffModeOnLaunch.ValueBool()
	req.AskExecutionEnvironmentOnLaunch = o.AskExecutionEnvironmentOnLaunch.ValueBool()
	req.AskForksOnLaunch = o.AskForksOnLaunch.ValueBool()
	req.AskInstanceGroupsOnLaunch = o.AskInstanceGroupsOnLaunch.ValueBool()
	req.AskInventoryOnLaunch = o.AskInventoryOnLaunch.ValueBool()
	req.AskJobSliceCountOnLaunch = o.AskJobSliceCountOnLaunch.ValueBool()
	req.AskJobTypeOnLaunch = o.AskJobTypeOnLaunch.ValueBool()
	req.AskLabelsOnLaunch = o.AskLabelsOnLaunch.ValueBool()
	req.AskLimitOnLaunch = o.AskLimitOnLaunch.ValueBool()
	req.AskScmBranchOnLaunch = o.AskScmBranchOnLaunch.ValueBool()
	req.AskSkipTagsOnLaunch = o.AskSkipTagsOnLaunch.ValueBool()
	req.AskTagsOnLaunch = o.AskTagsOnLaunch.ValueBool()
	req.AskTimeoutOnLaunch = o.AskTimeoutOnLaunch.ValueBool()
	req.AskVariablesOnLaunch = o.AskVariablesOnLaunch.ValueBool()
	req.AskVerbosityOnLaunch = o.AskVerbosityOnLaunch.ValueBool()
	req.BecomeEnabled = o.BecomeEnabled.ValueBool()
	req.Description = o.Description.ValueString()
	req.DiffMode = o.DiffMode.ValueBool()
	req.ExecutionEnvironment = o.ExecutionEnvironment.ValueInt64()
	req.ExtraVars = json.RawMessage(o.ExtraVars.String())
	req.ForceHandlers = o.ForceHandlers.ValueBool()
	req.Forks = o.Forks.ValueInt64()
	req.HostConfigKey = o.HostConfigKey.ValueString()
	req.Inventory = o.Inventory.ValueInt64()
	req.JobSliceCount = o.JobSliceCount.ValueInt64()
	req.JobTags = o.JobTags.ValueString()
	req.JobType = o.JobType.ValueString()
	req.Limit = o.Limit.ValueString()
	req.Name = o.Name.ValueString()
	req.Playbook = o.Playbook.ValueString()
	req.PreventInstanceGroupFallback = o.PreventInstanceGroupFallback.ValueBool()
	req.Project = o.Project.ValueInt64()
	req.ScmBranch = o.ScmBranch.ValueString()
	req.SkipTags = o.SkipTags.ValueString()
	req.StartAtTask = o.StartAtTask.ValueString()
	req.SurveyEnabled = o.SurveyEnabled.ValueBool()
	req.Timeout = o.Timeout.ValueInt64()
	req.UseFactCache = o.UseFactCache.ValueBool()
	req.Verbosity = o.Verbosity.ValueString()
	req.WebhookCredential = o.WebhookCredential.ValueInt64()
	req.WebhookService = o.WebhookService.ValueString()
	return
}

func (o *jobTemplateTerraformModel) setAllowSimultaneous(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.AllowSimultaneous, data)
}

func (o *jobTemplateTerraformModel) setAskCredentialOnLaunch(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.AskCredentialOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskDiffModeOnLaunch(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.AskDiffModeOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskExecutionEnvironmentOnLaunch(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.AskExecutionEnvironmentOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskForksOnLaunch(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.AskForksOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskInstanceGroupsOnLaunch(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.AskInstanceGroupsOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskInventoryOnLaunch(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.AskInventoryOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskJobSliceCountOnLaunch(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.AskJobSliceCountOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskJobTypeOnLaunch(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.AskJobTypeOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskLabelsOnLaunch(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.AskLabelsOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskLimitOnLaunch(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.AskLimitOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskScmBranchOnLaunch(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.AskScmBranchOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskSkipTagsOnLaunch(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.AskSkipTagsOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskTagsOnLaunch(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.AskTagsOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskTimeoutOnLaunch(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.AskTimeoutOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskVariablesOnLaunch(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.AskVariablesOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskVerbosityOnLaunch(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.AskVerbosityOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setBecomeEnabled(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.BecomeEnabled, data)
}

func (o *jobTemplateTerraformModel) setDescription(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *jobTemplateTerraformModel) setDiffMode(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.DiffMode, data)
}

func (o *jobTemplateTerraformModel) setExecutionEnvironment(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ExecutionEnvironment, data)
}

func (o *jobTemplateTerraformModel) setExtraVars(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetJsonString(&o.ExtraVars, data, false)
}

func (o *jobTemplateTerraformModel) setForceHandlers(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.ForceHandlers, data)
}

func (o *jobTemplateTerraformModel) setForks(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Forks, data)
}

func (o *jobTemplateTerraformModel) setHostConfigKey(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.HostConfigKey, data, false)
}

func (o *jobTemplateTerraformModel) setID(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *jobTemplateTerraformModel) setInventory(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Inventory, data)
}

func (o *jobTemplateTerraformModel) setJobSliceCount(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.JobSliceCount, data)
}

func (o *jobTemplateTerraformModel) setJobTags(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.JobTags, data, false)
}

func (o *jobTemplateTerraformModel) setJobType(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.JobType, data, false)
}

func (o *jobTemplateTerraformModel) setLimit(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Limit, data, false)
}

func (o *jobTemplateTerraformModel) setName(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *jobTemplateTerraformModel) setOrganization(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Organization, data)
}

func (o *jobTemplateTerraformModel) setPlaybook(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Playbook, data, false)
}

func (o *jobTemplateTerraformModel) setPreventInstanceGroupFallback(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.PreventInstanceGroupFallback, data)
}

func (o *jobTemplateTerraformModel) setProject(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Project, data)
}

func (o *jobTemplateTerraformModel) setScmBranch(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.ScmBranch, data, false)
}

func (o *jobTemplateTerraformModel) setSkipTags(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.SkipTags, data, false)
}

func (o *jobTemplateTerraformModel) setStartAtTask(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.StartAtTask, data, false)
}

func (o *jobTemplateTerraformModel) setSurveyEnabled(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.SurveyEnabled, data)
}

func (o *jobTemplateTerraformModel) setTimeout(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Timeout, data)
}

func (o *jobTemplateTerraformModel) setUseFactCache(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.UseFactCache, data)
}

func (o *jobTemplateTerraformModel) setVerbosity(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Verbosity, data, false)
}

func (o *jobTemplateTerraformModel) setWebhookCredential(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.WebhookCredential, data)
}

func (o *jobTemplateTerraformModel) setWebhookService(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.WebhookService, data, false)
}

func (o *jobTemplateTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setAllowSimultaneous(data["allow_simultaneous"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAskCredentialOnLaunch(data["ask_credential_on_launch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAskDiffModeOnLaunch(data["ask_diff_mode_on_launch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAskExecutionEnvironmentOnLaunch(data["ask_execution_environment_on_launch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAskForksOnLaunch(data["ask_forks_on_launch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAskInstanceGroupsOnLaunch(data["ask_instance_groups_on_launch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAskInventoryOnLaunch(data["ask_inventory_on_launch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAskJobSliceCountOnLaunch(data["ask_job_slice_count_on_launch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAskJobTypeOnLaunch(data["ask_job_type_on_launch"]); dg.HasError() {
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
	if dg, _ := o.setAskTimeoutOnLaunch(data["ask_timeout_on_launch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAskVariablesOnLaunch(data["ask_variables_on_launch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setAskVerbosityOnLaunch(data["ask_verbosity_on_launch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setBecomeEnabled(data["become_enabled"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDiffMode(data["diff_mode"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setExecutionEnvironment(data["execution_environment"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setExtraVars(data["extra_vars"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setForceHandlers(data["force_handlers"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setForks(data["forks"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setHostConfigKey(data["host_config_key"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setInventory(data["inventory"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setJobSliceCount(data["job_slice_count"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setJobTags(data["job_tags"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setJobType(data["job_type"]); dg.HasError() {
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
	if dg, _ := o.setPlaybook(data["playbook"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setPreventInstanceGroupFallback(data["prevent_instance_group_fallback"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setProject(data["project"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setScmBranch(data["scm_branch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSkipTags(data["skip_tags"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setStartAtTask(data["start_at_task"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSurveyEnabled(data["survey_enabled"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setTimeout(data["timeout"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setUseFactCache(data["use_fact_cache"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setVerbosity(data["verbosity"]); dg.HasError() {
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

// jobTemplateBodyRequestModel maps the schema for JobTemplate for creating and updating the data
type jobTemplateBodyRequestModel struct {
	// AllowSimultaneous ""
	AllowSimultaneous bool `json:"allow_simultaneous"`
	// AskCredentialOnLaunch ""
	AskCredentialOnLaunch bool `json:"ask_credential_on_launch"`
	// AskDiffModeOnLaunch ""
	AskDiffModeOnLaunch bool `json:"ask_diff_mode_on_launch"`
	// AskExecutionEnvironmentOnLaunch ""
	AskExecutionEnvironmentOnLaunch bool `json:"ask_execution_environment_on_launch"`
	// AskForksOnLaunch ""
	AskForksOnLaunch bool `json:"ask_forks_on_launch"`
	// AskInstanceGroupsOnLaunch ""
	AskInstanceGroupsOnLaunch bool `json:"ask_instance_groups_on_launch"`
	// AskInventoryOnLaunch ""
	AskInventoryOnLaunch bool `json:"ask_inventory_on_launch"`
	// AskJobSliceCountOnLaunch ""
	AskJobSliceCountOnLaunch bool `json:"ask_job_slice_count_on_launch"`
	// AskJobTypeOnLaunch ""
	AskJobTypeOnLaunch bool `json:"ask_job_type_on_launch"`
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
	// AskTimeoutOnLaunch ""
	AskTimeoutOnLaunch bool `json:"ask_timeout_on_launch"`
	// AskVariablesOnLaunch ""
	AskVariablesOnLaunch bool `json:"ask_variables_on_launch"`
	// AskVerbosityOnLaunch ""
	AskVerbosityOnLaunch bool `json:"ask_verbosity_on_launch"`
	// BecomeEnabled ""
	BecomeEnabled bool `json:"become_enabled"`
	// Description "Optional description of this job template."
	Description string `json:"description,omitempty"`
	// DiffMode "If enabled, textual changes made to any templated files on the host are shown in the standard output"
	DiffMode bool `json:"diff_mode"`
	// ExecutionEnvironment "The container image to be used for execution."
	ExecutionEnvironment int64 `json:"execution_environment,omitempty"`
	// ExtraVars ""
	ExtraVars json.RawMessage `json:"extra_vars,omitempty"`
	// ForceHandlers ""
	ForceHandlers bool `json:"force_handlers"`
	// Forks ""
	Forks int64 `json:"forks,omitempty"`
	// HostConfigKey ""
	HostConfigKey string `json:"host_config_key,omitempty"`
	// Inventory ""
	Inventory int64 `json:"inventory,omitempty"`
	// JobSliceCount "The number of jobs to slice into at runtime. Will cause the Job Template to launch a workflow if value is greater than 1."
	JobSliceCount int64 `json:"job_slice_count,omitempty"`
	// JobTags ""
	JobTags string `json:"job_tags,omitempty"`
	// JobType ""
	JobType string `json:"job_type,omitempty"`
	// Limit ""
	Limit string `json:"limit,omitempty"`
	// Name "Name of this job template."
	Name string `json:"name"`
	// Playbook ""
	Playbook string `json:"playbook,omitempty"`
	// PreventInstanceGroupFallback "If enabled, the job template will prevent adding any inventory or organization instance groups to the list of preferred instances groups to run on.If this setting is enabled and you provided an empty list, the global instance groups will be applied."
	PreventInstanceGroupFallback bool `json:"prevent_instance_group_fallback"`
	// Project ""
	Project int64 `json:"project,omitempty"`
	// ScmBranch "Branch to use in job run. Project default used if blank. Only allowed if project allow_override field is set to true."
	ScmBranch string `json:"scm_branch,omitempty"`
	// SkipTags ""
	SkipTags string `json:"skip_tags,omitempty"`
	// StartAtTask ""
	StartAtTask string `json:"start_at_task,omitempty"`
	// SurveyEnabled ""
	SurveyEnabled bool `json:"survey_enabled"`
	// Timeout "The amount of time (in seconds) to run before the task is canceled."
	Timeout int64 `json:"timeout,omitempty"`
	// UseFactCache "If enabled, the service will act as an Ansible Fact Cache Plugin; persisting facts at the end of a playbook run to the database and caching facts for use by Ansible."
	UseFactCache bool `json:"use_fact_cache"`
	// Verbosity ""
	Verbosity string `json:"verbosity,omitempty"`
	// WebhookCredential "Personal Access Token for posting back the status to the service API"
	WebhookCredential int64 `json:"webhook_credential,omitempty"`
	// WebhookService "Service that webhook requests will be accepted from"
	WebhookService string `json:"webhook_service,omitempty"`
}

type jobTemplateObjectRolesModel struct {
	ID    types.Int64 `tfsdk:"id"`
	Roles types.Map   `tfsdk:"roles"`
}
