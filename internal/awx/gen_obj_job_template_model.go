package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type jobTemplateTerraformModel struct {
	AllowSimultaneous               types.Bool   `tfsdk:"allow_simultaneous" json:"allow_simultaneous"`
	AskCredentialOnLaunch           types.Bool   `tfsdk:"ask_credential_on_launch" json:"ask_credential_on_launch"`
	AskDiffModeOnLaunch             types.Bool   `tfsdk:"ask_diff_mode_on_launch" json:"ask_diff_mode_on_launch"`
	AskExecutionEnvironmentOnLaunch types.Bool   `tfsdk:"ask_execution_environment_on_launch" json:"ask_execution_environment_on_launch"`
	AskForksOnLaunch                types.Bool   `tfsdk:"ask_forks_on_launch" json:"ask_forks_on_launch"`
	AskInstanceGroupsOnLaunch       types.Bool   `tfsdk:"ask_instance_groups_on_launch" json:"ask_instance_groups_on_launch"`
	AskInventoryOnLaunch            types.Bool   `tfsdk:"ask_inventory_on_launch" json:"ask_inventory_on_launch"`
	AskJobSliceCountOnLaunch        types.Bool   `tfsdk:"ask_job_slice_count_on_launch" json:"ask_job_slice_count_on_launch"`
	AskJobTypeOnLaunch              types.Bool   `tfsdk:"ask_job_type_on_launch" json:"ask_job_type_on_launch"`
	AskLabelsOnLaunch               types.Bool   `tfsdk:"ask_labels_on_launch" json:"ask_labels_on_launch"`
	AskLimitOnLaunch                types.Bool   `tfsdk:"ask_limit_on_launch" json:"ask_limit_on_launch"`
	AskScmBranchOnLaunch            types.Bool   `tfsdk:"ask_scm_branch_on_launch" json:"ask_scm_branch_on_launch"`
	AskSkipTagsOnLaunch             types.Bool   `tfsdk:"ask_skip_tags_on_launch" json:"ask_skip_tags_on_launch"`
	AskTagsOnLaunch                 types.Bool   `tfsdk:"ask_tags_on_launch" json:"ask_tags_on_launch"`
	AskTimeoutOnLaunch              types.Bool   `tfsdk:"ask_timeout_on_launch" json:"ask_timeout_on_launch"`
	AskVariablesOnLaunch            types.Bool   `tfsdk:"ask_variables_on_launch" json:"ask_variables_on_launch"`
	AskVerbosityOnLaunch            types.Bool   `tfsdk:"ask_verbosity_on_launch" json:"ask_verbosity_on_launch"`
	BecomeEnabled                   types.Bool   `tfsdk:"become_enabled" json:"become_enabled"`
	Description                     types.String `tfsdk:"description" json:"description"`
	DiffMode                        types.Bool   `tfsdk:"diff_mode" json:"diff_mode"`
	ExecutionEnvironment            types.Int64  `tfsdk:"execution_environment" json:"execution_environment"`
	ExtraVars                       types.String `tfsdk:"extra_vars" json:"extra_vars"`
	ForceHandlers                   types.Bool   `tfsdk:"force_handlers" json:"force_handlers"`
	Forks                           types.Int64  `tfsdk:"forks" json:"forks"`
	HostConfigKey                   types.String `tfsdk:"host_config_key" json:"host_config_key"`
	ID                              types.Int64  `tfsdk:"id" json:"id"`
	Inventory                       types.Int64  `tfsdk:"inventory" json:"inventory"`
	JobSliceCount                   types.Int64  `tfsdk:"job_slice_count" json:"job_slice_count"`
	JobTags                         types.String `tfsdk:"job_tags" json:"job_tags"`
	JobType                         types.String `tfsdk:"job_type" json:"job_type"`
	Limit                           types.String `tfsdk:"limit" json:"limit"`
	Name                            types.String `tfsdk:"name" json:"name"`
	Organization                    types.Int64  `tfsdk:"organization" json:"organization"`
	Playbook                        types.String `tfsdk:"playbook" json:"playbook"`
	PreventInstanceGroupFallback    types.Bool   `tfsdk:"prevent_instance_group_fallback" json:"prevent_instance_group_fallback"`
	Project                         types.Int64  `tfsdk:"project" json:"project"`
	ScmBranch                       types.String `tfsdk:"scm_branch" json:"scm_branch"`
	SkipTags                        types.String `tfsdk:"skip_tags" json:"skip_tags"`
	StartAtTask                     types.String `tfsdk:"start_at_task" json:"start_at_task"`
	SurveyEnabled                   types.Bool   `tfsdk:"survey_enabled" json:"survey_enabled"`
	Timeout                         types.Int64  `tfsdk:"timeout" json:"timeout"`
	UseFactCache                    types.Bool   `tfsdk:"use_fact_cache" json:"use_fact_cache"`
	Verbosity                       types.String `tfsdk:"verbosity" json:"verbosity"`
	WebhookCredential               types.Int64  `tfsdk:"webhook_credential" json:"webhook_credential"`
	WebhookService                  types.String `tfsdk:"webhook_service" json:"webhook_service"`
}

