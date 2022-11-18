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

// settingsAuthGithubOrgTerraformModel maps the schema for SettingsAuthGithubOrg when using Data Source
type settingsAuthGithubOrgTerraformModel struct {
	// SOCIAL_AUTH_GITHUB_ORG_CALLBACK_URL "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail."
	SOCIAL_AUTH_GITHUB_ORG_CALLBACK_URL types.String `tfsdk:"social_auth_github_org_callback_url" json:"SOCIAL_AUTH_GITHUB_ORG_CALLBACK_URL"`
	// SOCIAL_AUTH_GITHUB_ORG_KEY "The OAuth2 key (Client ID) from your GitHub organization application."
	SOCIAL_AUTH_GITHUB_ORG_KEY types.String `tfsdk:"social_auth_github_org_key" json:"SOCIAL_AUTH_GITHUB_ORG_KEY"`
	// SOCIAL_AUTH_GITHUB_ORG_NAME "The name of your GitHub organization, as used in your organization's URL: https://github.com/<yourorg>/."
	SOCIAL_AUTH_GITHUB_ORG_NAME types.String `tfsdk:"social_auth_github_org_name" json:"SOCIAL_AUTH_GITHUB_ORG_NAME"`
	// SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP types.String `tfsdk:"social_auth_github_org_organization_map" json:"SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP"`
	// SOCIAL_AUTH_GITHUB_ORG_SECRET "The OAuth2 secret (Client Secret) from your GitHub organization application."
	SOCIAL_AUTH_GITHUB_ORG_SECRET types.String `tfsdk:"social_auth_github_org_secret" json:"SOCIAL_AUTH_GITHUB_ORG_SECRET"`
	// SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP types.String `tfsdk:"social_auth_github_org_team_map" json:"SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP"`
}

// Clone the object
func (o settingsAuthGithubOrgTerraformModel) Clone() settingsAuthGithubOrgTerraformModel {
	return settingsAuthGithubOrgTerraformModel{
		SOCIAL_AUTH_GITHUB_ORG_CALLBACK_URL:     o.SOCIAL_AUTH_GITHUB_ORG_CALLBACK_URL,
		SOCIAL_AUTH_GITHUB_ORG_KEY:              o.SOCIAL_AUTH_GITHUB_ORG_KEY,
		SOCIAL_AUTH_GITHUB_ORG_NAME:             o.SOCIAL_AUTH_GITHUB_ORG_NAME,
		SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP: o.SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP,
		SOCIAL_AUTH_GITHUB_ORG_SECRET:           o.SOCIAL_AUTH_GITHUB_ORG_SECRET,
		SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP:         o.SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsAuthGithubOrg
func (o settingsAuthGithubOrgTerraformModel) BodyRequest() (req settingsAuthGithubOrgBodyRequestModel) {
	req.SOCIAL_AUTH_GITHUB_ORG_KEY = o.SOCIAL_AUTH_GITHUB_ORG_KEY.ValueString()
	req.SOCIAL_AUTH_GITHUB_ORG_NAME = o.SOCIAL_AUTH_GITHUB_ORG_NAME.ValueString()
	req.SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GITHUB_ORG_SECRET = o.SOCIAL_AUTH_GITHUB_ORG_SECRET.ValueString()
	req.SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP.ValueString())
	return
}

func (o *settingsAuthGithubOrgTerraformModel) setSocialAuthGithubOrgCallbackUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ORG_CALLBACK_URL, data, false)
}

func (o *settingsAuthGithubOrgTerraformModel) setSocialAuthGithubOrgKey(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ORG_KEY, data, false)
}

func (o *settingsAuthGithubOrgTerraformModel) setSocialAuthGithubOrgName(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ORG_NAME, data, false)
}

func (o *settingsAuthGithubOrgTerraformModel) setSocialAuthGithubOrgOrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthGithubOrgTerraformModel) setSocialAuthGithubOrgSecret(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GITHUB_ORG_SECRET, data, false)
}

func (o *settingsAuthGithubOrgTerraformModel) setSocialAuthGithubOrgTeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP, data, false)
}

func (o *settingsAuthGithubOrgTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setSocialAuthGithubOrgCallbackUrl(data["SOCIAL_AUTH_GITHUB_ORG_CALLBACK_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubOrgKey(data["SOCIAL_AUTH_GITHUB_ORG_KEY"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubOrgName(data["SOCIAL_AUTH_GITHUB_ORG_NAME"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubOrgOrganizationMap(data["SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubOrgSecret(data["SOCIAL_AUTH_GITHUB_ORG_SECRET"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGithubOrgTeamMap(data["SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsAuthGithubOrgBodyRequestModel maps the schema for SettingsAuthGithubOrg for creating and updating the data
type settingsAuthGithubOrgBodyRequestModel struct {
	// SOCIAL_AUTH_GITHUB_ORG_KEY "The OAuth2 key (Client ID) from your GitHub organization application."
	SOCIAL_AUTH_GITHUB_ORG_KEY string `json:"SOCIAL_AUTH_GITHUB_ORG_KEY,omitempty"`
	// SOCIAL_AUTH_GITHUB_ORG_NAME "The name of your GitHub organization, as used in your organization's URL: https://github.com/<yourorg>/."
	SOCIAL_AUTH_GITHUB_ORG_NAME string `json:"SOCIAL_AUTH_GITHUB_ORG_NAME,omitempty"`
	// SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ORG_ORGANIZATION_MAP,omitempty"`
	// SOCIAL_AUTH_GITHUB_ORG_SECRET "The OAuth2 secret (Client Secret) from your GitHub organization application."
	SOCIAL_AUTH_GITHUB_ORG_SECRET string `json:"SOCIAL_AUTH_GITHUB_ORG_SECRET,omitempty"`
	// SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP json.RawMessage `json:"SOCIAL_AUTH_GITHUB_ORG_TEAM_MAP,omitempty"`
}

var (
	_ datasource.DataSource              = &settingsAuthGithubOrgDataSource{}
	_ datasource.DataSourceWithConfigure = &settingsAuthGithubOrgDataSource{}
)

// NewSettingsAuthGithubOrgDataSource is a helper function to instantiate the SettingsAuthGithubOrg data source.
func NewSettingsAuthGithubOrgDataSource() datasource.DataSource {
	return &settingsAuthGithubOrgDataSource{}
}

// settingsAuthGithubOrgDataSource is the data source implementation.
type settingsAuthGithubOrgDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *settingsAuthGithubOrgDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/github-org/"
}

// Metadata returns the data source type name.
func (o *settingsAuthGithubOrgDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_auth_github_org"
}

// GetSchema defines the schema for the data source.
func (o *settingsAuthGithubOrgDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"SettingsAuthGithubOrg",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"social_auth_github_org_callback_url": {
					Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_org_key": {
					Description: "The OAuth2 key (Client ID) from your GitHub organization application.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_org_name": {
					Description: "The name of your GitHub organization, as used in your organization's URL: https://github.com/<yourorg>/.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_org_organization_map": {
					Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_org_secret": {
					Description: "The OAuth2 secret (Client Secret) from your GitHub organization application.",
					Type:        types.StringType,
					Sensitive:   true,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_github_org_team_map": {
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
func (o *settingsAuthGithubOrgDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsAuthGithubOrgTerraformModel
	var err error
	var endpoint = o.endpoint

	// Creates a new request for SettingsAuthGithubOrg
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubOrg on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for SettingsAuthGithubOrg
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthGithubOrg on %s", o.endpoint),
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
	if err = hookSettingsAuthGithubOrg(ctx, SourceData, CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthGithubOrg",
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
	_ resource.Resource              = &settingsAuthGithubOrgResource{}
	_ resource.ResourceWithConfigure = &settingsAuthGithubOrgResource{}
)

// NewSettingsAuthGithubOrgResource is a helper function to simplify the provider implementation.
func NewSettingsAuthGithubOrgResource() resource.Resource {
	return &settingsAuthGithubOrgResource{}
}

type settingsAuthGithubOrgResource struct {
	client   c.Client
	endpoint string
}

func (o *settingsAuthGithubOrgResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/github-org/"
}

func (o settingsAuthGithubOrgResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_settings_auth_github_org"
}

func (o settingsAuthGithubOrgResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"SettingsAuthGithubOrg",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"social_auth_github_org_key": {
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
				"social_auth_github_org_name": {
					Description: "The name of your GitHub organization, as used in your organization's URL: https://github.com/<yourorg>/.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_github_org_organization_map": {
					Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_github_org_secret": {
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
				"social_auth_github_org_team_map": {
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
				"social_auth_github_org_callback_url": {
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

func (o *settingsAuthGithubOrgResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state settingsAuthGithubOrgTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsAuthGithubOrg
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubOrg on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsAuthGithubOrg resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for SettingsAuthGithubOrg on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthGithubOrg(ctx, SourceResource, CalleeCreate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthGithubOrg",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthGithubOrgResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state settingsAuthGithubOrgTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
	var orig = state.Clone()

	// Creates a new request for SettingsAuthGithubOrg
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubOrg on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for SettingsAuthGithubOrg from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthGithubOrg on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthGithubOrg(ctx, SourceResource, CalleeRead, &orig, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthGithubOrg",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthGithubOrgResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state settingsAuthGithubOrgTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsAuthGithubOrg
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubOrg on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsAuthGithubOrg resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for SettingsAuthGithubOrg on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthGithubOrg(ctx, SourceResource, CalleeUpdate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthGithubOrg",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthGithubOrgResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
}
