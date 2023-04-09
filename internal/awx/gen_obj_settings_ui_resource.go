package awx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	p "path"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

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

func (o *settingsUIResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_settings_ui"
}

func (o *settingsUIResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Request elements
			"custom_login_info": schema.StringAttribute{
				Description: "If needed, you can add specific information (such as a legal notice or a disclaimer) to a text box in the login modal using this setting. Any content added must be in plain text or an HTML fragment, as other markup languages are not supported.",
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"custom_logo": schema.StringAttribute{
				Description: "To set up a custom logo, provide a file that you create. For the custom logo to look its best, use a .png file with a transparent background. GIF, PNG and JPEG formats are supported.",
				Required:    false,
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(``),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{},
			},
			"max_ui_job_events": schema.Int64Attribute{
				Description:   "Maximum number of job events for the UI to retrieve within a single request.",
				Required:      true,
				Optional:      false,
				Default:       int64default.StaticInt64(4000),
				PlanModifiers: []planmodifier.Int64{},
				Validators:    []validator.Int64{},
			},
			"ui_live_updates_enabled": schema.BoolAttribute{
				Description:   "If disabled, the page will not refresh when events are received. Reloading the page will be required to get the latest details.",
				Required:      true,
				Optional:      false,
				PlanModifiers: []planmodifier.Bool{},
				Validators:    []validator.Bool{},
			},
			// Write only elements
			// Data only elements
			"pendo_tracking_state": schema.StringAttribute{
				Description: "Enable or Disable User Analytics Tracking.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"off", "anonymous", "detailed"}...),
				},
			},
		},
	}
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
	tflog.Debug(ctx, "[SettingsUI/Create] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
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
	tflog.Debug(ctx, "[SettingsUI/Update] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
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
