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
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// settingsOpenIDConnectTerraformModel maps the schema for SettingsOpenIDConnect when using Data Source
type settingsOpenIDConnectTerraformModel struct {
	// SOCIAL_AUTH_OIDC_KEY "The OIDC key (Client ID) from your IDP."
	SOCIAL_AUTH_OIDC_KEY types.String `tfsdk:"social_auth_oidc_key" json:"SOCIAL_AUTH_OIDC_KEY"`
	// SOCIAL_AUTH_OIDC_OIDC_ENDPOINT "The URL for your OIDC provider including the path up to /.well-known/openid-configuration"
	SOCIAL_AUTH_OIDC_OIDC_ENDPOINT types.String `tfsdk:"social_auth_oidc_oidc_endpoint" json:"SOCIAL_AUTH_OIDC_OIDC_ENDPOINT"`
	// SOCIAL_AUTH_OIDC_SECRET "The OIDC secret (Client Secret) from your IDP."
	SOCIAL_AUTH_OIDC_SECRET types.String `tfsdk:"social_auth_oidc_secret" json:"SOCIAL_AUTH_OIDC_SECRET"`
	// SOCIAL_AUTH_OIDC_VERIFY_SSL "Verify the OIDV provider ssl certificate."
	SOCIAL_AUTH_OIDC_VERIFY_SSL types.Bool `tfsdk:"social_auth_oidc_verify_ssl" json:"SOCIAL_AUTH_OIDC_VERIFY_SSL"`
}

// Clone the object
func (o *settingsOpenIDConnectTerraformModel) Clone() settingsOpenIDConnectTerraformModel {
	return settingsOpenIDConnectTerraformModel{
		SOCIAL_AUTH_OIDC_KEY:           o.SOCIAL_AUTH_OIDC_KEY,
		SOCIAL_AUTH_OIDC_OIDC_ENDPOINT: o.SOCIAL_AUTH_OIDC_OIDC_ENDPOINT,
		SOCIAL_AUTH_OIDC_SECRET:        o.SOCIAL_AUTH_OIDC_SECRET,
		SOCIAL_AUTH_OIDC_VERIFY_SSL:    o.SOCIAL_AUTH_OIDC_VERIFY_SSL,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsOpenIDConnect
func (o *settingsOpenIDConnectTerraformModel) BodyRequest() (req settingsOpenIDConnectBodyRequestModel) {
	req.SOCIAL_AUTH_OIDC_KEY = o.SOCIAL_AUTH_OIDC_KEY.ValueString()
	req.SOCIAL_AUTH_OIDC_OIDC_ENDPOINT = o.SOCIAL_AUTH_OIDC_OIDC_ENDPOINT.ValueString()
	req.SOCIAL_AUTH_OIDC_SECRET = o.SOCIAL_AUTH_OIDC_SECRET.ValueString()
	req.SOCIAL_AUTH_OIDC_VERIFY_SSL = o.SOCIAL_AUTH_OIDC_VERIFY_SSL.ValueBool()
	return
}

func (o *settingsOpenIDConnectTerraformModel) setSocialAuthOidcKey(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_OIDC_KEY, data, false)
}

func (o *settingsOpenIDConnectTerraformModel) setSocialAuthOidcOidcEndpoint(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_OIDC_OIDC_ENDPOINT, data, false)
}

func (o *settingsOpenIDConnectTerraformModel) setSocialAuthOidcSecret(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.SOCIAL_AUTH_OIDC_SECRET, data, false)
}

func (o *settingsOpenIDConnectTerraformModel) setSocialAuthOidcVerifySsl(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.SOCIAL_AUTH_OIDC_VERIFY_SSL, data)
}

func (o *settingsOpenIDConnectTerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setSocialAuthOidcKey(data["SOCIAL_AUTH_OIDC_KEY"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthOidcOidcEndpoint(data["SOCIAL_AUTH_OIDC_OIDC_ENDPOINT"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthOidcSecret(data["SOCIAL_AUTH_OIDC_SECRET"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setSocialAuthOidcVerifySsl(data["SOCIAL_AUTH_OIDC_VERIFY_SSL"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsOpenIDConnectBodyRequestModel maps the schema for SettingsOpenIDConnect for creating and updating the data
type settingsOpenIDConnectBodyRequestModel struct {
	// SOCIAL_AUTH_OIDC_KEY "The OIDC key (Client ID) from your IDP."
	SOCIAL_AUTH_OIDC_KEY string `json:"SOCIAL_AUTH_OIDC_KEY,omitempty"`
	// SOCIAL_AUTH_OIDC_OIDC_ENDPOINT "The URL for your OIDC provider including the path up to /.well-known/openid-configuration"
	SOCIAL_AUTH_OIDC_OIDC_ENDPOINT string `json:"SOCIAL_AUTH_OIDC_OIDC_ENDPOINT,omitempty"`
	// SOCIAL_AUTH_OIDC_SECRET "The OIDC secret (Client Secret) from your IDP."
	SOCIAL_AUTH_OIDC_SECRET string `json:"SOCIAL_AUTH_OIDC_SECRET,omitempty"`
	// SOCIAL_AUTH_OIDC_VERIFY_SSL "Verify the OIDV provider ssl certificate."
	SOCIAL_AUTH_OIDC_VERIFY_SSL bool `json:"SOCIAL_AUTH_OIDC_VERIFY_SSL"`
}

var (
	_ datasource.DataSource              = &settingsOpenIDConnectDataSource{}
	_ datasource.DataSourceWithConfigure = &settingsOpenIDConnectDataSource{}
)

// NewSettingsOpenIDConnectDataSource is a helper function to instantiate the SettingsOpenIDConnect data source.
func NewSettingsOpenIDConnectDataSource() datasource.DataSource {
	return &settingsOpenIDConnectDataSource{}
}

// settingsOpenIDConnectDataSource is the data source implementation.
type settingsOpenIDConnectDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *settingsOpenIDConnectDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/oidc/"
}

// Metadata returns the data source type name.
func (o *settingsOpenIDConnectDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_oidc"
}

// GetSchema defines the schema for the data source.
func (o *settingsOpenIDConnectDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"SettingsOpenIDConnect",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"social_auth_oidc_key": {
					Description: "The OIDC key (Client ID) from your IDP.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_oidc_oidc_endpoint": {
					Description: "The URL for your OIDC provider including the path up to /.well-known/openid-configuration",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_oidc_secret": {
					Description: "The OIDC secret (Client Secret) from your IDP.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"social_auth_oidc_verify_ssl": {
					Description: "Verify the OIDV provider ssl certificate.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *settingsOpenIDConnectDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsOpenIDConnectTerraformModel
	var err error
	var endpoint = o.endpoint

	// Creates a new request for SettingsOpenIDConnect
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsOpenIDConnect on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for SettingsOpenIDConnect
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsOpenIDConnect on %s", o.endpoint),
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
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &settingsOpenIDConnectResource{}
	_ resource.ResourceWithConfigure = &settingsOpenIDConnectResource{}
)

// NewSettingsOpenIDConnectResource is a helper function to simplify the provider implementation.
func NewSettingsOpenIDConnectResource() resource.Resource {
	return &settingsOpenIDConnectResource{}
}

type settingsOpenIDConnectResource struct {
	client   c.Client
	endpoint string
}

func (o *settingsOpenIDConnectResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/oidc/"
}

func (o *settingsOpenIDConnectResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_settings_oidc"
}

func (o *settingsOpenIDConnectResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"SettingsOpenIDConnect",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"social_auth_oidc_key": {
					Description: "The OIDC key (Client ID) from your IDP.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_oidc_oidc_endpoint": {
					Description: "The URL for your OIDC provider including the path up to /.well-known/openid-configuration",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_oidc_secret": {
					Description: "The OIDC secret (Client Secret) from your IDP.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"social_auth_oidc_verify_ssl": {
					Description: "Verify the OIDV provider ssl certificate.",
					Type:        types.BoolType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				// Write only elements
				// Data only elements
			},
		}), nil
}

func (o *settingsOpenIDConnectResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state settingsOpenIDConnectTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsOpenIDConnect
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[SettingsOpenIDConnect/Create] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsOpenIDConnect on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsOpenIDConnect resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for SettingsOpenIDConnect on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsOpenIDConnectResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state settingsOpenIDConnectTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsOpenIDConnect
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsOpenIDConnect on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for SettingsOpenIDConnect from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsOpenIDConnect on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsOpenIDConnectResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state settingsOpenIDConnectTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsOpenIDConnect
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	tflog.Debug(ctx, "[SettingsOpenIDConnect/Update] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsOpenIDConnect on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsOpenIDConnect resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for SettingsOpenIDConnect on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics
	if d, err = state.updateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *settingsOpenIDConnectResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
}
