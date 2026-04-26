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

type labelDataSource = framework.GenericDataSource[labelTerraformModel, *labelTerraformModel]

// NewLabelDataSource is a helper function to instantiate the Label data source.
func NewLabelDataSource() datasource.DataSource {
	return &labelDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "label", Endpoint: "/api/v2/labels/"}},
		Cfg: framework.DataSourceCfg[labelTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.Int64Attribute{
						Description: "Database ID for this label.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ConflictsWith(
								path.MatchRoot("name"),
								path.MatchRoot("organization"),
							),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this label.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.AlsoRequires(
								path.MatchRoot("organization"),
							),
							stringvalidator.ConflictsWith(
								path.MatchRoot("id"),
							),
						},
					},
					"organization": schema.Int64Attribute{
						Description: "Organization this label belongs to.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.AlsoRequires(
								path.MatchRoot("name"),
							),
							int64validator.ConflictsWith(
								path.MatchRoot("id"),
							),
						},
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
				{Name: "by_name_organization", URLSuffix: "?name__exact=%s&organization=%d", Fields: []framework.SearchField{
					{Name: "name", Type: "string", URLEscape: true},
					{Name: "organization", Type: "int64", URLEscape: false},
				}},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "Label",
		},
	}
}
