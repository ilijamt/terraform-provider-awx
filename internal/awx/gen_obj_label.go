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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type labelTerraformModel struct {
	ID           types.Int64  `tfsdk:"id" json:"id"`
	Name         types.String `tfsdk:"name" json:"name"`
	Organization types.Int64  `tfsdk:"organization" json:"organization"`
}

func (o *labelTerraformModel) Clone() labelTerraformModel {
	return *o
}

func (o *labelTerraformModel) BodyRequest() *labelBodyRequestModel {
	var req labelBodyRequestModel
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	return &req
}

func (o *labelTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	return diags, nil
}

type labelBodyRequestModel struct {
	Name         string `json:"name"`
	Organization int64  `json:"organization"`
}

type labelResource = framework.GenericResource[labelTerraformModel, labelBodyRequestModel, *labelTerraformModel]

// NewLabelResource is a helper function to simplify the provider implementation.
func NewLabelResource() resource.Resource {
	return &labelResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "label", Endpoint: "/api/v2/labels/"}},
		Cfg: framework.ResourceCfg[labelTerraformModel, labelBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description: "Name of this label.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"organization": schema.Int64Attribute{
						Description: "Organization this label belongs to.",
						Required:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this label.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *labelTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			UnDeletable:  true,
			ApiVersion:   ApiVersion,
			ResourceName: "Label",
		},
	}
}

type labelDataSource = framework.GenericDataSource[labelTerraformModel, *labelTerraformModel]

// NewLabelDataSource is a helper function to instantiate the Label data source.
func NewLabelDataSource() datasource.DataSource {
	return &labelDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "label", Endpoint: "/api/v2/labels/"}},
		Cfg: framework.DataSourceCfg[labelTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"id": dschema.Int64Attribute{
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
					"name": dschema.StringAttribute{
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
					"organization": dschema.Int64Attribute{
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
