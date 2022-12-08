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
	"strings"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/mitchellh/mapstructure"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/schemavalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
	req.ExtraVars = json.RawMessage(o.ExtraVars.ValueString())
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

func (o *jobTemplateTerraformModel) setAllowSimultaneous(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AllowSimultaneous, data)
}

func (o *jobTemplateTerraformModel) setAskCredentialOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskCredentialOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskDiffModeOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskDiffModeOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskExecutionEnvironmentOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskExecutionEnvironmentOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskForksOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskForksOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskInstanceGroupsOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskInstanceGroupsOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskInventoryOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskInventoryOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskJobSliceCountOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskJobSliceCountOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskJobTypeOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskJobTypeOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskLabelsOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskLabelsOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskLimitOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskLimitOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskScmBranchOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskScmBranchOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskSkipTagsOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskSkipTagsOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskTagsOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskTagsOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskTimeoutOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskTimeoutOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskVariablesOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskVariablesOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setAskVerbosityOnLaunch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.AskVerbosityOnLaunch, data)
}

func (o *jobTemplateTerraformModel) setBecomeEnabled(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.BecomeEnabled, data)
}

func (o *jobTemplateTerraformModel) setDescription(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *jobTemplateTerraformModel) setDiffMode(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.DiffMode, data)
}

func (o *jobTemplateTerraformModel) setExecutionEnvironment(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ExecutionEnvironment, data)
}

func (o *jobTemplateTerraformModel) setExtraVars(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.ExtraVars, data, false)
}

func (o *jobTemplateTerraformModel) setForceHandlers(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.ForceHandlers, data)
}

func (o *jobTemplateTerraformModel) setForks(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Forks, data)
}

func (o *jobTemplateTerraformModel) setHostConfigKey(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.HostConfigKey, data, false)
}

func (o *jobTemplateTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *jobTemplateTerraformModel) setInventory(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Inventory, data)
}

func (o *jobTemplateTerraformModel) setJobSliceCount(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.JobSliceCount, data)
}

func (o *jobTemplateTerraformModel) setJobTags(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.JobTags, data, false)
}

func (o *jobTemplateTerraformModel) setJobType(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.JobType, data, false)
}

func (o *jobTemplateTerraformModel) setLimit(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Limit, data, false)
}

func (o *jobTemplateTerraformModel) setName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *jobTemplateTerraformModel) setOrganization(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Organization, data)
}

func (o *jobTemplateTerraformModel) setPlaybook(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Playbook, data, false)
}

func (o *jobTemplateTerraformModel) setPreventInstanceGroupFallback(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.PreventInstanceGroupFallback, data)
}

func (o *jobTemplateTerraformModel) setProject(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Project, data)
}

func (o *jobTemplateTerraformModel) setScmBranch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.ScmBranch, data, false)
}

func (o *jobTemplateTerraformModel) setSkipTags(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SkipTags, data, false)
}

func (o *jobTemplateTerraformModel) setStartAtTask(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.StartAtTask, data, false)
}

func (o *jobTemplateTerraformModel) setSurveyEnabled(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.SurveyEnabled, data)
}

func (o *jobTemplateTerraformModel) setTimeout(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Timeout, data)
}

func (o *jobTemplateTerraformModel) setUseFactCache(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.UseFactCache, data)
}

func (o *jobTemplateTerraformModel) setVerbosity(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Verbosity, data, false)
}

func (o *jobTemplateTerraformModel) setWebhookCredential(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.WebhookCredential, data)
}

func (o *jobTemplateTerraformModel) setWebhookService(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.WebhookService, data, false)
}

func (o *jobTemplateTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
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

var (
	_ datasource.DataSource              = &jobTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &jobTemplateDataSource{}
)

// NewJobTemplateDataSource is a helper function to instantiate the JobTemplate data source.
func NewJobTemplateDataSource() datasource.DataSource {
	return &jobTemplateDataSource{}
}

// jobTemplateDataSource is the data source implementation.
type jobTemplateDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *jobTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/job_templates/"
}

