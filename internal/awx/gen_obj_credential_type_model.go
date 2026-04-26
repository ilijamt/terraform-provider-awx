package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type credentialTypeTerraformModel struct {
	Description types.String `tfsdk:"description" json:"description"`
	ID          types.Int64  `tfsdk:"id" json:"id"`
	Injectors   types.String `tfsdk:"injectors" json:"injectors"`
	Inputs      types.String `tfsdk:"inputs" json:"inputs"`
	Kind        types.String `tfsdk:"kind" json:"kind"`
	Managed     types.Bool   `tfsdk:"managed" json:"managed"`
	Name        types.String `tfsdk:"name" json:"name"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace"`
}

func (o *credentialTypeTerraformModel) Clone() credentialTypeTerraformModel {
	return *o
}

func (o *credentialTypeTerraformModel) BodyRequest() *credentialTypeBodyRequestModel {
	var req credentialTypeBodyRequestModel
	req.Description = o.Description.ValueString()
	req.Injectors = json.RawMessage(o.Injectors.ValueString())
	req.Inputs = json.RawMessage(o.Inputs.ValueString())
	req.Kind = o.Kind.ValueString()
	req.Name = o.Name.ValueString()
	return &req
}

func (o *credentialTypeTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetJsonString(&o.Injectors, data["injectors"], false))
	collect(helpers.AttrValueSetJsonString(&o.Inputs, data["inputs"], false))
	collect(helpers.AttrValueSetString(&o.Kind, data["kind"], false))
	collect(helpers.AttrValueSetBool(&o.Managed, data["managed"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetString(&o.Namespace, data["namespace"], false))
	return diags, nil
}

type credentialTypeBodyRequestModel struct {
	Description string          `json:"description,omitempty"`
	Injectors   json.RawMessage `json:"injectors,omitempty"`
	Inputs      json.RawMessage `json:"inputs,omitempty"`
	Kind        string          `json:"kind"`
	Name        string          `json:"name"`
}
