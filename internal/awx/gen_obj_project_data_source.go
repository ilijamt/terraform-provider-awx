package awx

import (
	"context"
	"fmt"
	"net/http"
	p "path"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &projectDataSource{}
	_ datasource.DataSourceWithConfigure = &projectDataSource{}
)

// NewProjectDataSource is a helper function to instantiate the Project data source.
func NewProjectDataSource() datasource.DataSource {
	return &projectDataSource{}
}

// projectDataSource is the data source implementation.
type projectDataSource struct {
	client   c.Client
	endpoint string
}

// Configure adds the provider configured client to the data source.
func (o *projectDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	o.client = req.ProviderData.(c.Client)
	o.endpoint = "/api/v2/projects/"
}

// Metadata returns the data source type name.
func (o *projectDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

// Schema defines the schema for the data source.
func (o *projectDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// Data only elements
			"allow_override": schema.BoolAttribute{
				Description: "Allow changing the SCM branch or revision in a job template that uses this project.",
				Sensitive:   false,
				Computed:    true,
			},
			"credential": schema.Int64Attribute{
				Description: "Credential",
				Sensitive:   false,
				Computed:    true,
			},
			"default_environment": schema.Int64Attribute{
				Description: "The default execution environment for jobs run using this project.",
				Sensitive:   false,
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "Optional description of this project.",
				Sensitive:   false,
				Computed:    true,
			},
			"id": schema.Int64Attribute{
				Description: "Database ID for this project.",
				Sensitive:   false,
				Optional:    true,
				Computed:    true,
			},
			"local_path": schema.StringAttribute{
				Description: "Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project.",
				Sensitive:   false,
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of this project.",
				Sensitive:   false,
				Computed:    true,
			},
			"organization": schema.Int64Attribute{
				Description: "The organization used to determine access to this template.",
				Sensitive:   false,
				Computed:    true,
			},
			"scm_branch": schema.StringAttribute{
				Description: "Specific branch, tag or commit to checkout.",
				Sensitive:   false,
				Computed:    true,
			},
			"scm_clean": schema.BoolAttribute{
				Description: "Discard any local changes before syncing the project.",
				Sensitive:   false,
				Computed:    true,
			},
			"scm_delete_on_update": schema.BoolAttribute{
				Description: "Delete the project before syncing.",
				Sensitive:   false,
				Computed:    true,
			},
			"scm_refspec": schema.StringAttribute{
				Description: "For git projects, an additional refspec to fetch.",
				Sensitive:   false,
				Computed:    true,
			},
			"scm_revision": schema.StringAttribute{
				Description: "The last revision fetched by a project update",
				Sensitive:   false,
				Computed:    true,
			},
			"scm_track_submodules": schema.BoolAttribute{
				Description: "Track submodules latest commits on defined branch.",
				Sensitive:   false,
				Computed:    true,
			},
			"scm_type": schema.StringAttribute{
				Description: "Specifies the source control system used to store the project.",
				Sensitive:   false,
				Computed:    true,
			},
			"scm_update_cache_timeout": schema.Int64Attribute{
				Description: "The number of seconds after the last project update ran that a new project update will be launched as a job dependency.",
				Sensitive:   false,
				Computed:    true,
			},
			"scm_update_on_launch": schema.BoolAttribute{
				Description: "Update the project when a job is launched that uses the project.",
				Sensitive:   false,
				Computed:    true,
			},
			"scm_url": schema.StringAttribute{
				Description: "The location where the project is stored.",
				Sensitive:   false,
				Computed:    true,
			},
			"signature_validation_credential": schema.Int64Attribute{
				Description: "An optional credential used for validating files in the project against unexpected changes.",
				Sensitive:   false,
				Computed:    true,
			},
			"timeout": schema.Int64Attribute{
				Description: "The amount of time (in seconds) to run before the task is canceled.",
				Sensitive:   false,
				Computed:    true,
			},
			// Write only elements
		},
	}
}

func (o *projectDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("id"),
		),
	}
}

// Read refreshes the Terraform state with the latest data.
func (o *projectDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state projectTerraformModel
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

	// Creates a new request for Project
	var r *http.Request
	if r, err = o.client.NewRequest(ctx, http.MethodGet, endpoint, nil); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Project on %s for read", o.endpoint),
			err.Error(),
		)
		return
	}

	// Try and actually fetch the data for Project
	var data map[string]any
	if data, err = o.client.Do(ctx, r); err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to read resource for Project on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	var d diag.Diagnostics

	if data, d, err = extractDataIfSearchResult(data); err != nil {
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
