package provider

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = (*Provider)(nil)

// Provider defines the provider implementation.
type Provider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version    string
	config     Model
	httpClient *http.Client
	awxClient  c.Client

	fnResources   []func() resource.Resource
	fnDataSources []func() datasource.DataSource
}

// Model describes the provider data model.
type Model struct {
	Hostname  types.String `tfsdk:"hostname"`
	Username  types.String `tfsdk:"username"`
	Password  types.String `tfsdk:"password"`
	Token     types.String `tfsdk:"token"`
	VerifySSL types.Bool   `tfsdk:"verify_ssl"`
}

func (p *Provider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "awx"
	resp.Version = p.version
}

func (p *Provider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"hostname": schema.StringAttribute{
				Description: "The AWX Host that we connect to. (defaults to TOWER_HOST/AWX_HOST env variable if set)",
				Optional:    true,
				Required:    false,
			},
			"verify_ssl": schema.BoolAttribute{
				Description: "If you are using a self signed certificate this should be set to false (defaults to TOWER_VERIFY_SSL/VERIFY_SSL env variable if set) [default is true]",
				Optional:    true,
				Required:    false,
			},
			"username": schema.StringAttribute{
				Description: "The username to connect to the AWX host. (defaults to TOWER_USERNAME/AWX_USERNAME env variable if set) [must be used with password]",
				Optional:    true,
				Required:    false,
			},
			"password": schema.StringAttribute{
				Description: "The password to connect to the AWX host. (defaults to TOWER_PASSWORD/AWX_PASSWORD env variable if set) [must be used with username]",
				Optional:    true,
				Sensitive:   true,
				Required:    false,
				Validators: []validator.String{
					stringvalidator.AlsoRequires(
						path.Expressions{
							path.MatchRoot("username"),
						}...,
					),
				},
			},
			"token": schema.StringAttribute{
				Required:    false,
				Optional:    true,
				Sensitive:   true,
				Description: "The token to use to connect to the AWX host. (defaults to TOWER_AUTH_TOKEN/AWX_AUTH_TOKEN env variable if set) [conflicts with username/password]",
			},
		},
	}
}

func configureFromEnvironment(ctx context.Context, data *Model) {
	var envConfig = make(map[string]any)

	if val := helpers.GetFirstSetEnvVar("TOWER_HOST", "AWX_HOST"); val != "" && data.Hostname.IsNull() {
		data.Hostname = types.StringValue(val)
		envConfig["Hostname"] = val
	}

	if val := helpers.GetFirstSetEnvVar("TOWER_USERNAME", "AWX_USERNAME"); val != "" && data.Username.IsNull() {
		data.Username = types.StringValue(val)
		envConfig["Username"] = val
	}

	if val := helpers.GetFirstSetEnvVar("TOWER_PASSWORD", "AWX_PASSWORD"); val != "" && data.Password.IsNull() {
		data.Password = types.StringValue(val)
		envConfig["Password"] = strings.Repeat("*", len(val))
	}

	if val := helpers.GetFirstSetEnvVar("TOWER_AUTH_TOKEN", "AWX_AUTH_TOKEN"); val != "" && data.Token.IsNull() {
		data.Token = types.StringValue(val)
		envConfig["AuthToken"] = strings.Repeat("*", len(val))
	}

	if val := helpers.GetFirstSetEnvVar("TOWER_VERIFY_SSL", "AWX_VERIFY_SSL"); val != "" && data.VerifySSL.IsNull() {
		data.VerifySSL = types.BoolValue(helpers.Str2Bool(val))
		envConfig["VerifySSL"] = val
	}

	tflog.Debug(ctx, "Provider configuration from the environment", envConfig)
}

func configureDefaults(ctx context.Context, data *Model) {
	var defaults = make(map[string]any)
	if data.VerifySSL.IsNull() {
		data.VerifySSL = types.BoolValue(true)
		defaults["VerifySSL"] = data.VerifySSL.ValueBool()
	}
	tflog.Debug(ctx, "Defaults configured for provider", defaults)
}

func (p *Provider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config Model
	tflog.Debug(ctx, "Provider configuration started")

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	configureFromEnvironment(ctx, &config)
	configureDefaults(ctx, &config)

	var missingHostname = config.Hostname.ValueString() == "" || config.Hostname.IsUnknown()
	var noTokenAuth = config.Token.ValueString() == "" || config.Token.IsUnknown()
	var noBasicAuth = (config.Username.ValueString() == "" || config.Username.IsUnknown()) &&
		(config.Password.ValueString() == "" || config.Password.IsUnknown())

	if missingHostname {
		resp.Diagnostics.AddAttributeError(path.Root("host"), "Unknown AWX API Host", "The provider cannot create the AWX API client as there is an unknown configuration value for the AWX API host. "+
			"Set the host value in the configuration or use the TOWER_HOST or AWX_HOST environment variable."+
			"If either is already set, ensure the value is not empty.")
	}

	if (noTokenAuth && noBasicAuth) || (!noTokenAuth && !noBasicAuth) {
		resp.Diagnostics.AddError(
			fmt.Sprintf("must provide one of [%q, %q] or %q.", "username", "password", "token"),
			fmt.Sprintf("must provide one of [%q, %q] or %q.", "username", "password", "token"),
		)
	} else {
		if !noBasicAuth && noTokenAuth {
			if config.Username.ValueString() == "" || config.Username.IsUnknown() {
				resp.Diagnostics.AddAttributeError(path.Root("username"), "Unknown AWX API Username", "The provider cannot create the AWX API client as there is an unknown configuration value for the AWX API username. "+
					"Set the username value in the configuration or use the TOWER_USERNAME or AWX_USERNAME environment variable."+
					"If either is already set, ensure the value is not empty.")
			}

			if config.Password.ValueString() == "" || config.Password.IsUnknown() {
				resp.Diagnostics.AddAttributeError(path.Root("password"), "Unknown AWX API Password", "The provider cannot create the AWX API client as there is an unknown configuration value for the AWX API password. "+
					"Set the password value in the configuration or use the TOWER_PASSWORD or AWX_PASSWORD environment variable."+
					"If either is already set, ensure the value is not empty.")
			}
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}

	var client = p.awxClient
	if client == nil {
		if !noBasicAuth && noTokenAuth {
			client = c.NewClientWithBasicAuth(config.Username.ValueString(), config.Password.ValueString(), config.Hostname.ValueString(), p.version, !config.VerifySSL.ValueBool(), p.httpClient)
		} else {
			client = c.NewClientWithTokenAuth(config.Token.ValueString(), config.Hostname.ValueString(), p.version, !config.VerifySSL.ValueBool(), p.httpClient)
		}
		p.awxClient = client
	}
	resp.DataSourceData = client
	resp.ResourceData = client
	p.config = config
	tflog.Debug(ctx, "Provider configuration finished")
}

func (p *Provider) Resources(ctx context.Context) []func() resource.Resource {
	return p.fnResources
}

func (p *Provider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return p.fnDataSources
}

func NewFuncProvider(version string, httpClient *http.Client, awxClient c.Client, fnResources []func() resource.Resource, fnDataSources []func() datasource.DataSource) func() provider.Provider {
	return func() provider.Provider {
		return New(version, httpClient, awxClient, fnResources, fnDataSources)
	}
}

func New(version string, httpClient *http.Client, awxClient c.Client, fnResources []func() resource.Resource, fnDataSources []func() datasource.DataSource) provider.Provider {
	return &Provider{
		version:       version,
		fnResources:   fnResources,
		fnDataSources: fnDataSources,
		httpClient:    httpClient,
		awxClient:     awxClient,
	}
}
