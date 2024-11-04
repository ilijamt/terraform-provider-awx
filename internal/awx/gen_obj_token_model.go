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
	return tokensTerraformModel{
		Application:  o.Application,
		Description:  o.Description,
		Expires:      o.Expires,
		ID:           o.ID,
		RefreshToken: o.RefreshToken,
		Scope:        o.Scope,
		Token:        o.Token,
		User:         o.User,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Tokens
func (o *tokensTerraformModel) BodyRequest() (req tokensBodyRequestModel) {
	req.Application = o.Application.ValueInt64()
	req.Description = o.Description.ValueString()
	req.Scope = o.Scope.ValueString()
	return
}

func (o *tokensTerraformModel) setApplication(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.Application, data)
}

func (o *tokensTerraformModel) setDescription(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Description, data, false)
}

func (o *tokensTerraformModel) setExpires(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Expires, data, false)
}

func (o *tokensTerraformModel) setID(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *tokensTerraformModel) setRefreshToken(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.RefreshToken, data, false)
}

func (o *tokensTerraformModel) setScope(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Scope, data, false)
}

func (o *tokensTerraformModel) setToken(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Token, data, false)
}

func (o *tokensTerraformModel) setUser(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.User, data)
}

func (o *tokensTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setApplication(data["application"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setDescription(data["description"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setExpires(data["expires"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setRefreshToken(data["refresh_token"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setScope(data["scope"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setToken(data["token"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setUser(data["user"]); dg.HasError() {
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
