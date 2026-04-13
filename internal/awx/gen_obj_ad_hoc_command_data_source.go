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
					// Data only elements
					"become_enabled": schema.BoolAttribute{
						Description: "Become enabled",
						Sensitive:   false,
						Computed:    true,
					},
					"canceled_on": schema.StringAttribute{
						Description: "The date and time when the cancel request was sent.",
						Sensitive:   false,
						Computed:    true,
					},
					"controller_node": schema.StringAttribute{
						Description: "The instance that managed the execution environment.",
						Sensitive:   false,
						Computed:    true,
					},
					"credential": schema.Int64Attribute{
						Description: "Credential",
						Sensitive:   false,
						Computed:    true,
					},
					"diff_mode": schema.BoolAttribute{
						Description: "Diff mode",
						Sensitive:   false,
						Computed:    true,
					},
					"elapsed": schema.Float64Attribute{
						Description: "Elapsed time in seconds that the job ran.",
						Sensitive:   false,
						Computed:    true,
					},
					"execution_environment": schema.Int64Attribute{
						Description: "The container image to be used for execution.",
						Sensitive:   false,
						Computed:    true,
					},
					"execution_node": schema.StringAttribute{
						Description: "The node the job executed on.",
						Sensitive:   false,
						Computed:    true,
					},
					"extra_vars": schema.StringAttribute{
						Description: "Extra vars",
						Sensitive:   false,
						Computed:    true,
					},
					"failed": schema.BoolAttribute{
						Description: "Failed",
						Sensitive:   false,
						Computed:    true,
					},
					"finished": schema.StringAttribute{
						Description: "The date and time the job finished execution.",
						Sensitive:   false,
						Computed:    true,
					},
					"forks": schema.Int64Attribute{
						Description: "Forks",
						Sensitive:   false,
						Computed:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this ad hoc command.",
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
					},
					"job_explanation": schema.StringAttribute{
						Description: "A status field to indicate the state of the job if it wasn't able to run and capture stdout",
						Sensitive:   false,
						Computed:    true,
					},
					"job_type": schema.StringAttribute{
						Description: "Job type",
						Sensitive:   false,
						Computed:    true,
					},
					"launch_type": schema.StringAttribute{
						Description: "Launch type",
						Sensitive:   false,
						Computed:    true,
					},
					"launched_by": schema.Int64Attribute{
						Description: "Launched by",
						Sensitive:   false,
						Computed:    true,
					},
					"limit": schema.StringAttribute{
						Description: "Limit",
						Sensitive:   false,
						Computed:    true,
					},
					"module_args": schema.StringAttribute{
						Description: "Module args",
						Sensitive:   false,
						Computed:    true,
					},
					"module_name": schema.StringAttribute{
						Description: "Module name",
						Sensitive:   false,
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of this ad hoc command.",
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
					"started": schema.StringAttribute{
						Description: "The date and time the job was queued for starting.",
						Sensitive:   false,
						Computed:    true,
					},
					"status": schema.StringAttribute{
						Description: "Status",
						Sensitive:   false,
						Computed:    true,
					},
					"verbosity": schema.StringAttribute{
						Description: "Verbosity",
						Sensitive:   false,
						Computed:    true,
					},
					"work_unit_id": schema.StringAttribute{
						Description: "The Receptor work unit ID associated with this job.",
						Sensitive:   false,
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
