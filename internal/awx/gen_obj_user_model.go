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
	return userTerraformModel{
		Email:           o.Email,
		ExternalAccount: o.ExternalAccount,
		FirstName:       o.FirstName,
		ID:              o.ID,
		IsSuperuser:     o.IsSuperuser,
		IsSystemAuditor: o.IsSystemAuditor,
		LastLogin:       o.LastLogin,
		LastName:        o.LastName,
		LdapDn:          o.LdapDn,
		Password:        o.Password,
		Username:        o.Username,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for User
func (o *userTerraformModel) BodyRequest() (req userBodyRequestModel) {
	req.Email = o.Email.ValueString()
	req.FirstName = o.FirstName.ValueString()
	req.IsSuperuser = o.IsSuperuser.ValueBool()
	req.IsSystemAuditor = o.IsSystemAuditor.ValueBool()
	req.LastName = o.LastName.ValueString()
	req.Password = o.Password.ValueString()
	req.Username = o.Username.ValueString()
	return
}

func (o *userTerraformModel) setEmail(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Email, data, false)
}

func (o *userTerraformModel) setExternalAccount(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.ExternalAccount, data, false)
}

func (o *userTerraformModel) setFirstName(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.FirstName, data, false)
}

func (o *userTerraformModel) setID(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *userTerraformModel) setIsSuperuser(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.IsSuperuser, data)
}

func (o *userTerraformModel) setIsSystemAuditor(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetBool(&o.IsSystemAuditor, data)
}

func (o *userTerraformModel) setLastLogin(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.LastLogin, data, false)
}

func (o *userTerraformModel) setLastName(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.LastName, data, false)
}

func (o *userTerraformModel) setLdapDn(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.LdapDn, data, false)
}

func (o *userTerraformModel) setPassword(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Password, data, false)
}

func (o *userTerraformModel) setUsername(data any) (_ diag.Diagnostics, _ error) {
	return helpers.AttrValueSetString(&o.Username, data, false)
}

func (o *userTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setEmail(data["email"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setExternalAccount(data["external_account"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setFirstName(data["first_name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setID(data["id"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setIsSuperuser(data["is_superuser"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setIsSystemAuditor(data["is_system_auditor"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setLastLogin(data["last_login"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setLastName(data["last_name"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setLdapDn(data["ldap_dn"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setPassword(data["password"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setUsername(data["username"]); dg.HasError() {
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
