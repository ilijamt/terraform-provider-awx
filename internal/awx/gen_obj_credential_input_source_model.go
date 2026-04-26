package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type credentialInputSourceTerraformModel struct {
	Description      types.String `tfsdk:"description" json:"description"`
	ID               types.Int64  `tfsdk:"id" json:"id"`
	InputFieldName   types.String `tfsdk:"input_field_name" json:"input_field_name"`
	Metadata         types.String `tfsdk:"metadata" json:"metadata"`
	SourceCredential types.Int64  `tfsdk:"source_credential" json:"source_credential"`
	TargetCredential types.Int64  `tfsdk:"target_credential" json:"target_credential"`
}

func (o *credentialInputSourceTerraformModel) Clone() credentialInputSourceTerraformModel {
	return *o
}

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
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.InputFieldName, data["input_field_name"], false))
	collect(helpers.AttrValueSetJsonString(&o.Metadata, data["metadata"], false))
	collect(helpers.AttrValueSetInt64(&o.SourceCredential, data["source_credential"]))
	collect(helpers.AttrValueSetInt64(&o.TargetCredential, data["target_credential"]))
	return diags, nil
}

type credentialInputSourceBodyRequestModel struct {
	Description      string          `json:"description,omitempty"`
	InputFieldName   string          `json:"input_field_name"`
	Metadata         json.RawMessage `json:"metadata,omitempty"`
	SourceCredential int64           `json:"source_credential"`
	TargetCredential int64           `json:"target_credential"`
}
