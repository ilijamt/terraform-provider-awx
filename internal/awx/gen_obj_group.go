package awx

import (
	"encoding/json"
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

type groupTerraformModel struct {
	Description types.String `tfsdk:"description" json:"description"`
	ID          types.Int64  `tfsdk:"id" json:"id"`
	Inventory   types.Int64  `tfsdk:"inventory" json:"inventory"`
	Name        types.String `tfsdk:"name" json:"name"`
	Variables   types.String `tfsdk:"variables" json:"variables"`
}

func (o *groupTerraformModel) Clone() groupTerraformModel {
	return *o
}

func (o *groupTerraformModel) BodyRequest() *groupBodyRequestModel {
	var req groupBodyRequestModel
	req.Description = o.Description.ValueString()
	req.Inventory = o.Inventory.ValueInt64()
	req.Name = o.Name.ValueString()
	req.Variables = json.RawMessage(o.Variables.String())
	return &req
}

func (o *groupTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetInt64(&o.Inventory, data["inventory"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetJsonString(&o.Variables, data["variables"], false))
	return diags, nil
}

type groupBodyRequestModel struct {
	Description string          `json:"description,omitempty"`
	Inventory   int64           `json:"inventory"`
	Name        string          `json:"name"`
	Variables   json.RawMessage `json:"variables,omitempty"`
}

type groupResource = framework.GenericResource[groupTerraformModel, groupBodyRequestModel, *groupTerraformModel]

// NewGroupResource is a helper function to simplify the provider implementation.
func NewGroupResource() resource.Resource {
	return &groupResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "group", Endpoint: "/api/v2/groups/"}},
		Cfg: framework.ResourceCfg[groupTerraformModel, groupBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description: "Optional description of this group.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"inventory": schema.Int64Attribute{
						Description: "Inventory",
						Required:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of this group.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"variables": schema.StringAttribute{
						Description: "Group variables in JSON or YAML format.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this group.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *groupTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			ApiVersion:   ApiVersion,
			ResourceName: "Group",
		},
	}
}

type groupDataSource = framework.GenericDataSource[groupTerraformModel, *groupTerraformModel]

// NewGroupDataSource is a helper function to instantiate the Group data source.
func NewGroupDataSource() datasource.DataSource {
	return &groupDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "group", Endpoint: "/api/v2/groups/"}},
		Cfg: framework.DataSourceCfg[groupTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"description": dschema.StringAttribute{
						Description: "Optional description of this group.",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this group.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
							),
						},
					},
					"inventory": dschema.Int64Attribute{
						Description: "Inventory",
						Computed:    true,
					},
					"name": dschema.StringAttribute{
						Description: "Name of this group.",
						Computed:    true,
					},
					"variables": dschema.StringAttribute{
						Description: "Group variables in JSON or YAML format.",
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
			ResourceName: "Group",
		},
	}
}
