package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// hostTerraformModel maps the schema for Host when using Data Source
type hostTerraformModel struct {
	// Description "Optional description of this host."
	Description types.String `tfsdk:"description" json:"description"`
	// Enabled "Is this host online and available for running jobs?"
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
	// ID "Database ID for this host."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// InstanceId "The value used by the remote inventory source to uniquely identify the host"
	InstanceId types.String `tfsdk:"instance_id" json:"instance_id"`
	// Inventory ""
	Inventory types.Int64 `tfsdk:"inventory" json:"inventory"`
	// LastJob ""
	LastJob types.Int64 `tfsdk:"last_job" json:"last_job"`
	// LastJobHostSummary ""
	LastJobHostSummary types.Int64 `tfsdk:"last_job_host_summary" json:"last_job_host_summary"`
	// Name "Name of this host."
	Name types.String `tfsdk:"name" json:"name"`
	// Variables "Host variables in JSON or YAML format."
	Variables types.String `tfsdk:"variables" json:"variables"`
}

// Clone the object
func (o *hostTerraformModel) Clone() hostTerraformModel {
	return hostTerraformModel{
		Description:        o.Description,
		Enabled:            o.Enabled,
		ID:                 o.ID,
		InstanceId:         o.InstanceId,
		Inventory:          o.Inventory,
		LastJob:            o.LastJob,
		LastJobHostSummary: o.LastJobHostSummary,
		Name:               o.Name,
		Variables:          o.Variables,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Host
func (o *hostTerraformModel) BodyRequest() (req hostBodyRequestModel) {
	req.Description = o.Description.ValueString()
	req.Enabled = o.Enabled.ValueBool()
	req.InstanceId = o.InstanceId.ValueString()
	req.Inventory = o.Inventory.ValueInt64()
	req.Name = o.Name.ValueString()
	req.Variables = json.RawMessage(o.Variables.String())
	return
}

func (o *hostTerraformModel) setDescription(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *hostTerraformModel) setEnabled(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.Enabled, data)
}

func (o *hostTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *hostTerraformModel) setInstanceId(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.InstanceId, data, false)
}

func (o *hostTerraformModel) setInventory(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Inventory, data)
}

func (o *hostTerraformModel) setLastJob(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.LastJob, data)
}

func (o *hostTerraformModel) setLastJobHostSummary(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.LastJobHostSummary, data)
}

func (o *hostTerraformModel) setName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *hostTerraformModel) setVariables(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.Variables, data, false)
}

func (o *hostTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setEnabled(data["enabled"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setInstanceId(data["instance_id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setInventory(data["inventory"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setLastJob(data["last_job"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setLastJobHostSummary(data["last_job_host_summary"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setName(data["name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setVariables(data["variables"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// hostBodyRequestModel maps the schema for Host for creating and updating the data
type hostBodyRequestModel struct {
	// Description "Optional description of this host."
	Description string `json:"description,omitempty"`
	// Enabled "Is this host online and available for running jobs?"
	Enabled bool `json:"enabled"`
	// InstanceId "The value used by the remote inventory source to uniquely identify the host"
	InstanceId string `json:"instance_id,omitempty"`
	// Inventory ""
	Inventory int64 `json:"inventory"`
	// Name "Name of this host."
	Name string `json:"name"`
	// Variables "Host variables in JSON or YAML format."
	Variables json.RawMessage `json:"variables,omitempty"`
}

type hostObjectRolesModel struct {
	ID    types.Int64 `tfsdk:"id"`
	Roles types.Map   `tfsdk:"roles"`
}
