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

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/schemavalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// adHocCommandTerraformModel maps the schema for AdHocCommand when using Data Source
type adHocCommandTerraformModel struct {
	// BecomeEnabled ""
	BecomeEnabled types.Bool `tfsdk:"become_enabled" json:"become_enabled"`
	// CanceledOn "The date and time when the cancel request was sent."
	CanceledOn types.String `tfsdk:"canceled_on" json:"canceled_on"`
	// ControllerNode "The instance that managed the execution environment."
	ControllerNode types.String `tfsdk:"controller_node" json:"controller_node"`
	// Credential ""
	Credential types.Int64 `tfsdk:"credential" json:"credential"`
	// DiffMode ""
	DiffMode types.Bool `tfsdk:"diff_mode" json:"diff_mode"`
	// Elapsed "Elapsed time in seconds that the job ran."
	Elapsed types.Float64 `tfsdk:"elapsed" json:"elapsed"`
	// ExecutionEnvironment "The container image to be used for execution."
	ExecutionEnvironment types.Int64 `tfsdk:"execution_environment" json:"execution_environment"`
	// ExecutionNode "The node the job executed on."
	ExecutionNode types.String `tfsdk:"execution_node" json:"execution_node"`
	// ExtraVars ""
	ExtraVars types.String `tfsdk:"extra_vars" json:"extra_vars"`
	// Failed ""
	Failed types.Bool `tfsdk:"failed" json:"failed"`
	// Finished "The date and time the job finished execution."
	Finished types.String `tfsdk:"finished" json:"finished"`
	// Forks ""
	Forks types.Int64 `tfsdk:"forks" json:"forks"`
	// ID "Database ID for this ad hoc command."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// Inventory ""
	Inventory types.Int64 `tfsdk:"inventory" json:"inventory"`
	// JobExplanation "A status field to indicate the state of the job if it wasn't able to run and capture stdout"
	JobExplanation types.String `tfsdk:"job_explanation" json:"job_explanation"`
	// JobType ""
	JobType types.String `tfsdk:"job_type" json:"job_type"`
	// LaunchType ""
	LaunchType types.String `tfsdk:"launch_type" json:"launch_type"`
	// LaunchedBy ""
	LaunchedBy types.Int64 `tfsdk:"launched_by" json:"launched_by"`
	// Limit ""
	Limit types.String `tfsdk:"limit" json:"limit"`
	// ModuleArgs ""
	ModuleArgs types.String `tfsdk:"module_args" json:"module_args"`
	// ModuleName ""
	ModuleName types.String `tfsdk:"module_name" json:"module_name"`
	// Name "Name of this ad hoc command."
	Name types.String `tfsdk:"name" json:"name"`
	// Started "The date and time the job was queued for starting."
	Started types.String `tfsdk:"started" json:"started"`
	// Status ""
	Status types.String `tfsdk:"status" json:"status"`
	// Verbosity ""
	Verbosity types.String `tfsdk:"verbosity" json:"verbosity"`
	// WorkUnitId "The Receptor work unit ID associated with this job."
	WorkUnitId types.String `tfsdk:"work_unit_id" json:"work_unit_id"`
}

