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
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for AdHocCommand
func (o *adHocCommandTerraformModel) BodyRequest() *adHocCommandBodyRequestModel {
	var req adHocCommandBodyRequestModel
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
	return &req
}

func (o *adHocCommandTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.BecomeEnabled, data["become_enabled"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.CanceledOn, data["canceled_on"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.ControllerNode, data["controller_node"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.Credential, data["credential"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.DiffMode, data["diff_mode"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetFloat64(&o.Elapsed, data["elapsed"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.ExecutionEnvironment, data["execution_environment"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.ExecutionNode, data["execution_node"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.ExtraVars, data["extra_vars"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.Failed, data["failed"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Finished, data["finished"], false)
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
		dg, _ := helpers.AttrValueSetString(&o.JobExplanation, data["job_explanation"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.JobType, data["job_type"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.LaunchType, data["launch_type"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.LaunchedBy, data["launched_by"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Limit, data["limit"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.ModuleArgs, data["module_args"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.ModuleName, data["module_name"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Name, data["name"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Started, data["started"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Status, data["status"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Verbosity, data["verbosity"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.WorkUnitId, data["work_unit_id"], false)
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
