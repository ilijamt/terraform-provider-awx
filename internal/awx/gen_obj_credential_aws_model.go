package awx

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
	_ resource.Updater                             = &awsCredentialTerraformModel{}
	_ resource.Cloner[awsCredentialTerraformModel] = &awsCredentialTerraformModel{}
	_ resource.RequestBody                         = &awsCredentialTerraformModel{}
	_ resource.Credential                          = &awsCredentialTerraformModel{}
	_ resource.Id                                  = &awsCredentialTerraformModel{}
)

// awsCredentialTerraformModel maps the schema for Credential aws
type awsCredentialTerraformModel struct {
	// ID "Database ID for this credential."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// Name "Name of this credential"
	Name types.String `tfsdk:"name" json:"name"`
	// Description "Description of this credential"
	Description types.String `tfsdk:"description" json:"description"`
	// Organization "Organization of this credential"
	Organization types.Int64 `tfsdk:"organization" json:"organization"`
	// User "User of this credential"
	User types.Int64 `tfsdk:"user" json:"user"`
	// Team "Team of this credential"
	Team types.Int64 `tfsdk:"team" json:"team"`
	// Username "Access Key"
	Username types.String `tfsdk:"username" json:"username"`
	// Password "Secret Key"
	Password types.String `tfsdk:"password" json:"password"`
	// SecurityToken "STS Token"
	SecurityToken types.String `tfsdk:"security_token" json:"security_token"`
}

func (o *awsCredentialTerraformModel) GetId() (string, error) {
	if o.ID.IsNull() || o.ID.IsUnknown() {
		return "", fmt.Errorf("id not set")
	}
	return o.ID.String(), nil
}

func (o *awsCredentialTerraformModel) Data() models.Credential {
	var inputs = map[string]any{
		"username": o.Username.ValueString(),
		"password": o.Password.ValueString(),
	}
	if !o.SecurityToken.IsNull() && !o.SecurityToken.IsUnknown() {
		inputs["security_token"] = o.SecurityToken.ValueString()
	}

	return models.Credential{
		CredentialType: 5,
		Inputs:         inputs,
		Name:           o.Name.ValueString(),
		Description:    o.Description.ValueString(),
		Organization:   o.Organization.ValueInt64(),
		User:           o.User.ValueInt64(),
		Team:           o.Team.ValueInt64(),
	}
}

func (o *awsCredentialTerraformModel) RequestBody() ([]byte, error) {
	return json.Marshal(o.Data())
}

// Clone the object
func (o *awsCredentialTerraformModel) Clone() awsCredentialTerraformModel {
	return awsCredentialTerraformModel{
		ID:            o.ID,
		Name:          o.Name,
		Description:   o.Description,
		Organization:  o.Organization,
		User:          o.User,
		Team:          o.Team,
		Username:      o.Username,
		Password:      o.Password,
		SecurityToken: o.SecurityToken,
	}
}

func (o *awsCredentialTerraformModel) setName(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *awsCredentialTerraformModel) setDescription(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *awsCredentialTerraformModel) setOrganization(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Organization, data)
}

func (o *awsCredentialTerraformModel) setUser(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.User, data)
}

func (o *awsCredentialTerraformModel) setTeam(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Team, data)
}

func (o *awsCredentialTerraformModel) setUsername(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Username, data, false)
}

func (o *awsCredentialTerraformModel) setPassword(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Password, data, false)
}

func (o *awsCredentialTerraformModel) setSecurityToken(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.SecurityToken, data, false)
}

func (o *awsCredentialTerraformModel) setId(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *awsCredentialTerraformModel) UpdateWithApiData(callee resource.Callee, source resource.Source, data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("data is empty")
	}

	// Set the User, Team or Organization to the proper value
	// only one of them can be set for the credential, and they are write-only properties
	// setting them after creation should result in recreation of the resource
	var fieldUTO []helpers.FieldMapping
	if val, ok := data["organization"]; ok && val != nil {
		fieldUTO = append(
			fieldUTO,
			helpers.FieldMapping{APIField: "organization", Setter: o.setOrganization},
		)
	}
	if val, ok := data["user"]; ok && val != nil {
		fieldUTO = append(
			fieldUTO,
			helpers.FieldMapping{APIField: "user", Setter: o.setUser},
		)
	}
	if val, ok := data["team"]; ok && val != nil {
		fieldUTO = append(
			fieldUTO,
			helpers.FieldMapping{APIField: "team", Setter: o.setTeam},
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
