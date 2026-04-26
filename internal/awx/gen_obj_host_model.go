package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type hostTerraformModel struct {
	Description        types.String `tfsdk:"description" json:"description"`
	Enabled            types.Bool   `tfsdk:"enabled" json:"enabled"`
	ID                 types.Int64  `tfsdk:"id" json:"id"`
	InstanceId         types.String `tfsdk:"instance_id" json:"instance_id"`
	Inventory          types.Int64  `tfsdk:"inventory" json:"inventory"`
	LastJob            types.Int64  `tfsdk:"last_job" json:"last_job"`
	LastJobHostSummary types.Int64  `tfsdk:"last_job_host_summary" json:"last_job_host_summary"`
	Name               types.String `tfsdk:"name" json:"name"`
	Variables          types.String `tfsdk:"variables" json:"variables"`
}

func (o *hostTerraformModel) Clone() hostTerraformModel {
	return *o
}

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
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetBool(&o.Enabled, data["enabled"]))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.InstanceId, data["instance_id"], false))
	collect(helpers.AttrValueSetInt64(&o.Inventory, data["inventory"]))
	collect(helpers.AttrValueSetInt64(&o.LastJob, data["last_job"]))
	collect(helpers.AttrValueSetInt64(&o.LastJobHostSummary, data["last_job_host_summary"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetJsonString(&o.Variables, data["variables"], false))
	return diags, nil
}

type hostBodyRequestModel struct {
	Description string          `json:"description,omitempty"`
	Enabled     bool            `json:"enabled"`
	InstanceId  string          `json:"instance_id,omitempty"`
	Inventory   int64           `json:"inventory"`
	Name        string          `json:"name"`
	Variables   json.RawMessage `json:"variables,omitempty"`
}
