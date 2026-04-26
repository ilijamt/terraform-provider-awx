package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type adHocCommandTerraformModel struct {
	BecomeEnabled        types.Bool    `tfsdk:"become_enabled" json:"become_enabled"`
	CanceledOn           types.String  `tfsdk:"canceled_on" json:"canceled_on"`
	ControllerNode       types.String  `tfsdk:"controller_node" json:"controller_node"`
	Credential           types.Int64   `tfsdk:"credential" json:"credential"`
	DiffMode             types.Bool    `tfsdk:"diff_mode" json:"diff_mode"`
	Elapsed              types.Float64 `tfsdk:"elapsed" json:"elapsed"`
	ExecutionEnvironment types.Int64   `tfsdk:"execution_environment" json:"execution_environment"`
	ExecutionNode        types.String  `tfsdk:"execution_node" json:"execution_node"`
	ExtraVars            types.String  `tfsdk:"extra_vars" json:"extra_vars"`
	Failed               types.Bool    `tfsdk:"failed" json:"failed"`
	Finished             types.String  `tfsdk:"finished" json:"finished"`
	Forks                types.Int64   `tfsdk:"forks" json:"forks"`
	ID                   types.Int64   `tfsdk:"id" json:"id"`
	Inventory            types.Int64   `tfsdk:"inventory" json:"inventory"`
	JobExplanation       types.String  `tfsdk:"job_explanation" json:"job_explanation"`
	JobType              types.String  `tfsdk:"job_type" json:"job_type"`
	LaunchType           types.String  `tfsdk:"launch_type" json:"launch_type"`
	LaunchedBy           types.Int64   `tfsdk:"launched_by" json:"launched_by"`
	Limit                types.String  `tfsdk:"limit" json:"limit"`
	ModuleArgs           types.String  `tfsdk:"module_args" json:"module_args"`
	ModuleName           types.String  `tfsdk:"module_name" json:"module_name"`
	Name                 types.String  `tfsdk:"name" json:"name"`
	Started              types.String  `tfsdk:"started" json:"started"`
	Status               types.String  `tfsdk:"status" json:"status"`
	Verbosity            types.String  `tfsdk:"verbosity" json:"verbosity"`
	WorkUnitId           types.String  `tfsdk:"work_unit_id" json:"work_unit_id"`
}

func (o *adHocCommandTerraformModel) Clone() adHocCommandTerraformModel {
	return *o
}

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
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetBool(&o.BecomeEnabled, data["become_enabled"]))
	collect(helpers.AttrValueSetString(&o.CanceledOn, data["canceled_on"], false))
	collect(helpers.AttrValueSetString(&o.ControllerNode, data["controller_node"], false))
	collect(helpers.AttrValueSetInt64(&o.Credential, data["credential"]))
	collect(helpers.AttrValueSetBool(&o.DiffMode, data["diff_mode"]))
	collect(helpers.AttrValueSetFloat64(&o.Elapsed, data["elapsed"]))
	collect(helpers.AttrValueSetInt64(&o.ExecutionEnvironment, data["execution_environment"]))
	collect(helpers.AttrValueSetString(&o.ExecutionNode, data["execution_node"], false))
	collect(helpers.AttrValueSetJsonString(&o.ExtraVars, data["extra_vars"], false))
	collect(helpers.AttrValueSetBool(&o.Failed, data["failed"]))
	collect(helpers.AttrValueSetString(&o.Finished, data["finished"], false))
	collect(helpers.AttrValueSetInt64(&o.Forks, data["forks"]))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetInt64(&o.Inventory, data["inventory"]))
	collect(helpers.AttrValueSetString(&o.JobExplanation, data["job_explanation"], false))
	collect(helpers.AttrValueSetString(&o.JobType, data["job_type"], false))
	collect(helpers.AttrValueSetString(&o.LaunchType, data["launch_type"], false))
	collect(helpers.AttrValueSetInt64(&o.LaunchedBy, data["launched_by"]))
	collect(helpers.AttrValueSetString(&o.Limit, data["limit"], false))
	collect(helpers.AttrValueSetString(&o.ModuleArgs, data["module_args"], false))
	collect(helpers.AttrValueSetString(&o.ModuleName, data["module_name"], false))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetString(&o.Started, data["started"], false))
	collect(helpers.AttrValueSetString(&o.Status, data["status"], false))
	collect(helpers.AttrValueSetString(&o.Verbosity, data["verbosity"], false))
	collect(helpers.AttrValueSetString(&o.WorkUnitId, data["work_unit_id"], false))
	return diags, nil
}

type adHocCommandBodyRequestModel struct {
	BecomeEnabled        bool            `json:"become_enabled"`
	Credential           int64           `json:"credential,omitempty"`
	DiffMode             bool            `json:"diff_mode"`
	ExecutionEnvironment int64           `json:"execution_environment,omitempty"`
	ExtraVars            json.RawMessage `json:"extra_vars,omitempty"`
	Forks                int64           `json:"forks,omitempty"`
	Inventory            int64           `json:"inventory,omitempty"`
	JobType              string          `json:"job_type,omitempty"`
	Limit                string          `json:"limit,omitempty"`
	ModuleArgs           string          `json:"module_args,omitempty"`
	ModuleName           string          `json:"module_name,omitempty"`
	Verbosity            string          `json:"verbosity,omitempty"`
}
