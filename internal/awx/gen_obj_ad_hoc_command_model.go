package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
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
func (o *adHocCommandTerraformModel) Clone() adHocCommandTerraformModel {
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
func (o *adHocCommandTerraformModel) BodyRequest() (req adHocCommandBodyRequestModel) {
	req.BecomeEnabled = o.BecomeEnabled.ValueBool()
	req.Credential = o.Credential.ValueInt64()
	req.DiffMode = o.DiffMode.ValueBool()
	req.ExecutionEnvironment = o.ExecutionEnvironment.ValueInt64()
	req.ExtraVars = json.RawMessage(o.ExtraVars.String())
	req.Forks = o.Forks.ValueInt64()
	req.Inventory = o.Inventory.ValueInt64()
	req.JobType = o.JobType.ValueString()
	req.Limit = o.Limit.ValueString()
	req.ModuleArgs = o.ModuleArgs.ValueString()
	req.ModuleName = o.ModuleName.ValueString()
	req.Verbosity = o.Verbosity.ValueString()
	return
}

func (o *adHocCommandTerraformModel) setBecomeEnabled(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.BecomeEnabled, data)
}

func (o *adHocCommandTerraformModel) setCanceledOn(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.CanceledOn, data, false)
}

func (o *adHocCommandTerraformModel) setControllerNode(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.ControllerNode, data, false)
}

func (o *adHocCommandTerraformModel) setCredential(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Credential, data)
}

func (o *adHocCommandTerraformModel) setDiffMode(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.DiffMode, data)
}

func (o *adHocCommandTerraformModel) setElapsed(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetFloat64(&o.Elapsed, data)
}

func (o *adHocCommandTerraformModel) setExecutionEnvironment(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ExecutionEnvironment, data)
}

func (o *adHocCommandTerraformModel) setExecutionNode(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.ExecutionNode, data, false)
}

func (o *adHocCommandTerraformModel) setExtraVars(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetJsonString(&o.ExtraVars, data, false)
}

func (o *adHocCommandTerraformModel) setFailed(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.Failed, data)
}

func (o *adHocCommandTerraformModel) setFinished(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Finished, data, false)
}

func (o *adHocCommandTerraformModel) setForks(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Forks, data)
}

func (o *adHocCommandTerraformModel) setID(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *adHocCommandTerraformModel) setInventory(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Inventory, data)
}

func (o *adHocCommandTerraformModel) setJobExplanation(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.JobExplanation, data, false)
}

func (o *adHocCommandTerraformModel) setJobType(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.JobType, data, false)
}

func (o *adHocCommandTerraformModel) setLaunchType(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.LaunchType, data, false)
}

func (o *adHocCommandTerraformModel) setLaunchedBy(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.LaunchedBy, data)
}

func (o *adHocCommandTerraformModel) setLimit(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Limit, data, false)
}

func (o *adHocCommandTerraformModel) setModuleArgs(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.ModuleArgs, data, false)
}

func (o *adHocCommandTerraformModel) setModuleName(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.ModuleName, data, false)
}

func (o *adHocCommandTerraformModel) setName(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *adHocCommandTerraformModel) setStarted(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Started, data, false)
}

func (o *adHocCommandTerraformModel) setStatus(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Status, data, false)
}

func (o *adHocCommandTerraformModel) setVerbosity(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Verbosity, data, false)
}

func (o *adHocCommandTerraformModel) setWorkUnitId(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.WorkUnitId, data, false)
}

func (o *adHocCommandTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
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
	ExtraVars json.RawMessage `json:"extra_vars,omitempty"`
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
