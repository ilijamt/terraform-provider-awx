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

type credentialTypeTerraformModel struct {
	Description types.String `tfsdk:"description" json:"description"`
	ID          types.Int64  `tfsdk:"id" json:"id"`
	Injectors   types.String `tfsdk:"injectors" json:"injectors"`
	Inputs      types.String `tfsdk:"inputs" json:"inputs"`
	Kind        types.String `tfsdk:"kind" json:"kind"`
	Managed     types.Bool   `tfsdk:"managed" json:"managed"`
	Name        types.String `tfsdk:"name" json:"name"`
	Namespace   types.String `tfsdk:"namespace" json:"namespace"`
}

func (o *credentialTypeTerraformModel) Clone() credentialTypeTerraformModel {
	return *o
}

func (o *credentialTypeTerraformModel) BodyRequest() *credentialTypeBodyRequestModel {
	var req credentialTypeBodyRequestModel
	req.Description = o.Description.ValueString()
	req.Injectors = json.RawMessage(o.Injectors.ValueString())
	req.Inputs = json.RawMessage(o.Inputs.ValueString())
	req.Kind = o.Kind.ValueString()
	req.Name = o.Name.ValueString()
	return &req
}

func (o *credentialTypeTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetJsonString(&o.Injectors, data["injectors"], false))
	collect(helpers.AttrValueSetJsonString(&o.Inputs, data["inputs"], false))
	collect(helpers.AttrValueSetString(&o.Kind, data["kind"], false))
	collect(helpers.AttrValueSetBool(&o.Managed, data["managed"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetString(&o.Namespace, data["namespace"], false))
	return diags, nil
}

type credentialTypeBodyRequestModel struct {
	Description string          `json:"description,omitempty"`
	Injectors   json.RawMessage `json:"injectors,omitempty"`
	Inputs      json.RawMessage `json:"inputs,omitempty"`
	Kind        string          `json:"kind"`
	Name        string          `json:"name"`
}

type credentialTypeResource = framework.GenericResource[credentialTypeTerraformModel, credentialTypeBodyRequestModel, *credentialTypeTerraformModel]

// NewCredentialTypeResource is a helper function to simplify the provider implementation.
func NewCredentialTypeResource() resource.Resource {
	return &credentialTypeResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_type", Endpoint: "/api/v2/credential_types/"}},
		Cfg: framework.ResourceCfg[credentialTypeTerraformModel, credentialTypeBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description: "Optional description of this credential type.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"injectors": schema.StringAttribute{
						Description: "Enter injectors using either JSON or YAML syntax. Refer to the documentation for example syntax.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"inputs": schema.StringAttribute{
						Description: "Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"kind": schema.StringAttribute{
						Description: "The credential type",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"net",
								"cloud",
							),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this credential type.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(512),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this credential type.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"managed": schema.BoolAttribute{
						Description: "Is the resource managed",
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"namespace": schema.StringAttribute{
						Description: "The namespace to which the resource belongs to",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *credentialTypeTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialType",
		},
	}
}

type credentialTypeDataSource = framework.GenericDataSource[credentialTypeTerraformModel, *credentialTypeTerraformModel]

// NewCredentialTypeDataSource is a helper function to instantiate the CredentialType data source.
func NewCredentialTypeDataSource() datasource.DataSource {
	return &credentialTypeDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_type", Endpoint: "/api/v2/credential_types/"}},
		Cfg: framework.DataSourceCfg[credentialTypeTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"description": dschema.StringAttribute{
						Description: "Optional description of this credential type.",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
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
					"injectors": dschema.StringAttribute{
						Description: "Enter injectors using either JSON or YAML syntax. Refer to the documentation for example syntax.",
						Computed:    true,
					},
					"inputs": dschema.StringAttribute{
						Description: "Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax.",
						Computed:    true,
					},
					"kind": dschema.StringAttribute{
						Description: "The credential type",
						Computed:    true,
					},
					"managed": dschema.BoolAttribute{
						Description: "Is the resource managed",
						Computed:    true,
					},
					"name": dschema.StringAttribute{
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
					"namespace": dschema.StringAttribute{
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
