package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type userDataSource = framework.GenericDataSource[userTerraformModel, *userTerraformModel]

// NewUserDataSource is a helper function to instantiate the User data source.
func NewUserDataSource() datasource.DataSource {
	return &userDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "user", Endpoint: "/api/v2/users/"}},
		Cfg: framework.DataSourceCfg[userTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Data only elements
					"email": schema.StringAttribute{
						Description: "Email address",
						Sensitive:   false,
						Computed:    true,
					},
					"external_account": schema.StringAttribute{
						Description: "Set if the account is managed by an external service",
						Sensitive:   false,
						Computed:    true,
					},
					"first_name": schema.StringAttribute{
						Description: "First name",
						Sensitive:   false,
						Computed:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this user.",
						Sensitive:   false,
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("username"),
							),
						},
					},
					"is_superuser": schema.BoolAttribute{
						Description: "Designates that this user has all permissions without explicitly assigning them.",
						Sensitive:   false,
						Computed:    true,
					},
					"is_system_auditor": schema.BoolAttribute{
						Description: "Is system auditor",
						Sensitive:   false,
						Computed:    true,
					},
					"last_login": schema.StringAttribute{
						Description: "Last login",
						Sensitive:   false,
						Computed:    true,
					},
					"last_name": schema.StringAttribute{
						Description: "Last name",
						Sensitive:   false,
						Computed:    true,
					},
					"ldap_dn": schema.StringAttribute{
						Description: "Ldap dn",
						Sensitive:   false,
						Computed:    true,
					},
					"password": schema.StringAttribute{
						Description: "Field used to change the password.",
						Sensitive:   true,
						Computed:    true,
					},
					"username": schema.StringAttribute{
						Description: "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.",
						Sensitive:   false,
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
