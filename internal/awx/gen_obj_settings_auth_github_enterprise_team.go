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

// settingsAuthGithubEnterpriseTeamTerraformModel maps the schema for SettingsAuthGithubEnterpriseTeam when using Data Source
type settingsAuthGithubEnterpriseTeamTerraformModel struct {
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL types.String `tfsdk:"social_auth_github_enterprise_team_api_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL "Create an organization-owned application at https://github.com/organizations/<yourorg>/settings/applications and obtain an OAuth2 key (Client ID) and secret (Client Secret). Provide this URL as the callback URL for your application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL types.String `tfsdk:"social_auth_github_enterprise_team_callback_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID "Find the numeric team ID using the Github Enterprise API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID types.String `tfsdk:"social_auth_github_enterprise_team_id" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY "The OAuth2 key (Client ID) from your GitHub Enterprise organization application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY types.String `tfsdk:"social_auth_github_enterprise_team_key" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP types.String `tfsdk:"social_auth_github_enterprise_team_organization_map" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET "The OAuth2 secret (Client Secret) from your GitHub Enterprise organization application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET types.String `tfsdk:"social_auth_github_enterprise_team_secret" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP types.String `tfsdk:"social_auth_github_enterprise_team_team_map" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL types.String `tfsdk:"social_auth_github_enterprise_team_url" json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL"`
}

// Clone the object
func (o settingsAuthGithubEnterpriseTeamTerraformModel) Clone() settingsAuthGithubEnterpriseTeamTerraformModel {
	return settingsAuthGithubEnterpriseTeamTerraformModel{
		SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL:          o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL:     o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID:               o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY:              o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP: o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET:           o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP:         o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP,
		SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL:              o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsAuthGithubEnterpriseTeam
func (o settingsAuthGithubEnterpriseTeamTerraformModel) BodyRequest() (req settingsAuthGithubEnterpriseTeamBodyRequestModel) {
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET.ValueString()
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL = o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL.ValueString()
	return
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) setSocialAuthGithubEnterpriseTeamApiUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL, data, false)
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) setSocialAuthGithubEnterpriseTeamCallbackUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL, data, false)
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) setSocialAuthGithubEnterpriseTeamId(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID, data, false)
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) setSocialAuthGithubEnterpriseTeamKey(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY, data, false)
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) setSocialAuthGithubEnterpriseTeamOrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) setSocialAuthGithubEnterpriseTeamSecret(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET, data, false)
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) setSocialAuthGithubEnterpriseTeamTeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP, data, false)
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) setSocialAuthGithubEnterpriseTeamUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL, data, false)
}

