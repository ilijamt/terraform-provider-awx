package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// userTerraformModel maps the schema for User when using Data Source
type userTerraformModel struct {
	// Email ""
	Email types.String `tfsdk:"email" json:"email"`
	// ExternalAccount "Set if the account is managed by an external service"
	ExternalAccount types.String `tfsdk:"external_account" json:"external_account"`
	// FirstName ""
	FirstName types.String `tfsdk:"first_name" json:"first_name"`
	// ID "Database ID for this user."
	ID types.Int64 `tfsdk:"id" json:"id"`
	// IsSuperuser "Designates that this user has all permissions without explicitly assigning them."
	IsSuperuser types.Bool `tfsdk:"is_superuser" json:"is_superuser"`
	// IsSystemAuditor ""
	IsSystemAuditor types.Bool `tfsdk:"is_system_auditor" json:"is_system_auditor"`
	// LastLogin ""
	LastLogin types.String `tfsdk:"last_login" json:"last_login"`
	// LastName ""
	LastName types.String `tfsdk:"last_name" json:"last_name"`
	// LdapDn ""
	LdapDn types.String `tfsdk:"ldap_dn" json:"ldap_dn"`
	// Password "Field used to change the password."
	Password types.String `tfsdk:"password" json:"password"`
	// Username "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only."
	Username types.String `tfsdk:"username" json:"username"`
}

// Clone the object
func (o *userTerraformModel) Clone() userTerraformModel {
	return *o
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for User
func (o *userTerraformModel) BodyRequest() *userBodyRequestModel {
	var req userBodyRequestModel
	req.Email = o.Email.ValueString()
	req.FirstName = o.FirstName.ValueString()
	req.IsSuperuser = o.IsSuperuser.ValueBool()
	req.IsSystemAuditor = o.IsSystemAuditor.ValueBool()
	req.LastName = o.LastName.ValueString()
	req.Password = o.Password.ValueString()
	req.Username = o.Username.ValueString()
	return &req
}

func (o *userTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Email, data["email"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.ExternalAccount, data["external_account"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.FirstName, data["first_name"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetInt64(&o.ID, data["id"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.IsSuperuser, data["is_superuser"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetBool(&o.IsSystemAuditor, data["is_system_auditor"])
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.LastLogin, data["last_login"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.LastName, data["last_name"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.LdapDn, data["ldap_dn"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Password, data["password"], false)
		diags.Append(dg...)
	}
	{
		dg, _ := helpers.AttrValueSetString(&o.Username, data["username"], false)
		diags.Append(dg...)
	}
	return diags, nil
}

// userBodyRequestModel maps the schema for User for creating and updating the data
type userBodyRequestModel struct {
	// Email ""
	Email string `json:"email,omitempty"`
	// FirstName ""
	FirstName string `json:"first_name,omitempty"`
	// IsSuperuser "Designates that this user has all permissions without explicitly assigning them."
	IsSuperuser bool `json:"is_superuser"`
	// IsSystemAuditor ""
	IsSystemAuditor bool `json:"is_system_auditor"`
	// LastName ""
	LastName string `json:"last_name,omitempty"`
	// Password "Field used to change the password."
	Password string `json:"password,omitempty"`
	// Username "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only."
	Username string `json:"username"`
}