// Metadata returns the data source type name.
func (o *jobTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_job_template"
}

// GetSchema defines the schema for the data source.
func (o *jobTemplateDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"JobTemplate",
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
				"ask_credential_on_launch": {
					Description: "Ask credential on launch",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"ask_diff_mode_on_launch": {
					Description: "Ask diff mode on launch",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"ask_execution_environment_on_launch": {
					Description: "Ask execution environment on launch",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"ask_forks_on_launch": {
					Description: "Ask forks on launch",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"ask_instance_groups_on_launch": {
					Description: "Ask instance groups on launch",
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
				"ask_job_slice_count_on_launch": {
					Description: "Ask job slice count on launch",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"ask_job_type_on_launch": {
					Description: "Ask job type on launch",
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
				"ask_timeout_on_launch": {
					Description: "Ask timeout on launch",
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
				"ask_verbosity_on_launch": {
					Description: "Ask verbosity on launch",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"become_enabled": {
					Description: "Become enabled",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"description": {
					Description: "Optional description of this job template.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"diff_mode": {
					Description: "If enabled, textual changes made to any templated files on the host are shown in the standard output",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"execution_environment": {
					Description: "The container image to be used for execution.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"extra_vars": {
					Description: "Extra vars",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"force_handlers": {
					Description: "Force handlers",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"forks": {
					Description: "Forks",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"host_config_key": {
					Description: "Host config key",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"id": {
					Description: "Database ID for this job template.",
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
					Description: "Inventory",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"job_slice_count": {
					Description: "The number of jobs to slice into at runtime. Will cause the Job Template to launch a workflow if value is greater than 1.",
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
				"job_type": {
					Description: "Job type",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"run", "check"}...),
					},
				},
				"limit": {
					Description: "Limit",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"name": {
					Description: "Name of this job template.",
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
				"playbook": {
					Description: "Playbook",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"prevent_instance_group_fallback": {
					Description: "If enabled, the job template will prevent adding any inventory or organization instance groups to the list of preferred instances groups to run on.If this setting is enabled and you provided an empty list, the global instance groups will be applied.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"project": {
					Description: "Project",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"scm_branch": {
					Description: "Branch to use in job run. Project default used if blank. Only allowed if project allow_override field is set to true.",
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
				"start_at_task": {
					Description: "Start at task",
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
				"timeout": {
					Description: "The amount of time (in seconds) to run before the task is canceled.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"use_fact_cache": {
					Description: "If enabled, the service will act as an Ansible Fact Cache Plugin; persisting facts at the end of a playbook run to the database and caching facts for use by Ansible.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"verbosity": {
					Description: "Verbosity",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"0", "1", "2", "3", "4", "5"}...),
					},
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
func (o *jobTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state jobTemplateTerraformModel
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

	// Creates a new request for JobTemplate
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for JobTemplate
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for JobTemplate on %s", o.endpoint),
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
	if err = hookJobTemplate(ctx, ApiVersion, SourceData, CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			"Unable to process custom hook for the state on JobTemplate",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &jobTemplateResource{}
	_ resource.ResourceWithConfigure   = &jobTemplateResource{}
	_ resource.ResourceWithImportState = &jobTemplateResource{}
)

// NewJobTemplateResource is a helper function to simplify the provider implementation.
func NewJobTemplateResource() resource.Resource {
	return &jobTemplateResource{}
}

type jobTemplateResource struct {
	client   c.Client
	endpoint string
}

func (o *jobTemplateResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/job_templates/"
}

func (o *jobTemplateResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_job_template"
}

func (o *jobTemplateResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"JobTemplate",
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
				"ask_credential_on_launch": {
					Description: "Ask credential on launch",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"ask_diff_mode_on_launch": {
					Description: "Ask diff mode on launch",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"ask_execution_environment_on_launch": {
					Description: "Ask execution environment on launch",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"ask_forks_on_launch": {
					Description: "Ask forks on launch",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"ask_instance_groups_on_launch": {
					Description: "Ask instance groups on launch",
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
				"ask_job_slice_count_on_launch": {
					Description: "Ask job slice count on launch",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"ask_job_type_on_launch": {
					Description: "Ask job type on launch",
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
				"ask_timeout_on_launch": {
					Description: "Ask timeout on launch",
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
				"ask_verbosity_on_launch": {
					Description: "Ask verbosity on launch",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"become_enabled": {
					Description: "Become enabled",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"description": {
					Description: "Optional description of this job template.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"diff_mode": {
					Description: "If enabled, textual changes made to any templated files on the host are shown in the standard output",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"execution_environment": {
					Description: "The container image to be used for execution.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
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
				"force_handlers": {
					Description: "Force handlers",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"forks": {
					Description: "Forks",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.Int64Value(0)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						int64validator.Between(0, 2.147483647e+09),
					},
				},
				"host_config_key": {
					Description: "Host config key",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(1024),
					},
				},
				"inventory": {
					Description: "Inventory",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"job_slice_count": {
					Description: "The number of jobs to slice into at runtime. Will cause the Job Template to launch a workflow if value is greater than 1.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.Int64Value(1)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						int64validator.Between(0, 2.147483647e+09),
					},
				},
				"job_tags": {
					Description: "Job tags",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"job_type": {
					Description: "Job type",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`run`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"run", "check"}...),
					},
				},
				"limit": {
					Description: "Limit",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"name": {
					Description:   "Name of this job template.",
					Type:          types.StringType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(512),
					},
				},
				"playbook": {
					Description: "Playbook",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(1024),
					},
				},
				"prevent_instance_group_fallback": {
					Description: "If enabled, the job template will prevent adding any inventory or organization instance groups to the list of preferred instances groups to run on.If this setting is enabled and you provided an empty list, the global instance groups will be applied.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"project": {
					Description: "Project",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"scm_branch": {
					Description: "Branch to use in job run. Project default used if blank. Only allowed if project allow_override field is set to true.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(1024),
					},
				},
				"skip_tags": {
					Description: "Skip tags",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(1024),
					},
				},
				"start_at_task": {
					Description: "Start at task",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(1024),
					},
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
				"timeout": {
					Description: "The amount of time (in seconds) to run before the task is canceled.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.Int64Value(0)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						int64validator.Between(-2.147483648e+09, 2.147483647e+09),
					},
				},
				"use_fact_cache": {
					Description: "If enabled, the service will act as an Ansible Fact Cache Plugin; persisting facts at the end of a playbook run to the database and caching facts for use by Ansible.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"verbosity": {
					Description: "Verbosity",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`0`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"0", "1", "2", "3", "4", "5"}...),
					},
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
					Description: "Database ID for this job template.",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"organization": {
					Description: "The organization used to determine access to this template.",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
			},
		}), nil
}

func (o *jobTemplateResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the JobTemplate.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (o *jobTemplateResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state jobTemplateTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for JobTemplate
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[JobTemplate/Create] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new JobTemplate resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for JobTemplate on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookJobTemplate(ctx, ApiVersion, SourceResource, CalleeCreate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on JobTemplate",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *jobTemplateResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state jobTemplateTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
	var orig = state.Clone()

	// Creates a new request for JobTemplate
	var r *http.Request
	var id = state.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for JobTemplate from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for JobTemplate on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookJobTemplate(ctx, ApiVersion, SourceResource, CalleeRead, &orig, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on JobTemplate",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *jobTemplateResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state jobTemplateTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for JobTemplate
	var r *http.Request
	var id = plan.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[JobTemplate/Update] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new JobTemplate resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for JobTemplate on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookJobTemplate(ctx, ApiVersion, SourceResource, CalleeUpdate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on JobTemplate",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *jobTemplateResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state jobTemplateTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for JobTemplate
	var r *http.Request
	var id = state.ID
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	// Delete existing JobTemplate
	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to delete resource for JobTemplate on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}

var (
	_ datasource.DataSource              = &jobTemplateObjectRolesDataSource{}
	_ datasource.DataSourceWithConfigure = &jobTemplateObjectRolesDataSource{}
)

// NewJobTemplateObjectRolesDataSource is a helper function to instantiate the JobTemplate data source.
func NewJobTemplateObjectRolesDataSource() datasource.DataSource {
	return &jobTemplateObjectRolesDataSource{}
}

// jobTemplateObjectRolesDataSource is the data source implementation.
type jobTemplateObjectRolesDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *jobTemplateObjectRolesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/job_templates/%d/object_roles/"
}

// Metadata returns the data source type name.
func (o *jobTemplateObjectRolesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_job_template_object_roles"
}

// GetSchema defines the schema for the data source.
func (o *jobTemplateObjectRolesDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Version: helpers.SchemaVersion,
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description: "JobTemplate ID",
				Type:        types.Int64Type,
				Required:    true,
			},
			"roles": {
				Description: "Roles for jobtemplate",
				Type:        types.MapType{ElemType: types.Int64Type},
				Computed:    true,
			},
		},
	}, nil
}

// Read refreshes the Terraform state with the latest data.
func (o *jobTemplateObjectRolesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state jobTemplateObjectRolesModel
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
			"Unable to create a new request for jobTemplate",
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Credential
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			"Unable to fetch the request for jobtemplate object roles",
			err.Error(),
		)
		return
	}

	var sr searchResultObjectRole
	if err = mapstructure.Decode(data, &sr); err != nil {
		resp.Diagnostics.AddError(
			"Unable to decode the search result data for jobtemplate",
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
	_ resource.Resource                = &jobTemplateAssociateDisassociateCredential{}
	_ resource.ResourceWithConfigure   = &jobTemplateAssociateDisassociateCredential{}
	_ resource.ResourceWithImportState = &jobTemplateAssociateDisassociateCredential{}
)

type jobTemplateAssociateDisassociateCredentialTerraformModel struct {
	JobTemplateID types.Int64 `tfsdk:"job_template_id"`
	CredentialID  types.Int64 `tfsdk:"credential_id"`
}

// NewJobTemplateAssociateDisassociateCredentialResource is a helper function to simplify the provider implementation.
func NewJobTemplateAssociateDisassociateCredentialResource() resource.Resource {
	return &jobTemplateAssociateDisassociateCredential{}
}

type jobTemplateAssociateDisassociateCredential struct {
	client   c.Client
	endpoint string
}

func (o *jobTemplateAssociateDisassociateCredential) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/job_templates/%d/credentials/"
}

func (o jobTemplateAssociateDisassociateCredential) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_job_template_associate_credential"
}

func (o jobTemplateAssociateDisassociateCredential) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"JobTemplate/Associate",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				"job_template_id": {
					Description: "Database ID for this JobTemplate.",
					Required:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.RequiresReplace(),
					},
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.AlsoRequires(
							path.MatchRoot("credential_id"),
						),
					},
				},
				"credential_id": {
					Description: "Database ID of the credential to assign.",
					Required:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.RequiresReplace(),
					},
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.AlsoRequires(
							path.MatchRoot("job_template_id"),
						),
					},
				},
			},
		},
	), nil
}

func (o *jobTemplateAssociateDisassociateCredential) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var state jobTemplateAssociateDisassociateCredentialTerraformModel
	var parts = strings.Split(request.ID, "/")
	var err error
	if len(parts) != 2 {
		err = fmt.Errorf("requires the identifier to be set to <job_template_id>/<credential_id>, currently set to %s", request.ID)
		response.Diagnostics.AddError(
			"Unable to import state for JobTemplate association, invalid format.",
			err.Error(),
		)
		return
	}

	var jobTemplateId, credentialId int64

	jobTemplateId, err = strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the job_templateId for the JobTemplate association.", request.ID),
			err.Error(),
		)
		return
	}
	state.JobTemplateID = types.Int64Value(jobTemplateId)

	credentialId, err = strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the credential_id for the JobTemplate association.", request.ID),
			err.Error(),
		)
		return
	}
	state.CredentialID = types.Int64Value(credentialId)

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (o *jobTemplateAssociateDisassociateCredential) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error

	// Retrieve values from state
	var plan, state jobTemplateAssociateDisassociateCredentialTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for association of JobTemplate
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.JobTemplateID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = associateDisassociateRequestModel{ID: plan.CredentialID.ValueInt64(), Disassociate: false}
	tflog.Debug(ctx, "[JobTemplate/Create/Associate] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for create of type ", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to associate for JobTemplate on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	state.JobTemplateID = plan.JobTemplateID
	state.CredentialID = plan.CredentialID

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *jobTemplateAssociateDisassociateCredential) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state jobTemplateAssociateDisassociateCredentialTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for disassociation of JobTemplate
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.JobTemplateID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = associateDisassociateRequestModel{ID: state.CredentialID.ValueInt64(), Disassociate: true}
	tflog.Debug(ctx, "[JobTemplate/Delete/Disassociate] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for delete of type ", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to disassociate for JobTemplate on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}

func (o *jobTemplateAssociateDisassociateCredential) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
}

func (o *jobTemplateAssociateDisassociateCredential) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
}

var (
	_ resource.Resource              = &jobTemplateAssociateDisassociateNotificationTemplate{}
	_ resource.ResourceWithConfigure = &jobTemplateAssociateDisassociateNotificationTemplate{}
)

type jobTemplateAssociateDisassociateNotificationTemplateTerraformModel struct {
	JobTemplateID          types.Int64  `tfsdk:"job_template_id"`
	NotificationTemplateID types.Int64  `tfsdk:"notification_template_id"`
	Option                 types.String `tfsdk:"option"`
}

// NewJobTemplateAssociateDisassociateNotificationTemplateResource is a helper function to simplify the provider implementation.
func NewJobTemplateAssociateDisassociateNotificationTemplateResource() resource.Resource {
	return &jobTemplateAssociateDisassociateNotificationTemplate{}
}

type jobTemplateAssociateDisassociateNotificationTemplate struct {
	client   c.Client
	endpoint string
}

func (o *jobTemplateAssociateDisassociateNotificationTemplate) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/job_templates/%d/notification_templates_%s/"
}

func (o jobTemplateAssociateDisassociateNotificationTemplate) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_job_template_associate_notification_template"
}

func (o jobTemplateAssociateDisassociateNotificationTemplate) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"JobTemplate/Associate",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				"job_template_id": {
					Description: "Database ID for this JobTemplate.",
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
							path.MatchRoot("job_template_id"),
						),
						stringvalidator.OneOf([]string{"started", "success", "error"}...),
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
							path.MatchRoot("job_template_id"),
						),
					},
				},
			},
		},
	), nil
}

