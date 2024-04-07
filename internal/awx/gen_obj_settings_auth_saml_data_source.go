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
	_ datasource.DataSource              = &settingsAuthSamlDataSource{}
	_ datasource.DataSourceWithConfigure = &settingsAuthSamlDataSource{}
)

// NewSettingsAuthSAMLDataSource is a helper function to instantiate the SettingsAuthSAML data source.
func NewSettingsAuthSAMLDataSource() datasource.DataSource {
	return &settingsAuthSamlDataSource{}
}

// settingsAuthSamlDataSource is the data source implementation.
type settingsAuthSamlDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *settingsAuthSamlDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/saml/"
}

// Metadata returns the data source type name.
func (o *settingsAuthSamlDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_auth_saml"
}

// Schema defines the schema for the data source.
func (o *settingsAuthSamlDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"saml_auto_create_objects": schema.BoolAttribute{
				Description: "When enabled (the default), mapped Organizations and Teams will be created automatically on successful SAML login.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"social_auth_saml_callback_url": schema.StringAttribute{
				Description: "Register the service as a service provider (SP) with each identity provider (IdP) you have configured. Provide your SP Entity ID and this ACS URL for your application.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_saml_enabled_idps": schema.StringAttribute{
				Description: "Configure the Entity ID, SSO URL and certificate for each identity provider (IdP) in use. Multiple SAML IdPs are supported. Some IdPs may provide user data using attribute names that differ from the default OIDs. Attribute names may be overridden for each IdP. Refer to the Ansible documentation for additional details and syntax.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_saml_extra_data": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "A list of tuples that maps IDP attributes to extra_attributes. Each attribute will be a list of values, even if only 1 value.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.List{},
			},
			"social_auth_saml_metadata_url": schema.StringAttribute{
				Description: "If your identity provider (IdP) allows uploading an XML metadata file, you can download one from this URL.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_saml_organization_attr": schema.StringAttribute{
				Description: "Used to translate user organization membership.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_saml_organization_map": schema.StringAttribute{
				Description: "Mapping to organization admins/users from social auth accounts. This setting\ncontrols which users are placed into which organizations based on their\nusername and email address. Configuration details are available in the\ndocumentation.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_saml_org_info": schema.StringAttribute{
				Description: "Provide the URL, display name, and the name of your app. Refer to the documentation for example syntax.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_saml_security_config": schema.StringAttribute{
				Description: "A dict of key value pairs that are passed to the underlying python-saml security setting https://github.com/onelogin/python-saml#settings",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_saml_sp_entity_id": schema.StringAttribute{
				Description: "The application-defined unique identifier used as the audience of the SAML service provider (SP) configuration. This is usually the URL for the service.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_saml_sp_extra": schema.StringAttribute{
				Description: "A dict of key value pairs to be passed to the underlying python-saml Service Provider configuration setting.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_saml_sp_private_key": schema.StringAttribute{
				Description: "Create a keypair to use as a service provider (SP) and include the private key content here.",
				Sensitive:   true,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_saml_sp_public_cert": schema.StringAttribute{
				Description: "Create a keypair to use as a service provider (SP) and include the certificate content here.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_saml_support_contact": schema.StringAttribute{
				Description: "Provide the name and email address of the support contact for your service provider. Refer to the documentation for example syntax.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_saml_team_attr": schema.StringAttribute{
				Description: "Used to translate user team membership.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_saml_team_map": schema.StringAttribute{
				Description: "Mapping of team members (users) from social auth accounts. Configuration\ndetails are available in the documentation.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_saml_technical_contact": schema.StringAttribute{
				Description: "Provide the name and email address of the technical contact for your service provider. Refer to the documentation for example syntax.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"social_auth_saml_user_flags_by_attr": schema.StringAttribute{
				Description: "Used to map super users and system auditors from SAML.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
		},
	}
}

func (o *settingsAuthSamlDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// Read refreshes the Terraform state with the latest data.
func (o *settingsAuthSamlDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsAuthSamlTerraformModel
	var err error
	var endpoint = o.endpoint

	// Creates a new request for SettingsAuthSAML
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsAuthSAML on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for SettingsAuthSAML
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsAuthSAML on %s", endpoint),
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
	if err = hookSettingsSaml(ctx, ApiVersion, hooks.SourceData, hooks.CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			"Unable to process custom hook for the state on SettingsAuthSAML",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
