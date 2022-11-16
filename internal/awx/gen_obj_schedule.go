package awx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	p "path"
	"strconv"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"

	"github.com/hashicorp/terraform-plugin-framework-validators/schemavalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	// Until "The date this schedule will end. This field is computed from the RRULE. If the schedule does not end an emptry string will be returned"
	Until types.String `tfsdk:"until" json:"until"`
	// Verbosity ""
	Verbosity types.String `tfsdk:"verbosity" json:"verbosity"`
}

// Clone the object
func (o scheduleTerraformModel) Clone() scheduleTerraformModel {
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
func (o scheduleTerraformModel) BodyRequest() (req scheduleBodyRequestModel) {
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

func (o *scheduleTerraformModel) setDescription(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *scheduleTerraformModel) setDiffMode(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.DiffMode, data)
}

func (o *scheduleTerraformModel) setDtend(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Dtend, data, false)
}

func (o *scheduleTerraformModel) setDtstart(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Dtstart, data, false)
}

func (o *scheduleTerraformModel) setEnabled(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.Enabled, data)
}

func (o *scheduleTerraformModel) setExecutionEnvironment(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ExecutionEnvironment, data)
}

func (o *scheduleTerraformModel) setExtraData(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.ExtraData, data, false)
}

func (o *scheduleTerraformModel) setForks(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Forks, data)
}

func (o *scheduleTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *scheduleTerraformModel) setInventory(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Inventory, data)
}

func (o *scheduleTerraformModel) setJobSliceCount(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.JobSliceCount, data)
}

func (o *scheduleTerraformModel) setJobTags(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.JobTags, data, false)
}

func (o *scheduleTerraformModel) setJobType(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.JobType, data, false)
}

func (o *scheduleTerraformModel) setLimit(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Limit, data, false)
}

func (o *scheduleTerraformModel) setName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *scheduleTerraformModel) setNextRun(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.NextRun, data, false)
}

func (o *scheduleTerraformModel) setRrule(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Rrule, data, false)
}

func (o *scheduleTerraformModel) setScmBranch(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.ScmBranch, data, false)
}

func (o *scheduleTerraformModel) setSkipTags(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SkipTags, data, false)
}

func (o *scheduleTerraformModel) setTimeout(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Timeout, data)
}

func (o *scheduleTerraformModel) setTimezone(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Timezone, data, false)
}

func (o *scheduleTerraformModel) setUnifiedJobTemplate(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.UnifiedJobTemplate, data)
}

func (o *scheduleTerraformModel) setUntil(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Until, data, false)
}

func (o *scheduleTerraformModel) setVerbosity(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Verbosity, data, false)
}

func (o *scheduleTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
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

var (
	_ datasource.DataSource              = &scheduleDataSource{}
	_ datasource.DataSourceWithConfigure = &scheduleDataSource{}
)

// NewScheduleDataSource is a helper function to instantiate the Schedule data source.
func NewScheduleDataSource() datasource.DataSource {
	return &scheduleDataSource{}
}

// scheduleDataSource is the data source implementation.
type scheduleDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *scheduleDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/schedules/"
}

// Metadata returns the data source type name.
func (o *scheduleDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_schedule"
}

