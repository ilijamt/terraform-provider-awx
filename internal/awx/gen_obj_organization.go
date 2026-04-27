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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type organizationTerraformModel struct {
	DefaultEnvironment types.Int64  `tfsdk:"default_environment" json:"default_environment"`
	Description        types.String `tfsdk:"description" json:"description"`
	ID                 types.Int64  `tfsdk:"id" json:"id"`
	MaxHosts           types.Int64  `tfsdk:"max_hosts" json:"max_hosts"`
	Name               types.String `tfsdk:"name" json:"name"`
}

func (o *organizationTerraformModel) Clone() organizationTerraformModel {
	return *o
}

func (o *organizationTerraformModel) BodyRequest() *organizationBodyRequestModel {
	var req organizationBodyRequestModel
	req.DefaultEnvironment = o.DefaultEnvironment.ValueInt64()
	req.Description = o.Description.ValueString()
	req.MaxHosts = o.MaxHosts.ValueInt64()
	req.Name = o.Name.ValueString()
	return &req
}

func (o *organizationTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.DefaultEnvironment, data["default_environment"]))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetInt64(&o.MaxHosts, data["max_hosts"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	return diags, nil
}

type organizationBodyRequestModel struct {
	DefaultEnvironment int64  `json:"default_environment,omitempty"`
	Description        string `json:"description,omitempty"`
	MaxHosts           int64  `json:"max_hosts,omitempty"`
	Name               string `json:"name"`
}

type organizationResource = framework.GenericResource[organizationTerraformModel, organizationBodyRequestModel, *organizationTerraformModel]

// NewOrganizationResource is a helper function to simplify the provider implementation.
func NewOrganizationResource() resource.Resource {
	return &organizationResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "organization", Endpoint: "/api/v2/organizations/"}},
		Cfg: framework.ResourceCfg[organizationTerraformModel, organizationBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"default_environment": schema.Int64Attribute{
						Description: "The default execution environment for jobs run by this organization.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this organization.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"max_hosts": schema.Int64Attribute{
						Description: "Maximum number of hosts allowed to be managed by this organization.",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{
							int64validator.Between(0, 2147483647),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this organization.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this organization.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *organizationTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			ApiVersion:   ApiVersion,
			ResourceName: "Organization",
		},
	}
}

type organizationDataSource = framework.GenericDataSource[organizationTerraformModel, *organizationTerraformModel]

// NewOrganizationDataSource is a helper function to instantiate the Organization data source.
func NewOrganizationDataSource() datasource.DataSource {
	return &organizationDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "organization", Endpoint: "/api/v2/organizations/"}},
		Cfg: framework.DataSourceCfg[organizationTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"default_environment": dschema.Int64Attribute{
						Description: "The default execution environment for jobs run by this organization.",
						Computed:    true,
					},
					"description": dschema.StringAttribute{
						Description: "Optional description of this organization.",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this organization.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"max_hosts": dschema.Int64Attribute{
						Description: "Maximum number of hosts allowed to be managed by this organization.",
						Computed:    true,
					},
					"name": dschema.StringAttribute{
						Description: "Name of this organization.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
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
			ResourceName: "Organization",
		},
	}
}
