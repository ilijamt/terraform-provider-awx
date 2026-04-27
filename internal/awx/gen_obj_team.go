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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type teamTerraformModel struct {
	Description  types.String `tfsdk:"description" json:"description"`
	ID           types.Int64  `tfsdk:"id" json:"id"`
	Name         types.String `tfsdk:"name" json:"name"`
	Organization types.Int64  `tfsdk:"organization" json:"organization"`
}

func (o *teamTerraformModel) Clone() teamTerraformModel {
	return *o
}

func (o *teamTerraformModel) BodyRequest() *teamBodyRequestModel {
	var req teamBodyRequestModel
	req.Description = o.Description.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	return &req
}

func (o *teamTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	return diags, nil
}

type teamBodyRequestModel struct {
	Description  string `json:"description,omitempty"`
	Name         string `json:"name"`
	Organization int64  `json:"organization"`
}

type teamResource = framework.GenericResource[teamTerraformModel, teamBodyRequestModel, *teamTerraformModel]

// NewTeamResource is a helper function to simplify the provider implementation.
func NewTeamResource() resource.Resource {
	return &teamResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "team", Endpoint: "/api/v2/teams/"}},
		Cfg: framework.ResourceCfg[teamTerraformModel, teamBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description: "Optional description of this team.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this team.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"organization": schema.Int64Attribute{
						Description: "Organization",
						Required:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this team.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *teamTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			ApiVersion:   ApiVersion,
			ResourceName: "Team",
		},
	}
}

type teamDataSource = framework.GenericDataSource[teamTerraformModel, *teamTerraformModel]

// NewTeamDataSource is a helper function to instantiate the Team data source.
func NewTeamDataSource() datasource.DataSource {
	return &teamDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "team", Endpoint: "/api/v2/teams/"}},
		Cfg: framework.DataSourceCfg[teamTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"description": dschema.StringAttribute{
						Description: "Optional description of this team.",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this team.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"name": dschema.StringAttribute{
						Description: "Name of this team.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"organization": dschema.Int64Attribute{
						Description: "Organization",
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
			ResourceName: "Team",
		},
	}
}
