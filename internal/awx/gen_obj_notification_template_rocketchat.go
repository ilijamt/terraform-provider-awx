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

type notificationTemplateRocketchatTerraformModel struct {
	ID                    types.Int64  `tfsdk:"id" json:"id"`
	Name                  types.String `tfsdk:"name" json:"name"`
	Description           types.String `tfsdk:"description" json:"description"`
	Organization          types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType      types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages              types.String `tfsdk:"messages" json:"messages"`
	RocketchatNoVerifySsl types.Bool   `tfsdk:"rocketchat_no_verify_ssl" json:"-"`
	RocketchatUrl         types.String `tfsdk:"rocketchat_url" json:"-"`
}

func (o *notificationTemplateRocketchatTerraformModel) Clone() notificationTemplateRocketchatTerraformModel {
	return *o
}

type notificationTemplateRocketchatBodyRequestModel struct {
	Name                      string          `json:"name"`
	Description               string          `json:"description,omitempty"`
	Organization              int64           `json:"organization"`
	NotificationType          string          `json:"notification_type"`
	NotificationConfiguration map[string]any  `json:"notification_configuration"`
	Messages                  json.RawMessage `json:"messages,omitempty"`
}

// BodyRequest drops unset values rather than sending zero values, so AWX
// applies its own defaults.
func (o *notificationTemplateRocketchatTerraformModel) BodyRequest() *notificationTemplateRocketchatBodyRequestModel {
	req := &notificationTemplateRocketchatBodyRequestModel{
		Name:             o.Name.ValueString(),
		Description:      o.Description.ValueString(),
		Organization:     o.Organization.ValueInt64(),
		NotificationType: "rocketchat",
	}
	if v := o.Messages.ValueString(); v != "" {
		req.Messages = json.RawMessage(v)
	}

	config := map[string]any{}
	if !o.RocketchatNoVerifySsl.IsNull() && !o.RocketchatNoVerifySsl.IsUnknown() {
		config["rocketchat_no_verify_ssl"] = o.RocketchatNoVerifySsl.ValueBool()
	}
	if !o.RocketchatUrl.IsNull() && !o.RocketchatUrl.IsUnknown() {
		config["rocketchat_url"] = o.RocketchatUrl.ValueString()
	}
	req.NotificationConfiguration = config
	return req
}

func (o *notificationTemplateRocketchatTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetBool(&o.RocketchatNoVerifySsl, config["rocketchat_no_verify_ssl"]))
		collect(helpers.AttrValueSetString(&o.RocketchatUrl, config["rocketchat_url"], false))
	}
	return diags, nil
}

type notificationTemplateRocketchatDataSourceTerraformModel struct {
	ID                    types.Int64  `tfsdk:"id" json:"id"`
	Name                  types.String `tfsdk:"name" json:"name"`
	Description           types.String `tfsdk:"description" json:"description"`
	Organization          types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType      types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages              types.String `tfsdk:"messages" json:"messages"`
	RocketchatNoVerifySsl types.Bool   `tfsdk:"rocketchat_no_verify_ssl" json:"-"`
	RocketchatUrl         types.String `tfsdk:"rocketchat_url" json:"-"`
}

func (o *notificationTemplateRocketchatDataSourceTerraformModel) Clone() notificationTemplateRocketchatDataSourceTerraformModel {
	return *o
}

func (o *notificationTemplateRocketchatDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetBool(&o.RocketchatNoVerifySsl, config["rocketchat_no_verify_ssl"]))
		collect(helpers.AttrValueSetString(&o.RocketchatUrl, config["rocketchat_url"], false))
	}
	return diags, nil
}

type notificationTemplateRocketchatResource = framework.GenericResource[notificationTemplateRocketchatTerraformModel, notificationTemplateRocketchatBodyRequestModel, *notificationTemplateRocketchatTerraformModel]

func NewNotificationTemplateRocketchatResource() resource.Resource {
	attrs := framework.NotificationBaseResourceAttrs()
	attrs["rocketchat_no_verify_ssl"] = schema.BoolAttribute{
		Description: "Verify SSL.",
		Required:    true,
	}
	attrs["rocketchat_url"] = schema.StringAttribute{
		Description: "Target URL.",
		Required:    true,
	}
	return &notificationTemplateRocketchatResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template_rocketchat", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.ResourceCfg[notificationTemplateRocketchatTerraformModel, notificationTemplateRocketchatBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages an AWX `Rocket.Chat` (`rocketchat`) notification template with first-class typed configuration attributes. Equivalent to `awx_notification_template` with `notification_type = \"rocketchat\"`, but with per-field schema validation and sensitivity instead of a JSON `notification_configuration` string.",
				Attributes:          attrs,
			},
			IDAccessor: func(m *notificationTemplateRocketchatTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:      "id",
			WriteOnlyPlanToState: func(plan, state *notificationTemplateRocketchatTerraformModel) {
				if state.NotificationType.IsNull() || state.NotificationType.IsUnknown() {
					state.NotificationType = types.StringValue("rocketchat")
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "NotificationTemplateRocketchat",
		},
	}
}

type notificationTemplateRocketchatDataSource = framework.GenericDataSource[notificationTemplateRocketchatDataSourceTerraformModel, *notificationTemplateRocketchatDataSourceTerraformModel]

func NewNotificationTemplateRocketchatDataSource() datasource.DataSource {
	attrs := framework.NotificationBaseDataSourceAttrs()
	attrs["rocketchat_no_verify_ssl"] = dschema.BoolAttribute{
		Description: "Verify SSL.",
		Computed:    true,
	}
	attrs["rocketchat_url"] = dschema.StringAttribute{
		Description: "Target URL.",
		Computed:    true,
	}
	return &notificationTemplateRocketchatDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template_rocketchat", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.DataSourceCfg[notificationTemplateRocketchatDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `Rocket.Chat` (`rocketchat`) notification template by ID or name.",
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
			ResourceName: "NotificationTemplateRocketchat",
		},
	}
}
