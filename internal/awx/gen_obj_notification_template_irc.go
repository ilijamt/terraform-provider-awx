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

type notificationTemplateIrcTerraformModel struct {
	ID               types.Int64  `tfsdk:"id" json:"id"`
	Name             types.String `tfsdk:"name" json:"name"`
	Description      types.String `tfsdk:"description" json:"description"`
	Organization     types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages         types.String `tfsdk:"messages" json:"messages"`
	Nickname         types.String `tfsdk:"nickname" json:"-"`
	Password         types.String `tfsdk:"password" json:"-"`
	Port             types.Int64  `tfsdk:"port" json:"-"`
	Server           types.String `tfsdk:"server" json:"-"`
	Targets          types.List   `tfsdk:"targets" json:"-"`
	UseSsl           types.Bool   `tfsdk:"use_ssl" json:"-"`
}

func (o *notificationTemplateIrcTerraformModel) Clone() notificationTemplateIrcTerraformModel {
	return *o
}

type notificationTemplateIrcBodyRequestModel struct {
	Name                      string          `json:"name"`
	Description               string          `json:"description,omitempty"`
	Organization              int64           `json:"organization"`
	NotificationType          string          `json:"notification_type"`
	NotificationConfiguration map[string]any  `json:"notification_configuration"`
	Messages                  json.RawMessage `json:"messages,omitempty"`
}

// BodyRequest drops unset values rather than sending zero values, so AWX
// applies its own defaults.
func (o *notificationTemplateIrcTerraformModel) BodyRequest() *notificationTemplateIrcBodyRequestModel {
	req := &notificationTemplateIrcBodyRequestModel{
		Name:             o.Name.ValueString(),
		Description:      o.Description.ValueString(),
		Organization:     o.Organization.ValueInt64(),
		NotificationType: "irc",
	}
	if v := o.Messages.ValueString(); v != "" {
		req.Messages = json.RawMessage(v)
	}

	config := map[string]any{}
	if !o.Nickname.IsNull() && !o.Nickname.IsUnknown() {
		config["nickname"] = o.Nickname.ValueString()
	}
	if !o.Password.IsNull() && !o.Password.IsUnknown() {
		config["password"] = o.Password.ValueString()
	}
	if !o.Port.IsNull() && !o.Port.IsUnknown() {
		config["port"] = o.Port.ValueInt64()
	}
	if !o.Server.IsNull() && !o.Server.IsUnknown() {
		config["server"] = o.Server.ValueString()
	}
	if !o.Targets.IsNull() && !o.Targets.IsUnknown() {
		config["targets"] = helpers.ListAsStringSlice(o.Targets, false)
	}
	if !o.UseSsl.IsNull() && !o.UseSsl.IsUnknown() {
		config["use_ssl"] = o.UseSsl.ValueBool()
	}
	req.NotificationConfiguration = config
	return req
}

func (o *notificationTemplateIrcTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.Nickname, config["nickname"], false))
		collect(helpers.AttrValueSetString(&o.Password, config["password"], false))
		collect(helpers.AttrValueSetInt64(&o.Port, config["port"]))
		collect(helpers.AttrValueSetString(&o.Server, config["server"], false))
		collect(helpers.AttrValueSetListString(&o.Targets, config["targets"], false))
		collect(helpers.AttrValueSetBool(&o.UseSsl, config["use_ssl"]))
	}
	return diags, nil
}

// AWX answers every secret field with `$encrypted$`, which would otherwise read
// as drift on every plan.
func hookNotificationTemplateIrc(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *notificationTemplateIrcTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// AWX never returns this value, so state has to carry a known null.
		if orig.Password.IsNull() || orig.Password.IsUnknown() {
			state.Password = types.StringNull()
		} else {
			state.Password = orig.Password
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.Password, state.Password); subbed {
			state.Password = v
		}
	}
	return nil
}

