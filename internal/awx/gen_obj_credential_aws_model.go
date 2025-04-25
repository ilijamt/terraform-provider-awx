package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/resource"
)

var (
	_ resource.Updater                             = &awsCredentialTerraformModel{}
	_ resource.Cloner[awsCredentialTerraformModel] = &awsCredentialTerraformModel{}
	_ resource.Body                                = &awsCredentialBodyRequestModel{}
)

// awsCredentialBodyRequestModel maps the schema for Credential AWS for creating and updating the data
type awsCredentialBodyRequestModel struct {
	// Name "Name of this credential"
	Name string `json:"name"`
	// Description "Description of this credential"
	Description string `json:"description,omitempty"`
	// Organization "Organization of this credential"
	Organization int64 `json:"organization,omitempty"`
	// Username "Access Key"
	Username string `json:"username"`
	// Password "Secret Key"
	Password string `json:"password"`
	// SecurityToken "STS Token"
	SecurityToken string `json:"security_token,omitempty"`
}

func (o awsCredentialBodyRequestModel) MarshalJSON() ([]byte, error) { return json.Marshal(o) }

// awsCredentialTerraformModel maps the schema for Credential AWS
type awsCredentialTerraformModel struct {
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
}

// Clone the object
func (o *awsCredentialTerraformModel) Clone() awsCredentialTerraformModel {
	return awsCredentialTerraformModel{
		ID:            o.ID,
		Name:          o.Name,
		Description:   o.Description,
		Organization:  o.Organization,
		Username:      o.Username,
		Password:      o.Password,
		SecurityToken: o.SecurityToken,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for AWS
func (o *awsCredentialTerraformModel) BodyRequest() (req awsCredentialBodyRequestModel) {
	req.Name = o.Name.ValueString()
	req.Description = o.Description.ValueString()
	req.Organization = o.Organization.ValueInt64()
	req.Username = o.Username.ValueString()
	req.Password = o.Password.ValueString()
	req.SecurityToken = o.SecurityToken.ValueString()
	return
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

func (o *awsCredentialTerraformModel) UpdateWithApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("data is empty")
	}

	diags, _ = helpers.ApplyFieldMappings(
		data,
		[]helpers.FieldMapping{
			{APIField: "id", Setter: o.setId},
			{APIField: "name", Setter: o.setName},
			{APIField: "description", Setter: o.setDescription},
			{APIField: "organization", Setter: o.setOrganization},
			{APIField: "username", Setter: o.setUsername},
			{APIField: "password", Setter: o.setPassword},
			{APIField: "security_token", Setter: o.setSecurityToken},
		},
	)

	return diags, nil
}
