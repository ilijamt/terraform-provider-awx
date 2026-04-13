package awx

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type settingsMiscAuthenticationResource = framework.GenericResource[settingsMiscAuthenticationTerraformModel, settingsMiscAuthenticationBodyRequestModel, *settingsMiscAuthenticationTerraformModel]

// NewSettingsMiscAuthenticationResource is a helper function to simplify the provider implementation.
func NewSettingsMiscAuthenticationResource() resource.Resource {
	return &settingsMiscAuthenticationResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_misc_authentication", Endpoint: "/api/v2/settings/authentication/"}},
		Cfg: framework.ResourceCfg[settingsMiscAuthenticationTerraformModel, settingsMiscAuthenticationBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					// Request elements
					"allow_metrics_for_anonymous_users": schema.BoolAttribute{
						Description: "If true, anonymous users are allowed to poll metrics.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					"allow_oauth2_for_external_users": schema.BoolAttribute{
						Description: "For security reasons, users from external auth providers (LDAP, SAML, SSO, Radius, and others) are not allowed to create OAuth2 tokens. To change this behavior, enable this setting. Existing tokens will not be deleted when this setting is toggled off.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					"auth_basic_enabled": schema.BoolAttribute{
						Description: "Enable HTTP Basic Auth for the API Browser.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					"disable_local_auth": schema.BoolAttribute{
						Description: "Controls whether users are prevented from using the built-in authentication system. You probably want to do this if you are using an LDAP or SAML integration.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					"local_password_min_digits": schema.Int64Attribute{
						Description: "Minimum number of digit characters required in a local password. 0 means no minimum",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{},
					},
					"local_password_min_length": schema.Int64Attribute{
						Description: "Minimum number of characters required in a local password. 0 means no minimum",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{},
					},
					"local_password_min_special": schema.Int64Attribute{
						Description: "Minimum number of special characters required in a local password. 0 means no minimum",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{},
					},
					"local_password_min_upper": schema.Int64Attribute{
						Description: "Minimum number of uppercase characters required in a local password. 0 means no minimum",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(0),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{},
					},
					"login_redirect_override": schema.StringAttribute{
						Description: "URL to which unauthorized users will be redirected to log in.  If blank, users will be sent to the login page.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(``),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"oauth2_provider": schema.StringAttribute{
						Description: "Dictionary for customizing OAuth 2 timeouts, available items are `ACCESS_TOKEN_EXPIRE_SECONDS`, the duration of access tokens in the number of seconds, `AUTHORIZATION_CODE_EXPIRE_SECONDS`, the duration of authorization codes in the number of seconds, and `REFRESH_TOKEN_EXPIRE_SECONDS`, the duration of refresh tokens, after expired access tokens, in the number of seconds.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     stringdefault.StaticString(`{"ACCESS_TOKEN_EXPIRE_SECONDS":31536000000,"AUTHORIZATION_CODE_EXPIRE_SECONDS":600,"REFRESH_TOKEN_EXPIRE_SECONDS":2628000}`),
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"sessions_per_user": schema.Int64Attribute{
						Description: "Maximum number of simultaneous logged in sessions a user may have. To disable enter -1.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(-1),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{},
					},
					"session_cookie_age": schema.Int64Attribute{
						Description: "Number of seconds that a user is inactive before they will need to login again.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(1800),
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Int64{
							int64validator.Between(60, 30000000000),
						},
					},
					"social_auth_organization_map": schema.StringAttribute{
						Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"social_auth_team_map": schema.StringAttribute{
						Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{},
					},
					"social_auth_username_is_full_email": schema.BoolAttribute{
						Description: "Enabling this setting will tell social auth to use the full Email as username instead of the full name",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.Bool{},
					},
					"social_auth_user_fields": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "When set to an empty list `[]`, this setting prevents new user accounts from being created. Only users who have previously logged in using social auth or have a user account with a matching email address will be able to login.",
						Sensitive:   false,
						Required:    false,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.List{},
					},
					// Write only elements
					// Data only elements
					"authentication_backends": schema.ListAttribute{
						ElementType: types.StringType,
						Description: "List of authentication backends that are enabled based on license features and other authentication settings.",
						Required:    false,
						Optional:    false,
						Computed:    true,
						Sensitive:   false,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			NoId:         true,
			UnDeletable:  true,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsMiscAuthentication",
		},
	}
}
