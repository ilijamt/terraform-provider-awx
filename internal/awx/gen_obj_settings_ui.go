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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// settingsUITerraformModel maps the schema for SettingsUI when using Data Source
type settingsUITerraformModel struct {
	// CUSTOM_LOGIN_INFO "If needed, you can add specific information (such as a legal notice or a disclaimer) to a text box in the login modal using this setting. Any content added must be in plain text or an HTML fragment, as other markup languages are not supported."
	CUSTOM_LOGIN_INFO types.String `tfsdk:"custom_login_info" json:"CUSTOM_LOGIN_INFO"`
	// CUSTOM_LOGO "To set up a custom logo, provide a file that you create. For the custom logo to look its best, use a .png file with a transparent background. GIF, PNG and JPEG formats are supported."
	CUSTOM_LOGO types.String `tfsdk:"custom_logo" json:"CUSTOM_LOGO"`
	// MAX_UI_JOB_EVENTS "Maximum number of job events for the UI to retrieve within a single request."
	MAX_UI_JOB_EVENTS types.Int64 `tfsdk:"max_ui_job_events" json:"MAX_UI_JOB_EVENTS"`
	// PENDO_TRACKING_STATE "Enable or Disable User Analytics Tracking."
	PENDO_TRACKING_STATE types.String `tfsdk:"pendo_tracking_state" json:"PENDO_TRACKING_STATE"`
	// UI_LIVE_UPDATES_ENABLED "If disabled, the page will not refresh when events are received. Reloading the page will be required to get the latest details."
	UI_LIVE_UPDATES_ENABLED types.Bool `tfsdk:"ui_live_updates_enabled" json:"UI_LIVE_UPDATES_ENABLED"`
}

// Clone the object
func (o settingsUITerraformModel) Clone() settingsUITerraformModel {
	return settingsUITerraformModel{
		CUSTOM_LOGIN_INFO:       o.CUSTOM_LOGIN_INFO,
		CUSTOM_LOGO:             o.CUSTOM_LOGO,
		MAX_UI_JOB_EVENTS:       o.MAX_UI_JOB_EVENTS,
		PENDO_TRACKING_STATE:    o.PENDO_TRACKING_STATE,
		UI_LIVE_UPDATES_ENABLED: o.UI_LIVE_UPDATES_ENABLED,
	}
}

// BodyRequest returns the required data, so we can call the endpoint in AWX for SettingsUI
func (o settingsUITerraformModel) BodyRequest() (req settingsUIBodyRequestModel) {
	req.CUSTOM_LOGIN_INFO = o.CUSTOM_LOGIN_INFO.ValueString()
	req.CUSTOM_LOGO = o.CUSTOM_LOGO.ValueString()
	req.MAX_UI_JOB_EVENTS = o.MAX_UI_JOB_EVENTS.ValueInt64()
	req.UI_LIVE_UPDATES_ENABLED = o.UI_LIVE_UPDATES_ENABLED.ValueBool()
	return
}

func (o *settingsUITerraformModel) setCustomLoginInfo(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.CUSTOM_LOGIN_INFO, data, false)
}

func (o *settingsUITerraformModel) setCustomLogo(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.CUSTOM_LOGO, data, false)
}

func (o *settingsUITerraformModel) setMaxUiJobEvents(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetInt64(&o.MAX_UI_JOB_EVENTS, data)
}

func (o *settingsUITerraformModel) setPendoTrackingState(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetString(&o.PENDO_TRACKING_STATE, data, false)
}

func (o *settingsUITerraformModel) setUiLiveUpdatesEnabled(data any) (d diag.Diagnostics, err error) {
	return helpers.AttrValueSetBool(&o.UI_LIVE_UPDATES_ENABLED, data)
}

