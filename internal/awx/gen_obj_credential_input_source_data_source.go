package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type credentialInputSourceDataSource = framework.GenericDataSource[credentialInputSourceTerraformModel, *credentialInputSourceTerraformModel]

// NewCredentialInputSourceDataSource is a helper function to instantiate the CredentialInputSource data source.
func NewCredentialInputSourceDataSource() datasource.DataSource {
	return &credentialInputSourceDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_input_source", Endpoint: "/api/v2/credential_input_sources/"}},
		Cfg: framework.DataSourceCfg[credentialInputSourceTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description: "Optional description of this credential input source.",
						Computed:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this credential input source.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
							),
						},
					},
					"input_field_name": schema.StringAttribute{
						Description: "Input field name",
						Computed:    true,
					},
					"metadata": schema.StringAttribute{
						Description: "Metadata",
						Computed:    true,
					},
					"source_credential": schema.Int64Attribute{
						Description: "Source credential",
						Computed:    true,
					},
					"target_credential": schema.Int64Attribute{
						Description: "Target credential",
						Computed:    true,
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialInputSource",
		},
	}
}