// GetSchema defines the schema for the data source.
func (o *scheduleDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"Schedule",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"description": {
					Description: "Optional description of this schedule.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"diff_mode": {
					Description: "Diff mode",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"dtend": {
					Description: "The last occurrence of the schedule occurs before this time, aftewards the schedule expires.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"dtstart": {
					Description: "The first occurrence of the schedule occurs on or after this time.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"enabled": {
					Description: "Enables processing of this schedule.",
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
				"extra_data": {
					Description: "Extra data",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"forks": {
					Description: "Forks",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"id": {
					Description: "Database ID for this schedule.",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						schemavalidator.ExactlyOneOf(
							path.MatchRoot("id"),
						),
					},
				},
				"inventory": {
					Description: "Inventory applied as a prompt, assuming job template prompts for inventory",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"job_slice_count": {
					Description: "Job slice count",
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
						stringvalidator.OneOf([]string{"", "run", "check"}...),
					},
				},
				"limit": {
					Description: "Limit",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"name": {
					Description: "Name of this schedule.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"next_run": {
					Description: "The next time that the scheduled action will run.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"rrule": {
					Description: "A value representing the schedules iCal recurrence rule.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"scm_branch": {
					Description: "Scm branch",
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
				"timeout": {
					Description: "Timeout",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"timezone": {
					Description: "The timezone this schedule runs in. This field is extracted from the RRULE. If the timezone in the RRULE is a link to another timezone, the link will be reflected in this field.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"unified_job_template": {
					Description: "Unified job template",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"until": {
					Description: "The date this schedule will end. This field is computed from the RRULE. If the schedule does not end an emptry string will be returned",
					Type:        types.StringType,
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
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *scheduleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state scheduleTerraformModel
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

	if !searchDefined {
		var detailMessage string
		resp.Diagnostics.AddError(
			fmt.Sprintf("missing configuration for one of the predefined search groups"),
			detailMessage,
		)
		return
	}

	// Creates a new request for Schedule
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Schedule on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Schedule
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Schedule on %s", o.endpoint),
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
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &scheduleResource{}
	_ resource.ResourceWithConfigure   = &scheduleResource{}
	_ resource.ResourceWithImportState = &scheduleResource{}
)

// NewScheduleResource is a helper function to simplify the provider implementation.
func NewScheduleResource() resource.Resource {
	return &scheduleResource{}
}

type scheduleResource struct {
	client   c.Client
	endpoint string
}

func (o *scheduleResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/schedules/"
}

func (o scheduleResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_schedule"
}

func (o scheduleResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"Schedule",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"description": {
					Description: "Optional description of this schedule.",
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
					Description: "Diff mode",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"enabled": {
					Description: "Enables processing of this schedule.",
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
				"extra_data": {
					Description: "Extra data",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
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
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"inventory": {
					Description: "Inventory applied as a prompt, assuming job template prompts for inventory",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"job_slice_count": {
					Description: "Job slice count",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"job_tags": {
					Description: "Job tags",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
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
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"", "run", "check"}...),
					},
				},
				"limit": {
					Description: "Limit",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"name": {
					Description:   "Name of this schedule.",
					Type:          types.StringType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.LengthAtMost(512),
					},
				},
				"rrule": {
					Description:   "A value representing the schedules iCal recurrence rule.",
					Type:          types.StringType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators:    []tfsdk.AttributeValidator{},
				},
				"scm_branch": {
					Description: "Scm branch",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"skip_tags": {
					Description: "Skip tags",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"timeout": {
					Description: "Timeout",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"unified_job_template": {
					Description:   "Unified job template",
					Type:          types.Int64Type,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators:    []tfsdk.AttributeValidator{},
				},
				"verbosity": {
					Description: "Verbosity",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"0", "1", "2", "3", "4", "5"}...),
					},
				},
				// Write only elements
				// Data only elements
				"dtend": {
					Description: "The last occurrence of the schedule occurs before this time, aftewards the schedule expires.",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"dtstart": {
					Description: "The first occurrence of the schedule occurs on or after this time.",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"id": {
					Description: "Database ID for this schedule.",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"next_run": {
					Description: "The next time that the scheduled action will run.",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"timezone": {
					Description: "The timezone this schedule runs in. This field is extracted from the RRULE. If the timezone in the RRULE is a link to another timezone, the link will be reflected in this field.",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"until": {
					Description: "The date this schedule will end. This field is computed from the RRULE. If the schedule does not end an emptry string will be returned",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
			},
		}), nil
}

func (o *scheduleResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the Schedule.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (o *scheduleResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state scheduleTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Schedule
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Schedule on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new Schedule resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for Schedule on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *scheduleResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state scheduleTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Schedule
	var r *http.Request
	var id = state.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Schedule on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for Schedule from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Schedule on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *scheduleResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state scheduleTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Schedule
	var r *http.Request
	var id = plan.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Schedule on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new Schedule resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for Schedule on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *scheduleResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state scheduleTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for Schedule
	var r *http.Request
	var id = state.ID
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Schedule on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	// Delete existing Schedule
	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to delete resource for Schedule on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}