func (o *settingsUITerraformModel) updateFromApiData(data map[string]any) (diags diag.Diagnostics, err error) {
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	if dg, _ := o.setCustomLoginInfo(data["CUSTOM_LOGIN_INFO"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setCustomLogo(data["CUSTOM_LOGO"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setMaxUiJobEvents(data["MAX_UI_JOB_EVENTS"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setPendoTrackingState(data["PENDO_TRACKING_STATE"]); dg.HasError() {
		diags.Append(dg...)
	}
	if dg, _ := o.setUiLiveUpdatesEnabled(data["UI_LIVE_UPDATES_ENABLED"]); dg.HasError() {
		diags.Append(dg...)
	}
	return diags, nil
}

// settingsUIBodyRequestModel maps the schema for SettingsUI for creating and updating the data
type settingsUIBodyRequestModel struct {
	// CUSTOM_LOGIN_INFO "If needed, you can add specific information (such as a legal notice or a disclaimer) to a text box in the login modal using this setting. Any content added must be in plain text or an HTML fragment, as other markup languages are not supported."
	CUSTOM_LOGIN_INFO string `json:"CUSTOM_LOGIN_INFO,omitempty"`
	// CUSTOM_LOGO "To set up a custom logo, provide a file that you create. For the custom logo to look its best, use a .png file with a transparent background. GIF, PNG and JPEG formats are supported."
	CUSTOM_LOGO string `json:"CUSTOM_LOGO,omitempty"`
	// MAX_UI_JOB_EVENTS "Maximum number of job events for the UI to retrieve within a single request."
	MAX_UI_JOB_EVENTS int64 `json:"MAX_UI_JOB_EVENTS"`
	// UI_LIVE_UPDATES_ENABLED "If disabled, the page will not refresh when events are received. Reloading the page will be required to get the latest details."
	UI_LIVE_UPDATES_ENABLED bool `json:"UI_LIVE_UPDATES_ENABLED"`
}

var (
	_ datasource.DataSource              = &settingsUIDataSource{}
	_ datasource.DataSourceWithConfigure = &settingsUIDataSource{}
)

// NewSettingsUIDataSource is a helper function to instantiate the SettingsUI data source.
func NewSettingsUIDataSource() datasource.DataSource {
	return &settingsUIDataSource{}
}

// settingsUIDataSource is the data source implementation.
type settingsUIDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *settingsUIDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/ui/"
}

// Metadata returns the data source type name.
func (o *settingsUIDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_ui"
}

// GetSchema defines the schema for the data source.
func (o *settingsUIDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceData,
		"SettingsUI",
		tfsdk.Schema{
			Version: helpers.SchemaVersion,
			Attributes: map[string]tfsdk.Attribute{
				// Data only elements
				"custom_login_info": {
					Description: "If needed, you can add specific information (such as a legal notice or a disclaimer) to a text box in the login modal using this setting. Any content added must be in plain text or an HTML fragment, as other markup languages are not supported.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"custom_logo": {
					Description: "To set up a custom logo, provide a file that you create. For the custom logo to look its best, use a .png file with a transparent background. GIF, PNG and JPEG formats are supported.",
					Type:        types.StringType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"max_ui_job_events": {
					Description: "Maximum number of job events for the UI to retrieve within a single request.",
					Type:        types.Int64Type,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				"pendo_tracking_state": {
					Description: "Enable or Disable User Analytics Tracking.",
					Type:        types.StringType,
					Computed:    true,
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"off", "anonymous", "detailed"}...),
					},
				},
				"ui_live_updates_enabled": {
					Description: "If disabled, the page will not refresh when events are received. Reloading the page will be required to get the latest details.",
					Type:        types.BoolType,
					Computed:    true,
					Validators:  []tfsdk.AttributeValidator{},
				},
				// Write only elements
			},
		}), nil
}

// Read refreshes the Terraform state with the latest data.
func (o *settingsUIDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsUITerraformModel
	var err error
	var endpoint = o.endpoint

	// Creates a new request for SettingsUI
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsUI on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for SettingsUI
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsUI on %s", o.endpoint),
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
	_ resource.Resource              = &settingsUIResource{}
	_ resource.ResourceWithConfigure = &settingsUIResource{}
)

// NewSettingsUIResource is a helper function to simplify the provider implementation.
func NewSettingsUIResource() resource.Resource {
	return &settingsUIResource{}
}

type settingsUIResource struct {
	client   c.Client
	endpoint string
}

func (o *settingsUIResource) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/ui/"
}

func (o settingsUIResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_settings_ui"
}

func (o settingsUIResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return processSchema(
		SourceResource,
		"SettingsUI",
		tfsdk.Schema{
			Attributes: map[string]tfsdk.Attribute{
				// Request elements
				"custom_login_info": {
					Description: "If needed, you can add specific information (such as a legal notice or a disclaimer) to a text box in the login modal using this setting. Any content added must be in plain text or an HTML fragment, as other markup languages are not supported.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"custom_logo": {
					Description: "To set up a custom logo, provide a file that you create. For the custom logo to look its best, use a .png file with a transparent background. GIF, PNG and JPEG formats are supported.",
					Type:        types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.StringValue(``)),
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"max_ui_job_events": {
					Description: "Maximum number of job events for the UI to retrieve within a single request.",
					Type:        types.Int64Type,
					Required:    true,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						helpers.DefaultValue(types.Int64Value(4000)),
					},
					Validators: []tfsdk.AttributeValidator{},
				},
				"ui_live_updates_enabled": {
					Description:   "If disabled, the page will not refresh when events are received. Reloading the page will be required to get the latest details.",
					Type:          types.BoolType,
					Required:      true,
					PlanModifiers: []tfsdk.AttributePlanModifier{},
					Validators:    []tfsdk.AttributeValidator{},
				},
				// Write only elements
				// Data only elements
				"pendo_tracking_state": {
					Description: "Enable or Disable User Analytics Tracking.",
					Computed:    true,
					Type:        types.StringType,
					PlanModifiers: []tfsdk.AttributePlanModifier{
						resource.UseStateForUnknown(),
					},
					Validators: []tfsdk.AttributeValidator{
						stringvalidator.OneOf([]string{"off", "anonymous", "detailed"}...),
					},
				},
			},
		}), nil
}

func (o *settingsUIResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error
	var plan, state settingsUITerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsUI
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsUI on %s for create", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsUI resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create resource for SettingsUI on %s", o.endpoint),
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

func (o *settingsUIResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var err error

	// Get current state
	var state settingsUITerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsUI
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsUI on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Get refreshed values for SettingsUI from AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsUI on %s", o.endpoint),
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

func (o *settingsUIResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var err error
	var plan, state settingsUITerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for SettingsUI
	var r *http.Request
	var endpoint = p.Clean(o.endpoint) + "/"
	var buf bytes.Buffer
	var bodyRequest = plan.BodyRequest()
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPatch, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsUI on %s for update", o.endpoint),
			err.Error(),
		)
		return
	}

	// Create a new SettingsUI resource in AWX
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to update resource for SettingsUI on %s", o.endpoint),
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

func (o *settingsUIResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
}
