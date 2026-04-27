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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type executionEnvironmentTerraformModel struct {
	Credential   types.Int64  `tfsdk:"credential" json:"credential"`
	Description  types.String `tfsdk:"description" json:"description"`
	ID           types.Int64  `tfsdk:"id" json:"id"`
	Image        types.String `tfsdk:"image" json:"image"`
	Managed      types.Bool   `tfsdk:"managed" json:"managed"`
	Name         types.String `tfsdk:"name" json:"name"`
	Organization types.Int64  `tfsdk:"organization" json:"organization"`
	Pull         types.String `tfsdk:"pull" json:"pull"`
}

func (o *executionEnvironmentTerraformModel) Clone() executionEnvironmentTerraformModel {
	return *o
}

func (o *executionEnvironmentTerraformModel) BodyRequest() *executionEnvironmentBodyRequestModel {
	var req executionEnvironmentBodyRequestModel
	req.Credential = o.Credential.ValueInt64()
	req.Description = o.Description.ValueString()
	req.Image = o.Image.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	req.Pull = o.Pull.ValueString()
	return &req
}

func (o *executionEnvironmentTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.Credential, data["credential"]))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.Image, data["image"], false))
	collect(helpers.AttrValueSetBool(&o.Managed, data["managed"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	collect(helpers.AttrValueSetString(&o.Pull, data["pull"], false))
	return diags, nil
}

type executionEnvironmentBodyRequestModel struct {
	Credential   int64  `json:"credential,omitempty"`
	Description  string `json:"description,omitempty"`
	Image        string `json:"image"`
	Name         string `json:"name"`
	Organization int64  `json:"organization,omitempty"`
	Pull         string `json:"pull,omitempty"`
}

type executionEnvironmentResource = framework.GenericResource[executionEnvironmentTerraformModel, executionEnvironmentBodyRequestModel, *executionEnvironmentTerraformModel]

// NewExecutionEnvironmentResource is a helper function to simplify the provider implementation.
func NewExecutionEnvironmentResource() resource.Resource {
	return &executionEnvironmentResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "execution_environment", Endpoint: "/api/v2/execution_environments/"}},
		Cfg: framework.ResourceCfg[executionEnvironmentTerraformModel, executionEnvironmentBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"credential": schema.Int64Attribute{
						Description: "Credential",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this execution environment.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"image": schema.StringAttribute{
						Description: "The full image location, including the container registry, image name, and version tag.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(1024),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this execution environment.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"organization": schema.Int64Attribute{
						Description: "The organization used to determine access to this execution environment.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"pull": schema.StringAttribute{
						Description: "Pull image before running?",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"",
								"always",
								"missing",
								"never",
							),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this execution environment.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"managed": schema.BoolAttribute{
						Description: "Managed",
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *executionEnvironmentTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			ApiVersion:   ApiVersion,
			ResourceName: "ExecutionEnvironment",
		},
	}
}

type executionEnvironmentDataSource = framework.GenericDataSource[executionEnvironmentTerraformModel, *executionEnvironmentTerraformModel]

// NewExecutionEnvironmentDataSource is a helper function to instantiate the ExecutionEnvironment data source.
func NewExecutionEnvironmentDataSource() datasource.DataSource {
	return &executionEnvironmentDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "execution_environment", Endpoint: "/api/v2/execution_environments/"}},
		Cfg: framework.DataSourceCfg[executionEnvironmentTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"credential": dschema.Int64Attribute{
						Description: "Credential",
						Computed:    true,
					},
					"description": dschema.StringAttribute{
						Description: "Optional description of this execution environment.",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this execution environment.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"image": dschema.StringAttribute{
						Description: "The full image location, including the container registry, image name, and version tag.",
						Computed:    true,
					},
					"managed": dschema.BoolAttribute{
						Description: "Managed",
						Computed:    true,
					},
					"name": dschema.StringAttribute{
						Description: "Name of this execution environment.",
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
						Description: "The organization used to determine access to this execution environment.",
						Computed:    true,
					},
					"pull": dschema.StringAttribute{
						Description: "Pull image before running?",
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
			ResourceName: "ExecutionEnvironment",
		},
	}
}
