package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

type meDataSource = framework.GenericDataSource[meTerraformModel, *meTerraformModel]

// NewMeDataSource is a helper function to instantiate the Me data source.
func NewMeDataSource() datasource.DataSource {
	return &meDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "me", Endpoint: "/api/v2/me/"}},
		Cfg: framework.DataSourceCfg[meTerraformModel]{
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
						Computed:    true,
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
					"username": schema.StringAttribute{
						Description: "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.",
						Sensitive:   false,
						Computed:    true,
					},
				},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "Me",
		},
	}
}
