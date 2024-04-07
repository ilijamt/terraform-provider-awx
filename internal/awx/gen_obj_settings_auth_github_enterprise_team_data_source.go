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
)

var (
	_ datasource.DataSource              = &settingsAuthGithubEnterpriseTeamDataSource{}
	_ datasource.DataSourceWithConfigure = &settingsAuthGithubEnterpriseTeamDataSource{}
)

// NewSettingsAuthGithubEnterpriseTeamDataSource is a helper function to instantiate the SettingsAuthGithubEnterpriseTeam data source.
func NewSettingsAuthGithubEnterpriseTeamDataSource() datasource.DataSource {
	return &settingsAuthGithubEnterpriseTeamDataSource{}
}

// settingsAuthGithubEnterpriseTeamDataSource is the data source implementation.
type settingsAuthGithubEnterpriseTeamDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *settingsAuthGithubEnterpriseTeamDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/github-enterprise-team/"
}

// Metadata returns the data source type name.
func (o *settingsAuthGithubEnterpriseTeamDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_auth_github_enterprise_team"
}

// Schema defines the schema for the data source.
func (o *settingsAuthGithubEnterpriseTeamDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"social_auth_github_enterprise_team_api_url": schema.StringAttribute{
				Description: "The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_github_enterprise_team_callback_url": schema.StringAttribute{
				Description: "Create an organization-owned application at https://github.com/organizations/<yourorg>/settings/applications and obtain an OAuth2 key (Client ID) and secret (Client Secret). Provide this URL as the callback URL for your application.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_github_enterprise_team_id": schema.StringAttribute{
				Description: "Find the numeric team ID using the Github Enterprise API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_github_enterprise_team_key": schema.StringAttribute{
				Description: "The OAuth2 key (Client ID) from your GitHub Enterprise organization application.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_github_enterprise_team_organization_map": schema.StringAttribute{
				Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_github_enterprise_team_secret": schema.StringAttribute{
				Description: "The OAuth2 secret (Client Secret) from your GitHub Enterprise organization application.",
				Sensitive:   true,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_github_enterprise_team_team_map": schema.StringAttribute{
				Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_github_enterprise_team_url": schema.StringAttribute{
				Description: "The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
		},
	}
}

func (o *settingsAuthGithubEnterpriseTeamDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// Read refreshes the Terraform state with the latest data.
func (o *settingsAuthGithubEnterpriseTeamDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsAuthGithubEnterpriseTeamTerraformModel
	var err error
	var endpoint = o.endpoint

	// Creates a new request for SettingsAuthGithubEnterpriseTeam
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthGithubEnterpriseTeam on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for SettingsAuthGithubEnterpriseTeam
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthGithubEnterpriseTeam on %s", endpoint),
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
	if err = hookSettingsAuthGithubEnterpriseTeam(ctx, ApiVersion, hooks.SourceData, hooks.CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthGithubEnterpriseTeam",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
