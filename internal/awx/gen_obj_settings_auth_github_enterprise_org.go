package awx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	p "path"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// settingsAuthGithubEnterpriseOrgTerraformModel maps the schema for SettingsAuthGithubEnterpriseOrg when using Data Source
type settingsAuthGithubEnterpriseOrgTerraformModel struct {
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL types.String `tfsdk:"social_auth_github_enterprise_org_api_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_CALLBACK_URL "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_CALLBACK_URL types.String `tfsdk:"social_auth_github_enterprise_org_callback_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_CALLBACK_URL"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY "The OAuth2 key (Client ID) from your GitHub Enterprise organization application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY types.String `tfsdk:"social_auth_github_enterprise_org_key" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME "The name of your GitHub Enterprise organization, as used in your organization's URL: https://github.com/<yourorg>/."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME types.String `tfsdk:"social_auth_github_enterprise_org_name" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP types.String `tfsdk:"social_auth_github_enterprise_org_organization_map" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET "The OAuth2 secret (Client Secret) from your GitHub Enterprise organization application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET types.String `tfsdk:"social_auth_github_enterprise_org_secret" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP types.String `tfsdk:"social_auth_github_enterprise_org_team_map" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL types.String `tfsdk:"social_auth_github_enterprise_org_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL"`
}

// Clone the object
func (o settingsAuthGithubEnterpriseOrgTerraformModel) Clone() settingsAuthGithubEnterpriseOrgTerraformModel {
	return settingsAuthGithubEnterpriseOrgTerraformModel{
		SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL:          o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_CALLBACK_URL:     o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_CALLBACK_URL,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY:              o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME:             o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP: o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET:           o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP:         o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL:              o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsAuthGithubEnterpriseOrg
func (o settingsAuthGithubEnterpriseOrgTerraformModel) BodyRequest() (req settingsAuthGithubEnterpriseOrgBodyRequestModel) {
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL.ValueString()
	return
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) setSocialAuthGithubEnterpriseOrgApiUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL, data, false)
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) setSocialAuthGithubEnterpriseOrgCallbackUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_CALLBACK_URL, data, false)
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) setSocialAuthGithubEnterpriseOrgKey(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY, data, false)
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) setSocialAuthGithubEnterpriseOrgName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME, data, false)
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) setSocialAuthGithubEnterpriseOrgOrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) setSocialAuthGithubEnterpriseOrgSecret(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET, data, false)
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) setSocialAuthGithubEnterpriseOrgTeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP, data, false)
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) setSocialAuthGithubEnterpriseOrgUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL, data, false)
}

