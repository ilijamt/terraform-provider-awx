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

type credentialInputSourceTerraformModel struct {
	Description      types.String `tfsdk:"description" json:"description"`
	ID               types.Int64  `tfsdk:"id" json:"id"`
	InputFieldName   types.String `tfsdk:"input_field_name" json:"input_field_name"`
	Metadata         types.String `tfsdk:"metadata" json:"metadata"`
	SourceCredential types.Int64  `tfsdk:"source_credential" json:"source_credential"`
	TargetCredential types.Int64  `tfsdk:"target_credential" json:"target_credential"`
}

func (o *credentialInputSourceTerraformModel) Clone() credentialInputSourceTerraformModel {
	return *o
}

func (o *credentialInputSourceTerraformModel) BodyRequest() *credentialInputSourceBodyRequestModel {
	var req credentialInputSourceBodyRequestModel
	req.Description = o.Description.ValueString()
	req.InputFieldName = o.InputFieldName.ValueString()
	req.Metadata = json.RawMessage(o.Metadata.ValueString())
	req.SourceCredential = o.SourceCredential.ValueInt64()
	req.TargetCredential = o.TargetCredential.ValueInt64()
	return &req
}

func (o *credentialInputSourceTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.InputFieldName, data["input_field_name"], false))
	collect(helpers.AttrValueSetJsonString(&o.Metadata, data["metadata"], false))
	collect(helpers.AttrValueSetInt64(&o.SourceCredential, data["source_credential"]))
	collect(helpers.AttrValueSetInt64(&o.TargetCredential, data["target_credential"]))
	return diags, nil
}

type credentialInputSourceBodyRequestModel struct {
	Description      string          `json:"description,omitempty"`
	InputFieldName   string          `json:"input_field_name"`
	Metadata         json.RawMessage `json:"metadata,omitempty"`
	SourceCredential int64           `json:"source_credential"`
	TargetCredential int64           `json:"target_credential"`
}

type credentialInputSourceResource = framework.GenericResource[credentialInputSourceTerraformModel, credentialInputSourceBodyRequestModel, *credentialInputSourceTerraformModel]

// NewCredentialInputSourceResource is a helper function to simplify the provider implementation.
func NewCredentialInputSourceResource() resource.Resource {
	return &credentialInputSourceResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_input_source", Endpoint: "/api/v2/credential_input_sources/"}},
		Cfg: framework.ResourceCfg[credentialInputSourceTerraformModel, credentialInputSourceBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description: "Optional description of this credential input source.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"input_field_name": schema.StringAttribute{
						Description: "Input field name",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(1024),
						},
					},
					"metadata": schema.StringAttribute{
						Description: "Metadata",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"source_credential": schema.Int64Attribute{
						Description: "Source credential",
						Required:    true,
					},
					"target_credential": schema.Int64Attribute{
						Description: "Target credential",
						Required:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this credential input source.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *credentialInputSourceTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			ApiVersion:   ApiVersion,
			ResourceName: "CredentialInputSource",
		},
	}
}

type credentialInputSourceDataSource = framework.GenericDataSource[credentialInputSourceTerraformModel, *credentialInputSourceTerraformModel]

// NewCredentialInputSourceDataSource is a helper function to instantiate the CredentialInputSource data source.
func NewCredentialInputSourceDataSource() datasource.DataSource {
	return &credentialInputSourceDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "credential_input_source", Endpoint: "/api/v2/credential_input_sources/"}},
		Cfg: framework.DataSourceCfg[credentialInputSourceTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"description": dschema.StringAttribute{
						Description: "Optional description of this credential input source.",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this credential input source.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
							),
						},
					},
					"input_field_name": dschema.StringAttribute{
						Description: "Input field name",
						Computed:    true,
					},
					"metadata": dschema.StringAttribute{
						Description: "Metadata",
						Computed:    true,
					},
					"source_credential": dschema.Int64Attribute{
						Description: "Source credential",
						Computed:    true,
					},
					"target_credential": dschema.Int64Attribute{
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
