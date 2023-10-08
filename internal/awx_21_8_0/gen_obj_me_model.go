package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

// meTerraformModel maps the schema for Me when using Data Source
type meTerraformModel struct {
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
	// Username "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only."
	Username types.String `tfsdk:"username" json:"username"`
}

// Clone the object
func (o *meTerraformModel) Clone() meTerraformModel {
	return meTerraformModel{
		Email:           o.Email,
		ExternalAccount: o.ExternalAccount,
		FirstName:       o.FirstName,
		ID:              o.ID,
		IsSuperuser:     o.IsSuperuser,
		IsSystemAuditor: o.IsSystemAuditor,
		LastLogin:       o.LastLogin,
		LastName:        o.LastName,
		LdapDn:          o.LdapDn,
		Username:        o.Username,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for Me
func (o *meTerraformModel) BodyRequest() (req meBodyRequestModel) {
	return
}

func (o *meTerraformModel) setEmail(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Email, data, false)
}

func (o *meTerraformModel) setExternalAccount(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.ExternalAccount, data, false)
}

func (o *meTerraformModel) setFirstName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.FirstName, data, false)
}

func (o *meTerraformModel) setID(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.ID, data)
}

func (o *meTerraformModel) setIsSuperuser(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.IsSuperuser, data)
}

func (o *meTerraformModel) setIsSystemAuditor(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.IsSystemAuditor, data)
}

func (o *meTerraformModel) setLastLogin(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.LastLogin, data, false)
}

func (o *meTerraformModel) setLastName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.LastName, data, false)
}

func (o *meTerraformModel) setLdapDn(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.LdapDn, data, false)
}

func (o *meTerraformModel) setUsername(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.Username, data, false)
}

func (o *meTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
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
	if dg, _ := o.setUsername(data["username"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// meBodyRequestModel maps the schema for Me for creating and updating the data
type meBodyRequestModel struct {
}