func (o *jobTemplateAssociateDisassociateNotificationTemplate) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error

	// Retrieve values from state
	var plan, state jobTemplateAssociateDisassociateNotificationTemplateTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for association of JobTemplate
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.JobTemplateID.ValueInt64(), plan.Option.ValueString())) + "/"
	var buf bytes.Buffer
	var bodyRequest = associateDisassociateRequestModel{ID: plan.NotificationTemplateID.ValueInt64(), Disassociate: false}
	tflog.Debug(ctx, "[JobTemplate/Create/Associate] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for create of type notification_job_template", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to associate for JobTemplate on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	state.JobTemplateID = plan.JobTemplateID
	state.NotificationTemplateID = plan.NotificationTemplateID
	state.Option = plan.Option

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *jobTemplateAssociateDisassociateNotificationTemplate) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state jobTemplateAssociateDisassociateNotificationTemplateTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for disassociation of JobTemplate
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.JobTemplateID.ValueInt64(), state.Option.ValueString())) + "/"
	var buf bytes.Buffer
	var bodyRequest = associateDisassociateRequestModel{ID: state.NotificationTemplateID.ValueInt64(), Disassociate: true}
	tflog.Debug(ctx, "[JobTemplate/Delete/Disassociate] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for delete of type notification_job_template", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to disassociate for JobTemplate on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}

