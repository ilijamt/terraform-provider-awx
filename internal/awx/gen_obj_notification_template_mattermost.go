package awx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
)

type notificationTemplateMattermostTerraformModel struct {
	ID                    types.Int64  `tfsdk:"id" json:"id"`
	Name                  types.String `tfsdk:"name" json:"name"`
	Description           types.String `tfsdk:"description" json:"description"`
	Organization          types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType      types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages              types.String `tfsdk:"messages" json:"messages"`
	MattermostNoVerifySsl types.Bool   `tfsdk:"mattermost_no_verify_ssl" json:"-"`
	MattermostUrl         types.String `tfsdk:"mattermost_url" json:"-"`
}

func (o *notificationTemplateMattermostTerraformModel) Clone() notificationTemplateMattermostTerraformModel {
	return *o
}

type notificationTemplateMattermostBodyRequestModel struct {
	Name                      string          `json:"name"`
	Description               string          `json:"description,omitempty"`
	Organization              int64           `json:"organization"`
	NotificationType          string          `json:"notification_type"`
	NotificationConfiguration map[string]any  `json:"notification_configuration"`
	Messages                  json.RawMessage `json:"messages,omitempty"`
}

// BodyRequest drops unset values rather than sending zero values, so AWX
// applies its own defaults.
func (o *notificationTemplateMattermostTerraformModel) BodyRequest() *notificationTemplateMattermostBodyRequestModel {
	req := &notificationTemplateMattermostBodyRequestModel{
		Name:             o.Name.ValueString(),
		Description:      o.Description.ValueString(),
		Organization:     o.Organization.ValueInt64(),
		NotificationType: "mattermost",
	}
	if v := o.Messages.ValueString(); v != "" {
		req.Messages = json.RawMessage(v)
	}

	config := map[string]any{}
	if !o.MattermostNoVerifySsl.IsNull() && !o.MattermostNoVerifySsl.IsUnknown() {
		config["mattermost_no_verify_ssl"] = o.MattermostNoVerifySsl.ValueBool()
	}
	if !o.MattermostUrl.IsNull() && !o.MattermostUrl.IsUnknown() {
		config["mattermost_url"] = o.MattermostUrl.ValueString()
	}
	req.NotificationConfiguration = config
	return req
}

func (o *notificationTemplateMattermostTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
	diags := diag.Diagnostics{}
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	collect(helpers.AttrValueSetString(&o.NotificationType, data["notification_type"], false))
	collect(helpers.AttrValueSetJsonString(&o.Messages, data["messages"], false))

	if config, ok := data["notification_configuration"].(map[string]any); ok {
		collect(helpers.AttrValueSetBool(&o.MattermostNoVerifySsl, config["mattermost_no_verify_ssl"]))
		collect(helpers.AttrValueSetString(&o.MattermostUrl, config["mattermost_url"], false))
	}
	return diags, nil
}

type notificationTemplateMattermostDataSourceTerraformModel struct {
	ID                    types.Int64  `tfsdk:"id" json:"id"`
	Name                  types.String `tfsdk:"name" json:"name"`
	Description           types.String `tfsdk:"description" json:"description"`
	Organization          types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType      types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages              types.String `tfsdk:"messages" json:"messages"`
	MattermostNoVerifySsl types.Bool   `tfsdk:"mattermost_no_verify_ssl" json:"-"`
	MattermostUrl         types.String `tfsdk:"mattermost_url" json:"-"`
}

func (o *notificationTemplateMattermostDataSourceTerraformModel) Clone() notificationTemplateMattermostDataSourceTerraformModel {
	return *o
}

func (o *notificationTemplateMattermostDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
	diags := diag.Diagnostics{}
	if data == nil {
		return diags, fmt.Errorf("no data passed")
	}
	collect := func(d diag.Diagnostics, _ error) { diags.Append(d...) }
	collect(helpers.AttrValueSetInt64(&o.ID, data["id"]))
	collect(helpers.AttrValueSetString(&o.Name, data["name"], false))
	collect(helpers.AttrValueSetString(&o.Description, data["description"], false))
	collect(helpers.AttrValueSetInt64(&o.Organization, data["organization"]))
	collect(helpers.AttrValueSetString(&o.NotificationType, data["notification_type"], false))
	collect(helpers.AttrValueSetJsonString(&o.Messages, data["messages"], false))

	if config, ok := data["notification_configuration"].(map[string]any); ok {
		collect(helpers.AttrValueSetBool(&o.MattermostNoVerifySsl, config["mattermost_no_verify_ssl"]))
		collect(helpers.AttrValueSetString(&o.MattermostUrl, config["mattermost_url"], false))
	}
	return diags, nil
}

type notificationTemplateMattermostResource = framework.GenericResource[notificationTemplateMattermostTerraformModel, notificationTemplateMattermostBodyRequestModel, *notificationTemplateMattermostTerraformModel]

func NewNotificationTemplateMattermostResource() resource.Resource {
	attrs := framework.NotificationBaseResourceAttrs()
	attrs["mattermost_no_verify_ssl"] = schema.BoolAttribute{
		Description: "Verify SSL.",
		Required:    true,
	}
	attrs["mattermost_url"] = schema.StringAttribute{
		Description: "Target URL.",
		Required:    true,
	}
	return &notificationTemplateMattermostResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template_mattermost", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.ResourceCfg[notificationTemplateMattermostTerraformModel, notificationTemplateMattermostBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages an AWX `Mattermost` (`mattermost`) notification template with first-class typed configuration attributes. Equivalent to `awx_notification_template` with `notification_type = \"mattermost\"`, but with per-field schema validation and sensitivity instead of a JSON `notification_configuration` string.",
				Attributes:          attrs,
			},
			IDAccessor: func(m *notificationTemplateMattermostTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:      "id",
			WriteOnlyPlanToState: func(plan, state *notificationTemplateMattermostTerraformModel) {
				if state.NotificationType.IsNull() || state.NotificationType.IsUnknown() {
					state.NotificationType = types.StringValue("mattermost")
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "NotificationTemplateMattermost",
		},
	}
}

type notificationTemplateMattermostDataSource = framework.GenericDataSource[notificationTemplateMattermostDataSourceTerraformModel, *notificationTemplateMattermostDataSourceTerraformModel]

func NewNotificationTemplateMattermostDataSource() datasource.DataSource {
	attrs := framework.NotificationBaseDataSourceAttrs()
	attrs["mattermost_no_verify_ssl"] = dschema.BoolAttribute{
		Description: "Verify SSL.",
		Computed:    true,
	}
	attrs["mattermost_url"] = dschema.StringAttribute{
		Description: "Target URL.",
		Computed:    true,
	}
	return &notificationTemplateMattermostDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template_mattermost", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.DataSourceCfg[notificationTemplateMattermostDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `Mattermost` (`mattermost`) notification template by ID or name.",
				Attributes:          attrs,
			},
			SearchGroups: []framework.SearchGroup{
				{Name: "by_id", URLSuffix: "%d/", Fields: []framework.SearchField{
					{Name: "id", Type: "int64", URLEscape: false},
				}},
				{Name: "by_name", URLSuffix: "?name__exact=%s", Fields: []framework.SearchField{
					{Name: "name", Type: "string", URLEscape: true},
				}},
			},
			ApiVersion:   ApiVersion,
			ResourceName: "NotificationTemplateMattermost",
		},
	}
}
