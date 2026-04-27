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

type hostTerraformModel struct {
	Description        types.String `tfsdk:"description" json:"description"`
	Enabled            types.Bool   `tfsdk:"enabled" json:"enabled"`
	ID                 types.Int64  `tfsdk:"id" json:"id"`
	InstanceId         types.String `tfsdk:"instance_id" json:"instance_id"`
	Inventory          types.Int64  `tfsdk:"inventory" json:"inventory"`
	LastJob            types.Int64  `tfsdk:"last_job" json:"last_job"`
	LastJobHostSummary types.Int64  `tfsdk:"last_job_host_summary" json:"last_job_host_summary"`
	Name               types.String `tfsdk:"name" json:"name"`
	Variables          types.String `tfsdk:"variables" json:"variables"`
}

func (o *hostTerraformModel) Clone() hostTerraformModel {
	return *o
}

func (o *hostTerraformModel) BodyRequest() *hostBodyRequestModel {
	var req hostBodyRequestModel
	req.Description = o.Description.ValueString()
	req.Enabled = o.Enabled.ValueBool()
	req.InstanceId = o.InstanceId.ValueString()
	req.Inventory = o.Inventory.ValueInt64()
	req.Name = o.Name.ValueString()
	req.Variables = json.RawMessage(o.Variables.String())
	return &req
}

func (o *hostTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetBool(&o.Enabled, data["enabled"]))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.InstanceId, data["instance_id"], false))
	collect(helpers.AttrValueSetInt64(&o.Inventory, data["inventory"]))
	collect(helpers.AttrValueSetInt64(&o.LastJob, data["last_job"]))
	collect(helpers.AttrValueSetInt64(&o.LastJobHostSummary, data["last_job_host_summary"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetJsonString(&o.Variables, data["variables"], false))
	return diags, nil
}

type hostBodyRequestModel struct {
	Description string          `json:"description,omitempty"`
	Enabled     bool            `json:"enabled"`
	InstanceId  string          `json:"instance_id,omitempty"`
	Inventory   int64           `json:"inventory"`
	Name        string          `json:"name"`
	Variables   json.RawMessage `json:"variables,omitempty"`
}

type hostResource = framework.GenericResource[hostTerraformModel, hostBodyRequestModel, *hostTerraformModel]

// NewHostResource is a helper function to simplify the provider implementation.
func NewHostResource() resource.Resource {
	return &hostResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "host", Endpoint: "/api/v2/hosts/"}},
		Cfg: framework.ResourceCfg[hostTerraformModel, hostBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description: "Optional description of this host.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"enabled": schema.BoolAttribute{
						Description: "Is this host online and available for running jobs?",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"instance_id": schema.StringAttribute{
						Description: "The value used by the remote inventory source to uniquely identify the host",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.LengthAtMost(1024),
						},
					},
					"inventory": schema.Int64Attribute{
						Description: "Inventory",
						Required:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of this host.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"variables": schema.StringAttribute{
						Description: "Host variables in JSON or YAML format.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this host.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"last_job": schema.Int64Attribute{
						Description: "Last job",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"last_job_host_summary": schema.Int64Attribute{
						Description: "Last job host summary",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *hostTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			ApiVersion:   ApiVersion,
			ResourceName: "Host",
		},
	}
}

type hostDataSource = framework.GenericDataSource[hostTerraformModel, *hostTerraformModel]

// NewHostDataSource is a helper function to instantiate the Host data source.
func NewHostDataSource() datasource.DataSource {
	return &hostDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "host", Endpoint: "/api/v2/hosts/"}},
		Cfg: framework.DataSourceCfg[hostTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"description": dschema.StringAttribute{
						Description: "Optional description of this host.",
						Computed:    true,
					},
					"enabled": dschema.BoolAttribute{
						Description: "Is this host online and available for running jobs?",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this host.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"instance_id": dschema.StringAttribute{
						Description: "The value used by the remote inventory source to uniquely identify the host",
						Computed:    true,
					},
					"inventory": dschema.Int64Attribute{
						Description: "Inventory",
						Computed:    true,
					},
					"last_job": dschema.Int64Attribute{
						Description: "Last job",
						Computed:    true,
					},
					"last_job_host_summary": dschema.Int64Attribute{
						Description: "Last job host summary",
						Computed:    true,
					},
					"name": dschema.StringAttribute{
						Description: "Name of this host.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"variables": dschema.StringAttribute{
						Description: "Host variables in JSON or YAML format.",
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
			ResourceName: "Host",
		},
	}
}
