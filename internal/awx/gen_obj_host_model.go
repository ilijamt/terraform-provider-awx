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
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Host
func (o *hostTerraformModel) BodyRequest() *hostBodyRequestModel {
	var req hostBodyRequestModel
	req.Description = o.Description.ValueString()
	req.Enabled = o.Enabled.ValueBool()
	req.InstanceId = o.InstanceId.ValueString()
	req.Inventory = o.Inventory.ValueInt64()
	req.Name = o.Name.ValueString()
	req.Variables = json.RawMessage(o.Variables.String())
	return &req
}

func (o *hostTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Description, data["description"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.Enabled, data["enabled"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.ID, data["id"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.InstanceId, data["instance_id"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.Inventory, data["inventory"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.LastJob, data["last_job"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.LastJobHostSummary, data["last_job_host_summary"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Name, data["name"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.Variables, data["variables"], false)
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
