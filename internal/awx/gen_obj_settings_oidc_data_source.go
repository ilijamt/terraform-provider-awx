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
	_ datasource.DataSource              = &settingsOpenIDConnectDataSource{}
	_ datasource.DataSourceWithConfigure = &settingsOpenIDConnectDataSource{}
)

// NewSettingsOpenIDConnectDataSource is a helper function to instantiate the SettingsOpenIDConnect data source.
func NewSettingsOpenIDConnectDataSource() datasource.DataSource {
	return &settingsOpenIDConnectDataSource{}
}

// settingsOpenIDConnectDataSource is the data source implementation.
type settingsOpenIDConnectDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *settingsOpenIDConnectDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/settings/oidc/"
}

// Metadata returns the data source type name.
func (o *settingsOpenIDConnectDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings_oidc"
}

// Schema defines the schema for the data source.
func (o *settingsOpenIDConnectDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"social_auth_oidc_key": schema.StringAttribute{
				Description: "The OIDC key (Client ID) from your IDP.",
				Sensitive:   false,
				Computed:    true,
			},
			"social_auth_oidc_oidc_endpoint": schema.StringAttribute{
				Description: "The URL for your OIDC provider including the path up to /.well-known/openid-configuration",
				Sensitive:   false,
				Computed:    true,
			},
			"social_auth_oidc_secret": schema.StringAttribute{
				Description: "The OIDC secret (Client Secret) from your IDP.",
				Sensitive:   false,
				Computed:    true,
			},
			"social_auth_oidc_verify_ssl": schema.BoolAttribute{
				Description: "Verify the OIDV provider ssl certificate.",
				Sensitive:   false,
				Computed:    true,
			},
			// Write only elements
		},
	}
}

func (o *settingsOpenIDConnectDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(),
	}
}

// Read refreshes the Terraform state with the latest data.
func (o *settingsOpenIDConnectDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state settingsOpenIDConnectTerraformModel
	var err error
	var endpoint = o.endpoint

	// Creates a new request for SettingsOpenIDConnect
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for SettingsOpenIDConnect on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for SettingsOpenIDConnect
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for SettingsOpenIDConnect on %s", o.endpoint),
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