// Clone the object
func (o adHocCommandTerraformModel) Clone() adHocCommandTerraformModel {
	return adHocCommandTerraformModel{
		BecomeEnabled:        o.BecomeEnabled,
		CanceledOn:           o.CanceledOn,
		ControllerNode:       o.ControllerNode,
		Credential:           o.Credential,
		DiffMode:             o.DiffMode,
		Elapsed:              o.Elapsed,
		ExecutionEnvironment: o.ExecutionEnvironment,
		ExecutionNode:        o.ExecutionNode,
		ExtraVars:            o.ExtraVars,
		Failed:               o.Failed,
		Finished:             o.Finished,
		Forks:                o.Forks,
		ID:                   o.ID,
		Inventory:            o.Inventory,
		JobExplanation:       o.JobExplanation,
		JobType:              o.JobType,
		LaunchType:           o.LaunchType,
		LaunchedBy:           o.LaunchedBy,
		Limit:                o.Limit,
		ModuleArgs:           o.ModuleArgs,
		ModuleName:           o.ModuleName,
		Name:                 o.Name,
		Started:              o.Started,
		Status:               o.Status,
		Verbosity:            o.Verbosity,
		WorkUnitId:           o.WorkUnitId,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for AdHocCommand
func (o adHocCommandTerraformModel) BodyRequest() (req adHocCommandBodyRequestModel) {
	req.BecomeEnabled = o.BecomeEnabled.ValueBool()
	req.Credential = o.Credential.ValueInt64()
	req.DiffMode = o.DiffMode.ValueBool()
	req.ExecutionEnvironment = o.ExecutionEnvironment.ValueInt64()
	req.ExtraVars = o.ExtraVars.ValueString()
	req.Forks = o.Forks.ValueInt64()
	req.Inventory = o.Inventory.ValueInt64()
	req.JobType = o.JobType.ValueString()
	req.Limit = o.Limit.ValueString()
	req.ModuleArgs = o.ModuleArgs.ValueString()
	req.ModuleName = o.ModuleName.ValueString()
	req.Verbosity = o.Verbosity.ValueString()
	return
}

func (o *adHocCommandTerraformModel) setBecomeEnabled(data any) (d diag.Diagnostics, err error) {
	// Decode "become_enabled"
	if val, ok := data.(bool); ok {
		o.BecomeEnabled = types.BoolValue(val)
	} else {
		o.BecomeEnabled = types.BoolNull()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setCanceledOn(data any) (d diag.Diagnostics, err error) {
	// Decode "canceled_on"
	if val, ok := data.(string); ok {
		o.CanceledOn = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.CanceledOn = types.StringValue(val.String())
	} else {
		o.CanceledOn = types.StringNull()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setControllerNode(data any) (d diag.Diagnostics, err error) {
	// Decode "controller_node"
	if val, ok := data.(string); ok {
		o.ControllerNode = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.ControllerNode = types.StringValue(val.String())
	} else {
		o.ControllerNode = types.StringNull()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setCredential(data any) (d diag.Diagnostics, err error) {
	// Decode "credential"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.Credential = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.Credential = types.Int64Value(val)
	} else {
		o.Credential = types.Int64Null()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setDiffMode(data any) (d diag.Diagnostics, err error) {
	// Decode "diff_mode"
	if val, ok := data.(bool); ok {
		o.DiffMode = types.BoolValue(val)
	} else {
		o.DiffMode = types.BoolNull()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setElapsed(data any) (d diag.Diagnostics, err error) {
	// Decode "elapsed"
	if val, ok := data.(json.Number); ok {
		v, err := val.Float64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to float64", val),
				err.Error(),
			)
			return d, err
		}
		o.Elapsed = types.Float64Value(v)
	} else if val, ok := data.(float64); ok {
		o.Elapsed = types.Float64Value(val)
	} else {
		o.Elapsed = types.Float64Null()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setExecutionEnvironment(data any) (d diag.Diagnostics, err error) {
	// Decode "execution_environment"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.ExecutionEnvironment = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.ExecutionEnvironment = types.Int64Value(val)
	} else {
		o.ExecutionEnvironment = types.Int64Null()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setExecutionNode(data any) (d diag.Diagnostics, err error) {
	// Decode "execution_node"
	if val, ok := data.(string); ok {
		o.ExecutionNode = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.ExecutionNode = types.StringValue(val.String())
	} else {
		o.ExecutionNode = types.StringNull()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setExtraVars(data any) (d diag.Diagnostics, err error) {
	// Decode "extra_vars"
	if val, ok := data.(string); ok {
		o.ExtraVars = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.ExtraVars = types.StringValue(val.String())
	} else {
		o.ExtraVars = types.StringNull()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setFailed(data any) (d diag.Diagnostics, err error) {
	// Decode "failed"
	if val, ok := data.(bool); ok {
		o.Failed = types.BoolValue(val)
	} else {
		o.Failed = types.BoolNull()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setFinished(data any) (d diag.Diagnostics, err error) {
	// Decode "finished"
	if val, ok := data.(string); ok {
		o.Finished = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.Finished = types.StringValue(val.String())
	} else {
		o.Finished = types.StringNull()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setForks(data any) (d diag.Diagnostics, err error) {
	// Decode "forks"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.Forks = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.Forks = types.Int64Value(val)
	} else {
		o.Forks = types.Int64Null()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	// Decode "id"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.ID = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.ID = types.Int64Value(val)
	} else {
		o.ID = types.Int64Null()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setInventory(data any) (d diag.Diagnostics, err error) {
	// Decode "inventory"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.Inventory = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.Inventory = types.Int64Value(val)
	} else {
		o.Inventory = types.Int64Null()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setJobExplanation(data any) (d diag.Diagnostics, err error) {
	// Decode "job_explanation"
	if val, ok := data.(string); ok {
		o.JobExplanation = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.JobExplanation = types.StringValue(val.String())
	} else {
		o.JobExplanation = types.StringNull()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setJobType(data any) (d diag.Diagnostics, err error) {
	// Decode "job_type"
	if val, ok := data.(string); ok {
		o.JobType = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.JobType = types.StringValue(val.String())
	} else {
		o.JobType = types.StringNull()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setLaunchType(data any) (d diag.Diagnostics, err error) {
	// Decode "launch_type"
	if val, ok := data.(string); ok {
		o.LaunchType = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.LaunchType = types.StringValue(val.String())
	} else {
		o.LaunchType = types.StringNull()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setLaunchedBy(data any) (d diag.Diagnostics, err error) {
	// Decode "launched_by"
	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		o.LaunchedBy = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		o.LaunchedBy = types.Int64Value(val)
	} else {
		o.LaunchedBy = types.Int64Null()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setLimit(data any) (d diag.Diagnostics, err error) {
	// Decode "limit"
	if val, ok := data.(string); ok {
		o.Limit = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.Limit = types.StringValue(val.String())
	} else {
		o.Limit = types.StringNull()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setModuleArgs(data any) (d diag.Diagnostics, err error) {
	// Decode "module_args"
	if val, ok := data.(string); ok {
		o.ModuleArgs = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.ModuleArgs = types.StringValue(val.String())
	} else {
		o.ModuleArgs = types.StringNull()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setModuleName(data any) (d diag.Diagnostics, err error) {
	// Decode "module_name"
	if val, ok := data.(string); ok {
		o.ModuleName = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.ModuleName = types.StringValue(val.String())
	} else {
		o.ModuleName = types.StringNull()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setName(data any) (d diag.Diagnostics, err error) {
	// Decode "name"
	if val, ok := data.(string); ok {
		o.Name = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.Name = types.StringValue(val.String())
	} else {
		o.Name = types.StringNull()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setStarted(data any) (d diag.Diagnostics, err error) {
	// Decode "started"
	if val, ok := data.(string); ok {
		o.Started = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.Started = types.StringValue(val.String())
	} else {
		o.Started = types.StringNull()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setStatus(data any) (d diag.Diagnostics, err error) {
	// Decode "status"
	if val, ok := data.(string); ok {
		o.Status = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.Status = types.StringValue(val.String())
	} else {
		o.Status = types.StringNull()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setVerbosity(data any) (d diag.Diagnostics, err error) {
	// Decode "verbosity"
	if val, ok := data.(string); ok {
		o.Verbosity = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.Verbosity = types.StringValue(val.String())
	} else {
		o.Verbosity = types.StringNull()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) setWorkUnitId(data any) (d diag.Diagnostics, err error) {
	// Decode "work_unit_id"
	if val, ok := data.(string); ok {
		o.WorkUnitId = types.StringValue(helpers.TrimString(false, false, val))
	} else if val, ok := data.(json.Number); ok {
		o.WorkUnitId = types.StringValue(val.String())
	} else {
		o.WorkUnitId = types.StringNull()
	}
	return d, nil
}

func (o *adHocCommandTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setBecomeEnabled(data["become_enabled"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setCanceledOn(data["canceled_on"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setControllerNode(data["controller_node"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setCredential(data["credential"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDiffMode(data["diff_mode"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setElapsed(data["elapsed"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setExecutionEnvironment(data["execution_environment"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setExecutionNode(data["execution_node"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setExtraVars(data["extra_vars"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setFailed(data["failed"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setFinished(data["finished"]); dg.HasError() {
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
	if dg, _ := o.setJobExplanation(data["job_explanation"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setJobType(data["job_type"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setLaunchType(data["launch_type"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setLaunchedBy(data["launched_by"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setLimit(data["limit"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setModuleArgs(data["module_args"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setModuleName(data["module_name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setName(data["name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setStarted(data["started"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setStatus(data["status"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setVerbosity(data["verbosity"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setWorkUnitId(data["work_unit_id"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// adHocCommandBodyRequestModel maps the schema for AdHocCommand for creating and updating the data
type adHocCommandBodyRequestModel struct {
	// BecomeEnabled ""
	BecomeEnabled bool `json:"become_enabled"`
	// Credential ""
	Credential int64 `json:"credential,omitempty"`
	// DiffMode ""
	DiffMode bool `json:"diff_mode"`
	// ExecutionEnvironment "The container image to be used for execution."
	ExecutionEnvironment int64 `json:"execution_environment,omitempty"`
	// ExtraVars ""
	ExtraVars string `json:"extra_vars,omitempty"`
	// Forks ""
	Forks int64 `json:"forks,omitempty"`
	// Inventory ""
	Inventory int64 `json:"inventory,omitempty"`
	// JobType ""
	JobType string `json:"job_type,omitempty"`
	// Limit ""
	Limit string `json:"limit,omitempty"`
	// ModuleArgs ""
	ModuleArgs string `json:"module_args,omitempty"`
	// ModuleName ""
	ModuleName string `json:"module_name,omitempty"`
	// Verbosity ""
	Verbosity string `json:"verbosity,omitempty"`
}

var (
	_ datasource.DataSource              = &adHocCommandDataSource{}
	_ datasource.DataSourceWithConfigure = &adHocCommandDataSource{}
)

// NewAdHocCommandDataSource is a helper function to instantiate the AdHocCommand data source.
func NewAdHocCommandDataSource() datasource.DataSource {
	return &adHocCommandDataSource{}
}

// adHocCommandDataSource is the data source implementation.
type adHocCommandDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *adHocCommandDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/ad_hoc_commands/"
}

// Metadata returns the data source type name.
func (o *adHocCommandDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ad_hoc_command"
}

// GetSchema defines the schema for the data source.
func (o *adHocCommandDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"AdHocCommand",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"become_enabled": {
					Description: "Become enabled",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"canceled_on": {
					Description: "The date and time when the cancel request was sent.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"controller_node": {
					Description: "The instance that managed the execution environment.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"credential": {
					Description: "Credential",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"diff_mode": {
					Description: "Diff mode",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"elapsed": {
					Description: "Elapsed time in seconds that the job ran.",
					Type:        types.Float64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"execution_environment": {
					Description: "The container image to be used for execution.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"execution_node": {
					Description: "The node the job executed on.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"extra_vars": {
					Description: "Extra vars",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"failed": {
					Description: "Failed",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"finished": {
					Description: "The date and time the job finished execution.",
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
					Description: "Database ID for this ad hoc command.",
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
				"job_explanation": {
					Description: "A status field to indicate the state of the job if it wasn't able to run and capture stdout",
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
				"launch_type": {
					Description: "Launch type",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"manual", "relaunch", "callback", "scheduled", "dependency", "workflow", "webhook", "sync", "scm"}...),
					},
				},
				"launched_by": {
					Description: "Launched by",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"limit": {
					Description: "Limit",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"module_args": {
					Description: "Module args",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"module_name": {
					Description: "Module name",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"command", "shell", "yum", "apt", "apt_key", "apt_repository", "apt_rpm", "service", "group", "user", "mount", "ping", "selinux", "setup", "win_ping", "win_service", "win_updates", "win_group", "win_user"}...),
					},
				},
				"name": {
					Description: "Name of this ad hoc command.",
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
				"started": {
					Description: "The date and time the job was queued for starting.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"status": {
					Description: "Status",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"new", "pending", "waiting", "running", "successful", "failed", "error", "canceled"}...),
					},
				},
				"verbosity": {
					Description: "Verbosity",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"0", "1", "2", "3", "4", "5"}...),
					},
				},
				"work_unit_id": {
					Description: "The Receptor work unit ID associated with this job.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *adHocCommandDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state adHocCommandTerraformModel
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
			fmt.Sprintf("missing configuration for one of the predefined search groups"),
			detailMessage,
		)
		return
	}

	// Creates a new request for AdHocCommand
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for AdHocCommand on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for AdHocCommand
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for AdHocCommand on %s", o.endpoint),
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
	_ resource.Resource                = &adHocCommandResource{}
	_ resource.ResourceWithConfigure   = &adHocCommandResource{}
	_ resource.ResourceWithImportState = &adHocCommandResource{}
)

// NewAdHocCommandResource is a helper function to simplify the provider implementation.
func NewAdHocCommandResource() resource.Resource {
	return &adHocCommandResource{}
}

type adHocCommandResource struct {
	client   c.Client
	endpoint string
}

func (o *adHocCommandResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/ad_hoc_commands/"
}

func (o adHocCommandResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ad_hoc_command"
}

func (o adHocCommandResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"AdHocCommand",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
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
				"credential": {
					Description: "Credential",
					Type:        types.Int64Type,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
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
				"module_args": {
					Description: "Module args",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"module_name": {
					Description: "Module name",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`command`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"command", "shell", "yum", "apt", "apt_key", "apt_repository", "apt_rpm", "service", "group", "user", "mount", "ping", "selinux", "setup", "win_ping", "win_service", "win_updates", "win_group", "win_user"}...),
					},
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
				// Write only elements
				// Data only elements
				"canceled_on": {
					Description: "The date and time when the cancel request was sent.",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"controller_node": {
					Description: "The instance that managed the execution environment.",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"elapsed": {
					Description: "Elapsed time in seconds that the job ran.",
					Computed:    true,
					Type:        types.Float64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"execution_node": {
					Description: "The node the job executed on.",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"failed": {
					Description: "",
					Computed:    true,
					Type:        types.BoolType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"finished": {
					Description: "The date and time the job finished execution.",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"id": {
					Description: "Database ID for this ad hoc command.",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"job_explanation": {
					Description: "A status field to indicate the state of the job if it wasn't able to run and capture stdout",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"launch_type": {
					Description: "",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"manual", "relaunch", "callback", "scheduled", "dependency", "workflow", "webhook", "sync", "scm"}...),
					},
				},
				"launched_by": {
					Description: "",
					Computed:    true,
					Type:        types.Int64Type,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"name": {
					Description: "Name of this ad hoc command.",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"started": {
					Description: "The date and time the job was queued for starting.",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
				"status": {
					Description: "",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"new", "pending", "waiting", "running", "successful", "failed", "error", "canceled"}...),
					},
				},
				"work_unit_id": {
					Description: "The Receptor work unit ID associated with this job.",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
			},
		}), nil
}

func (o *adHocCommandResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var id, err = strconv.ParseInt(request.ID, 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the AdHocCommand.", request.ID),
			err.Error(),
		)
		return
	}
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (o *adHocCommandResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state adHocCommandTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for AdHocCommand
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for AdHocCommand on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new AdHocCommand resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for AdHocCommand on %s", o.endpoint),
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

func (o *adHocCommandResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state adHocCommandTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for AdHocCommand
	var r *http.Request
	var id = state.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for AdHocCommand on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for AdHocCommand from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for AdHocCommand on %s", o.endpoint),
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

func (o *adHocCommandResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state adHocCommandTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for AdHocCommand
	var r *http.Request
	var id = plan.ID.ValueInt64()
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for AdHocCommand on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new AdHocCommand resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for AdHocCommand on %s", o.endpoint),
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

func (o *adHocCommandResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state adHocCommandTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for AdHocCommand
	var r *http.Request
	var id = state.ID
	var endpoint = p.Clean(fmt.Sprintf("%s/%v", o.endpoint, id.ValueInt64())) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodDelete, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for AdHocCommand on %s for delete", o.endpoint),
			err.Error(),
		)
		return
	}

	// Delete existing AdHocCommand
	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to delete resource for AdHocCommand on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}
