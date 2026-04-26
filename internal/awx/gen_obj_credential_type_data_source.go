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

type credentialTypeDataSource = framework.GenericDataSource[credentialTypeTerraformModel, *credentialTypeTerraformModel]

// NewCredentialTypeDataSource is a helper function to instantiate the CredentialType data source.
func NewCredentialTypeDataSource() datasource.DataSource {
	return &credentialTypeDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_type", Endpoint: "/api/v2/credential_types/"}},
		Cfg: framework.DataSourceCfg[credentialTypeTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description: "Optional description of this credential type.",
						Computed:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this credential type.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"injectors": schema.StringAttribute{
						Description: "Enter injectors using either JSON or YAML syntax. Refer to the documentation for example syntax.",
						Computed:    true,
					},
					"inputs": schema.StringAttribute{
						Description: "Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax.",
						Computed:    true,
					},
					"kind": schema.StringAttribute{
						Description: "The credential type",
						Computed:    true,
					},
					"managed": schema.BoolAttribute{
						Description: "Is the resource managed",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of this credential type.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"namespace": schema.StringAttribute{
						Description: "The namespace to which the resource belongs to",
						Computed:    true,
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
				{Name: "by_name", URLSuffix: "?name__exact=%s", Fields: []framework.SearchField{
					{Name: "name", Type: "string", URLEscape: true},
				}},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialType",
		},
	}
}
