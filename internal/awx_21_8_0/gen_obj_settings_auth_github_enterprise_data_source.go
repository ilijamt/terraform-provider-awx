package awx_21_8_0

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
)

var (
	_ datasource.DataSource              = &settingsAuthGithubEnterpriseDataSource{}
	_ datasource.DataSourceWithConfigure = &settingsAuthGithubEnterpriseDataSource{}
)

// NewSettingsAuthGithubEnterpriseDataSource is a helper function to instantiate the SettingsAuthGithubEnterprise data source.
func NewSettingsAuthGithubEnterpriseDataSource() datasource.DataSource {
	return &settingsAuthGithubEnterpriseDataSource{}
}

// settingsAuthGithubEnterpriseDataSource is the data source implementation.
type settingsAuthGithubEnterpriseDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *settingsAuthGithubEnterpriseDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/github-enterprise/"
}

// Metadata returns the data source type name.
func (o *settingsAuthGithubEnterpriseDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_auth_github_enterprise"
}

// Schema defines the schema for the data source.
func (o *settingsAuthGithubEnterpriseDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"social_auth_github_enterprise_api_url": schema.StringAttribute{
				Description: "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_github_enterprise_callback_url": schema.StringAttribute{
				Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_github_enterprise_key": schema.StringAttribute{
				Description: "The OAuth2 key (Client ID) from your GitHub Enterprise developer application.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_github_enterprise_organization_map": schema.StringAttribute{
				Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_github_enterprise_secret": schema.StringAttribute{
				Description: "The OAuth2 secret (Client Secret) from your GitHub Enterprise developer application.",
				Sensitive:   true,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_github_enterprise_team_map": schema.StringAttribute{
				Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_github_enterprise_url": schema.StringAttribute{
				Description: "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			// Write only elements
		},
	}
}

func (o *settingsAuthGithubEnterpriseDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// Read refreshes the Terraform state with the latest data.
func (o *settingsAuthGithubEnterpriseDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsAuthGithubEnterpriseTerraformModel
	var err error
	var endpoint = o.endpoint

	// Creates a new request for SettingsAuthGithubEnterprise
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubEnterprise on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for SettingsAuthGithubEnterprise
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthGithubEnterprise on %s", endpoint),
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
	if err = hookSettingsAuthGithubEnterprise(ctx, ApiVersion, hooks.SourceData, hooks.CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthGithubEnterprise",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
