package awx

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	p "path"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &inventorySourceDataSource{}
	_ datasource.DataSourceWithConfigure = &inventorySourceDataSource{}
)

// NewInventorySourceDataSource is a helper function to instantiate the InventorySource data source.
func NewInventorySourceDataSource() datasource.DataSource {
	return &inventorySourceDataSource{}
}

// inventorySourceDataSource is the data source implementation.
type inventorySourceDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *inventorySourceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/inventory_sources/"
}

// Metadata returns the data source type name.
func (o *inventorySourceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_inventory_source"
}

// Schema defines the schema for the data source.
func (o *inventorySourceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"credential": schema.Int64Attribute{
				Description: "Cloud credential to use for inventory updates.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"description": schema.StringAttribute{
				Description: "Optional description of this inventory source.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"enabled_value": schema.StringAttribute{
				Description: "Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=\"status.power_state\"and enabled_value=\"powered_on\" with host variables:{   \"status\": {     \"power_state\": \"powered_on\",     \"created\": \"2020-08-04T18:13:04+00:00\",     \"healthy\": true    },    \"name\": \"foobar\",    \"ip_address\": \"192.168.2.1\"}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"enabled_var": schema.StringAttribute{
				Description: "Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as \"foo.bar\", in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(\"foo\", {}).get(\"bar\", default)",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"execution_environment": schema.Int64Attribute{
				Description: "The container image to be used for execution.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"host_filter": schema.StringAttribute{
				Description: "This field is deprecated and will be removed in a future release. Regex where only matching hosts will be imported.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"id": schema.Int64Attribute{
				Description: "Database ID for this inventory source.",
				Sensitive:   false,
				Optional:    true,
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.ExactlyOneOf(
						path.MatchRoot("id"),
						path.MatchRoot("name"),
					),
				},
			},
			"inventory": schema.Int64Attribute{
				Description: "Inventory",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"limit": schema.StringAttribute{
				Description: "Enter host, group or pattern match",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"name": schema.StringAttribute{
				Description: "Name of this inventory source.",
				Sensitive:   false,
				Optional:    true,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.ExactlyOneOf(
						path.MatchRoot("id"),
						path.MatchRoot("name"),
					),
				},
			},
			"overwrite": schema.BoolAttribute{
				Description: "Overwrite local groups and hosts from remote inventory source.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"overwrite_vars": schema.BoolAttribute{
				Description: "Overwrite local variables from remote inventory source.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"scm_branch": schema.StringAttribute{
				Description: "Inventory source SCM branch. Project default used if blank. Only allowed if project allow_override field is set to true.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"source": schema.StringAttribute{
				Description: "Source",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"file", "constructed", "scm", "ec2", "gce", "azure_rm", "vmware", "satellite6", "openstack", "rhv", "controller", "insights"}...),
				},
			},
			"source_path": schema.StringAttribute{
				Description: "Source path",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"source_project": schema.Int64Attribute{
				Description: "Project containing inventory file used as source.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"source_vars": schema.StringAttribute{
				Description: "Inventory source variables in YAML or JSON format.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"timeout": schema.Int64Attribute{
				Description: "The amount of time (in seconds) to run before the task is canceled.",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(-2147483648, 2147483647),
				},
			},
			"update_cache_timeout": schema.Int64Attribute{
				Description: "Update cache timeout",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"update_on_launch": schema.BoolAttribute{
				Description: "Update on launch",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"verbosity": schema.StringAttribute{
				Description: "Verbosity",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"0", "1", "2"}...),
				},
			},
		},
	}
}

func (o *inventorySourceDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// Read refreshes the Terraform state with the latest data.
func (o *inventorySourceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state inventorySourceTerraformModel
	var err error
	var endpoint string
	var searchDefined bool

	// Only one group should evaluate to True, terraform should prevent from being able to set
	// the conflicting groups

	// Evaluate group 'by_id' based on the schema definition
	var groupByIdExists = func() bool {
		var groupByIdExists = true
		var paramsById = []any{o.endpoint}
		var attrID types.Int64
		req.Config.GetAttribute(ctx, path.Root("id"), &attrID)
		groupByIdExists = groupByIdExists && (!attrID.IsNull() && !attrID.IsUnknown())
		paramsById = append(paramsById, attrID.ValueInt64())
		if groupByIdExists {
			endpoint = p.Clean(fmt.Sprintf("%s/%d/", paramsById...))
		}
		return groupByIdExists
	}()
	searchDefined = searchDefined || groupByIdExists

	// Evaluate group 'by_name' based on the schema definition
	var groupByNameExists = func() bool {
		var groupByNameExists = true
		var paramsByName = []any{o.endpoint}
		var attrName types.String
		req.Config.GetAttribute(ctx, path.Root("name"), &attrName)
		groupByNameExists = groupByNameExists && (!attrName.IsNull() && !attrName.IsUnknown())
		paramsByName = append(paramsByName, url.PathEscape(attrName.ValueString()))
		if groupByNameExists {
			endpoint = p.Clean(fmt.Sprintf("%s/?name__exact=%s", paramsByName...))
		}
		return groupByNameExists
	}()
	searchDefined = searchDefined || groupByNameExists

	if !searchDefined {
		var detailMessage string
		resp.Diagnostics.AddError(
			"missing configuration for one of the predefined search groups",
			detailMessage,
		)
		return
	}

	// Creates a new request for InventorySource
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for InventorySource on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for InventorySource
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for InventorySource on %s", endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics

	if data, d, err = helpers.ExtractDataIfSearchResult(data); err != nil {
		resp.Diagnostics.Append(d...)
		return
	}

	if d, err = state.updateFromApiData(data); err != nil {
		resp.Diagnostics.Append(d...)
		return
	}

	// Set state
	if err = hooks.RequireResourceStateOrOrig(ctx, ApiVersion, hooks.SourceData, hooks.CalleeRead, nil, &state); err != nil {
		resp.Diagnostics.AddError(
			"Unable to process custom hook for the state on InventorySource",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
