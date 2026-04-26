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

type adHocCommandDataSource = framework.GenericDataSource[adHocCommandTerraformModel, *adHocCommandTerraformModel]

// NewAdHocCommandDataSource is a helper function to instantiate the AdHocCommand data source.
func NewAdHocCommandDataSource() datasource.DataSource {
	return &adHocCommandDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "ad_hoc_command", Endpoint: "/api/v2/ad_hoc_commands/"}},
		Cfg: framework.DataSourceCfg[adHocCommandTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"become_enabled": schema.BoolAttribute{
						Description: "Become enabled",
						Computed:    true,
					},
					"canceled_on": schema.StringAttribute{
						Description: "The date and time when the cancel request was sent.",
						Computed:    true,
					},
					"controller_node": schema.StringAttribute{
						Description: "The instance that managed the execution environment.",
						Computed:    true,
					},
					"credential": schema.Int64Attribute{
						Description: "Credential",
						Computed:    true,
					},
					"diff_mode": schema.BoolAttribute{
						Description: "Diff mode",
						Computed:    true,
					},
					"elapsed": schema.Float64Attribute{
						Description: "Elapsed time in seconds that the job ran.",
						Computed:    true,
					},
					"execution_environment": schema.Int64Attribute{
						Description: "The container image to be used for execution.",
						Computed:    true,
					},
					"execution_node": schema.StringAttribute{
						Description: "The node the job executed on.",
						Computed:    true,
					},
					"extra_vars": schema.StringAttribute{
						Description: "Extra vars",
						Computed:    true,
					},
					"failed": schema.BoolAttribute{
						Description: "Failed",
						Computed:    true,
					},
					"finished": schema.StringAttribute{
						Description: "The date and time the job finished execution.",
						Computed:    true,
					},
					"forks": schema.Int64Attribute{
						Description: "Forks",
						Computed:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this ad hoc command.",
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
					"job_explanation": schema.StringAttribute{
						Description: "A status field to indicate the state of the job if it wasn't able to run and capture stdout",
						Computed:    true,
					},
					"job_type": schema.StringAttribute{
						Description: "Job type",
						Computed:    true,
					},
					"launch_type": schema.StringAttribute{
						Description: "Launch type",
						Computed:    true,
					},
					"launched_by": schema.Int64Attribute{
						Description: "Launched by",
						Computed:    true,
					},
					"limit": schema.StringAttribute{
						Description: "Limit",
						Computed:    true,
					},
					"module_args": schema.StringAttribute{
						Description: "Module args",
						Computed:    true,
					},
					"module_name": schema.StringAttribute{
						Description: "Module name",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of this ad hoc command.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.ExactlyOneOf(
								path.MatchRoot("id"),
								path.MatchRoot("name"),
							),
						},
					},
					"started": schema.StringAttribute{
						Description: "The date and time the job was queued for starting.",
						Computed:    true,
					},
					"status": schema.StringAttribute{
						Description: "Status",
						Computed:    true,
					},
					"verbosity": schema.StringAttribute{
						Description: "Verbosity",
						Computed:    true,
					},
					"work_unit_id": schema.StringAttribute{
						Description: "The Receptor work unit ID associated with this job.",
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
			Hook: func(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *adHocCommandTerraformModel) error {
				return hooks.RequireResourceStateOrOrig(ctx, apiVersion, source, callee, orig, state)
			},
			ApiVersion:   ApiVersion,
			ResourceName: "AdHocCommand",
		},
	}
}
