package awx

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
	// Team "Write-only field used to add team to owner role. If provided, do not give either user or organization. Only valid for creation."
	Team types.Int64 `tfsdk:"team" json:"team"`
	// User "Write-only field used to add user to owner role. If provided, do not give either team or organization. Only valid for creation."
	User types.Int64 `tfsdk:"user" json:"user"`
}

// Clone the object
func (o *credentialTerraformModel) Clone() credentialTerraformModel {
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Credential
func (o *credentialTerraformModel) BodyRequest() *credentialBodyRequestModel {
	var req credentialBodyRequestModel
	req.CredentialType = o.CredentialType.ValueInt64()
	req.Description = o.Description.ValueString()
	req.Inputs = json.RawMessage(o.Inputs.ValueString())
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	return &req
}

func (o *credentialTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.Cloud, data["cloud"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.CredentialType, data["credential_type"])
		diags.Append(dg...)
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
		dg, _ := helpers.AttrValueSetJsonString(&o.Inputs, data["inputs"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Kind, data["kind"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.Kubernetes, data["kubernetes"])
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
		dg, _ := helpers.AttrValueSetInt64(&o.Organization, data["organization"])
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
	// Team "Write-only field used to add team to owner role. If provided, do not give either user or organization. Only valid for creation."
	Team int64 `json:"team,omitempty"`
	// User "Write-only field used to add user to owner role. If provided, do not give either team or organization. Only valid for creation."
	User int64 `json:"user,omitempty"`
}
