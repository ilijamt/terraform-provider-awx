package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// credentialInputSourceTerraformModel maps the schema for CredentialInputSource when using Data Source
type credentialInputSourceTerraformModel struct {
	// Description "Optional description of this credential input source."
	Description types.String `tfsdk:"description" json:"description"`
	// ID "Database ID for this credential input source."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// InputFieldName ""
	InputFieldName types.String `tfsdk:"input_field_name" json:"input_field_name"`
	// Metadata ""
	Metadata types.String `tfsdk:"metadata" json:"metadata"`
	// SourceCredential ""
	SourceCredential types.Int64 `tfsdk:"source_credential" json:"source_credential"`
	// TargetCredential ""
	TargetCredential types.Int64 `tfsdk:"target_credential" json:"target_credential"`
}

// Clone the object
func (o *credentialInputSourceTerraformModel) Clone() credentialInputSourceTerraformModel {
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for CredentialInputSource
func (o *credentialInputSourceTerraformModel) BodyRequest() *credentialInputSourceBodyRequestModel {
	var req credentialInputSourceBodyRequestModel
	req.Description = o.Description.ValueString()
	req.InputFieldName = o.InputFieldName.ValueString()
	req.Metadata = json.RawMessage(o.Metadata.ValueString())
	req.SourceCredential = o.SourceCredential.ValueInt64()
	req.TargetCredential = o.TargetCredential.ValueInt64()
	return &req
}

func (o *credentialInputSourceTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
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
		dg, _ := helpers.AttrValueSetString(&o.InputFieldName, data["input_field_name"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetJsonString(&o.Metadata, data["metadata"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.SourceCredential, data["source_credential"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.TargetCredential, data["target_credential"])
		diags.Append(dg...)
	}
	return diags, nil
}

// credentialInputSourceBodyRequestModel maps the schema for CredentialInputSource for creating and updating the data
type credentialInputSourceBodyRequestModel struct {
	// Description "Optional description of this credential input source."
	Description string `json:"description,omitempty"`
	// InputFieldName ""
	InputFieldName string `json:"input_field_name"`
	// Metadata ""
	Metadata json.RawMessage `json:"metadata,omitempty"`
	// SourceCredential ""
	SourceCredential int64 `json:"source_credential"`
	// TargetCredential ""
	TargetCredential int64 `json:"target_credential"`
}
