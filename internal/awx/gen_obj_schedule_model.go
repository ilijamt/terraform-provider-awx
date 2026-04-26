package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type scheduleTerraformModel struct {
	Description          types.String `tfsdk:"description" json:"description"`
	DiffMode             types.Bool   `tfsdk:"diff_mode" json:"diff_mode"`
	Dtend                types.String `tfsdk:"dtend" json:"dtend"`
	Dtstart              types.String `tfsdk:"dtstart" json:"dtstart"`
	Enabled              types.Bool   `tfsdk:"enabled" json:"enabled"`
	ExecutionEnvironment types.Int64  `tfsdk:"execution_environment" json:"execution_environment"`
	ExtraData            types.String `tfsdk:"extra_data" json:"extra_data"`
	Forks                types.Int64  `tfsdk:"forks" json:"forks"`
	ID                   types.Int64  `tfsdk:"id" json:"id"`
	Inventory            types.Int64  `tfsdk:"inventory" json:"inventory"`
	JobSliceCount        types.Int64  `tfsdk:"job_slice_count" json:"job_slice_count"`
	JobTags              types.String `tfsdk:"job_tags" json:"job_tags"`
	JobType              types.String `tfsdk:"job_type" json:"job_type"`
	Limit                types.String `tfsdk:"limit" json:"limit"`
	Name                 types.String `tfsdk:"name" json:"name"`
	NextRun              types.String `tfsdk:"next_run" json:"next_run"`
	Rrule                types.String `tfsdk:"rrule" json:"rrule"`
	ScmBranch            types.String `tfsdk:"scm_branch" json:"scm_branch"`
	SkipTags             types.String `tfsdk:"skip_tags" json:"skip_tags"`
	Timeout              types.Int64  `tfsdk:"timeout" json:"timeout"`
	Timezone             types.String `tfsdk:"timezone" json:"timezone"`
	UnifiedJobTemplate   types.Int64  `tfsdk:"unified_job_template" json:"unified_job_template"`
	Until                types.String `tfsdk:"until" json:"until"`
	Verbosity            types.String `tfsdk:"verbosity" json:"verbosity"`
}

func (o *scheduleTerraformModel) Clone() scheduleTerraformModel {
	return *o
}

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
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetBool(&o.DiffMode, data["diff_mode"]))
	collect(helpers.AttrValueSetString(&o.Dtend, data["dtend"], false))
	collect(helpers.AttrValueSetString(&o.Dtstart, data["dtstart"], false))
	collect(helpers.AttrValueSetBool(&o.Enabled, data["enabled"]))
	collect(helpers.AttrValueSetInt64(&o.ExecutionEnvironment, data["execution_environment"]))
	collect(helpers.AttrValueSetJsonString(&o.ExtraData, data["extra_data"], false))
	collect(helpers.AttrValueSetInt64(&o.Forks, data["forks"]))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetInt64(&o.Inventory, data["inventory"]))
	collect(helpers.AttrValueSetInt64(&o.JobSliceCount, data["job_slice_count"]))
	collect(helpers.AttrValueSetString(&o.JobTags, data["job_tags"], false))
	collect(helpers.AttrValueSetString(&o.JobType, data["job_type"], false))
	collect(helpers.AttrValueSetString(&o.Limit, data["limit"], false))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetString(&o.NextRun, data["next_run"], false))
	collect(helpers.AttrValueSetString(&o.Rrule, data["rrule"], false))
	collect(helpers.AttrValueSetString(&o.ScmBranch, data["scm_branch"], false))
	collect(helpers.AttrValueSetString(&o.SkipTags, data["skip_tags"], false))
	collect(helpers.AttrValueSetInt64(&o.Timeout, data["timeout"]))
	collect(helpers.AttrValueSetString(&o.Timezone, data["timezone"], false))
	collect(helpers.AttrValueSetInt64(&o.UnifiedJobTemplate, data["unified_job_template"]))
	collect(helpers.AttrValueSetString(&o.Until, data["until"], false))
	collect(helpers.AttrValueSetString(&o.Verbosity, data["verbosity"], false))
	return diags, nil
}

type scheduleBodyRequestModel struct {
	Description          string          `json:"description,omitempty"`
	DiffMode             bool            `json:"diff_mode"`
	Enabled              bool            `json:"enabled"`
	ExecutionEnvironment int64           `json:"execution_environment,omitempty"`
	ExtraData            json.RawMessage `json:"extra_data,omitempty"`
	Forks                int64           `json:"forks,omitempty"`
	Inventory            int64           `json:"inventory,omitempty"`
	JobSliceCount        int64           `json:"job_slice_count,omitempty"`
	JobTags              string          `json:"job_tags,omitempty"`
	JobType              string          `json:"job_type,omitempty"`
	Limit                string          `json:"limit,omitempty"`
	Name                 string          `json:"name"`
	Rrule                string          `json:"rrule"`
	ScmBranch            string          `json:"scm_branch,omitempty"`
	SkipTags             string          `json:"skip_tags,omitempty"`
	Timeout              int64           `json:"timeout,omitempty"`
	UnifiedJobTemplate   int64           `json:"unified_job_template"`
	Verbosity            string          `json:"verbosity,omitempty"`
}
