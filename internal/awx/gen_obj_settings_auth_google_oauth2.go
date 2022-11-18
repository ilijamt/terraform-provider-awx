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

// settingsAuthGoogleOauth2TerraformModel maps the schema for SettingsAuthGoogleOauth2 when using Data Source
type settingsAuthGoogleOauth2TerraformModel struct {
	// SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS "Extra arguments for Google OAuth2 login. You can restrict it to only allow a single domain to authenticate, even if the user is logged in with multple Google accounts. Refer to the documentation for more detail."
	SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS types.String `tfsdk:"social_auth_google_oauth2_auth_extra_arguments" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail."
	SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL types.String `tfsdk:"social_auth_google_oauth2_callback_url" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_KEY "The OAuth2 key from your web application."
	SOCIAL_AUTH_GOOGLE_OAUTH2_KEY types.String `tfsdk:"social_auth_google_oauth2_key" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_KEY"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP types.String `tfsdk:"social_auth_google_oauth2_organization_map" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET "The OAuth2 secret from your web application."
	SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET types.String `tfsdk:"social_auth_google_oauth2_secret" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP types.String `tfsdk:"social_auth_google_oauth2_team_map" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS "Update this setting to restrict the domains who are allowed to login using Google OAuth2."
	SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS types.List `tfsdk:"social_auth_google_oauth2_whitelisted_domains" json:"SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS"`
}

// Clone the object
func (o settingsAuthGoogleOauth2TerraformModel) Clone() settingsAuthGoogleOauth2TerraformModel {
	return settingsAuthGoogleOauth2TerraformModel{
		SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS: o.SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS,
		SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL:         o.SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL,
		SOCIAL_AUTH_GOOGLE_OAUTH2_KEY:                  o.SOCIAL_AUTH_GOOGLE_OAUTH2_KEY,
		SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP:     o.SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP,
		SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET:               o.SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET,
		SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP:             o.SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP,
		SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS:  o.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsAuthGoogleOauth2
func (o settingsAuthGoogleOauth2TerraformModel) BodyRequest() (req settingsAuthGoogleOauth2BodyRequestModel) {
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS = json.RawMessage(o.SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS.ValueString())
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_KEY = o.SOCIAL_AUTH_GOOGLE_OAUTH2_KEY.ValueString()
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP = json.RawMessage(o.SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP.ValueString())
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET = o.SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET.ValueString()
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP = json.RawMessage(o.SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP.ValueString())
	req.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS = []string{}
	for _, val := range o.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS.Elements() {
		if _, ok := val.(types.String); ok {
			req.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS = append(req.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS, val.(types.String).ValueString())
		} else {
			req.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS = append(req.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS, val.String())
		}
	}
	return
}

func (o *settingsAuthGoogleOauth2TerraformModel) setSocialAuthGoogleOauth2AuthExtraArguments(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS, data, false)
}

func (o *settingsAuthGoogleOauth2TerraformModel) setSocialAuthGoogleOauth2CallbackUrl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL, data, false)
}

func (o *settingsAuthGoogleOauth2TerraformModel) setSocialAuthGoogleOauth2Key(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_KEY, data, false)
}

func (o *settingsAuthGoogleOauth2TerraformModel) setSocialAuthGoogleOauth2OrganizationMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP, data, false)
}

func (o *settingsAuthGoogleOauth2TerraformModel) setSocialAuthGoogleOauth2Secret(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET, data, false)
}

func (o *settingsAuthGoogleOauth2TerraformModel) setSocialAuthGoogleOauth2TeamMap(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetJsonString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP, data, false)
}

func (o *settingsAuthGoogleOauth2TerraformModel) setSocialAuthGoogleOauth2WhitelistedDomains(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetListString(&o.SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS, data, false)
}

