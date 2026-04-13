package awx

import (
	"context"
	"fmt"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/mitchellh/mapstructure"
)

var (
	_ datasource.DataSource              = &workflowJobTemplateObjectRolesDataSource{}
	_ datasource.DataSourceWithConfigure = &workflowJobTemplateObjectRolesDataSource{}
)

// NewWorkflowJobTemplateObjectRolesDataSource is a helper function to instantiate the WorkflowJobTemplate data source.
func NewWorkflowJobTemplateObjectRolesDataSource() datasource.DataSource {
	return &workflowJobTemplateObjectRolesDataSource{DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "workflow_job_template_object_roles", Endpoint: "/api/v2/workflow_job_templates/%d/object_roles/"}}}
}

// workflowJobTemplateObjectRolesDataSource is the data source implementation.
type workflowJobTemplateObjectRolesDataSource struct {
	framework.DataSourceBase
}

// Schema defines the schema for the data source.
func (o *workflowJobTemplateObjectRolesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		DeprecationMessage: "This data source has been deprecated and will be removed in a future release.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "WorkflowJobTemplate ID",
				Required:    true,
			},
			"roles": schema.MapAttribute{
				Description: "Roles for workflow_job_template",
				ElementType: types.Int64Type,
				Computed:    true,
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (o *workflowJobTemplateObjectRolesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state workflowJobTemplateObjectRolesModel
	var err error
	var id types.Int64

	if framework.DiagnosticsHasError(&resp.Diagnostics, req.Config.GetAttribute(ctx, path.Root("id"), &id)...) {
		return
	}
	state.ID = types.Int64Value(id.ValueInt64())

	var endpoint = fmt.Sprintf(o.Endpoint, id.ValueInt64())
	data, d := framework.ReadRequest(ctx, o.Client, endpoint, "WorkflowJobTemplate/ObjectRoles")
	if framework.DiagnosticsHasError(&resp.Diagnostics, d...) {
		return
	}

	var sr models.SearchResultObjectRole
	if err = mapstructure.Decode(data, &sr); err != nil {
		resp.Diagnostics.AddError(
			"Unable to decode the search result data for workflowjobtemplate",
			err.Error(),
		)
		return
	}

	var in = make(map[string]attr.Value, sr.Count)
	for _, role := range sr.Results {
		in[role.Name] = types.Int64Value(role.ID)
	}

	var dg diag.Diagnostics
	if state.Roles, dg = types.MapValue(types.Int64Type, in); framework.DiagnosticsHasError(&resp.Diagnostics, dg...) {
		return
	}

	if framework.DiagnosticsHasError(&resp.Diagnostics, resp.State.Set(ctx, &state)...) {
		return
	}
}
