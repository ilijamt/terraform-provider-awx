package aws

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/models"
	"github.com/ilijamt/terraform-provider-awx/internal/resource"
)

var (
	_ resource.Updater                = &TerraformModel{}
	_ resource.Cloner[TerraformModel] = &TerraformModel{}
	_ resource.RequestBody            = &TerraformModel{}
	_ resource.Credential             = &TerraformModel{}
	_ resource.Id                     = &TerraformModel{}
)

// TerraformModel maps the schema for Credential aws
type TerraformModel struct {
	// ID "Database ID for this credential."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// Name "Name of this credential"
	Name types.String `tfsdk:"name" json:"name"`
	// Description "Description of this credential"
	Description types.String `tfsdk:"description" json:"description"`
	// Organization "Organization of this credential"
	Organization types.Int64 `tfsdk:"organization" json:"organization"`
	// Username "Access Key"
	Username types.String `tfsdk:"username" json:"username"`
	// Password "Secret Key"
	Password types.String `tfsdk:"password" json:"password"`
	// SecurityToken "STS Token"
	SecurityToken types.String `tfsdk:"security_token" json:"security_token"`

	// internal variables that are required for the request to finish
	// successfully
	userId           int64
	credentialTypeId int64
}

func (o *TerraformModel) GetId() (string, error) {
	if o.ID.IsNull() || o.ID.IsUnknown() {
		return "", fmt.Errorf("id not set")
	}
	return o.ID.String(), nil
}

func (o *TerraformModel) Data() models.Credential {
	var inputs = map[string]any{
		"username": o.Username.ValueString(),
		"password": o.Password.ValueString(),
	}
	if !o.SecurityToken.IsNull() && !o.SecurityToken.IsUnknown() {
		inputs["security_token"] = o.SecurityToken.ValueString()
	}

	return models.Credential{
		CredentialType: o.credentialTypeId,
		Inputs:         inputs,
		User:           o.userId,
		Name:           o.Name.ValueString(),
		Description:    o.Description.ValueString(),
		Organization:   o.Organization.ValueInt64Pointer(),
	}
}

func (o *TerraformModel) RequestBody() ([]byte, error) {
	return json.Marshal(o.Data())
}

// Clone the object
func (o *TerraformModel) Clone() TerraformModel {
	return TerraformModel{
		ID:            o.ID,
		Name:          o.Name,
		Description:   o.Description,
		Organization:  o.Organization,
		Username:      o.Username,
		Password:      o.Password,
		SecurityToken: o.SecurityToken,
	}
}

func (o *TerraformModel) setName(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *TerraformModel) setDescription(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *TerraformModel) setOrganization(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Organization, data)
}

func (o *TerraformModel) setUsername(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Username, data, false)
}

func (o *TerraformModel) setId(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *TerraformModel) UpdateWithApiData(callee resource.Callee, source resource.Source, data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("data is empty")
	}

	var fieldUTO []helpers.FieldMapping
	o.Organization = types.Int64Null()
	if val, ok := data["organization"]; ok && val != nil {
		fieldUTO = append(
			fieldUTO,
			helpers.FieldMapping{APIField: "organization", Setter: o.setOrganization},
		)
	}

	// Set the default items to the values in the API payload
	var fieldMappings = append(
		[]helpers.FieldMapping{
			{APIField: "id", Setter: o.setId},
			{APIField: "name", Setter: o.setName},
			{APIField: "description", Setter: o.setDescription},
		},
		fieldUTO...,
	)

	// We need to process all the inputs that are not a secret
	// if an input is a secret, then the value will be $encrypted$ which is not useful
	// so we skip those fields
	if inputs, ok := data["inputs"].(map[string]any); ok {
		fieldMappings = append(
			fieldMappings,
			helpers.FieldMapping{
				APIField: "username",
				Setter:   o.setUsername,
				Data:     inputs,
			},
		)
	}

	diags, _ = helpers.ApplyFieldMappings(data, fieldMappings...)
	return diags, nil
}