func (o *jobTemplateTerraformModel) Clone() jobTemplateTerraformModel {
	return *o
}

func (o *jobTemplateTerraformModel) BodyRequest() *jobTemplateBodyRequestModel {
	var req jobTemplateBodyRequestModel
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
	return &req
}

func (o *jobTemplateTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetBool(&o.AllowSimultaneous, data["allow_simultaneous"]))
	collect(helpers.AttrValueSetBool(&o.AskCredentialOnLaunch, data["ask_credential_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskDiffModeOnLaunch, data["ask_diff_mode_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskExecutionEnvironmentOnLaunch, data["ask_execution_environment_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskForksOnLaunch, data["ask_forks_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskInstanceGroupsOnLaunch, data["ask_instance_groups_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskInventoryOnLaunch, data["ask_inventory_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskJobSliceCountOnLaunch, data["ask_job_slice_count_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskJobTypeOnLaunch, data["ask_job_type_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskLabelsOnLaunch, data["ask_labels_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskLimitOnLaunch, data["ask_limit_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskScmBranchOnLaunch, data["ask_scm_branch_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskSkipTagsOnLaunch, data["ask_skip_tags_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskTagsOnLaunch, data["ask_tags_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskTimeoutOnLaunch, data["ask_timeout_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskVariablesOnLaunch, data["ask_variables_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.AskVerbosityOnLaunch, data["ask_verbosity_on_launch"]))
	collect(helpers.AttrValueSetBool(&o.BecomeEnabled, data["become_enabled"]))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetBool(&o.DiffMode, data["diff_mode"]))
	collect(helpers.AttrValueSetInt64(&o.ExecutionEnvironment, data["execution_environment"]))
	collect(helpers.AttrValueSetJsonString(&o.ExtraVars, data["extra_vars"], false))
	collect(helpers.AttrValueSetBool(&o.ForceHandlers, data["force_handlers"]))
	collect(helpers.AttrValueSetInt64(&o.Forks, data["forks"]))
	collect(helpers.AttrValueSetString(&o.HostConfigKey, data["host_config_key"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetInt64(&o.Inventory, data["inventory"]))
	collect(helpers.AttrValueSetInt64(&o.JobSliceCount, data["job_slice_count"]))
	collect(helpers.AttrValueSetString(&o.JobTags, data["job_tags"], false))
	collect(helpers.AttrValueSetString(&o.JobType, data["job_type"], false))
	collect(helpers.AttrValueSetString(&o.Limit, data["limit"], false))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	collect(helpers.AttrValueSetString(&o.Playbook, data["playbook"], false))
	collect(helpers.AttrValueSetBool(&o.PreventInstanceGroupFallback, data["prevent_instance_group_fallback"]))
	collect(helpers.AttrValueSetInt64(&o.Project, data["project"]))
	collect(helpers.AttrValueSetString(&o.ScmBranch, data["scm_branch"], false))
	collect(helpers.AttrValueSetString(&o.SkipTags, data["skip_tags"], false))
	collect(helpers.AttrValueSetString(&o.StartAtTask, data["start_at_task"], false))
	collect(helpers.AttrValueSetBool(&o.SurveyEnabled, data["survey_enabled"]))
	collect(helpers.AttrValueSetInt64(&o.Timeout, data["timeout"]))
	collect(helpers.AttrValueSetBool(&o.UseFactCache, data["use_fact_cache"]))
	collect(helpers.AttrValueSetString(&o.Verbosity, data["verbosity"], false))
	collect(helpers.AttrValueSetInt64(&o.WebhookCredential, data["webhook_credential"]))
	collect(helpers.AttrValueSetString(&o.WebhookService, data["webhook_service"], false))
	return diags, nil
}

type jobTemplateBodyRequestModel struct {
	AllowSimultaneous               bool            `json:"allow_simultaneous"`
	AskCredentialOnLaunch           bool            `json:"ask_credential_on_launch"`
	AskDiffModeOnLaunch             bool            `json:"ask_diff_mode_on_launch"`
	AskExecutionEnvironmentOnLaunch bool            `json:"ask_execution_environment_on_launch"`
	AskForksOnLaunch                bool            `json:"ask_forks_on_launch"`
	AskInstanceGroupsOnLaunch       bool            `json:"ask_instance_groups_on_launch"`
	AskInventoryOnLaunch            bool            `json:"ask_inventory_on_launch"`
	AskJobSliceCountOnLaunch        bool            `json:"ask_job_slice_count_on_launch"`
	AskJobTypeOnLaunch              bool            `json:"ask_job_type_on_launch"`
	AskLabelsOnLaunch               bool            `json:"ask_labels_on_launch"`
	AskLimitOnLaunch                bool            `json:"ask_limit_on_launch"`
	AskScmBranchOnLaunch            bool            `json:"ask_scm_branch_on_launch"`
	AskSkipTagsOnLaunch             bool            `json:"ask_skip_tags_on_launch"`
	AskTagsOnLaunch                 bool            `json:"ask_tags_on_launch"`
	AskTimeoutOnLaunch              bool            `json:"ask_timeout_on_launch"`
	AskVariablesOnLaunch            bool            `json:"ask_variables_on_launch"`
	AskVerbosityOnLaunch            bool            `json:"ask_verbosity_on_launch"`
	BecomeEnabled                   bool            `json:"become_enabled"`
	Description                     string          `json:"description,omitempty"`
	DiffMode                        bool            `json:"diff_mode"`
	ExecutionEnvironment            int64           `json:"execution_environment,omitempty"`
	ExtraVars                       json.RawMessage `json:"extra_vars,omitempty"`
	ForceHandlers                   bool            `json:"force_handlers"`
	Forks                           int64           `json:"forks,omitempty"`
	HostConfigKey                   string          `json:"host_config_key,omitempty"`
	Inventory                       int64           `json:"inventory,omitempty"`
	JobSliceCount                   int64           `json:"job_slice_count,omitempty"`
	JobTags                         string          `json:"job_tags,omitempty"`
	JobType                         string          `json:"job_type,omitempty"`
	Limit                           string          `json:"limit,omitempty"`
	Name                            string          `json:"name"`
	Playbook                        string          `json:"playbook,omitempty"`
	PreventInstanceGroupFallback    bool            `json:"prevent_instance_group_fallback"`
	Project                         int64           `json:"project,omitempty"`
	ScmBranch                       string          `json:"scm_branch,omitempty"`
	SkipTags                        string          `json:"skip_tags,omitempty"`
	StartAtTask                     string          `json:"start_at_task,omitempty"`
	SurveyEnabled                   bool            `json:"survey_enabled"`
	Timeout                         int64           `json:"timeout,omitempty"`
	UseFactCache                    bool            `json:"use_fact_cache"`
	Verbosity                       string          `json:"verbosity,omitempty"`
	WebhookCredential               int64           `json:"webhook_credential,omitempty"`
	WebhookService                  string          `json:"webhook_service,omitempty"`
}