func (o *jobTemplateAssociateDisassociateNotificationTemplate) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
}

func (o *jobTemplateAssociateDisassociateNotificationTemplate) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
}

var (
	_ resource.Resource                = &jobTemplateSurvey{}
	_ resource.ResourceWithConfigure   = &jobTemplateSurvey{}
	_ resource.ResourceWithImportState = &jobTemplateSurvey{}
)

type jobTemplateSurveyTerraformModel struct {
	JobTemplateID types.Int64  `tfsdk:"job_template_id"`
	Spec          types.String `tfsdk:"spec"`
}

func (o jobTemplateSurveyTerraformModel) Clone() jobTemplateSurveyTerraformModel {
	return jobTemplateSurveyTerraformModel{
		JobTemplateID: types.Int64Value(o.JobTemplateID.ValueInt64()),
		Spec:          types.StringValue(o.Spec.ValueString()),
	}
}

func (o jobTemplateSurveyTerraformModel) BodyRequest() jobTemplateSurveyModel {
	return jobTemplateSurveyModel{
		Spec: json.RawMessage(o.Spec.ValueString()),
	}
}

type jobTemplateSurveyModel struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Spec        json.RawMessage `json:"spec"`
}

// NewJobTemplateSurveyResource is a helper function to simplify the provider implementation.
func NewJobTemplateSurveyResource() resource.Resource {
	return &jobTemplateSurvey{}
}

