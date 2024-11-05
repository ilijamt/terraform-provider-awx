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
	return credentialTypeTerraformModel{
		Description: o.Description,
		ID:          o.ID,
		Injectors:   o.Injectors,
		Inputs:      o.Inputs,
		Kind:        o.Kind,
		Managed:     o.Managed,
		Name:        o.Name,
		Namespace:   o.Namespace,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for CredentialType
func (o *credentialTypeTerraformModel) BodyRequest() (req credentialTypeBodyRequestModel) {
	req.Description = o.Description.ValueString()
	req.Injectors = json.RawMessage(o.Injectors.ValueString())
	req.Inputs = json.RawMessage(o.Inputs.ValueString())
	req.Kind = o.Kind.ValueString()
	req.Name = o.Name.ValueString()
	return
}

func (o *credentialTypeTerraformModel) setDescription(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *credentialTypeTerraformModel) setID(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *credentialTypeTerraformModel) setInjectors(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetJsonString(&o.Injectors, data, false)
}

func (o *credentialTypeTerraformModel) setInputs(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetJsonString(&o.Inputs, data, false)
}

func (o *credentialTypeTerraformModel) setKind(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Kind, data, false)
}

func (o *credentialTypeTerraformModel) setManaged(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.Managed, data)
}

func (o *credentialTypeTerraformModel) setName(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *credentialTypeTerraformModel) setNamespace(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Namespace, data, false)
}

func (o *credentialTypeTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setInjectors(data["injectors"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setInputs(data["inputs"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setKind(data["kind"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setManaged(data["managed"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setName(data["name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setNamespace(data["namespace"]); dg.HasError() {
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
