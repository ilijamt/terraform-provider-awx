package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// tokensTerraformModel maps the schema for Tokens when using Data Source
type tokensTerraformModel struct {
	// Application ""
	Application types.Int64 `tfsdk:"application" json:"application"`
	// Description "Optional description of this access token."
	Description types.String `tfsdk:"description" json:"description"`
	// Expires ""
	Expires types.String `tfsdk:"expires" json:"expires"`
	// ID "Database ID for this access token."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// RefreshToken ""
	RefreshToken types.String `tfsdk:"refresh_token" json:"refresh_token"`
	// Scope "Allowed scopes, further restricts user's permissions. Must be a simple space-separated string with allowed scopes ['read', 'write']."
	Scope types.String `tfsdk:"scope" json:"scope"`
	// Token ""
	Token types.String `tfsdk:"token" json:"token"`
	// User "The user representing the token owner"
	User types.Int64 `tfsdk:"user" json:"user"`
}

// Clone the object
func (o *tokensTerraformModel) Clone() tokensTerraformModel {
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Tokens
func (o *tokensTerraformModel) BodyRequest() *tokensBodyRequestModel {
	var req tokensBodyRequestModel
	req.Application = o.Application.ValueInt64()
	req.Description = o.Description.ValueString()
	req.Scope = o.Scope.ValueString()
	return &req
}

func (o *tokensTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.Application, data["application"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Description, data["description"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Expires, data["expires"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.ID, data["id"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.RefreshToken, data["refresh_token"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Scope, data["scope"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Token, data["token"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.User, data["user"])
		diags.Append(dg...)
	}
	return diags, nil
}

// tokensBodyRequestModel maps the schema for Tokens for creating and updating the data
type tokensBodyRequestModel struct {
	// Application ""
	Application int64 `json:"application,omitempty"`
	// Description "Optional description of this access token."
	Description string `json:"description,omitempty"`
	// Scope "Allowed scopes, further restricts user's permissions. Must be a simple space-separated string with allowed scopes ['read', 'write']."
	Scope string `json:"scope,omitempty"`
}
