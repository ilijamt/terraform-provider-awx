package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type userTerraformModel struct {
	Email           types.String `tfsdk:"email" json:"email"`
	ExternalAccount types.String `tfsdk:"external_account" json:"external_account"`
	FirstName       types.String `tfsdk:"first_name" json:"first_name"`
	ID              types.Int64  `tfsdk:"id" json:"id"`
	IsSuperuser     types.Bool   `tfsdk:"is_superuser" json:"is_superuser"`
	IsSystemAuditor types.Bool   `tfsdk:"is_system_auditor" json:"is_system_auditor"`
	LastLogin       types.String `tfsdk:"last_login" json:"last_login"`
	LastName        types.String `tfsdk:"last_name" json:"last_name"`
	LdapDn          types.String `tfsdk:"ldap_dn" json:"ldap_dn"`
	Password        types.String `tfsdk:"password" json:"password"`
	Username        types.String `tfsdk:"username" json:"username"`
}

func (o *userTerraformModel) Clone() userTerraformModel {
	return *o
}

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
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.Email, data["email"], false))
	collect(helpers.AttrValueSetString(&o.ExternalAccount, data["external_account"], false))
	collect(helpers.AttrValueSetString(&o.FirstName, data["first_name"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetBool(&o.IsSuperuser, data["is_superuser"]))
	collect(helpers.AttrValueSetBool(&o.IsSystemAuditor, data["is_system_auditor"]))
	collect(helpers.AttrValueSetString(&o.LastLogin, data["last_login"], false))
	collect(helpers.AttrValueSetString(&o.LastName, data["last_name"], false))
	collect(helpers.AttrValueSetString(&o.LdapDn, data["ldap_dn"], false))
	collect(helpers.AttrValueSetString(&o.Password, data["password"], false))
	collect(helpers.AttrValueSetString(&o.Username, data["username"], false))
	return diags, nil
}

type userBodyRequestModel struct {
	Email           string `json:"email,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	IsSuperuser     bool   `json:"is_superuser"`
	IsSystemAuditor bool   `json:"is_system_auditor"`
	LastName        string `json:"last_name,omitempty"`
	Password        string `json:"password,omitempty"`
	Username        string `json:"username"`
}
