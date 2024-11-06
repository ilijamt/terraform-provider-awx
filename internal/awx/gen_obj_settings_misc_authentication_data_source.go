package awx

import (
	"context"
	"fmt"
	"net/http"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &settingsMiscAuthenticationDataSource{}
	_ datasource.DataSourceWithConfigure = &settingsMiscAuthenticationDataSource{}
)

// NewSettingsMiscAuthenticationDataSource is a helper function to instantiate the SettingsMiscAuthentication data source.
func NewSettingsMiscAuthenticationDataSource() datasource.DataSource {
	return &settingsMiscAuthenticationDataSource{}
}

// settingsMiscAuthenticationDataSource is the data source implementation.
type settingsMiscAuthenticationDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *settingsMiscAuthenticationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/authentication/"
}

// Metadata returns the data source type name.
func (o *settingsMiscAuthenticationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_misc_authentication"
}

// Schema defines the schema for the data source.
func (o *settingsMiscAuthenticationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"allow_metrics_for_anonymous_users": schema.BoolAttribute{
				Description: "If true, anonymous users are allowed to poll metrics.",
				Sensitive:   false,
				Computed:    true,
			},
			"allow_oauth2_for_external_users": schema.BoolAttribute{
				Description: "For security reasons, users from external auth providers (LDAP, SAML, SSO, Radius, and others) are not allowed to create OAuth2 tokens. To change this behavior, enable this setting. Existing tokens will not be deleted when this setting is toggled off.",
				Sensitive:   false,
				Computed:    true,
			},
			"authentication_backends": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "List of authentication backends that are enabled based on license features and other authentication settings.",
				Sensitive:   false,
				Computed:    true,
			},
			"auth_basic_enabled": schema.BoolAttribute{
				Description: "Enable HTTP Basic Auth for the API Browser.",
				Sensitive:   false,
				Computed:    true,
			},
			"disable_local_auth": schema.BoolAttribute{
				Description: "Controls whether users are prevented from using the built-in authentication system. You probably want to do this if you are using an LDAP or SAML integration.",
				Sensitive:   false,
				Computed:    true,
			},
			"local_password_min_digits": schema.Int64Attribute{
				Description: "Minimum number of digit characters required in a local password. 0 means no minimum",
				Sensitive:   false,
				Computed:    true,
			},
			"local_password_min_length": schema.Int64Attribute{
				Description: "Minimum number of characters required in a local password. 0 means no minimum",
				Sensitive:   false,
				Computed:    true,
			},
			"local_password_min_special": schema.Int64Attribute{
				Description: "Minimum number of special characters required in a local password. 0 means no minimum",
				Sensitive:   false,
				Computed:    true,
			},
			"local_password_min_upper": schema.Int64Attribute{
				Description: "Minimum number of uppercase characters required in a local password. 0 means no minimum",
				Sensitive:   false,
				Computed:    true,
			},
			"login_redirect_override": schema.StringAttribute{
				Description: "URL to which unauthorized users will be redirected to log in.  If blank, users will be sent to the login page.",
				Sensitive:   false,
				Computed:    true,
			},
			"oauth2_provider": schema.StringAttribute{
				Description: "Dictionary for customizing OAuth 2 timeouts, available items are `ACCESS_TOKEN_EXPIRE_SECONDS`, the duration of access tokens in the number of seconds, `AUTHORIZATION_CODE_EXPIRE_SECONDS`, the duration of authorization codes in the number of seconds, and `REFRESH_TOKEN_EXPIRE_SECONDS`, the duration of refresh tokens, after expired access tokens, in the number of seconds.",
				Sensitive:   false,
				Computed:    true,
			},
			"sessions_per_user": schema.Int64Attribute{
				Description: "Maximum number of simultaneous logged in sessions a user may have. To disable enter -1.",
				Sensitive:   false,
				Computed:    true,
			},
			"session_cookie_age": schema.Int64Attribute{
				Description: "Number of seconds that a user is inactive before they will need to login again.",
				Sensitive:   false,
				Computed:    true,
			},
			"social_auth_organization_map": schema.StringAttribute{
				Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
				Sensitive:   false,
				Computed:    true,
			},
			"social_auth_team_map": schema.StringAttribute{
				Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
				Sensitive:   false,
				Computed:    true,
			},
			"social_auth_username_is_full_email": schema.BoolAttribute{
				Description: "Enabling this setting will tell social auth to use the full Email as username instead of the full name",
				Sensitive:   false,
				Computed:    true,
			},
			"social_auth_user_fields": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "When set to an empty list `[]`, this setting prevents new user accounts from being created. Only users who have previously logged in using social auth or have a user account with a matching email address will be able to login.",
				Sensitive:   false,
				Computed:    true,
			},
		},
	}
}

func (o *settingsMiscAuthenticationDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// Read refreshes the Terraform state with the latest data.
func (o *settingsMiscAuthenticationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsMiscAuthenticationTerraformModel
	var err error
	var endpoint = o.endpoint

	// Creates a new request for SettingsMiscAuthentication
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsMiscAuthentication on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for SettingsMiscAuthentication
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsMiscAuthentication on %s", endpoint),
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
