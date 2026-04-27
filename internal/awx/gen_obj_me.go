package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type meTerraformModel struct {
	Email           types.String `tfsdk:"email" json:"email"`
	ExternalAccount types.String `tfsdk:"external_account" json:"external_account"`
	FirstName       types.String `tfsdk:"first_name" json:"first_name"`
	ID              types.Int64  `tfsdk:"id" json:"id"`
	IsSuperuser     types.Bool   `tfsdk:"is_superuser" json:"is_superuser"`
	IsSystemAuditor types.Bool   `tfsdk:"is_system_auditor" json:"is_system_auditor"`
	LastLogin       types.String `tfsdk:"last_login" json:"last_login"`
	LastName        types.String `tfsdk:"last_name" json:"last_name"`
	LdapDn          types.String `tfsdk:"ldap_dn" json:"ldap_dn"`
	Username        types.String `tfsdk:"username" json:"username"`
}

func (o *meTerraformModel) Clone() meTerraformModel {
	return *o
}

func (o *meTerraformModel) BodyRequest() *meBodyRequestModel {
	var req meBodyRequestModel
	return &req
}

func (o *meTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
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
	collect(helpers.AttrValueSetString(&o.Username, data["username"], false))
	return diags, nil
}

type meBodyRequestModel struct {
}

type meDataSource = framework.GenericDataSource[meTerraformModel, *meTerraformModel]

// NewMeDataSource is a helper function to instantiate the Me data source.
func NewMeDataSource() datasource.DataSource {
	return &meDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "me", Endpoint: "/api/v2/me/"}},
		Cfg: framework.DataSourceCfg[meTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"email": dschema.StringAttribute{
						Description: "Email address",
						Computed:    true,
					},
					"external_account": dschema.StringAttribute{
						Description: "Set if the account is managed by an external service",
						Computed:    true,
					},
					"first_name": dschema.StringAttribute{
						Description: "First name",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this user.",
						Computed:    true,
					},
					"is_superuser": dschema.BoolAttribute{
						Description: "Designates that this user has all permissions without explicitly assigning them.",
						Computed:    true,
					},
					"is_system_auditor": dschema.BoolAttribute{
						Description: "Is system auditor",
						Computed:    true,
					},
					"last_login": dschema.StringAttribute{
						Description: "Last login",
						Computed:    true,
					},
					"last_name": dschema.StringAttribute{
						Description: "Last name",
						Computed:    true,
					},
					"ldap_dn": dschema.StringAttribute{
						Description: "Ldap dn",
						Computed:    true,
					},
					"username": dschema.StringAttribute{
						Description: "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.",
						Computed:    true,
					},
				},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "Me",
		},
	}
}
