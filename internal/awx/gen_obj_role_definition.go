package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type roleDefinitionTerraformModel struct {
	ContentType types.String `tfsdk:"content_type" json:"content_type"`
	CreatedBy   types.Int64  `tfsdk:"created_by" json:"created_by"`
	Description types.String `tfsdk:"description" json:"description"`
	ID          types.Int64  `tfsdk:"id" json:"id"`
	Managed     types.Bool   `tfsdk:"managed" json:"managed"`
	ModifiedBy  types.Int64  `tfsdk:"modified_by" json:"modified_by"`
	Name        types.String `tfsdk:"name" json:"name"`
	Permissions types.List   `tfsdk:"permissions" json:"permissions"`
}

func (o *roleDefinitionTerraformModel) Clone() roleDefinitionTerraformModel {
	return *o
}

func (o *roleDefinitionTerraformModel) BodyRequest() *roleDefinitionBodyRequestModel {
	var req roleDefinitionBodyRequestModel
	req.ContentType = o.ContentType.ValueString()
	req.Description = o.Description.ValueString()
	req.Name = o.Name.ValueString()
	req.Permissions = helpers.ListAsStringSlice(o.Permissions, false)
	return &req
}

func (o *roleDefinitionTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.ContentType, data["content_type"], false))
	collect(helpers.AttrValueSetInt64(&o.CreatedBy, data["created_by"]))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetBool(&o.Managed, data["managed"]))
	collect(helpers.AttrValueSetInt64(&o.ModifiedBy, data["modified_by"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetListString(&o.Permissions, data["permissions"], false))
	return diags, nil
}

type roleDefinitionBodyRequestModel struct {
	ContentType string   `json:"content_type,omitempty"`
	Description string   `json:"description,omitempty"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

type roleDefinitionDataSource = framework.GenericDataSource[roleDefinitionTerraformModel, *roleDefinitionTerraformModel]

// NewRoleDefinitionDataSource is a helper function to instantiate the RoleDefinition data source.
func NewRoleDefinitionDataSource() datasource.DataSource {
	return &roleDefinitionDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "role_definition", Endpoint: "/api/v2/role_definitions/"}},
		Cfg: framework.DataSourceCfg[roleDefinitionTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"content_type": dschema.StringAttribute{
						Description: "The type of resource this applies to",
						Computed:    true,
					},
					"created_by": dschema.Int64Attribute{
						Description: "The user who created this resource",
						Computed:    true,
					},
					"description": dschema.StringAttribute{
						Description: "Optional description of this role definition.",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this role definition.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"managed": dschema.BoolAttribute{
						Description: "Managed",
						Computed:    true,
					},
					"modified_by": dschema.Int64Attribute{
						Description: "The user who last modified this resource",
						Computed:    true,
					},
					"name": dschema.StringAttribute{
						Description: "Name of this role definition.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"permissions": dschema.ListAttribute{
						ElementType: types.StringType,
						Description: "Permissions",
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
			ResourceName: "RoleDefinition",
		},
	}
}
