package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"strings"

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
	version string
	config  Model

	fnResources   []func() resource.Resource
	fnDataSources []func() datasource.DataSource
}

// Model describes the provider data model.
type Model struct {
	Hostname           types.String `tfsdk:"hostname"`
	Username           types.String `tfsdk:"username"`
	Password           types.String `tfsdk:"password"`
	InsecureSkipVerify types.Bool   `tfsdk:"insecure_skip_verify"`
}

func (p *Provider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "awx"
	resp.Version = p.version
}

func (p *Provider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"hostname": schema.StringAttribute{
				Description: "The AWX Host that we connect to. (defaults to TOWER_HOST env variable if set)",
				Optional:    true,
			},
			"insecure_skip_verify": schema.BoolAttribute{
				Description: "Are we using a self signed certificate? [false] (defaults to a negation of TOWER_VERIFY_SSL env variable if set)",
				Optional:    true,
			},
			"username": schema.StringAttribute{
				Description: "The username to connect to the AWX host. (defaults to TOWER_USERNAME env variable if set)",
				Optional:    true,
			},
			"password": schema.StringAttribute{
				Description: "The password to connect to the AWX host. (defaults to TOWER_PASSWORD env variable if set)",
				Optional:    true,
				Sensitive:   true,
			},
		},
	}
}

func configureFromEnvironment(ctx context.Context, data *Model) {
	var envConfig = make(map[string]interface{})

	if val := helpers.GetFirstSetEnvVar("TOWER_HOST", "AWX_HOST"); val != "" && data.Hostname.IsNull() {
		data.Hostname = types.StringValue(val)
		envConfig["Hostname"] = data.Hostname.String()
	}

	if val := helpers.GetFirstSetEnvVar("TOWER_USERNAME", "AWX_USERNAME"); val != "" && data.Username.IsNull() {
		data.Username = types.StringValue(val)
		envConfig["Username"] = data.Username.String()
	}

	if val := helpers.GetFirstSetEnvVar("TOWER_PASSWORD", "AWX_PASSWORD"); val != "" && data.Password.IsNull() {
		data.Password = types.StringValue(val)
		envConfig["Password"] = types.StringValue(strings.Repeat("*", len(val)))
	}

	if val := helpers.GetFirstSetEnvVar("TOWER_VERIFY_SSL", "AWX_VERIFY_SSL"); val != "" && data.InsecureSkipVerify.IsNull() {
		data.InsecureSkipVerify = types.BoolValue(helpers.Str2Bool(val))
		envConfig["InsecureSkipVerify"] = data.InsecureSkipVerify.String()
	}

	tflog.Debug(ctx, "Provider configuration from the environment", envConfig)
}

func configureDefaults(ctx context.Context, data *Model) {
	var defaults = make(map[string]interface{})
	if data.InsecureSkipVerify.IsNull() {
		data.InsecureSkipVerify = types.BoolValue(false)
		defaults["InsecureSkipVerify"] = data.InsecureSkipVerify.ValueBool()
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

	if "" == config.Hostname.String() {
		resp.Diagnostics.AddAttributeError(path.Root("host"), "Unknown AWX API Host", "The provider cannot create the AWX API client as there is an unknown configuration value for the AWX API host. "+
			"Set the host value in the configuration or use the TOWER_HOST or AWX_HOST environment variable."+
			"If either is already set, ensure the value is not empty.")
	}

	if "" == config.Username.String() {
		resp.Diagnostics.AddAttributeError(path.Root("username"), "Unknown AWX API Username", "The provider cannot create the AWX API client as there is an unknown configuration value for the AWX API username. "+
			"Set the host value in the configuration or use the TOWER_USERNAME or AWX_USERNAME environment variable."+
			"If either is already set, ensure the value is not empty.")
	}

	if "" == config.Username.String() {
		resp.Diagnostics.AddAttributeError(path.Root("password"), "Unknown AWX API Password", "The provider cannot create the AWX API client as there is an unknown configuration value for the AWX API password. "+
			"Set the host value in the configuration or use the TOWER_PASSWORD or AWX_PASSWORD environment variable."+
			"If either is already set, ensure the value is not empty.")
	}

	if resp.Diagnostics.HasError() {
		return
	}

	var client = c.NewClient(config.Username.String(), config.Password.String(), config.Hostname.String(), p.version, config.InsecureSkipVerify.ValueBool())
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

func NewFuncProvider(version string, fnResources []func() resource.Resource, fnDataSources []func() datasource.DataSource) func() provider.Provider {
	return func() provider.Provider {
		return New(version, fnResources, fnDataSources)
	}
}

func New(version string, fnResources []func() resource.Resource, fnDataSources []func() datasource.DataSource) provider.Provider {
	return &Provider{
		version:       version,
		fnResources:   fnResources,
		fnDataSources: fnDataSources,
	}
}
