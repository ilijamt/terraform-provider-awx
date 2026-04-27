package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
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

type userResource = framework.GenericResource[userTerraformModel, userBodyRequestModel, *userTerraformModel]

// NewUserResource is a helper function to simplify the provider implementation.
func NewUserResource() resource.Resource {
	return &userResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "user", Endpoint: "/api/v2/users/"}},
		Cfg: framework.ResourceCfg[userTerraformModel, userBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"email": schema.StringAttribute{
						Description: "Email address",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(254),
						},
					},
					"first_name": schema.StringAttribute{
						Description: "First name",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(150),
						},
					},
					"is_superuser": schema.BoolAttribute{
						Description: "Designates that this user has all permissions without explicitly assigning them.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"is_system_auditor": schema.BoolAttribute{
						Description: "Is system auditor",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"last_name": schema.StringAttribute{
						Description: "Last name",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(150),
						},
					},
					"password": schema.StringAttribute{
						Description: "Field used to change the password.",
						Sensitive:   true,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"username": schema.StringAttribute{
						Description: "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(150),
						},
					},
					"external_account": schema.StringAttribute{
						Description: "Set if the account is managed by an external service",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this user.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"last_login": schema.StringAttribute{
						Description: "Last login",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"ldap_dn": schema.StringAttribute{
						Description: "Ldap dn",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *userTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			Hook:         hookUser,
			ApiVersion:   ApiVersion,
			ResourceName: "User",
		},
	}
}

type userDataSource = framework.GenericDataSource[userTerraformModel, *userTerraformModel]

// NewUserDataSource is a helper function to instantiate the User data source.
func NewUserDataSource() datasource.DataSource {
	return &userDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "user", Endpoint: "/api/v2/users/"}},
		Cfg: framework.DataSourceCfg[userTerraformModel]{
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
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("username"),
							),
						},
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
					"password": dschema.StringAttribute{
						Description: "Field used to change the password.",
						Sensitive:   true,
						Computed:    true,
					},
					"username": dschema.StringAttribute{
						Description: "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("username"),
							),
						},
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
				{Name: "by_username", URLSuffix: "?username__exact=%s", Fields: []framework.SearchField{
					{Name: "username", Type: "string", URLEscape: true},
				}},
			},
			Hook:         hookUser,
			ApiVersion:   ApiVersion,
			ResourceName: "User",
		},
	}
}
