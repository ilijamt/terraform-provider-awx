package awx

import (
	"context"
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
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
)

type notificationTemplateGrafanaTerraformModel struct {
	ID               types.Int64  `tfsdk:"id" json:"id"`
	Name             types.String `tfsdk:"name" json:"name"`
	Description      types.String `tfsdk:"description" json:"description"`
	Organization     types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages         types.String `tfsdk:"messages" json:"messages"`
	GrafanaKey       types.String `tfsdk:"grafana_key" json:"-"`
	GrafanaUrl       types.String `tfsdk:"grafana_url" json:"-"`
}

func (o *notificationTemplateGrafanaTerraformModel) Clone() notificationTemplateGrafanaTerraformModel {
	return *o
}

type notificationTemplateGrafanaBodyRequestModel struct {
	Name                      string          `json:"name"`
	Description               string          `json:"description,omitempty"`
	Organization              int64           `json:"organization"`
	NotificationType          string          `json:"notification_type"`
	NotificationConfiguration map[string]any  `json:"notification_configuration"`
	Messages                  json.RawMessage `json:"messages,omitempty"`
}

// BodyRequest drops unset values rather than sending zero values, so AWX
// applies its own defaults.
func (o *notificationTemplateGrafanaTerraformModel) BodyRequest() *notificationTemplateGrafanaBodyRequestModel {
	req := &notificationTemplateGrafanaBodyRequestModel{
		Name:             o.Name.ValueString(),
		Description:      o.Description.ValueString(),
		Organization:     o.Organization.ValueInt64(),
		NotificationType: "grafana",
	}
	if v := o.Messages.ValueString(); v != "" {
		req.Messages = json.RawMessage(v)
	}

	config := map[string]any{}
	if !o.GrafanaKey.IsNull() && !o.GrafanaKey.IsUnknown() {
		config["grafana_key"] = o.GrafanaKey.ValueString()
	}
	if !o.GrafanaUrl.IsNull() && !o.GrafanaUrl.IsUnknown() {
		config["grafana_url"] = o.GrafanaUrl.ValueString()
	}
	req.NotificationConfiguration = config
	return req
}

func (o *notificationTemplateGrafanaTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.GrafanaKey, config["grafana_key"], false))
		collect(helpers.AttrValueSetString(&o.GrafanaUrl, config["grafana_url"], false))
	}
	return diags, nil
}

// AWX answers every secret field with `$encrypted$`, which would otherwise read
// as drift on every plan.
func hookNotificationTemplateGrafana(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *notificationTemplateGrafanaTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// AWX never returns this value, so state has to carry a known null.
		if orig.GrafanaKey.IsNull() || orig.GrafanaKey.IsUnknown() {
			state.GrafanaKey = types.StringNull()
		} else {
			state.GrafanaKey = orig.GrafanaKey
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.GrafanaKey, state.GrafanaKey); subbed {
			state.GrafanaKey = v
		}
	}
	return nil
}

type notificationTemplateGrafanaDataSourceTerraformModel struct {
	ID               types.Int64  `tfsdk:"id" json:"id"`
	Name             types.String `tfsdk:"name" json:"name"`
	Description      types.String `tfsdk:"description" json:"description"`
	Organization     types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages         types.String `tfsdk:"messages" json:"messages"`
	GrafanaUrl       types.String `tfsdk:"grafana_url" json:"-"`
}

func (o *notificationTemplateGrafanaDataSourceTerraformModel) Clone() notificationTemplateGrafanaDataSourceTerraformModel {
	return *o
}

func (o *notificationTemplateGrafanaDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.GrafanaUrl, config["grafana_url"], false))
	}
	return diags, nil
}

type notificationTemplateGrafanaResource = framework.GenericResource[notificationTemplateGrafanaTerraformModel, notificationTemplateGrafanaBodyRequestModel, *notificationTemplateGrafanaTerraformModel]

func NewNotificationTemplateGrafanaResource() resource.Resource {
	attrs := framework.NotificationBaseResourceAttrs()
	attrs["grafana_key"] = schema.StringAttribute{
		Description: "Grafana API Key.",
		Required:    true,
		Sensitive:   true,
	}
	attrs["grafana_url"] = schema.StringAttribute{
		Description: "Grafana URL.",
		Required:    true,
	}
	return &notificationTemplateGrafanaResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template_grafana", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.ResourceCfg[notificationTemplateGrafanaTerraformModel, notificationTemplateGrafanaBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages an AWX `Grafana` (`grafana`) notification template with first-class typed configuration attributes. Equivalent to `awx_notification_template` with `notification_type = \"grafana\"`, but with per-field schema validation and sensitivity instead of a JSON `notification_configuration` string.",
				Attributes:          attrs,
			},
			IDAccessor: func(m *notificationTemplateGrafanaTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:      "id",
			Hook:       hookNotificationTemplateGrafana,
			WriteOnlyPlanToState: func(plan, state *notificationTemplateGrafanaTerraformModel) {
				if state.NotificationType.IsNull() || state.NotificationType.IsUnknown() {
					state.NotificationType = types.StringValue("grafana")
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "NotificationTemplateGrafana",
		},
	}
}

type notificationTemplateGrafanaDataSource = framework.GenericDataSource[notificationTemplateGrafanaDataSourceTerraformModel, *notificationTemplateGrafanaDataSourceTerraformModel]

func NewNotificationTemplateGrafanaDataSource() datasource.DataSource {
	attrs := framework.NotificationBaseDataSourceAttrs()
	attrs["grafana_url"] = dschema.StringAttribute{
		Description: "Grafana URL.",
		Computed:    true,
	}
	return &notificationTemplateGrafanaDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template_grafana", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.DataSourceCfg[notificationTemplateGrafanaDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `Grafana` (`grafana`) notification template by ID or name.",
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
			ResourceName: "NotificationTemplateGrafana",
		},
	}
}
