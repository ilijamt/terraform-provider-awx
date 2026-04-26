package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type credentialTerraformModel struct {
	Cloud          types.Bool   `tfsdk:"cloud" json:"cloud"`
	CredentialType types.Int64  `tfsdk:"credential_type" json:"credential_type"`
	Description    types.String `tfsdk:"description" json:"description"`
	ID             types.Int64  `tfsdk:"id" json:"id"`
	Inputs         types.String `tfsdk:"inputs" json:"inputs"`
	Kind           types.String `tfsdk:"kind" json:"kind"`
	Kubernetes     types.Bool   `tfsdk:"kubernetes" json:"kubernetes"`
	Managed        types.Bool   `tfsdk:"managed" json:"managed"`
	Name           types.String `tfsdk:"name" json:"name"`
	Organization   types.Int64  `tfsdk:"organization" json:"organization"`
	Team           types.Int64  `tfsdk:"team" json:"team"`
	User           types.Int64  `tfsdk:"user" json:"user"`
}

func (o *credentialTerraformModel) Clone() credentialTerraformModel {
	return *o
}

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
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetBool(&o.Cloud, data["cloud"]))
	collect(helpers.AttrValueSetInt64(&o.CredentialType, data["credential_type"]))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetJsonString(&o.Inputs, data["inputs"], false))
	collect(helpers.AttrValueSetString(&o.Kind, data["kind"], false))
	collect(helpers.AttrValueSetBool(&o.Kubernetes, data["kubernetes"]))
	collect(helpers.AttrValueSetBool(&o.Managed, data["managed"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	return diags, nil
}

type credentialBodyRequestModel struct {
	CredentialType int64           `json:"credential_type"`
	Description    string          `json:"description,omitempty"`
	Inputs         json.RawMessage `json:"inputs,omitempty"`
	Name           string          `json:"name"`
	Organization   int64           `json:"organization,omitempty"`
	Team           int64           `json:"team,omitempty"`
	User           int64           `json:"user,omitempty"`
}
