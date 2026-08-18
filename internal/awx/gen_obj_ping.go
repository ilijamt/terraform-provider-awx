package awx

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type pingTerraformModel struct {
	ActiveNode     types.String `tfsdk:"active_node" json:"active_node"`
	Ha             types.Bool   `tfsdk:"ha" json:"ha"`
	InstallUuid    types.String `tfsdk:"install_uuid" json:"install_uuid"`
	InstanceGroups types.String `tfsdk:"instance_groups" json:"instance_groups"`
	Instances      types.String `tfsdk:"instances" json:"instances"`
	Version        types.String `tfsdk:"version" json:"version"`
}

func (o *pingTerraformModel) Clone() pingTerraformModel {
	return *o
}

func (o *pingTerraformModel) BodyRequest() *pingBodyRequestModel {
	var req pingBodyRequestModel
	return &req
}

func (o *pingTerraformModel) UpdateFromApiData(data map[string]any) (diags diag.Diagnostics, _ error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetString(&o.ActiveNode, data["active_node"], false))
	collect(helpers.AttrValueSetBool(&o.Ha, data["ha"]))
	collect(helpers.AttrValueSetString(&o.InstallUuid, data["install_uuid"], false))
	collect(helpers.AttrValueSetJsonString(&o.InstanceGroups, data["instance_groups"], false))
	collect(helpers.AttrValueSetJsonString(&o.Instances, data["instances"], false))
	collect(helpers.AttrValueSetString(&o.Version, data["version"], false))
	return diags, nil
}

type pingBodyRequestModel struct {
}

type pingDataSource = framework.GenericDataSource[pingTerraformModel, *pingTerraformModel]

// NewPingDataSource is a helper function to instantiate the Ping data source.
func NewPingDataSource() datasource.DataSource {
	return &pingDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "ping", Endpoint: "/api/v2/ping/"}},
		Cfg: framework.DataSourceCfg[pingTerraformModel]{
			Schema: dschema.Schema{
				Attributes: map[string]dschema.Attribute{
					"active_node": dschema.StringAttribute{
						Description: "Hostname of the node that answered this request.",
						Computed:    true,
					},
					"ha": dschema.BoolAttribute{
						Description: "Whether the installation runs more than one node.",
						Computed:    true,
					},
					"install_uuid": dschema.StringAttribute{
						Description: "Identifier unique to this installation, stable across restarts.",
						Computed:    true,
					},
					"instance_groups": dschema.StringAttribute{
						Description: "The instance groups in the mesh, each with its capacity and member nodes. Capacity shifts as jobs run, so a config that reads this diffs on every plan.",
						Computed:    true,
					},
					"instances": dschema.StringAttribute{
						Description: "The nodes in the mesh, each with its type, capacity and last heartbeat. Heartbeats move on their own, so a config that reads this diffs on every plan.",
						Computed:    true,
					},
					"version": dschema.StringAttribute{
						Description: "The AWX version this installation runs, which is not necessarily the version the provider was generated against.",
						Computed:    true,
					},
				},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "Ping",
		},
	}
}