func (o *settingsAuthGoogleOauth2TerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setSocialAuthGoogleOauth2AuthExtraArguments(data["SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGoogleOauth2CallbackUrl(data["SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGoogleOauth2Key(data["SOCIAL_AUTH_GOOGLE_OAUTH2_KEY"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGoogleOauth2OrganizationMap(data["SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGoogleOauth2Secret(data["SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGoogleOauth2TeamMap(data["SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthGoogleOauth2WhitelistedDomains(data["SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsAuthGoogleOauth2BodyRequestModel maps the schema for SettingsAuthGoogleOauth2 for creating and updating the data
type settingsAuthGoogleOauth2BodyRequestModel struct {
	// SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS "Extra arguments for Google OAuth2 login. You can restrict it to only allow a single domain to authenticate, even if the user is logged in with multple Google accounts. Refer to the documentation for more detail."
	SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS json.RawMessage `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS,omitempty"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_KEY "The OAuth2 key from your web application."
	SOCIAL_AUTH_GOOGLE_OAUTH2_KEY string `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_KEY,omitempty"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation."
	SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP json.RawMessage `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP,omitempty"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET "The OAuth2 secret from your web application."
	SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET string `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET,omitempty"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation."
	SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP json.RawMessage `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP,omitempty"`
	// SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS "Update this setting to restrict the domains who are allowed to login using Google OAuth2."
	SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS []string `json:"SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS,omitempty"`
}

var (
	_ datasource.DataSource              = &settingsAuthGoogleOauth2DataSource{}
	_ datasource.DataSourceWithConfigure = &settingsAuthGoogleOauth2DataSource{}
)

// NewSettingsAuthGoogleOauth2DataSource is a helper function to instantiate the SettingsAuthGoogleOauth2 data source.
func NewSettingsAuthGoogleOauth2DataSource() datasource.DataSource {
	return &settingsAuthGoogleOauth2DataSource{}
}

// settingsAuthGoogleOauth2DataSource is the data source implementation.
type settingsAuthGoogleOauth2DataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *settingsAuthGoogleOauth2DataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/google-oauth2/"
}

// Metadata returns the data source type name.
func (o *settingsAuthGoogleOauth2DataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_auth_google_oauth2"
}

// GetSchema defines the schema for the data source.
func (o *settingsAuthGoogleOauth2DataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"SettingsAuthGoogleOauth2",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"social_auth_google_oauth2_auth_extra_arguments": {
					Description: "Extra arguments for Google OAuth2 login. You can restrict it to only allow a single domain to authenticate, even if the user is logged in with multple Google accounts. Refer to the documentation for more detail.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_google_oauth2_callback_url": {
					Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_google_oauth2_key": {
					Description: "The OAuth2 key from your web application.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_google_oauth2_organization_map": {
					Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_google_oauth2_secret": {
					Description: "The OAuth2 secret from your web application.",
					Type:        types.StringType,
					Sensitive:   true,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_google_oauth2_team_map": {
					Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_google_oauth2_whitelisted_domains": {
					Description: "Update this setting to restrict the domains who are allowed to login using Google OAuth2.",
					Type:        types.ListType{ElemType: types.StringType},
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *settingsAuthGoogleOauth2DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsAuthGoogleOauth2TerraformModel
	var err error
	var endpoint = o.endpoint

	// Creates a new request for SettingsAuthGoogleOauth2
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGoogleOauth2 on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for SettingsAuthGoogleOauth2
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthGoogleOauth2 on %s", o.endpoint),
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
	if err = hookSettingsAuthGoogleOauth2(ctx, SourceData, CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthGoogleOauth2",
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
	_ resource.Resource              = &settingsAuthGoogleOauth2Resource{}
	_ resource.ResourceWithConfigure = &settingsAuthGoogleOauth2Resource{}
)

// NewSettingsAuthGoogleOauth2Resource is a helper function to simplify the provider implementation.
func NewSettingsAuthGoogleOauth2Resource() resource.Resource {
	return &settingsAuthGoogleOauth2Resource{}
}

type settingsAuthGoogleOauth2Resource struct {
	client   c.Client
	endpoint string
}

func (o *settingsAuthGoogleOauth2Resource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/google-oauth2/"
}

func (o settingsAuthGoogleOauth2Resource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_settings_auth_google_oauth2"
}

func (o settingsAuthGoogleOauth2Resource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"SettingsAuthGoogleOauth2",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"social_auth_google_oauth2_auth_extra_arguments": {
					Description: "Extra arguments for Google OAuth2 login. You can restrict it to only allow a single domain to authenticate, even if the user is logged in with multple Google accounts. Refer to the documentation for more detail.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(`{}`)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_google_oauth2_key": {
					Description: "The OAuth2 key from your web application.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_google_oauth2_organization_map": {
					Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_google_oauth2_secret": {
					Description: "The OAuth2 secret from your web application.",
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
				"social_auth_google_oauth2_team_map": {
					Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_google_oauth2_whitelisted_domains": {
					Description: "Update this setting to restrict the domains who are allowed to login using Google OAuth2.",
					Type:        types.ListType{ElemType: types.StringType},
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				// Write only elements
				// Data only elements
				"social_auth_google_oauth2_callback_url": {
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

func (o *settingsAuthGoogleOauth2Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state settingsAuthGoogleOauth2TerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsAuthGoogleOauth2
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGoogleOauth2 on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsAuthGoogleOauth2 resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for SettingsAuthGoogleOauth2 on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthGoogleOauth2(ctx, SourceResource, CalleeCreate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthGoogleOauth2",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthGoogleOauth2Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state settingsAuthGoogleOauth2TerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
	var orig = state.Clone()

	// Creates a new request for SettingsAuthGoogleOauth2
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGoogleOauth2 on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for SettingsAuthGoogleOauth2 from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthGoogleOauth2 on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthGoogleOauth2(ctx, SourceResource, CalleeRead, &orig, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthGoogleOauth2",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthGoogleOauth2Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state settingsAuthGoogleOauth2TerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsAuthGoogleOauth2
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGoogleOauth2 on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsAuthGoogleOauth2 resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for SettingsAuthGoogleOauth2 on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if err = hookSettingsAuthGoogleOauth2(ctx, SourceResource, CalleeUpdate, &plan, &state); err != nil {
		response.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthGoogleOauth2",
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsAuthGoogleOauth2Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
}
