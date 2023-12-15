package awx

import (
	"context"
	"fmt"
	"net/http"
	p "path"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &instanceGroupDataSource{}
	_ datasource.DataSourceWithConfigure = &instanceGroupDataSource{}
)

// NewInstanceGroupDataSource is a helper function to instantiate the InstanceGroup data source.
func NewInstanceGroupDataSource() datasource.DataSource {
	return &instanceGroupDataSource{}
}

// instanceGroupDataSource is the data source implementation.
type instanceGroupDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *instanceGroupDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/instance_groups/"
}

// Metadata returns the data source type name.
func (o *instanceGroupDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_instance_group"
}

// Schema defines the schema for the data source.
func (o *instanceGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"capacity": schema.Int64Attribute{
				Description: "Capacity",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"consumed_capacity": schema.Float64Attribute{
				Description: "Consumed capacity",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Float64{},
			},
			"credential": schema.Int64Attribute{
				Description: "Credential",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"id": schema.Int64Attribute{
				Description: "Database ID for this instance group.",
				Sensitive:   false,
				Optional:    true,
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.ExactlyOneOf(
						path.MatchRoot("id"),
					),
				},
			},
			"instances": schema.Int64Attribute{
				Description: "Instances",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"is_container_group": schema.BoolAttribute{
				Description: "Indicates whether instances in this group are containerized.Containerized groups have a designated Openshift or Kubernetes cluster.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Bool{},
			},
			"jobs_running": schema.Int64Attribute{
				Description: "Jobs running",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"jobs_total": schema.Int64Attribute{
				Description: "Count of all jobs that target this instance group",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"max_concurrent_jobs": schema.Int64Attribute{
				Description: "Maximum number of concurrent jobs to run on a group. When set to zero, no maximum is enforced.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"max_forks": schema.Int64Attribute{
				Description: "Maximum number of forks to execute concurrently on a group. When set to zero, no maximum is enforced.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"name": schema.StringAttribute{
				Description: "Name of this instance group.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"percent_capacity_remaining": schema.Float64Attribute{
				Description: "Percent capacity remaining",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Float64{},
			},
			"pod_spec_override": schema.StringAttribute{
				Description: "Pod spec override",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"policy_instance_list": schema.StringAttribute{
				Description: "List of exact-match Instances that will be assigned to this group",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.String{},
			},
			"policy_instance_minimum": schema.Int64Attribute{
				Description: "Static minimum number of Instances that will be automatically assign to this group when new instances come online.",
				Sensitive:   false,
				Computed:    true,
				Validators:  []validator.Int64{},
			},
			"policy_instance_percentage": schema.Int64Attribute{
				Description: "Minimum percentage of all instances that will be automatically assigned to this group when new instances come online.",
				Sensitive:   false,
				Computed:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 100),
				},
			},
			// Write only elements
		},
	}
}

func (o *instanceGroupDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// Read refreshes the Terraform state with the latest data.
func (o *instanceGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state instanceGroupTerraformModel
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

	if !searchDefined {
		var detailMessage string
		resp.Diagnostics.AddError(
			"missing configuration for one of the predefined search groups",
			detailMessage,
		)
		return
	}

	// Creates a new request for InstanceGroup
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for InstanceGroup on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for InstanceGroup
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for InstanceGroup on %s", endpoint),
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
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
