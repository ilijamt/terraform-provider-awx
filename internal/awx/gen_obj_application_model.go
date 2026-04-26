package awx

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
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Application
func (o *applicationTerraformModel) BodyRequest() *applicationBodyRequestModel {
	var req applicationBodyRequestModel
	req.AuthorizationGrantType = o.AuthorizationGrantType.ValueString()
	req.ClientType = o.ClientType.ValueString()
	req.Description = o.Description.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	req.RedirectUris = o.RedirectUris.ValueString()
	req.SkipAuthorization = o.SkipAuthorization.ValueBool()
	return &req
}

func (o *applicationTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.AuthorizationGrantType, data["authorization_grant_type"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.ClientId, data["client_id"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.ClientSecret, data["client_secret"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.ClientType, data["client_type"], false)
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
		dg, _ := helpers.AttrValueSetString(&o.Name, data["name"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.Organization, data["organization"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.RedirectUris, data["redirect_uris"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.SkipAuthorization, data["skip_authorization"])
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
