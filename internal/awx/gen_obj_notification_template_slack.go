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

type notificationTemplateSlackTerraformModel struct {
	ID               types.Int64  `tfsdk:"id" json:"id"`
	Name             types.String `tfsdk:"name" json:"name"`
	Description      types.String `tfsdk:"description" json:"description"`
	Organization     types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages         types.String `tfsdk:"messages" json:"messages"`
	Channels         types.List   `tfsdk:"channels" json:"-"`
	Token            types.String `tfsdk:"token" json:"-"`
}

func (o *notificationTemplateSlackTerraformModel) Clone() notificationTemplateSlackTerraformModel {
	return *o
}

type notificationTemplateSlackBodyRequestModel struct {
	Name                      string          `json:"name"`
	Description               string          `json:"description,omitempty"`
	Organization              int64           `json:"organization"`
	NotificationType          string          `json:"notification_type"`
	NotificationConfiguration map[string]any  `json:"notification_configuration"`
	Messages                  json.RawMessage `json:"messages,omitempty"`
}

// BodyRequest drops unset values rather than sending zero values, so AWX
// applies its own defaults.
func (o *notificationTemplateSlackTerraformModel) BodyRequest() *notificationTemplateSlackBodyRequestModel {
	req := &notificationTemplateSlackBodyRequestModel{
		Name:             o.Name.ValueString(),
		Description:      o.Description.ValueString(),
		Organization:     o.Organization.ValueInt64(),
		NotificationType: "slack",
	}
	if v := o.Messages.ValueString(); v != "" {
		req.Messages = json.RawMessage(v)
	}

	config := map[string]any{}
	if !o.Channels.IsNull() && !o.Channels.IsUnknown() {
		config["channels"] = helpers.ListAsStringSlice(o.Channels, false)
	}
	if !o.Token.IsNull() && !o.Token.IsUnknown() {
		config["token"] = o.Token.ValueString()
	}
	req.NotificationConfiguration = config
	return req
}

func (o *notificationTemplateSlackTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetListString(&o.Channels, config["channels"], false))
		collect(helpers.AttrValueSetString(&o.Token, config["token"], false))
	}
	return diags, nil
}

// AWX answers every secret field with `$encrypted$`, which would otherwise read
// as drift on every plan.
func hookNotificationTemplateSlack(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *notificationTemplateSlackTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// AWX never returns this value, so state has to carry a known null.
		if orig.Token.IsNull() || orig.Token.IsUnknown() {
			state.Token = types.StringNull()
		} else {
			state.Token = orig.Token
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.Token, state.Token); subbed {
			state.Token = v
		}
	}
	return nil
}

type notificationTemplateSlackDataSourceTerraformModel struct {
	ID               types.Int64  `tfsdk:"id" json:"id"`
	Name             types.String `tfsdk:"name" json:"name"`
	Description      types.String `tfsdk:"description" json:"description"`
	Organization     types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages         types.String `tfsdk:"messages" json:"messages"`
	Channels         types.List   `tfsdk:"channels" json:"-"`
}

func (o *notificationTemplateSlackDataSourceTerraformModel) Clone() notificationTemplateSlackDataSourceTerraformModel {
	return *o
}

func (o *notificationTemplateSlackDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetListString(&o.Channels, config["channels"], false))
	}
	return diags, nil
}

type notificationTemplateSlackResource = framework.GenericResource[notificationTemplateSlackTerraformModel, notificationTemplateSlackBodyRequestModel, *notificationTemplateSlackTerraformModel]

func NewNotificationTemplateSlackResource() resource.Resource {
	attrs := framework.NotificationBaseResourceAttrs()
	attrs["channels"] = schema.ListAttribute{
		Description: "Destination Channels.",
		ElementType: types.StringType,
		Required:    true,
	}
	attrs["token"] = schema.StringAttribute{
		Description: "Token.",
		Required:    true,
		Sensitive:   true,
	}
	return &notificationTemplateSlackResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template_slack", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.ResourceCfg[notificationTemplateSlackTerraformModel, notificationTemplateSlackBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages an AWX `Slack` (`slack`) notification template with first-class typed configuration attributes. Equivalent to `awx_notification_template` with `notification_type = \"slack\"`, but with per-field schema validation and sensitivity instead of a JSON `notification_configuration` string.",
				Attributes:          attrs,
			},
			IDAccessor: func(m *notificationTemplateSlackTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:      "id",
			Hook:       hookNotificationTemplateSlack,
			WriteOnlyPlanToState: func(plan, state *notificationTemplateSlackTerraformModel) {
				if state.NotificationType.IsNull() || state.NotificationType.IsUnknown() {
					state.NotificationType = types.StringValue("slack")
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "NotificationTemplateSlack",
		},
	}
}

type notificationTemplateSlackDataSource = framework.GenericDataSource[notificationTemplateSlackDataSourceTerraformModel, *notificationTemplateSlackDataSourceTerraformModel]

func NewNotificationTemplateSlackDataSource() datasource.DataSource {
	attrs := framework.NotificationBaseDataSourceAttrs()
	attrs["channels"] = dschema.ListAttribute{
		Description: "Destination Channels.",
		ElementType: types.StringType,
		Computed:    true,
	}
	return &notificationTemplateSlackDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template_slack", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.DataSourceCfg[notificationTemplateSlackDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `Slack` (`slack`) notification template by ID or name.",
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
			ResourceName: "NotificationTemplateSlack",
		},
	}
}
