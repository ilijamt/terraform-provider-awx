package provider

import (
	"context"
	"github.com/ilijamt/terraform-provider-awx/internal/awx"
	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &Provider{}
var _ provider.ProviderWithMetadata = &Provider{}

// Provider defines the provider implementation.
type Provider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
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

func (p *Provider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"hostname": {
				Description: "The AWX Host that we connect to.",
				Optional:    true,
				Type:        types.StringType,
			},
			"insecure_skip_verify": {
				Description: "Are we using a self signed certificate?",
				Optional:    true,
				Type:        types.BoolType,
			},
			"username": {
				Description: "The username to connect to the AWX host",
				Optional:    true,
				Type:        types.StringType,
			},
			"password": {
				Description: "The password to connect to the AWX host",
				Optional:    true,
				Sensitive:   true,
				Type:        types.StringType,
			},
		},
	}, nil
}

func (p *Provider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config Model

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if config.Hostname.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("hostname"),
			"Unknown AWX API Host",
			"The provider cannot create the AWX API client as there is an unknown configuration value for the AWX API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the TOWER_HOST environment variable.",
		)
	}

	if config.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown AWX API username",
			"The provider cannot create the AWX API client as there is an unknown configuration value for the AWX API username. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the TOWER_USERNAME environment variable.",
		)
	}

	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown AWX API password",
			"The provider cannot create the AWX API client as there is an unknown configuration value for the AWX API password. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the TOWER_PASSWORD environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	hostname := os.Getenv("TOWER_HOST")
	username := os.Getenv("TOWER_USERNAME")
	password := os.Getenv("TOWER_PASSWORD")
	insecureSkipVerify := false

	if !config.Hostname.IsNull() {
		hostname = config.Hostname.ValueString()
	}

	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}

	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	if "" == hostname {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown AWX API Host",
			"The provider cannot create the AWX API client as there is an unknown configuration value for the AWX API host. "+
				"Set the host value in the configuration or use the TOWER_HOST environment variable."+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if "" == username {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown AWX API Username",
			"The provider cannot create the AWX API client as there is an unknown configuration value for the AWX API username. "+
				"Set the host value in the configuration or use the TOWER_USERNAME environment variable."+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if "" == password {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown AWX API Password",
			"The provider cannot create the AWX API client as there is an unknown configuration value for the AWX API password. "+
				"Set the host value in the configuration or use the TOWER_PASSWORD environment variable."+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	var client = c.NewClient(username, password, hostname, insecureSkipVerify)
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *Provider) Resources(ctx context.Context) []func() resource.Resource {
	return awx.Resources()
}

func (p *Provider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return awx.DataSources()
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &Provider{
			version: version,
		}
	}
}
