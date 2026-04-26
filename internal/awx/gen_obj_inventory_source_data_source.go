package awx

import (
	"context"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type inventorySourceDataSource = framework.GenericDataSource[inventorySourceTerraformModel, *inventorySourceTerraformModel]

// NewInventorySourceDataSource is a helper function to instantiate the InventorySource data source.
func NewInventorySourceDataSource() datasource.DataSource {
	return &inventorySourceDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "inventory_source", Endpoint: "/api/v2/inventory_sources/"}},
		Cfg: framework.DataSourceCfg[inventorySourceTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"credential": schema.Int64Attribute{
						Description: "Cloud credential to use for inventory updates.",
						Computed:    true,
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this inventory source.",
						Computed:    true,
					},
					"enabled_value": schema.StringAttribute{
						Description: "Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=\"status.power_state\"and enabled_value=\"powered_on\" with host variables:{   \"status\": {     \"power_state\": \"powered_on\",     \"created\": \"2020-08-04T18:13:04+00:00\",     \"healthy\": true    },    \"name\": \"foobar\",    \"ip_address\": \"192.168.2.1\"}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled",
						Computed:    true,
					},
					"enabled_var": schema.StringAttribute{
						Description: "Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as \"foo.bar\", in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(\"foo\", {}).get(\"bar\", default)",
						Computed:    true,
					},
					"execution_environment": schema.Int64Attribute{
						Description: "The container image to be used for execution.",
						Computed:    true,
					},
					"host_filter": schema.StringAttribute{
						DeprecationMessage: "This field is deprecated and will be removed in a future release.",
						Description:        "Regex where only matching hosts will be imported.",
						Computed:           true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this inventory source.",
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
						Computed:    true,
					},
					"limit": schema.StringAttribute{
						Description: "Enter host, group or pattern match",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of this inventory source.",
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
						Computed:    true,
					},
					"overwrite_vars": schema.BoolAttribute{
						Description: "Overwrite local variables from remote inventory source.",
						Computed:    true,
					},
					"scm_branch": schema.StringAttribute{
						Description: "Inventory source SCM branch. Project default used if blank. Only allowed if project allow_override field is set to true.",
						Computed:    true,
					},
					"source": schema.StringAttribute{
						Description: "Source",
						Computed:    true,
					},
					"source_path": schema.StringAttribute{
						Description: "Source path",
						Computed:    true,
					},
					"source_project": schema.Int64Attribute{
						Description: "Project containing inventory file used as source.",
						Computed:    true,
					},
					"source_vars": schema.StringAttribute{
						Description: "Inventory source variables in YAML or JSON format.",
						Computed:    true,
					},
					"timeout": schema.Int64Attribute{
						Description: "The amount of time (in seconds) to run before the task is canceled.",
						Computed:    true,
					},
					"update_cache_timeout": schema.Int64Attribute{
						Description: "Update cache timeout",
						Computed:    true,
					},
					"update_on_launch": schema.BoolAttribute{
						Description: "Update on launch",
						Computed:    true,
					},
					"verbosity": schema.StringAttribute{
						Description: "Verbosity",
						Computed:    true,
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
				{Name: "by_name", URLSuffix: "?name__exact=%s", Fields: []framework.SearchField{
					{Name: "name", Type: "string", URLEscape: true},
				}},
			},
			Hook: func(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *inventorySourceTerraformModel) error {
				return hooks.RequireResourceStateOrOrig(ctx, apiVersion, source, callee, orig, state)
			},
			ApiVersion:   ApiVersion,
			ResourceName: "InventorySource",
		},
	}
}
