package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type settingsBulkTerraformModel struct {
	BULK_HOST_MAX_CREATE types.Int64 `tfsdk:"bulk_host_max_create" json:"BULK_HOST_MAX_CREATE"`
	BULK_HOST_MAX_DELETE types.Int64 `tfsdk:"bulk_host_max_delete" json:"BULK_HOST_MAX_DELETE"`
	BULK_JOB_MAX_LAUNCH  types.Int64 `tfsdk:"bulk_job_max_launch" json:"BULK_JOB_MAX_LAUNCH"`
}

func (o *settingsBulkTerraformModel) Clone() settingsBulkTerraformModel {
	return *o
}

func (o *settingsBulkTerraformModel) BodyRequest() *settingsBulkBodyRequestModel {
	var req settingsBulkBodyRequestModel
	req.BULK_HOST_MAX_CREATE = o.BULK_HOST_MAX_CREATE.ValueInt64()
	req.BULK_HOST_MAX_DELETE = o.BULK_HOST_MAX_DELETE.ValueInt64()
	req.BULK_JOB_MAX_LAUNCH = o.BULK_JOB_MAX_LAUNCH.ValueInt64()
	return &req
}

func (o *settingsBulkTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.BULK_HOST_MAX_CREATE, data["BULK_HOST_MAX_CREATE"]))
	collect(helpers.AttrValueSetInt64(&o.BULK_HOST_MAX_DELETE, data["BULK_HOST_MAX_DELETE"]))
	collect(helpers.AttrValueSetInt64(&o.BULK_JOB_MAX_LAUNCH, data["BULK_JOB_MAX_LAUNCH"]))
	return diags, nil
}

type settingsBulkBodyRequestModel struct {
	BULK_HOST_MAX_CREATE int64 `json:"BULK_HOST_MAX_CREATE"`
	BULK_HOST_MAX_DELETE int64 `json:"BULK_HOST_MAX_DELETE"`
	BULK_JOB_MAX_LAUNCH  int64 `json:"BULK_JOB_MAX_LAUNCH"`
}

type settingsBulkResource = framework.GenericResource[settingsBulkTerraformModel, settingsBulkBodyRequestModel, *settingsBulkTerraformModel]

// NewSettingsBulkResource is a helper function to simplify the provider implementation.
func NewSettingsBulkResource() resource.Resource {
	return &settingsBulkResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_bulk", Endpoint: "/api/v2/settings/bulk/"}},
		Cfg: framework.ResourceCfg[settingsBulkTerraformModel, settingsBulkBodyRequestModel]{
			Schema: schema.Schema{
				Attributes: map[string]schema.Attribute{
					"bulk_host_max_create": schema.Int64Attribute{
						Description: "Max number of hosts to allow to be created in a single bulk action",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(100),
					},
					"bulk_host_max_delete": schema.Int64Attribute{
						Description: "Max number of hosts to allow to be deleted in a single bulk action",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(250),
					},
					"bulk_job_max_launch": schema.Int64Attribute{
						Description: "Max jobs to allow bulk jobs to launch",
						Optional:    true,
						Computed:    true,
						Default:     int64default.StaticInt64(100),
					},
				},
			},
			NoId:         true,
			UnDeletable:  true,
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsBulk",
		},
	}
}

type settingsBulkDataSource = framework.GenericDataSource[settingsBulkTerraformModel, *settingsBulkTerraformModel]

// NewSettingsBulkDataSource is a helper function to instantiate the SettingsBulk data source.
func NewSettingsBulkDataSource() datasource.DataSource {
	return &settingsBulkDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "settings_bulk", Endpoint: "/api/v2/settings/bulk/"}},
		Cfg: framework.DataSourceCfg[settingsBulkTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"bulk_host_max_create": dschema.Int64Attribute{
						Description: "Max number of hosts to allow to be created in a single bulk action",
						Computed:    true,
					},
					"bulk_host_max_delete": dschema.Int64Attribute{
						Description: "Max number of hosts to allow to be deleted in a single bulk action",
						Computed:    true,
					},
					"bulk_job_max_launch": dschema.Int64Attribute{
						Description: "Max jobs to allow bulk jobs to launch",
						Computed:    true,
					},
				},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "SettingsBulk",
		},
	}
}
