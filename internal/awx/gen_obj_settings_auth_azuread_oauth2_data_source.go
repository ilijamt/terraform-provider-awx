package awx

import (
	"context"
	"fmt"
	"net/http"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

var (
	_ datasource.DataSource              = &settingsAuthAzureAdoauth2DataSource{}
	_ datasource.DataSourceWithConfigure = &settingsAuthAzureAdoauth2DataSource{}
)

// NewSettingsAuthAzureADOauth2DataSource is a helper function to instantiate the SettingsAuthAzureADOauth2 data source.
func NewSettingsAuthAzureADOauth2DataSource() datasource.DataSource {
	return &settingsAuthAzureAdoauth2DataSource{}
}

// settingsAuthAzureAdoauth2DataSource is the data source implementation.
type settingsAuthAzureAdoauth2DataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *settingsAuthAzureAdoauth2DataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/azuread-oauth2/"
}

// Metadata returns the data source type name.
func (o *settingsAuthAzureAdoauth2DataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_auth_azuread_oauth2"
}

// Schema defines the schema for the data source.
func (o *settingsAuthAzureAdoauth2DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"social_auth_azuread_oauth2_callback_url": schema.StringAttribute{
				Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail. ",
				Sensitive:   false,
				Computed:    true,
			},
			"social_auth_azuread_oauth2_key": schema.StringAttribute{
				Description: "The OAuth2 key (Client ID) from your Azure AD application.",
				Sensitive:   false,
				Computed:    true,
			},
			"social_auth_azuread_oauth2_organization_map": schema.StringAttribute{
				Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
				Sensitive:   false,
				Computed:    true,
			},
			"social_auth_azuread_oauth2_secret": schema.StringAttribute{
				Description: "The OAuth2 secret (Client Secret) from your Azure AD application.",
				Sensitive:   true,
				Computed:    true,
			},
			"social_auth_azuread_oauth2_team_map": schema.StringAttribute{
				Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
				Sensitive:   false,
				Computed:    true,
			},
			// Write only elements
		},
	}
}

func (o *settingsAuthAzureAdoauth2DataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(),
	}
}

// Read refreshes the Terraform state with the latest data.
func (o *settingsAuthAzureAdoauth2DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsAuthAzureAdoauth2TerraformModel
	var err error
	var endpoint = o.endpoint

	// Creates a new request for SettingsAuthAzureADOauth2
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthAzureADOauth2 on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for SettingsAuthAzureADOauth2
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthAzureADOauth2 on %s", o.endpoint),
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
	if err = hookSettingsAuthAzureADOauth2(ctx, ApiVersion, SourceData, CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthAzureADOauth2",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
