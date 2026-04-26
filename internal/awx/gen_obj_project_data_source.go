package awx

import (
	"github.com/ilijamt/terraform-provider-awx/internal/framework"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type projectDataSource = framework.GenericDataSource[projectTerraformModel, *projectTerraformModel]

// NewProjectDataSource is a helper function to instantiate the Project data source.
func NewProjectDataSource() datasource.DataSource {
	return &projectDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "project", Endpoint: "/api/v2/projects/"}},
		Cfg: framework.DataSourceCfg[projectTerraformModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"allow_override": schema.BoolAttribute{
						Description: "Allow changing the SCM branch or revision in a job template that uses this project.",
						Computed:    true,
					},
					"credential": schema.Int64Attribute{
						Description: "Credential",
						Computed:    true,
					},
					"default_environment": schema.Int64Attribute{
						Description: "The default execution environment for jobs run using this project.",
						Computed:    true,
					},
					"description": schema.StringAttribute{
						Description: "Optional description of this project.",
						Computed:    true,
					},
					"id": schema.Int64Attribute{
						Description: "Database ID for this project.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.ExactlyOneOf(
								path.MatchRoot("id"),
							),
						},
					},
					"local_path": schema.StringAttribute{
						Description: "Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project.",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of this project.",
						Computed:    true,
					},
					"organization": schema.Int64Attribute{
						Description: "The organization used to determine access to this template.",
						Computed:    true,
					},
					"scm_branch": schema.StringAttribute{
						Description: "Specific branch, tag or commit to checkout.",
						Computed:    true,
					},
					"scm_clean": schema.BoolAttribute{
						Description: "Discard any local changes before syncing the project.",
						Computed:    true,
					},
					"scm_delete_on_update": schema.BoolAttribute{
						Description: "Delete the project before syncing.",
						Computed:    true,
					},
					"scm_refspec": schema.StringAttribute{
						Description: "For git projects, an additional refspec to fetch.",
						Computed:    true,
					},
					"scm_track_submodules": schema.BoolAttribute{
						Description: "Track submodules latest commits on defined branch.",
						Computed:    true,
					},
					"scm_type": schema.StringAttribute{
						Description: "Specifies the source control system used to store the project.",
						Computed:    true,
					},
					"scm_update_cache_timeout": schema.Int64Attribute{
						Description: "The number of seconds after the last project update ran that a new project update will be launched as a job dependency.",
						Computed:    true,
					},
					"scm_update_on_launch": schema.BoolAttribute{
						Description: "Update the project when a job is launched that uses the project.",
						Computed:    true,
					},
					"scm_url": schema.StringAttribute{
						Description: "The location where the project is stored.",
						Computed:    true,
					},
					"signature_validation_credential": schema.Int64Attribute{
						Description: "An optional credential used for validating files in the project against unexpected changes.",
						Computed:    true,
					},
					"timeout": schema.Int64Attribute{
						Description: "The amount of time (in seconds) to run before the task is canceled.",
						Computed:    true,
					},
				},
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "Project",
		},
	}
}
