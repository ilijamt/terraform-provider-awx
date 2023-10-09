package awx

import (
	"context"
	"fmt"
	"net/http"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

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

// Schema defines the schema for the data source.
func (o *settingsAuthGoogleOauth2DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"social_auth_google_oauth2_auth_extra_arguments": schema.StringAttribute{
				Description: "Extra arguments for Google OAuth2 login. You can restrict it to only allow a single domain to authenticate, even if the user is logged in with multple Google accounts. Refer to the documentation for more detail.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_google_oauth2_callback_url": schema.StringAttribute{
				Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_google_oauth2_key": schema.StringAttribute{
				Description: "The OAuth2 key from your web application.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_google_oauth2_organization_map": schema.StringAttribute{
				Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_google_oauth2_secret": schema.StringAttribute{
				Description: "The OAuth2 secret from your web application.",
				Sensitive:   true,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_google_oauth2_team_map": schema.StringAttribute{
				Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_google_oauth2_whitelisted_domains": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "Update this setting to restrict the domains who are allowed to login using Google OAuth2.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.List{},
			},
			// Write only elements
		},
	}
}

func (o *settingsAuthGoogleOauth2DataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
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
			fmt.Sprintf("Unable to read resource for SettingsAuthGoogleOauth2 on %s", endpoint),
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
	if err = hookSettingsAuthGoogleOauth2(ctx, ApiVersion, hooks.SourceData, hooks.CalleeRead, nil, &state); err != nil {
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
