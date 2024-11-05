package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// scheduleTerraformModel maps the schema for Schedule when using Data Source
type scheduleTerraformModel struct {
	// Description "Optional description of this schedule."
	Description types.String `tfsdk:"description" json:"description"`
	// DiffMode ""
	DiffMode types.Bool `tfsdk:"diff_mode" json:"diff_mode"`
	// Dtend "The last occurrence of the schedule occurs before this time, aftewards the schedule expires."
	Dtend types.String `tfsdk:"dtend" json:"dtend"`
	// Dtstart "The first occurrence of the schedule occurs on or after this time."
	Dtstart types.String `tfsdk:"dtstart" json:"dtstart"`
	// Enabled "Enables processing of this schedule."
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
	// ExecutionEnvironment "The container image to be used for execution."
	ExecutionEnvironment types.Int64 `tfsdk:"execution_environment" json:"execution_environment"`
	// ExtraData ""
	ExtraData types.String `tfsdk:"extra_data" json:"extra_data"`
	// Forks ""
	Forks types.Int64 `tfsdk:"forks" json:"forks"`
	// ID "Database ID for this schedule."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// Inventory "Inventory applied as a prompt, assuming job template prompts for inventory"
	Inventory types.Int64 `tfsdk:"inventory" json:"inventory"`
	// JobSliceCount ""
	JobSliceCount types.Int64 `tfsdk:"job_slice_count" json:"job_slice_count"`
	// JobTags ""
	JobTags types.String `tfsdk:"job_tags" json:"job_tags"`
	// JobType ""
	JobType types.String `tfsdk:"job_type" json:"job_type"`
	// Limit ""
	Limit types.String `tfsdk:"limit" json:"limit"`
	// Name "Name of this schedule."
	Name types.String `tfsdk:"name" json:"name"`
	// NextRun "The next time that the scheduled action will run."
	NextRun types.String `tfsdk:"next_run" json:"next_run"`
	// Rrule "A value representing the schedules iCal recurrence rule."
	Rrule types.String `tfsdk:"rrule" json:"rrule"`
	// ScmBranch ""
	ScmBranch types.String `tfsdk:"scm_branch" json:"scm_branch"`
	// SkipTags ""
	SkipTags types.String `tfsdk:"skip_tags" json:"skip_tags"`
	// Timeout ""
	Timeout types.Int64 `tfsdk:"timeout" json:"timeout"`
	// Timezone "The timezone this schedule runs in. This field is extracted from the RRULE. If the timezone in the RRULE is a link to another timezone, the link will be reflected in this field."
	Timezone types.String `tfsdk:"timezone" json:"timezone"`
	// UnifiedJobTemplate ""
	UnifiedJobTemplate types.Int64 `tfsdk:"unified_job_template" json:"unified_job_template"`
	// Until "The date this schedule will end. This field is computed from the RRULE. If the schedule does not end an empty string will be returned"
	Until types.String `tfsdk:"until" json:"until"`
	// Verbosity ""
	Verbosity types.String `tfsdk:"verbosity" json:"verbosity"`
}

// Clone the object
func (o *scheduleTerraformModel) Clone() scheduleTerraformModel {
	return scheduleTerraformModel{
		Description:          o.Description,
		DiffMode:             o.DiffMode,
		Dtend:                o.Dtend,
		Dtstart:              o.Dtstart,
		Enabled:              o.Enabled,
		ExecutionEnvironment: o.ExecutionEnvironment,
		ExtraData:            o.ExtraData,
		Forks:                o.Forks,
		ID:                   o.ID,
		Inventory:            o.Inventory,
		JobSliceCount:        o.JobSliceCount,
		JobTags:              o.JobTags,
		JobType:              o.JobType,
		Limit:                o.Limit,
		Name:                 o.Name,
		NextRun:              o.NextRun,
		Rrule:                o.Rrule,
		ScmBranch:            o.ScmBranch,
		SkipTags:             o.SkipTags,
		Timeout:              o.Timeout,
		Timezone:             o.Timezone,
		UnifiedJobTemplate:   o.UnifiedJobTemplate,
		Until:                o.Until,
		Verbosity:            o.Verbosity,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Schedule
func (o *scheduleTerraformModel) BodyRequest() (req scheduleBodyRequestModel) {
	req.Description = o.Description.ValueString()
	req.DiffMode = o.DiffMode.ValueBool()
	req.Enabled = o.Enabled.ValueBool()
	req.ExecutionEnvironment = o.ExecutionEnvironment.ValueInt64()
	req.ExtraData = json.RawMessage(o.ExtraData.ValueString())
	req.Forks = o.Forks.ValueInt64()
	req.Inventory = o.Inventory.ValueInt64()
	req.JobSliceCount = o.JobSliceCount.ValueInt64()
	req.JobTags = o.JobTags.ValueString()
	req.JobType = o.JobType.ValueString()
	req.Limit = o.Limit.ValueString()
	req.Name = o.Name.ValueString()
	req.Rrule = o.Rrule.ValueString()
	req.ScmBranch = o.ScmBranch.ValueString()
	req.SkipTags = o.SkipTags.ValueString()
	req.Timeout = o.Timeout.ValueInt64()
	req.UnifiedJobTemplate = o.UnifiedJobTemplate.ValueInt64()
	req.Verbosity = o.Verbosity.ValueString()
	return
}

func (o *scheduleTerraformModel) setDescription(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *scheduleTerraformModel) setDiffMode(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.DiffMode, data)
}

func (o *scheduleTerraformModel) setDtend(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Dtend, data, false)
}

