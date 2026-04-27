package awx

import (
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

type scheduleResource = framework.GenericResource[scheduleTerraformModel, scheduleBodyRequestModel, *scheduleTerraformModel]

// NewScheduleResource is a helper function to simplify the provider implementation.
func NewScheduleResource() resource.Resource {
	return &scheduleResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "schedule", Endpoint: "/api/v2/schedules/"}},
		Cfg: framework.ResourceCfg[scheduleTerraformModel, scheduleBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description: "Optional description of this schedule.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"diff_mode": schema.BoolAttribute{
						Description: "Diff mode",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"enabled": schema.BoolAttribute{
						Description: "Enables processing of this schedule.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"execution_environment": schema.Int64Attribute{
						Description: "The container image to be used for execution.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"extra_data": schema.StringAttribute{
						Description: "Extra data",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"forks": schema.Int64Attribute{
						Description: "Forks",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
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
					"job_slice_count": schema.Int64Attribute{
						Description: "Job slice count",
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
					"job_type": schema.StringAttribute{
						Description: "Job type",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"",
								"run",
								"check",
							),
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
						Description: "Name of this schedule.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"rrule": schema.StringAttribute{
						Description: "A value representing the schedules iCal recurrence rule.",
						Required:    true,
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
					"timeout": schema.Int64Attribute{
						Description: "Timeout",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"unified_job_template": schema.Int64Attribute{
						Description: "Unified job template",
						Required:    true,
					},
					"verbosity": schema.StringAttribute{
						Description: "Verbosity",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"0",
								"1",
								"2",
								"3",
								"4",
								"5",
							),
						},
					},
					"dtend": schema.StringAttribute{
						Description: "The last occurrence of the schedule occurs before this time, aftewards the schedule expires.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"dtstart": schema.StringAttribute{
						Description: "The first occurrence of the schedule occurs on or after this time.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this schedule.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"next_run": schema.StringAttribute{
						Description: "The next time that the scheduled action will run.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"timezone": schema.StringAttribute{
						Description: "The timezone this schedule runs in. This field is extracted from the RRULE. If the timezone in the RRULE is a link to another timezone, the link will be reflected in this field.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"until": schema.StringAttribute{
						Description: "The date this schedule will end. This field is computed from the RRULE. If the schedule does not end an empty string will be returned",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *scheduleTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			ApiVersion:   ApiVersion,
			ResourceName: "Schedule",
		},
	}
}

type scheduleDataSource = framework.GenericDataSource[scheduleTerraformModel, *scheduleTerraformModel]

// NewScheduleDataSource is a helper function to instantiate the Schedule data source.
func NewScheduleDataSource() datasource.DataSource {
	return &scheduleDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "schedule", Endpoint: "/api/v2/schedules/"}},
		Cfg: framework.DataSourceCfg[scheduleTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"description": dschema.StringAttribute{
						Description: "Optional description of this schedule.",
						Computed:    true,
					},
					"diff_mode": dschema.BoolAttribute{
						Description: "Diff mode",
						Computed:    true,
					},
					"dtend": dschema.StringAttribute{
						Description: "The last occurrence of the schedule occurs before this time, aftewards the schedule expires.",
						Computed:    true,
					},
					"dtstart": dschema.StringAttribute{
						Description: "The first occurrence of the schedule occurs on or after this time.",
						Computed:    true,
					},
					"enabled": dschema.BoolAttribute{
						Description: "Enables processing of this schedule.",
						Computed:    true,
					},
					"execution_environment": dschema.Int64Attribute{
						Description: "The container image to be used for execution.",
						Computed:    true,
					},
					"extra_data": dschema.StringAttribute{
						Description: "Extra data",
						Computed:    true,
					},
					"forks": dschema.Int64Attribute{
						Description: "Forks",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this schedule.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
							),
						},
					},
					"inventory": dschema.Int64Attribute{
						Description: "Inventory applied as a prompt, assuming job template prompts for inventory",
						Computed:    true,
					},
					"job_slice_count": dschema.Int64Attribute{
						Description: "Job slice count",
						Computed:    true,
					},
					"job_tags": dschema.StringAttribute{
						Description: "Job tags",
						Computed:    true,
					},
					"job_type": dschema.StringAttribute{
						Description: "Job type",
						Computed:    true,
					},
					"limit": dschema.StringAttribute{
						Description: "Limit",
						Computed:    true,
					},
					"name": dschema.StringAttribute{
						Description: "Name of this schedule.",
						Computed:    true,
					},
					"next_run": dschema.StringAttribute{
						Description: "The next time that the scheduled action will run.",
						Computed:    true,
					},
					"rrule": dschema.StringAttribute{
						Description: "A value representing the schedules iCal recurrence rule.",
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
					"timeout": dschema.Int64Attribute{
						Description: "Timeout",
						Computed:    true,
					},
					"timezone": dschema.StringAttribute{
						Description: "The timezone this schedule runs in. This field is extracted from the RRULE. If the timezone in the RRULE is a link to another timezone, the link will be reflected in this field.",
						Computed:    true,
					},
					"unified_job_template": dschema.Int64Attribute{
						Description: "Unified job template",
						Computed:    true,
					},
					"until": dschema.StringAttribute{
						Description: "The date this schedule will end. This field is computed from the RRULE. If the schedule does not end an empty string will be returned",
						Computed:    true,
					},
					"verbosity": dschema.StringAttribute{
						Description: "Verbosity",
						Computed:    true,
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "Schedule",
		},
	}
}
