package awx_21_8_0

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// credentialTerraformModel maps the schema for Credential when using Data Source
type credentialTerraformModel struct {
	// Cloud ""
	Cloud types.Bool `tfsdk:"cloud" json:"cloud"`
	// CredentialType "Specify the type of credential you want to create. Refer to the documentation for details on each type."
	CredentialType types.Int64 `tfsdk:"credential_type" json:"credential_type"`
	// Description "Optional description of this credential."
	Description types.String `tfsdk:"description" json:"description"`
	// ID "Database ID for this credential."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// Inputs "Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax."
	Inputs types.String `tfsdk:"inputs" json:"inputs"`
	// Kind ""
	Kind types.String `tfsdk:"kind" json:"kind"`
	// Kubernetes ""
	Kubernetes types.Bool `tfsdk:"kubernetes" json:"kubernetes"`
	// Managed ""
	Managed types.Bool `tfsdk:"managed" json:"managed"`
	// Name "Name of this credential."
	Name types.String `tfsdk:"name" json:"name"`
	// Organization "Inherit permissions from organization roles. If provided on creation, do not give either user or team."
	Organization types.Int64 `tfsdk:"organization" json:"organization"`
}

// Clone the object
func (o *credentialTerraformModel) Clone() credentialTerraformModel {
	return credentialTerraformModel{
		Cloud:          o.Cloud,
		CredentialType: o.CredentialType,
		Description:    o.Description,
		ID:             o.ID,
		Inputs:         o.Inputs,
		Kind:           o.Kind,
		Kubernetes:     o.Kubernetes,
		Managed:        o.Managed,
		Name:           o.Name,
		Organization:   o.Organization,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Credential
func (o *credentialTerraformModel) BodyRequest() (req credentialBodyRequestModel) {
	req.CredentialType = o.CredentialType.ValueInt64()
	req.Description = o.Description.ValueString()
	req.Inputs = json.RawMessage(o.Inputs.ValueString())
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	return
}

func (o *credentialTerraformModel) setCloud(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.Cloud, data)
}

func (o *credentialTerraformModel) setCredentialType(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.CredentialType, data)
}

func (o *credentialTerraformModel) setDescription(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *credentialTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *credentialTerraformModel) setInputs(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.Inputs, data, false)
}

func (o *credentialTerraformModel) setKind(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Kind, data, false)
}

func (o *credentialTerraformModel) setKubernetes(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.Kubernetes, data)
}

func (o *credentialTerraformModel) setManaged(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.Managed, data)
}

func (o *credentialTerraformModel) setName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *credentialTerraformModel) setOrganization(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Organization, data)
}

func (o *credentialTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setCloud(data["cloud"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setCredentialType(data["credential_type"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setInputs(data["inputs"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setKind(data["kind"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setKubernetes(data["kubernetes"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setManaged(data["managed"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setName(data["name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setOrganization(data["organization"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// credentialBodyRequestModel maps the schema for Credential for creating and updating the data
type credentialBodyRequestModel struct {
	// CredentialType "Specify the type of credential you want to create. Refer to the documentation for details on each type."
	CredentialType int64 `json:"credential_type"`
	// Description "Optional description of this credential."
	Description string `json:"description,omitempty"`
	// Inputs "Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax."
	Inputs json.RawMessage `json:"inputs,omitempty"`
	// Name "Name of this credential."
	Name string `json:"name"`
	// Organization "Inherit permissions from organization roles. If provided on creation, do not give either user or team."
	Organization int64 `json:"organization,omitempty"`
}

type credentialObjectRolesModel struct {
	ID    types.Int64 `tfsdk:"id"`
	Roles types.Map   `tfsdk:"roles"`
}
