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

// settingsAuthGithubTeamTerraformModel maps the schema for SettingsAuthGithubTeam when using Data Source
type settingsAuthGithubTeamTerraformModel struct {
	// SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL "Create an organization-owned application at https://github.com/organizations/<yourorg>/settings/applications and obtain an OAuth2 key (Client ID) and secret (Client Secret). Provide this URL as the callback URL for your application."
	SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL types.String `tfsdk:"social_auth_github_team_callback_url" json:"SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL"`
	// SOCIAL_AUTH_GITHUB_TEAM_ID "Find the numeric team ID using the Github API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/."
	SOCIAL_AUTH_GITHUB_TEAM_ID types.String `tfsdk:"social_auth_github_team_id" json:"SOCIAL_AUTH_GITHUB_TEAM_ID"`
	// SOCIAL_AUTH_GITHUB_TEAM_KEY "The OAuth2 key (Client ID) from your GitHub organization application."
	SOCIAL_AUTH_GITHUB_TEAM_KEY types.String `tfsdk:"social_auth_github_team_key" json:"SOCIAL_AUTH_GITHUB_TEAM_KEY"`
	// SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP types.String `tfsdk:"social_auth_github_team_organization_map" json:"SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP"`
	// SOCIAL_AUTH_GITHUB_TEAM_SECRET "The OAuth2 secret (Client Secret) from your GitHub organization application."
	SOCIAL_AUTH_GITHUB_TEAM_SECRET types.String `tfsdk:"social_auth_github_team_secret" json:"SOCIAL_AUTH_GITHUB_TEAM_SECRET"`
	// SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP types.String `tfsdk:"social_auth_github_team_team_map" json:"SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP"`
}

// Clone the object
func (o settingsAuthGithubTeamTerraformModel) Clone() settingsAuthGithubTeamTerraformModel {
	return settingsAuthGithubTeamTerraformModel{
		SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL:     o.SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL,
		SOCIAL_AUTH_GITHUB_TEAM_ID:               o.SOCIAL_AUTH_GITHUB_TEAM_ID,
		SOCIAL_AUTH_GITHUB_TEAM_KEY:              o.SOCIAL_AUTH_GITHUB_TEAM_KEY,
		SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP: o.SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP,
		SOCIAL_AUTH_GITHUB_TEAM_SECRET:           o.SOCIAL_AUTH_GITHUB_TEAM_SECRET,
		SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP:         o.SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsAuthGithubTeam
func (o settingsAuthGithubTeamTerraformModel) BodyRequest() (req settingsAuthGithubTeamBodyRequestModel) {
	req.SOCIAL_AUTH_GITHUB_TEAM_ID = o.SOCIAL_AUTH_GITHUB_TEAM_ID.ValueString()
	req.SOCIAL_AUTH_GITHUB_TEAM_KEY = o.SOCIAL_AUTH_GITHUB_TEAM_KEY.ValueString()
	req.SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_TEAM_SECRET = o.SOCIAL_AUTH_GITHUB_TEAM_SECRET.ValueString()
	req.SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP.ValueString())
	return
}

func (o *settingsAuthGithubTeamTerraformModel) setSocialAuthGithubTeamCallbackUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL, data, false)
}

func (o *settingsAuthGithubTeamTerraformModel) setSocialAuthGithubTeamId(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_TEAM_ID, data, false)
}

func (o *settingsAuthGithubTeamTerraformModel) setSocialAuthGithubTeamKey(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_TEAM_KEY, data, false)
}

func (o *settingsAuthGithubTeamTerraformModel) setSocialAuthGithubTeamOrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthGithubTeamTerraformModel) setSocialAuthGithubTeamSecret(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_TEAM_SECRET, data, false)
}

func (o *settingsAuthGithubTeamTerraformModel) setSocialAuthGithubTeamTeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP, data, false)
}

func (o *settingsAuthGithubTeamTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setSocialAuthGithubTeamCallbackUrl(data["SOCIAL_AUTH_GITHUB_TEAM_CALLBACK_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubTeamId(data["SOCIAL_AUTH_GITHUB_TEAM_ID"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubTeamKey(data["SOCIAL_AUTH_GITHUB_TEAM_KEY"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubTeamOrganizationMap(data["SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubTeamSecret(data["SOCIAL_AUTH_GITHUB_TEAM_SECRET"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubTeamTeamMap(data["SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsAuthGithubTeamBodyRequestModel maps the schema for SettingsAuthGithubTeam for creating and updating the data
type settingsAuthGithubTeamBodyRequestModel struct {
	// SOCIAL_AUTH_GITHUB_TEAM_ID "Find the numeric team ID using the Github API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/."
	SOCIAL_AUTH_GITHUB_TEAM_ID string `json:"SOCIAL_AUTH_GITHUB_TEAM_ID,omitempty"`
	// SOCIAL_AUTH_GITHUB_TEAM_KEY "The OAuth2 key (Client ID) from your GitHub organization application."
	SOCIAL_AUTH_GITHUB_TEAM_KEY string `json:"SOCIAL_AUTH_GITHUB_TEAM_KEY,omitempty"`
	// SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_TEAM_ORGANIZATION_MAP,omitempty"`
	// SOCIAL_AUTH_GITHUB_TEAM_SECRET "The OAuth2 secret (Client Secret) from your GitHub organization application."
	SOCIAL_AUTH_GITHUB_TEAM_SECRET string `json:"SOCIAL_AUTH_GITHUB_TEAM_SECRET,omitempty"`
	// SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_TEAM_TEAM_MAP,omitempty"`
}

var (
	_ datasource.DataSource              = &settingsAuthGithubTeamDataSource{}
	_ datasource.DataSourceWithConfigure = &settingsAuthGithubTeamDataSource{}
)

// NewSettingsAuthGithubTeamDataSource is a helper function to instantiate the SettingsAuthGithubTeam data source.
func NewSettingsAuthGithubTeamDataSource() datasource.DataSource {
	return &settingsAuthGithubTeamDataSource{}
}

// settingsAuthGithubTeamDataSource is the data source implementation.
type settingsAuthGithubTeamDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *settingsAuthGithubTeamDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/github-team/"
}

// Metadata returns the data source type name.
func (o *settingsAuthGithubTeamDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_auth_github_team"
}

// GetSchema defines the schema for the data source.
func (o *settingsAuthGithubTeamDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"SettingsAuthGithubTeam",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"social_auth_github_team_callback_url": {
					Description: "Create an organization-owned application at https://github.com/organizations/<yourorg>/settings/applications and obtain an OAuth2 key (Client ID) and secret (Client Secret). Provide this URL as the callback URL for your application.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_team_id": {
					Description: "Find the numeric team ID using the Github API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_team_key": {
					Description: "The OAuth2 key (Client ID) from your GitHub organization application.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_team_organization_map": {
					Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_team_secret": {
					Description: "The OAuth2 secret (Client Secret) from your GitHub organization application.",
					Type:        types.StringType,
					Sensitive:   true,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_team_team_map": {
					Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *settingsAuthGithubTeamDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsAuthGithubTeamTerraformModel
	var err error
	var endpoint string
	endpoint = o.endpoint

	// Creates a new request for SettingsAuthGithubTeam
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubTeam on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for SettingsAuthGithubTeam
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthGithubTeam on %s", o.endpoint),
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
	if err = hookSettingsAuthGithubTeam(SourceData, CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to process custom hook for the state on SettingsAuthGithubTeam"),
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
	_ resource.Resource              = &settingsAuthGithubTeamResource{}
	_ resource.ResourceWithConfigure = &settingsAuthGithubTeamResource{}
)

// NewSettingsAuthGithubTeamResource is a helper function to simplify the provider implementation.
func NewSettingsAuthGithubTeamResource() resource.Resource {
	return &settingsAuthGithubTeamResource{}
}

type settingsAuthGithubTeamResource struct {
	client   c.Client
	endpoint string
}

func (o *settingsAuthGithubTeamResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/github-team/"
}

func (o settingsAuthGithubTeamResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_settings_auth_github_team"
}

func (o settingsAuthGithubTeamResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"SettingsAuthGithubTeam",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"social_auth_github_team_id": {
					Description: "Find the numeric team ID using the Github API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_github_team_key": {
					Description: "The OAuth2 key (Client ID) from your GitHub organization application.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_github_team_organization_map": {
					Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_github_team_secret": {
					Description: "The OAuth2 secret (Client Secret) from your GitHub organization application.",
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
				"social_auth_github_team_team_map": {
					Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				// Write only elements
				// Data only elements
				"social_auth_github_team_callback_url": {
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

func (o *settingsAuthGithubTeamResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state settingsAuthGithubTeamTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsAuthGithubTeam
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubTeam on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsAuthGithubTeam resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for SettingsAuthGithubTeam on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthGithubTeam(SourceResource, CalleeCreate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to process custom hook for the state on SettingsAuthGithubTeam"),
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthGithubTeamResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state settingsAuthGithubTeamTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
	var orig = state.Clone()

	// Creates a new request for SettingsAuthGithubTeam
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubTeam on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for SettingsAuthGithubTeam from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthGithubTeam on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthGithubTeam(SourceResource, CalleeRead, &orig, &state); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to process custom hook for the state on SettingsAuthGithubTeam"),
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthGithubTeamResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state settingsAuthGithubTeamTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsAuthGithubTeam
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf("%s", o.endpoint)) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubTeam on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsAuthGithubTeam resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for SettingsAuthGithubTeam on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthGithubTeam(SourceResource, CalleeUpdate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to process custom hook for the state on SettingsAuthGithubTeam"),
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthGithubTeamResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	return
}