func (o *settingsAuthGithubEnterpriseOrgTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseOrgApiUrl(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseOrgCallbackUrl(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_CALLBACK_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseOrgKey(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseOrgName(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseOrgOrganizationMap(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseOrgSecret(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseOrgTeamMap(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseOrgUrl(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsAuthGithubEnterpriseOrgBodyRequestModel maps the schema for SettingsAuthGithubEnterpriseOrg for creating and updating the data
type settingsAuthGithubEnterpriseOrgBodyRequestModel struct {
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_API_URL,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY "The OAuth2 key (Client ID) from your GitHub Enterprise organization application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_KEY,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME "The name of your GitHub Enterprise organization, as used in your organization's URL: https://github.com/<yourorg>/."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_NAME,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_ORGANIZATION_MAP,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET "The OAuth2 secret (Client Secret) from your GitHub Enterprise organization application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_SECRET,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_TEAM_MAP,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORG_URL,omitempty"`
}

var (
	_ datasource.DataSource              = &settingsAuthGithubEnterpriseOrgDataSource{}
	_ datasource.DataSourceWithConfigure = &settingsAuthGithubEnterpriseOrgDataSource{}
)

// NewSettingsAuthGithubEnterpriseOrgDataSource is a helper function to instantiate the SettingsAuthGithubEnterpriseOrg data source.
func NewSettingsAuthGithubEnterpriseOrgDataSource() datasource.DataSource {
	return &settingsAuthGithubEnterpriseOrgDataSource{}
}

// settingsAuthGithubEnterpriseOrgDataSource is the data source implementation.
type settingsAuthGithubEnterpriseOrgDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *settingsAuthGithubEnterpriseOrgDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/github-enterprise-org/"
}

// Metadata returns the data source type name.
func (o *settingsAuthGithubEnterpriseOrgDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_auth_github_enterprise_org"
}

// GetSchema defines the schema for the data source.
func (o *settingsAuthGithubEnterpriseOrgDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"SettingsAuthGithubEnterpriseOrg",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"social_auth_github_enterprise_org_api_url": {
					Description: "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_org_callback_url": {
					Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_org_key": {
					Description: "The OAuth2 key (Client ID) from your GitHub Enterprise organization application.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_org_name": {
					Description: "The name of your GitHub Enterprise organization, as used in your organization's URL: https://github.com/<yourorg>/.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_org_organization_map": {
					Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_org_secret": {
					Description: "The OAuth2 secret (Client Secret) from your GitHub Enterprise organization application.",
					Type:        types.StringType,
					Sensitive:   true,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_org_team_map": {
					Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_org_url": {
					Description: "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *settingsAuthGithubEnterpriseOrgDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsAuthGithubEnterpriseOrgTerraformModel
	var err error
	var endpoint string
	endpoint = o.endpoint

	// Creates a new request for SettingsAuthGithubEnterpriseOrg
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubEnterpriseOrg on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for SettingsAuthGithubEnterpriseOrg
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthGithubEnterpriseOrg on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics

	if d, err = state.updateFromApiData(data); err != nil {
		resp.Diagnostics.Append(d...)
		return
	}

	// Set state
	if err = hookSettingsAuthGithubEnterpriseOrg(SourceData, CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to process custom hook for the state on SettingsAuthGithubEnterpriseOrg"),
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &settingsAuthGithubEnterpriseOrgResource{}
	_ resource.ResourceWithConfigure = &settingsAuthGithubEnterpriseOrgResource{}
)

// NewSettingsAuthGithubEnterpriseOrgResource is a helper function to simplify the provider implementation.
func NewSettingsAuthGithubEnterpriseOrgResource() resource.Resource {
	return &settingsAuthGithubEnterpriseOrgResource{}
}

type settingsAuthGithubEnterpriseOrgResource struct {
	client   c.Client
	endpoint string
}

func (o *settingsAuthGithubEnterpriseOrgResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/github-enterprise-org/"
}

func (o settingsAuthGithubEnterpriseOrgResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_settings_auth_github_enterprise_org"
}

func (o settingsAuthGithubEnterpriseOrgResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"SettingsAuthGithubEnterpriseOrg",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"social_auth_github_enterprise_org_api_url": {
					Description: "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_org_key": {
					Description: "The OAuth2 key (Client ID) from your GitHub Enterprise organization application.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_org_name": {
					Description: "The name of your GitHub Enterprise organization, as used in your organization's URL: https://github.com/<yourorg>/.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_org_organization_map": {
					Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_org_secret": {
					Description: "The OAuth2 secret (Client Secret) from your GitHub Enterprise organization application.",
					Type:        types.StringType,
					Sensitive:   true,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_org_team_map": {
					Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_org_url": {
					Description: "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				// Write only elements
				// Data only elements
				"social_auth_github_enterprise_org_callback_url": {
					Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail.",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
			},
		}), nil
}

func (o *settingsAuthGithubEnterpriseOrgResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state settingsAuthGithubEnterpriseOrgTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsAuthGithubEnterpriseOrg
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubEnterpriseOrg on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsAuthGithubEnterpriseOrg resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for SettingsAuthGithubEnterpriseOrg on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthGithubEnterpriseOrg(SourceResource, CalleeCreate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to process custom hook for the state on SettingsAuthGithubEnterpriseOrg"),
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthGithubEnterpriseOrgResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state settingsAuthGithubEnterpriseOrgTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
	var orig = state.Clone()

	// Creates a new request for SettingsAuthGithubEnterpriseOrg
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubEnterpriseOrg on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for SettingsAuthGithubEnterpriseOrg from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthGithubEnterpriseOrg on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthGithubEnterpriseOrg(SourceResource, CalleeRead, &orig, &state); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to process custom hook for the state on SettingsAuthGithubEnterpriseOrg"),
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthGithubEnterpriseOrgResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state settingsAuthGithubEnterpriseOrgTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsAuthGithubEnterpriseOrg
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubEnterpriseOrg on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsAuthGithubEnterpriseOrg resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for SettingsAuthGithubEnterpriseOrg on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthGithubEnterpriseOrg(SourceResource, CalleeUpdate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to process custom hook for the state on SettingsAuthGithubEnterpriseOrg"),
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthGithubEnterpriseOrgResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	return
}
