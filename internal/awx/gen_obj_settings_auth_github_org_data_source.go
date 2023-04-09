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

// Schema defines the schema for the data source.
func (o *settingsAuthGithubOrgDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"social_auth_github_org_callback_url": schema.StringAttribute{
				Description: "Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail.",
				Sensitive:   false,
				Computed:    true,
			},
			"social_auth_github_org_key": schema.StringAttribute{
				Description: "The OAuth2 key (Client ID) from your GitHub organization application.",
				Sensitive:   false,
				Computed:    true,
			},
			"social_auth_github_org_name": schema.StringAttribute{
				Description: "The name of your GitHub organization, as used in your organization's URL: https://github.com/<yourorg>/.",
				Sensitive:   false,
				Computed:    true,
			},
			"social_auth_github_org_organization_map": schema.StringAttribute{
				Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
				Sensitive:   false,
				Computed:    true,
			},
			"social_auth_github_org_secret": schema.StringAttribute{
				Description: "The OAuth2 secret (Client Secret) from your GitHub organization application.",
				Sensitive:   true,
				Computed:    true,
			},
			"social_auth_github_org_team_map": schema.StringAttribute{
				Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
				Sensitive:   false,
				Computed:    true,
			},
			// Write only elements
		},
	}
}

func (o *settingsAuthGithubOrgDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(),
	}
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
	if err = hookSettingsAuthGithubOrg(ctx, ApiVersion, SourceData, CalleeRead, nil, &state); err != nil {
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
