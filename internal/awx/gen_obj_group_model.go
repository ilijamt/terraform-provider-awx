package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type groupTerraformModel struct {
	Description types.String `tfsdk:"description" json:"description"`
	ID          types.Int64  `tfsdk:"id" json:"id"`
	Inventory   types.Int64  `tfsdk:"inventory" json:"inventory"`
	Name        types.String `tfsdk:"name" json:"name"`
	Variables   types.String `tfsdk:"variables" json:"variables"`
}

func (o *groupTerraformModel) Clone() groupTerraformModel {
	return *o
}

func (o *groupTerraformModel) BodyRequest() *groupBodyRequestModel {
	var req groupBodyRequestModel
	req.Description = o.Description.ValueString()
	req.Inventory = o.Inventory.ValueInt64()
	req.Name = o.Name.ValueString()
	req.Variables = json.RawMessage(o.Variables.String())
	return &req
}

func (o *groupTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetInt64(&o.Inventory, data["inventory"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetJsonString(&o.Variables, data["variables"], false))
	return diags, nil
}

type groupBodyRequestModel struct {
	Description string          `json:"description,omitempty"`
	Inventory   int64           `json:"inventory"`
	Name        string          `json:"name"`
	Variables   json.RawMessage `json:"variables,omitempty"`
}
