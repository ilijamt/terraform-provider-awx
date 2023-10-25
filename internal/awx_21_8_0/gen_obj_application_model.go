package awx_21_8_0

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// applicationTerraformModel maps the schema for Application when using Data Source
type applicationTerraformModel struct {
	// AuthorizationGrantType "The Grant type the user must use for acquire tokens for this application."
	AuthorizationGrantType types.String `tfsdk:"authorization_grant_type" json:"authorization_grant_type"`
	// ClientId ""
	ClientId types.String `tfsdk:"client_id" json:"client_id"`
	// ClientSecret "Used for more stringent verification of access to an application when creating a token."
	ClientSecret types.String `tfsdk:"client_secret" json:"client_secret"`
	// ClientType "Set to Public or Confidential depending on how secure the client device is."
	ClientType types.String `tfsdk:"client_type" json:"client_type"`
	// Description "Optional description of this application."
	Description types.String `tfsdk:"description" json:"description"`
	// ID "Database ID for this application."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// Name "Name of this application."
	Name types.String `tfsdk:"name" json:"name"`
	// Organization "Organization containing this application."
	Organization types.Int64 `tfsdk:"organization" json:"organization"`
	// RedirectUris "Allowed URIs list, space separated"
	RedirectUris types.String `tfsdk:"redirect_uris" json:"redirect_uris"`
	// SkipAuthorization "Set True to skip authorization step for completely trusted applications."
	SkipAuthorization types.Bool `tfsdk:"skip_authorization" json:"skip_authorization"`
}

// Clone the object
func (o *applicationTerraformModel) Clone() applicationTerraformModel {
	return applicationTerraformModel{
		AuthorizationGrantType: o.AuthorizationGrantType,
		ClientId:               o.ClientId,
		ClientSecret:           o.ClientSecret,
		ClientType:             o.ClientType,
		Description:            o.Description,
		ID:                     o.ID,
		Name:                   o.Name,
		Organization:           o.Organization,
		RedirectUris:           o.RedirectUris,
		SkipAuthorization:      o.SkipAuthorization,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Application
func (o *applicationTerraformModel) BodyRequest() (req applicationBodyRequestModel) {
	req.AuthorizationGrantType = o.AuthorizationGrantType.ValueString()
	req.ClientType = o.ClientType.ValueString()
	req.Description = o.Description.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	req.RedirectUris = o.RedirectUris.ValueString()
	req.SkipAuthorization = o.SkipAuthorization.ValueBool()
	return
}

func (o *applicationTerraformModel) setAuthorizationGrantType(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.AuthorizationGrantType, data, false)
}

func (o *applicationTerraformModel) setClientId(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.ClientId, data, false)
}

func (o *applicationTerraformModel) setClientSecret(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.ClientSecret, data, false)
}

func (o *applicationTerraformModel) setClientType(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.ClientType, data, false)
}

func (o *applicationTerraformModel) setDescription(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *applicationTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *applicationTerraformModel) setName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Name, data, false)
}

func (o *applicationTerraformModel) setOrganization(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.Organization, data)
}

func (o *applicationTerraformModel) setRedirectUris(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.RedirectUris, data, false)
}

func (o *applicationTerraformModel) setSkipAuthorization(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.SkipAuthorization, data)
}

func (o *applicationTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setAuthorizationGrantType(data["authorization_grant_type"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setClientId(data["client_id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setClientSecret(data["client_secret"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setClientType(data["client_type"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setName(data["name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setOrganization(data["organization"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setRedirectUris(data["redirect_uris"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSkipAuthorization(data["skip_authorization"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// applicationBodyRequestModel maps the schema for Application for creating and updating the data
type applicationBodyRequestModel struct {
	// AuthorizationGrantType "The Grant type the user must use for acquire tokens for this application."
	AuthorizationGrantType string `json:"authorization_grant_type"`
	// ClientType "Set to Public or Confidential depending on how secure the client device is."
	ClientType string `json:"client_type"`
	// Description "Optional description of this application."
	Description string `json:"description,omitempty"`
	// Name "Name of this application."
	Name string `json:"name"`
	// Organization "Organization containing this application."
	Organization int64 `json:"organization"`
	// RedirectUris "Allowed URIs list, space separated"
	RedirectUris string `json:"redirect_uris,omitempty"`
	// SkipAuthorization "Set True to skip authorization step for completely trusted applications."
	SkipAuthorization bool `json:"skip_authorization"`
}