func (o *scheduleTerraformModel) setDtstart(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Dtstart, data, false)
}

func (o *scheduleTerraformModel) setEnabled(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.Enabled, data)
}

func (o *scheduleTerraformModel) setExecutionEnvironment(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ExecutionEnvironment, data)
}

func (o *scheduleTerraformModel) setExtraData(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetJsonString(&o.ExtraData, data, false)
}

func (o *scheduleTerraformModel) setForks(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Forks, data)
}

func (o *scheduleTerraformModel) setID(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *scheduleTerraformModel) setInventory(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Inventory, data)
}

func (o *scheduleTerraformModel) setJobSliceCount(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.JobSliceCount, data)
}

func (o *scheduleTerraformModel) setJobTags(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.JobTags, data, false)
}

func (o *scheduleTerraformModel) setJobType(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.JobType, data, false)
}

func (o *scheduleTerraformModel) setLimit(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Limit, data, false)
}

func (o *scheduleTerraformModel) setName(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *scheduleTerraformModel) setNextRun(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.NextRun, data, false)
}

func (o *scheduleTerraformModel) setRrule(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Rrule, data, false)
}

func (o *scheduleTerraformModel) setScmBranch(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.ScmBranch, data, false)
}

func (o *scheduleTerraformModel) setSkipTags(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.SkipTags, data, false)
}

func (o *scheduleTerraformModel) setTimeout(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Timeout, data)
}

func (o *scheduleTerraformModel) setTimezone(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Timezone, data, false)
}

func (o *scheduleTerraformModel) setUnifiedJobTemplate(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.UnifiedJobTemplate, data)
}

func (o *scheduleTerraformModel) setUntil(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Until, data, false)
}

func (o *scheduleTerraformModel) setVerbosity(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Verbosity, data, false)
}

func (o *scheduleTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDiffMode(data["diff_mode"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDtend(data["dtend"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDtstart(data["dtstart"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setEnabled(data["enabled"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setExecutionEnvironment(data["execution_environment"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setExtraData(data["extra_data"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setForks(data["forks"]); dg.HasError() {
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
	if dg, _ := o.setNextRun(data["next_run"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setRrule(data["rrule"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setScmBranch(data["scm_branch"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSkipTags(data["skip_tags"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setTimeout(data["timeout"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setTimezone(data["timezone"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setUnifiedJobTemplate(data["unified_job_template"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setUntil(data["until"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setVerbosity(data["verbosity"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// scheduleBodyRequestModel maps the schema for Schedule for creating and updating the data
type scheduleBodyRequestModel struct {
	// Description "Optional description of this schedule."
	Description string `json:"description,omitempty"`
	// DiffMode ""
	DiffMode bool `json:"diff_mode"`
	// Enabled "Enables processing of this schedule."
	Enabled bool `json:"enabled"`
	// ExecutionEnvironment "The container image to be used for execution."
	ExecutionEnvironment int64 `json:"execution_environment,omitempty"`
	// ExtraData ""
	ExtraData json.RawMessage `json:"extra_data,omitempty"`
	// Forks ""
	Forks int64 `json:"forks,omitempty"`
	// Inventory "Inventory applied as a prompt, assuming job template prompts for inventory"
	Inventory int64 `json:"inventory,omitempty"`
	// JobSliceCount ""
	JobSliceCount int64 `json:"job_slice_count,omitempty"`
	// JobTags ""
	JobTags string `json:"job_tags,omitempty"`
	// JobType ""
	JobType string `json:"job_type,omitempty"`
	// Limit ""
	Limit string `json:"limit,omitempty"`
	// Name "Name of this schedule."
	Name string `json:"name"`
	// Rrule "A value representing the schedules iCal recurrence rule."
	Rrule string `json:"rrule"`
	// ScmBranch ""
	ScmBranch string `json:"scm_branch,omitempty"`
	// SkipTags ""
	SkipTags string `json:"skip_tags,omitempty"`
	// Timeout ""
	Timeout int64 `json:"timeout,omitempty"`
	// UnifiedJobTemplate ""
	UnifiedJobTemplate int64 `json:"unified_job_template"`
	// Verbosity ""
	Verbosity string `json:"verbosity,omitempty"`
}
