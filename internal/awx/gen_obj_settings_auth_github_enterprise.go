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

// settingsAuthGithubEnterpriseTerraformModel maps the schema for SettingsAuthGithubEnterprise when using Data Source
type settingsAuthGithubEnterpriseTerraformModel struct {
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL types.String `tfsdk:"social_auth_github_enterprise_api_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_CALLBACK_URL "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_CALLBACK_URL types.String `tfsdk:"social_auth_github_enterprise_callback_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_CALLBACK_URL"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY "The OAuth2 key (Client ID) from your GitHub Enterprise developer application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY types.String `tfsdk:"social_auth_github_enterprise_key" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP types.String `tfsdk:"social_auth_github_enterprise_organization_map" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET "The OAuth2 secret (Client Secret) from your GitHub Enterprise developer application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET types.String `tfsdk:"social_auth_github_enterprise_secret" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP types.String `tfsdk:"social_auth_github_enterprise_team_map" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_URL "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_URL types.String `tfsdk:"social_auth_github_enterprise_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_URL"`
}

// Clone the object
func (o settingsAuthGithubEnterpriseTerraformModel) Clone() settingsAuthGithubEnterpriseTerraformModel {
	return settingsAuthGithubEnterpriseTerraformModel{
		SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL:          o.SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_CALLBACK_URL:     o.SOCIAL_AUTH_GITHUB_ENTERPRISE_CALLBACK_URL,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY:              o.SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP: o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET:           o.SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP:         o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_URL:              o.SOCIAL_AUTH_GITHUB_ENTERPRISE_URL,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsAuthGithubEnterprise
func (o settingsAuthGithubEnterpriseTerraformModel) BodyRequest() (req settingsAuthGithubEnterpriseBodyRequestModel) {
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_URL = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_URL.ValueString()
	return
}

func (o *settingsAuthGithubEnterpriseTerraformModel) setSocialAuthGithubEnterpriseApiUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL, data, false)
}

func (o *settingsAuthGithubEnterpriseTerraformModel) setSocialAuthGithubEnterpriseCallbackUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_CALLBACK_URL, data, false)
}

func (o *settingsAuthGithubEnterpriseTerraformModel) setSocialAuthGithubEnterpriseKey(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY, data, false)
}

func (o *settingsAuthGithubEnterpriseTerraformModel) setSocialAuthGithubEnterpriseOrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthGithubEnterpriseTerraformModel) setSocialAuthGithubEnterpriseSecret(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET, data, false)
}

func (o *settingsAuthGithubEnterpriseTerraformModel) setSocialAuthGithubEnterpriseTeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP, data, false)
}

func (o *settingsAuthGithubEnterpriseTerraformModel) setSocialAuthGithubEnterpriseUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_URL, data, false)
}

func (o *settingsAuthGithubEnterpriseTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseApiUrl(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseCallbackUrl(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_CALLBACK_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseKey(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseOrganizationMap(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseSecret(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseTeamMap(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseUrl(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsAuthGithubEnterpriseBodyRequestModel maps the schema for SettingsAuthGithubEnterprise for creating and updating the data
type settingsAuthGithubEnterpriseBodyRequestModel struct {
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_API_URL,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY "The OAuth2 key (Client ID) from your GitHub Enterprise developer application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_KEY,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_ORGANIZATION_MAP,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET "The OAuth2 secret (Client Secret) from your GitHub Enterprise developer application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_SECRET,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_MAP,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_URL "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_URL string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_URL,omitempty"`
}

var (
	_ datasource.DataSource              = &settingsAuthGithubEnterpriseDataSource{}
	_ datasource.DataSourceWithConfigure = &settingsAuthGithubEnterpriseDataSource{}
)

// NewSettingsAuthGithubEnterpriseDataSource is a helper function to instantiate the SettingsAuthGithubEnterprise data source.
func NewSettingsAuthGithubEnterpriseDataSource() datasource.DataSource {
	return &settingsAuthGithubEnterpriseDataSource{}
}

// settingsAuthGithubEnterpriseDataSource is the data source implementation.
type settingsAuthGithubEnterpriseDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *settingsAuthGithubEnterpriseDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/github-enterprise/"
}

// Metadata returns the data source type name.
func (o *settingsAuthGithubEnterpriseDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_auth_github_enterprise"
}

// GetSchema defines the schema for the data source.
func (o *settingsAuthGithubEnterpriseDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"SettingsAuthGithubEnterprise",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"social_auth_github_enterprise_api_url": {
					Description: "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_callback_url": {
					Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_key": {
					Description: "The OAuth2 key (Client ID) from your GitHub Enterprise developer application.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_organization_map": {
					Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_secret": {
					Description: "The OAuth2 secret (Client Secret) from your GitHub Enterprise developer application.",
					Type:        types.StringType,
					Sensitive:   true,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_team_map": {
					Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_url": {
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
func (o *settingsAuthGithubEnterpriseDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsAuthGithubEnterpriseTerraformModel
	var err error
	var endpoint = o.endpoint

	// Creates a new request for SettingsAuthGithubEnterprise
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubEnterprise on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for SettingsAuthGithubEnterprise
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthGithubEnterprise on %s", o.endpoint),
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
	if err = hookSettingsAuthGithubEnterprise(SourceData, CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthGithubEnterprise",
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
	_ resource.Resource              = &settingsAuthGithubEnterpriseResource{}
	_ resource.ResourceWithConfigure = &settingsAuthGithubEnterpriseResource{}
)

// NewSettingsAuthGithubEnterpriseResource is a helper function to simplify the provider implementation.
func NewSettingsAuthGithubEnterpriseResource() resource.Resource {
	return &settingsAuthGithubEnterpriseResource{}
}

type settingsAuthGithubEnterpriseResource struct {
	client   c.Client
	endpoint string
}

func (o *settingsAuthGithubEnterpriseResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/github-enterprise/"
}

func (o settingsAuthGithubEnterpriseResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_settings_auth_github_enterprise"
}

func (o settingsAuthGithubEnterpriseResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"SettingsAuthGithubEnterprise",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"social_auth_github_enterprise_api_url": {
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
				"social_auth_github_enterprise_key": {
					Description: "The OAuth2 key (Client ID) from your GitHub Enterprise developer application.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_organization_map": {
					Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_secret": {
					Description: "The OAuth2 secret (Client Secret) from your GitHub Enterprise developer application.",
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
				"social_auth_github_enterprise_team_map": {
					Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_url": {
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
				"social_auth_github_enterprise_callback_url": {
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

func (o *settingsAuthGithubEnterpriseResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state settingsAuthGithubEnterpriseTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsAuthGithubEnterprise
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubEnterprise on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsAuthGithubEnterprise resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for SettingsAuthGithubEnterprise on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthGithubEnterprise(SourceResource, CalleeCreate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthGithubEnterprise",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthGithubEnterpriseResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state settingsAuthGithubEnterpriseTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
	var orig = state.Clone()

	// Creates a new request for SettingsAuthGithubEnterprise
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubEnterprise on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for SettingsAuthGithubEnterprise from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthGithubEnterprise on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthGithubEnterprise(SourceResource, CalleeRead, &orig, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthGithubEnterprise",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthGithubEnterpriseResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state settingsAuthGithubEnterpriseTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsAuthGithubEnterprise
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubEnterprise on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsAuthGithubEnterprise resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for SettingsAuthGithubEnterprise on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthGithubEnterprise(SourceResource, CalleeUpdate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthGithubEnterprise",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthGithubEnterpriseResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
}
