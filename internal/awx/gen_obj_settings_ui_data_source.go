package awx

import (
	"context"
	"fmt"
	"net/http"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

var (
	_ datasource.DataSource              = &settingsUiDataSource{}
	_ datasource.DataSourceWithConfigure = &settingsUiDataSource{}
)

// NewSettingsUIDataSource is a helper function to instantiate the SettingsUI data source.
func NewSettingsUIDataSource() datasource.DataSource {
	return &settingsUiDataSource{}
}

// settingsUiDataSource is the data source implementation.
type settingsUiDataSource struct {
	client   c.Client
	endpoint string
	name     string
}

// Configure adds the provider configured client to the data source.
func (o *settingsUiDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.name = "SettingsUI"
	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/ui/"
}

// Metadata returns the data source type name.
func (o *settingsUiDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_ui"
}

// Schema defines the schema for the data source.
func (o *settingsUiDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"custom_login_info": schema.StringAttribute{
				Description: "If needed, you can add specific information (such as a legal notice or a disclaimer) to a text box in the login modal using this setting. Any content added must be in plain text or an HTML fragment, as other markup languages are not supported.",
				Sensitive:   false,
				Computed:    true,
			},
			"custom_logo": schema.StringAttribute{
				Description: "To set up a custom logo, provide a file that you create. For the custom logo to look its best, use a .png file with a transparent background. GIF, PNG and JPEG formats are supported.",
				Sensitive:   false,
				Computed:    true,
			},
			"max_ui_job_events": schema.Int64Attribute{
				Description: "Maximum number of job events for the UI to retrieve within a single request.",
				Sensitive:   false,
				Computed:    true,
			},
			"pendo_tracking_state": schema.StringAttribute{
				Description: "Enable or Disable User Analytics Tracking.",
				Sensitive:   false,
				Computed:    true,
			},
			"ui_live_updates_enabled": schema.BoolAttribute{
				Description: "If disabled, the page will not refresh when events are received. Reloading the page will be required to get the latest details.",
				Sensitive:   false,
				Computed:    true,
			},
		},
	}
}

func (o *settingsUiDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// Read refreshes the Terraform state with the latest data.
func (o *settingsUiDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsUiTerraformModel
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
			fmt.Sprintf("Unable to read resource for SettingsUI on %s", endpoint),
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
