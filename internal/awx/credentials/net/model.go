package net

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
	_ resource.Updater                = &terraformModel{}
	_ resource.Cloner[terraformModel] = &terraformModel{}
	_ resource.RequestBody            = &terraformModel{}
	_ resource.Credential             = &terraformModel{}
	_ resource.Id                     = &terraformModel{}
)

// terraformModel maps the schema for Credential net
type terraformModel struct {
	// ID "Database ID for this credential."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// Name "Name of this credential"
	Name types.String `tfsdk:"name" json:"name"`
	// Description "Description of this credential"
	Description types.String `tfsdk:"description" json:"description"`
	// Organization "Organization of this credential"
	Organization types.Int64 `tfsdk:"organization" json:"organization"`
	// Username "Username"
	Username types.String `tfsdk:"username" json:"username"`
	// Password "Password"
	Password types.String `tfsdk:"password" json:"password"`
	// SshKeyData "SSH Private Key"
	SshKeyData types.String `tfsdk:"ssh_key_data" json:"ssh_key_data"`
	// SshKeyUnlock "Private Key Passphrase"
	SshKeyUnlock types.String `tfsdk:"ssh_key_unlock" json:"ssh_key_unlock"`
	// Authorize "Authorize"
	Authorize types.Bool `tfsdk:"authorize" json:"authorize"`
	// AuthorizePassword "Authorize Password"
	AuthorizePassword types.String `tfsdk:"authorize_password" json:"authorize_password"`

	// internal variables that are required for the request to finish
	// successfully
	userId           int64
	credentialTypeId int64
}

func (o *terraformModel) GetId() (string, error) {
	if o.ID.IsNull() || o.ID.IsUnknown() {
		return "", fmt.Errorf("id not set")
	}
	return o.ID.String(), nil
}

func (o *terraformModel) Data() models.Credential {
	var inputs = map[string]any{
		"username": o.Username.ValueString(),
	}
	if !o.Password.IsNull() && !o.Password.IsUnknown() {
		inputs["password"] = o.Password.ValueString()
	}
	if !o.SshKeyData.IsNull() && !o.SshKeyData.IsUnknown() {
		inputs["ssh_key_data"] = o.SshKeyData.ValueString()
	}
	if !o.SshKeyUnlock.IsNull() && !o.SshKeyUnlock.IsUnknown() {
		inputs["ssh_key_unlock"] = o.SshKeyUnlock.ValueString()
	}
	if !o.Authorize.IsNull() && !o.Authorize.IsUnknown() {
		inputs["authorize"] = o.Authorize.ValueBool()
	}
	if !o.AuthorizePassword.IsNull() && !o.AuthorizePassword.IsUnknown() {
		inputs["authorize_password"] = o.AuthorizePassword.ValueString()
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

func (o *terraformModel) RequestBody() ([]byte, error) {
	return json.Marshal(o.Data())
}

// Clone the object
func (o *terraformModel) Clone() terraformModel {
	return terraformModel{
		ID:                o.ID,
		Name:              o.Name,
		Description:       o.Description,
		Organization:      o.Organization,
		Username:          o.Username,
		Password:          o.Password,
		SshKeyData:        o.SshKeyData,
		SshKeyUnlock:      o.SshKeyUnlock,
		Authorize:         o.Authorize,
		AuthorizePassword: o.AuthorizePassword,
	}
}

func (o *terraformModel) setName(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *terraformModel) setDescription(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *terraformModel) setOrganization(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Organization, data)
}

func (o *terraformModel) setUsername(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Username, data, false)
}

func (o *terraformModel) setAuthorize(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.Authorize, data)
}

func (o *terraformModel) setId(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *terraformModel) UpdateWithApiData(callee resource.Callee, source resource.Source, data map[string]any) (diags diag.Diagnostics, _ error) {
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
			helpers.FieldMapping{
				APIField: "authorize",
				Setter:   o.setAuthorize,
				Data:     inputs,
			},
		)
	}

	diags, _ = helpers.ApplyFieldMappings(data, fieldMappings...)
	return diags, nil
}
