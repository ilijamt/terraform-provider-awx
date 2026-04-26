package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// credentialTypeTerraformModel maps the schema for CredentialType when using Data Source
type credentialTypeTerraformModel struct {
	// Description "Optional description of this credential type."
	Description types.String `tfsdk:"description" json:"description"`
	// ID "Database ID for this credential type."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// Injectors "Enter injectors using either JSON or YAML syntax. Refer to the documentation for example syntax."
	Injectors types.String `tfsdk:"injectors" json:"injectors"`
	// Inputs "Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax."
	Inputs types.String `tfsdk:"inputs" json:"inputs"`
	// Kind "The credential type"
	Kind types.String `tfsdk:"kind" json:"kind"`
	// Managed "Is the resource managed"
	Managed types.Bool `tfsdk:"managed" json:"managed"`
	// Name "Name of this credential type."
	Name types.String `tfsdk:"name" json:"name"`
	// Namespace "The namespace to which the resource belongs to"
	Namespace types.String `tfsdk:"namespace" json:"namespace"`
}

// Clone the object
func (o *credentialTypeTerraformModel) Clone() credentialTypeTerraformModel {
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for CredentialType
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
	{
		dg, _ := helpers.AttrValueSetString(&o.Description, data["description"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.ID, data["id"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.Injectors, data["injectors"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.Inputs, data["inputs"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Kind, data["kind"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.Managed, data["managed"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Name, data["name"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Namespace, data["namespace"], false)
		diags.Append(dg...)
	}
	return diags, nil
}

// credentialTypeBodyRequestModel maps the schema for CredentialType for creating and updating the data
type credentialTypeBodyRequestModel struct {
	// Description "Optional description of this credential type."
	Description string `json:"description,omitempty"`
	// Injectors "Enter injectors using either JSON or YAML syntax. Refer to the documentation for example syntax."
	Injectors json.RawMessage `json:"injectors,omitempty"`
	// Inputs "Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax."
	Inputs json.RawMessage `json:"inputs,omitempty"`
	// Kind "The credential type"
	Kind string `json:"kind"`
	// Name "Name of this credential type."
	Name string `json:"name"`
}
