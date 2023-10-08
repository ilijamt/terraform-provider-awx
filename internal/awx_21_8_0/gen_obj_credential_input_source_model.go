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
	return credentialInputSourceTerraformModel{
		Description:      o.Description,
		ID:               o.ID,
		InputFieldName:   o.InputFieldName,
		Metadata:         o.Metadata,
		SourceCredential: o.SourceCredential,
		TargetCredential: o.TargetCredential,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for CredentialInputSource
func (o *credentialInputSourceTerraformModel) BodyRequest() (req credentialInputSourceBodyRequestModel) {
	req.Description = o.Description.ValueString()
	req.InputFieldName = o.InputFieldName.ValueString()
	req.Metadata = json.RawMessage(o.Metadata.ValueString())
	req.SourceCredential = o.SourceCredential.ValueInt64()
	req.TargetCredential = o.TargetCredential.ValueInt64()
	return
}

func (o *credentialInputSourceTerraformModel) setDescription(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *credentialInputSourceTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *credentialInputSourceTerraformModel) setInputFieldName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.InputFieldName, data, false)
}

func (o *credentialInputSourceTerraformModel) setMetadata(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.Metadata, data, false)
}

func (o *credentialInputSourceTerraformModel) setSourceCredential(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.SourceCredential, data)
}

func (o *credentialInputSourceTerraformModel) setTargetCredential(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.TargetCredential, data)
}

func (o *credentialInputSourceTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setInputFieldName(data["input_field_name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setMetadata(data["metadata"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSourceCredential(data["source_credential"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setTargetCredential(data["target_credential"]); dg.HasError() {
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
