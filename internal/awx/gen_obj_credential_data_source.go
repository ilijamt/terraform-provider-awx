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

type credentialDataSource = framework.GenericDataSource[credentialTerraformModel, *credentialTerraformModel]

// NewCredentialDataSource is a helper function to instantiate the Credential data source.
func NewCredentialDataSource() datasource.DataSource {
	return &credentialDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential", Endpoint: "/api/v2/credentials/"}},
		Cfg: framework.DataSourceCfg[credentialTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"cloud": schema.BoolAttribute{
						Description: "Cloud",
						Computed:    true,
					},
					"credential_type": schema.Int64Attribute{
						Description: "Specify the type of credential you want to create. Refer to the documentation for details on each type.",
						Computed:    true,
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this credential.",
						Computed:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this credential.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"inputs": schema.StringAttribute{
						Description: "Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax.",
						Computed:    true,
					},
					"kind": schema.StringAttribute{
						Description: "Kind",
						Computed:    true,
					},
					"kubernetes": schema.BoolAttribute{
						Description: "Kubernetes",
						Computed:    true,
					},
					"managed": schema.BoolAttribute{
						Description: "Managed",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of this credential.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"organization": schema.Int64Attribute{
						Description: "Inherit permissions from organization roles. If provided on creation, do not give either user or team.",
						Computed:    true,
					},
					"team": schema.Int64Attribute{
						Description: "Write-only field used to add team to owner role. If provided, do not give either user or organization. Only valid for creation.",
						Optional:    true,
						Computed:    true,
					},
					"user": schema.Int64Attribute{
						Description: "Write-only field used to add user to owner role. If provided, do not give either team or organization. Only valid for creation.",
						Optional:    true,
						Computed:    true,
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
				{Name: "by_name", URLSuffix: "/?name__exact=%s", Fields: []framework.SearchField{
					{Name: "name", Type: "string", URLEscape: true},
				}},
			},
			Hook:         hookCredential,
			ApiVersion:   ApiVersion,
			ResourceName: "Credential",
		},
	}
}