type jobTemplateSurvey struct {
	client   c.Client
	endpoint string
}

func (o *jobTemplateSurvey) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/job_templates/%d/survey_spec/"
}

func (o jobTemplateSurvey) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_job_template_survey_spec"
}

func (o jobTemplateSurvey) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"JobTemplate/Survey",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				"job_template_id": {
					Description: "Database ID for this JobTemplate.",
					Required:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.RequiresReplace(),
					},
				},
				"spec": {
					Description: "The survey spec for this JobTemplate.",
					Required:    true,
					Type:        types.StringType,
				},
			},
		}), nil
}

// ImportState imports the survey spec for JobTemplate
func (o *jobTemplateSurvey) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the JobTemplate.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("job_template_id"), id)...)
}

// Delete the survey spec for JobTemplate
func (o *jobTemplateSurvey) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	var state jobTemplateSurveyTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.JobTemplateID.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to delete resource for JobTemplate/Survey on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}

// Read the survey spec for JobTemplate
func (o *jobTemplateSurvey) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error
	var state jobTemplateSurveyTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.JobTemplateID.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for JobTemplate/Survey on %s", o.endpoint),
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

// Create the survey spec for JobTemplate
func (o *jobTemplateSurvey) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state jobTemplateSurveyTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.JobTemplateID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[JobTemplate/Create/Survey] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for JobTemplate/Survey on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	state.Spec = types.StringValue(plan.Spec.ValueString())
	state.JobTemplateID = types.Int64Value(plan.JobTemplateID.ValueInt64())
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

// Update the survey spec for JobTemplate
func (o *jobTemplateSurvey) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state jobTemplateSurveyTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.JobTemplateID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[JobTemplate/Update/SurveySpec] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for JobTemplate/Survey on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	state.Spec = types.StringValue(plan.Spec.ValueString())
	state.JobTemplateID = types.Int64Value(plan.JobTemplateID.ValueInt64())
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}
