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
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Schedule
func (o *scheduleTerraformModel) BodyRequest() *scheduleBodyRequestModel {
	var req scheduleBodyRequestModel
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
	return &req
}

func (o *scheduleTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Description, data["description"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.DiffMode, data["diff_mode"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Dtend, data["dtend"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Dtstart, data["dtstart"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.Enabled, data["enabled"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.ExecutionEnvironment, data["execution_environment"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.ExtraData, data["extra_data"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.Forks, data["forks"])
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
		dg, _ := helpers.AttrValueSetInt64(&o.JobSliceCount, data["job_slice_count"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.JobTags, data["job_tags"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.JobType, data["job_type"], false)
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
		dg, _ := helpers.AttrValueSetString(&o.NextRun, data["next_run"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Rrule, data["rrule"], false)
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
		dg, _ := helpers.AttrValueSetInt64(&o.Timeout, data["timeout"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Timezone, data["timezone"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.UnifiedJobTemplate, data["unified_job_template"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Until, data["until"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Verbosity, data["verbosity"], false)
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