func (o *settingsAuthGithubEnterpriseTeamTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseTeamApiUrl(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseTeamCallbackUrl(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseTeamId(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseTeamKey(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseTeamOrganizationMap(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseTeamSecret(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseTeamTeamMap(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubEnterpriseTeamUrl(data["SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsAuthGithubEnterpriseTeamBodyRequestModel maps the schema for SettingsAuthGithubEnterpriseTeam for creating and updating the data
type settingsAuthGithubEnterpriseTeamBodyRequestModel struct {
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID "Find the numeric team ID using the Github Enterprise API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY "The OAuth2 key (Client ID) from your GitHub Enterprise organization application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET "The OAuth2 secret (Client Secret) from your GitHub Enterprise organization application."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP,omitempty"`
	// SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details."
	SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL string `json:"SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL,omitempty"`
}

var (
	_ datasource.DataSource              = &settingsAuthGithubEnterpriseTeamDataSource{}
	_ datasource.DataSourceWithConfigure = &settingsAuthGithubEnterpriseTeamDataSource{}
)

// NewSettingsAuthGithubEnterpriseTeamDataSource is a helper function to instantiate the SettingsAuthGithubEnterpriseTeam data source.
func NewSettingsAuthGithubEnterpriseTeamDataSource() datasource.DataSource {
	return &settingsAuthGithubEnterpriseTeamDataSource{}
}

// settingsAuthGithubEnterpriseTeamDataSource is the data source implementation.
type settingsAuthGithubEnterpriseTeamDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *settingsAuthGithubEnterpriseTeamDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/github-enterprise-team/"
}

// Metadata returns the data source type name.
func (o *settingsAuthGithubEnterpriseTeamDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_auth_github_enterprise_team"
}

// GetSchema defines the schema for the data source.
func (o *settingsAuthGithubEnterpriseTeamDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"SettingsAuthGithubEnterpriseTeam",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"social_auth_github_enterprise_team_api_url": {
					Description: "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_team_callback_url": {
					Description: "Create an organization-owned application at https://github.com/organizations/<yourorg>/settings/applications and obtain an OAuth2 key (Client ID) and secret (Client Secret). Provide this URL as the callback URL for your application.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_team_id": {
					Description: "Find the numeric team ID using the Github Enterprise API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_team_key": {
					Description: "The OAuth2 key (Client ID) from your GitHub Enterprise organization application.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_team_organization_map": {
					Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_team_secret": {
					Description: "The OAuth2 secret (Client Secret) from your GitHub Enterprise organization application.",
					Type:        types.StringType,
					Sensitive:   true,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_team_team_map": {
					Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_team_url": {
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
func (o *settingsAuthGithubEnterpriseTeamDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsAuthGithubEnterpriseTeamTerraformModel
	var err error
	var endpoint = o.endpoint

	// Creates a new request for SettingsAuthGithubEnterpriseTeam
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubEnterpriseTeam on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for SettingsAuthGithubEnterpriseTeam
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthGithubEnterpriseTeam on %s", o.endpoint),
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
	if err = hookSettingsAuthGithubEnterpriseTeam(SourceData, CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthGithubEnterpriseTeam",
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
	_ resource.Resource              = &settingsAuthGithubEnterpriseTeamResource{}
	_ resource.ResourceWithConfigure = &settingsAuthGithubEnterpriseTeamResource{}
)

// NewSettingsAuthGithubEnterpriseTeamResource is a helper function to simplify the provider implementation.
func NewSettingsAuthGithubEnterpriseTeamResource() resource.Resource {
	return &settingsAuthGithubEnterpriseTeamResource{}
}

type settingsAuthGithubEnterpriseTeamResource struct {
	client   c.Client
	endpoint string
}

func (o *settingsAuthGithubEnterpriseTeamResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/github-enterprise-team/"
}

func (o settingsAuthGithubEnterpriseTeamResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_settings_auth_github_enterprise_team"
}

func (o settingsAuthGithubEnterpriseTeamResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"SettingsAuthGithubEnterpriseTeam",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"social_auth_github_enterprise_team_api_url": {
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
				"social_auth_github_enterprise_team_id": {
					Description: "Find the numeric team ID using the Github Enterprise API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_team_key": {
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
				"social_auth_github_enterprise_team_organization_map": {
					Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_team_secret": {
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
				"social_auth_github_enterprise_team_team_map": {
					Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_github_enterprise_team_url": {
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
				"social_auth_github_enterprise_team_callback_url": {
					Description: "Create an organization-owned application at https://github.com/organizations/<yourorg>/settings/applications and obtain an OAuth2 key (Client ID) and secret (Client Secret). Provide this URL as the callback URL for your application.",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
				},
			},
		}), nil
}

func (o *settingsAuthGithubEnterpriseTeamResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state settingsAuthGithubEnterpriseTeamTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsAuthGithubEnterpriseTeam
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubEnterpriseTeam on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsAuthGithubEnterpriseTeam resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for SettingsAuthGithubEnterpriseTeam on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthGithubEnterpriseTeam(SourceResource, CalleeCreate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthGithubEnterpriseTeam",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthGithubEnterpriseTeamResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state settingsAuthGithubEnterpriseTeamTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
	var orig = state.Clone()

	// Creates a new request for SettingsAuthGithubEnterpriseTeam
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubEnterpriseTeam on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for SettingsAuthGithubEnterpriseTeam from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthGithubEnterpriseTeam on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthGithubEnterpriseTeam(SourceResource, CalleeRead, &orig, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthGithubEnterpriseTeam",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthGithubEnterpriseTeamResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state settingsAuthGithubEnterpriseTeamTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsAuthGithubEnterpriseTeam
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubEnterpriseTeam on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsAuthGithubEnterpriseTeam resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for SettingsAuthGithubEnterpriseTeam on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthGithubEnterpriseTeam(SourceResource, CalleeUpdate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthGithubEnterpriseTeam",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthGithubEnterpriseTeamResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
}
