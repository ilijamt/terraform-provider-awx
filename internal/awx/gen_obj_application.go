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

type applicationTerraformModel struct {
	AuthorizationGrantType types.String `tfsdk:"authorization_grant_type" json:"authorization_grant_type"`
	ClientId               types.String `tfsdk:"client_id" json:"client_id"`
	ClientSecret           types.String `tfsdk:"client_secret" json:"client_secret"`
	ClientType             types.String `tfsdk:"client_type" json:"client_type"`
	Description            types.String `tfsdk:"description" json:"description"`
	ID                     types.Int64  `tfsdk:"id" json:"id"`
	Name                   types.String `tfsdk:"name" json:"name"`
	Organization           types.Int64  `tfsdk:"organization" json:"organization"`
	RedirectUris           types.String `tfsdk:"redirect_uris" json:"redirect_uris"`
	SkipAuthorization      types.Bool   `tfsdk:"skip_authorization" json:"skip_authorization"`
}

func (o *applicationTerraformModel) Clone() applicationTerraformModel {
	return *o
}

func (o *applicationTerraformModel) BodyRequest() *applicationBodyRequestModel {
	var req applicationBodyRequestModel
	req.AuthorizationGrantType = o.AuthorizationGrantType.ValueString()
	req.ClientType = o.ClientType.ValueString()
	req.Description = o.Description.ValueString()
	req.Name = o.Name.ValueString()
	req.Organization = o.Organization.ValueInt64()
	req.RedirectUris = o.RedirectUris.ValueString()
	req.SkipAuthorization = o.SkipAuthorization.ValueBool()
	return &req
}

func (o *applicationTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.AuthorizationGrantType, data["authorization_grant_type"], false))
	collect(helpers.AttrValueSetString(&o.ClientId, data["client_id"], false))
	collect(helpers.AttrValueSetString(&o.ClientSecret, data["client_secret"], false))
	collect(helpers.AttrValueSetString(&o.ClientType, data["client_type"], false))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	collect(helpers.AttrValueSetString(&o.RedirectUris, data["redirect_uris"], false))
	collect(helpers.AttrValueSetBool(&o.SkipAuthorization, data["skip_authorization"]))
	return diags, nil
}

type applicationBodyRequestModel struct {
	AuthorizationGrantType string `json:"authorization_grant_type"`
	ClientType             string `json:"client_type"`
	Description            string `json:"description,omitempty"`
	Name                   string `json:"name"`
	Organization           int64  `json:"organization"`
	RedirectUris           string `json:"redirect_uris,omitempty"`
	SkipAuthorization      bool   `json:"skip_authorization"`
}

type applicationResource = framework.GenericResource[applicationTerraformModel, applicationBodyRequestModel, *applicationTerraformModel]

// NewApplicationResource is a helper function to simplify the provider implementation.
func NewApplicationResource() resource.Resource {
	return &applicationResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "application", Endpoint: "/api/v2/applications/"}},
		Cfg: framework.ResourceCfg[applicationTerraformModel, applicationBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"authorization_grant_type": schema.StringAttribute{
						Description: "The Grant type the user must use for acquire tokens for this application.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"authorization-code",
								"password",
							),
						},
					},
					"client_type": schema.StringAttribute{
						Description: "Set to Public or Confidential depending on how secure the client device is.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"confidential",
								"public",
							),
						},
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this application.",
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"name": schema.StringAttribute{
						Description: "Name of this application.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(255),
						},
					},
					"organization": schema.Int64Attribute{
						Description: "Organization containing this application.",
						Required:    true,
					},
					"redirect_uris": schema.StringAttribute{
						Description: "Allowed URIs list, space separated",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"skip_authorization": schema.BoolAttribute{
						Description: "Set True to skip authorization step for completely trusted applications.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"client_id": schema.StringAttribute{
						Description: "Client id",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"client_secret": schema.StringAttribute{
						Description: "Used for more stringent verification of access to an application when creating a token.",
						Sensitive:   true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this application.",
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			IDAccessor:   func(m *applicationTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:        "id",
			Hook:         hookApplication,
			ApiVersion:   ApiVersion,
			ResourceName: "Application",
		},
	}
}

type applicationDataSource = framework.GenericDataSource[applicationTerraformModel, *applicationTerraformModel]

// NewApplicationDataSource is a helper function to instantiate the Application data source.
func NewApplicationDataSource() datasource.DataSource {
	return &applicationDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "application", Endpoint: "/api/v2/applications/"}},
		Cfg: framework.DataSourceCfg[applicationTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"authorization_grant_type": dschema.StringAttribute{
						Description: "The Grant type the user must use for acquire tokens for this application.",
						Computed:    true,
					},
					"client_id": dschema.StringAttribute{
						Description: "Client id",
						Computed:    true,
					},
					"client_secret": dschema.StringAttribute{
						Description: "Used for more stringent verification of access to an application when creating a token.",
						Sensitive:   true,
						Computed:    true,
					},
					"client_type": dschema.StringAttribute{
						Description: "Set to Public or Confidential depending on how secure the client device is.",
						Computed:    true,
					},
					"description": dschema.StringAttribute{
						Description: "Optional description of this application.",
						Computed:    true,
					},
					"id": dschema.Int64Attribute{
						Description: "Database ID for this application.",
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
						Description: "Name of this application.",
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
						Description: "Organization containing this application.",
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
					"redirect_uris": dschema.StringAttribute{
						Description: "Allowed URIs list, space separated",
						Computed:    true,
					},
					"skip_authorization": dschema.BoolAttribute{
						Description: "Set True to skip authorization step for completely trusted applications.",
						Computed:    true,
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
			Hook:         hookApplication,
			ApiVersion:   ApiVersion,
			ResourceName: "Application",
		},
	}
}