type notificationTemplateIrcDataSourceTerraformModel struct {
	ID               types.Int64  `tfsdk:"id" json:"id"`
	Name             types.String `tfsdk:"name" json:"name"`
	Description      types.String `tfsdk:"description" json:"description"`
	Organization     types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages         types.String `tfsdk:"messages" json:"messages"`
	Nickname         types.String `tfsdk:"nickname" json:"-"`
	Port             types.Int64  `tfsdk:"port" json:"-"`
	Server           types.String `tfsdk:"server" json:"-"`
	Targets          types.List   `tfsdk:"targets" json:"-"`
	UseSsl           types.Bool   `tfsdk:"use_ssl" json:"-"`
}

func (o *notificationTemplateIrcDataSourceTerraformModel) Clone() notificationTemplateIrcDataSourceTerraformModel {
	return *o
}

func (o *notificationTemplateIrcDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.Nickname, config["nickname"], false))
		collect(helpers.AttrValueSetInt64(&o.Port, config["port"]))
		collect(helpers.AttrValueSetString(&o.Server, config["server"], false))
		collect(helpers.AttrValueSetListString(&o.Targets, config["targets"], false))
		collect(helpers.AttrValueSetBool(&o.UseSsl, config["use_ssl"]))
	}
	return diags, nil
}

type notificationTemplateIrcResource = framework.GenericResource[notificationTemplateIrcTerraformModel, notificationTemplateIrcBodyRequestModel, *notificationTemplateIrcTerraformModel]

func NewNotificationTemplateIrcResource() resource.Resource {
	attrs := framework.NotificationBaseResourceAttrs()
	attrs["nickname"] = schema.StringAttribute{
		Description: "IRC Nick.",
		Required:    true,
	}
	attrs["password"] = schema.StringAttribute{
		Description: "IRC Server Password.",
		Required:    true,
		Sensitive:   true,
	}
	attrs["port"] = schema.Int64Attribute{
		Description: "IRC Server Port.",
		Required:    true,
	}
	attrs["server"] = schema.StringAttribute{
		Description: "IRC Server Address.",
		Required:    true,
	}
	attrs["targets"] = schema.ListAttribute{
		Description: "Destination Channels or Users.",
		ElementType: types.StringType,
		Required:    true,
	}
	attrs["use_ssl"] = schema.BoolAttribute{
		Description: "SSL Connection.",
		Required:    true,
	}
	return &notificationTemplateIrcResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template_irc", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.ResourceCfg[notificationTemplateIrcTerraformModel, notificationTemplateIrcBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages an AWX `IRC` (`irc`) notification template with first-class typed configuration attributes. Equivalent to `awx_notification_template` with `notification_type = \"irc\"`, but with per-field schema validation and sensitivity instead of a JSON `notification_configuration` string.",
				Attributes:          attrs,
			},
			IDAccessor: func(m *notificationTemplateIrcTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:      "id",
			Hook:       hookNotificationTemplateIrc,
			WriteOnlyPlanToState: func(plan, state *notificationTemplateIrcTerraformModel) {
				if state.NotificationType.IsNull() || state.NotificationType.IsUnknown() {
					state.NotificationType = types.StringValue("irc")
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "NotificationTemplateIrc",
		},
	}
}

type notificationTemplateIrcDataSource = framework.GenericDataSource[notificationTemplateIrcDataSourceTerraformModel, *notificationTemplateIrcDataSourceTerraformModel]

func NewNotificationTemplateIrcDataSource() datasource.DataSource {
	attrs := framework.NotificationBaseDataSourceAttrs()
	attrs["nickname"] = dschema.StringAttribute{
		Description: "IRC Nick.",
		Computed:    true,
	}
	attrs["port"] = dschema.Int64Attribute{
		Description: "IRC Server Port.",
		Computed:    true,
	}
	attrs["server"] = dschema.StringAttribute{
		Description: "IRC Server Address.",
		Computed:    true,
	}
	attrs["targets"] = dschema.ListAttribute{
		Description: "Destination Channels or Users.",
		ElementType: types.StringType,
		Computed:    true,
	}
	attrs["use_ssl"] = dschema.BoolAttribute{
		Description: "SSL Connection.",
		Computed:    true,
	}
	return &notificationTemplateIrcDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template_irc", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.DataSourceCfg[notificationTemplateIrcDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `IRC` (`irc`) notification template by ID or name.",
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
			ResourceName: "NotificationTemplateIrc",
		},
	}
}
