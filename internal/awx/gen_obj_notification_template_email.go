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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
)

type notificationTemplateEmailTerraformModel struct {
	ID               types.Int64  `tfsdk:"id" json:"id"`
	Name             types.String `tfsdk:"name" json:"name"`
	Description      types.String `tfsdk:"description" json:"description"`
	Organization     types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages         types.String `tfsdk:"messages" json:"messages"`
	Host             types.String `tfsdk:"host" json:"-"`
	Password         types.String `tfsdk:"password" json:"-"`
	Port             types.Int64  `tfsdk:"port" json:"-"`
	Recipients       types.List   `tfsdk:"recipients" json:"-"`
	Sender           types.String `tfsdk:"sender" json:"-"`
	Timeout          types.Int64  `tfsdk:"timeout" json:"-"`
	UseSsl           types.Bool   `tfsdk:"use_ssl" json:"-"`
	UseTls           types.Bool   `tfsdk:"use_tls" json:"-"`
	Username         types.String `tfsdk:"username" json:"-"`
}

func (o *notificationTemplateEmailTerraformModel) Clone() notificationTemplateEmailTerraformModel {
	return *o
}

type notificationTemplateEmailBodyRequestModel struct {
	Name                      string          `json:"name"`
	Description               string          `json:"description,omitempty"`
	Organization              int64           `json:"organization"`
	NotificationType          string          `json:"notification_type"`
	NotificationConfiguration map[string]any  `json:"notification_configuration"`
	Messages                  json.RawMessage `json:"messages,omitempty"`
}

// BodyRequest drops unset values rather than sending zero values, so AWX
// applies its own defaults.
func (o *notificationTemplateEmailTerraformModel) BodyRequest() *notificationTemplateEmailBodyRequestModel {
	req := &notificationTemplateEmailBodyRequestModel{
		Name:             o.Name.ValueString(),
		Description:      o.Description.ValueString(),
		Organization:     o.Organization.ValueInt64(),
		NotificationType: "email",
	}
	if v := o.Messages.ValueString(); v != "" {
		req.Messages = json.RawMessage(v)
	}

	config := map[string]any{}
	if !o.Host.IsNull() && !o.Host.IsUnknown() {
		config["host"] = o.Host.ValueString()
	}
	if !o.Password.IsNull() && !o.Password.IsUnknown() {
		config["password"] = o.Password.ValueString()
	}
	if !o.Port.IsNull() && !o.Port.IsUnknown() {
		config["port"] = o.Port.ValueInt64()
	}
	if !o.Recipients.IsNull() && !o.Recipients.IsUnknown() {
		config["recipients"] = helpers.ListAsStringSlice(o.Recipients, false)
	}
	if !o.Sender.IsNull() && !o.Sender.IsUnknown() {
		config["sender"] = o.Sender.ValueString()
	}
	if !o.Timeout.IsNull() && !o.Timeout.IsUnknown() {
		config["timeout"] = o.Timeout.ValueInt64()
	}
	if !o.UseSsl.IsNull() && !o.UseSsl.IsUnknown() {
		config["use_ssl"] = o.UseSsl.ValueBool()
	}
	if !o.UseTls.IsNull() && !o.UseTls.IsUnknown() {
		config["use_tls"] = o.UseTls.ValueBool()
	}
	if !o.Username.IsNull() && !o.Username.IsUnknown() {
		config["username"] = o.Username.ValueString()
	}
	req.NotificationConfiguration = config
	return req
}

func (o *notificationTemplateEmailTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.Host, config["host"], false))
		collect(helpers.AttrValueSetString(&o.Password, config["password"], false))
		collect(helpers.AttrValueSetInt64(&o.Port, config["port"]))
		collect(helpers.AttrValueSetListString(&o.Recipients, config["recipients"], false))
		collect(helpers.AttrValueSetString(&o.Sender, config["sender"], false))
		collect(helpers.AttrValueSetInt64(&o.Timeout, config["timeout"]))
		collect(helpers.AttrValueSetBool(&o.UseSsl, config["use_ssl"]))
		collect(helpers.AttrValueSetBool(&o.UseTls, config["use_tls"]))
		collect(helpers.AttrValueSetString(&o.Username, config["username"], false))
	}
	return diags, nil
}

// AWX answers every secret field with `$encrypted$`, which would otherwise read
// as drift on every plan.
func hookNotificationTemplateEmail(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *notificationTemplateEmailTerraformModel) error {
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

type notificationTemplateEmailDataSourceTerraformModel struct {
	ID               types.Int64  `tfsdk:"id" json:"id"`
	Name             types.String `tfsdk:"name" json:"name"`
	Description      types.String `tfsdk:"description" json:"description"`
	Organization     types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages         types.String `tfsdk:"messages" json:"messages"`
	Host             types.String `tfsdk:"host" json:"-"`
	Port             types.Int64  `tfsdk:"port" json:"-"`
	Recipients       types.List   `tfsdk:"recipients" json:"-"`
	Sender           types.String `tfsdk:"sender" json:"-"`
	Timeout          types.Int64  `tfsdk:"timeout" json:"-"`
	UseSsl           types.Bool   `tfsdk:"use_ssl" json:"-"`
	UseTls           types.Bool   `tfsdk:"use_tls" json:"-"`
	Username         types.String `tfsdk:"username" json:"-"`
}

func (o *notificationTemplateEmailDataSourceTerraformModel) Clone() notificationTemplateEmailDataSourceTerraformModel {
	return *o
}

func (o *notificationTemplateEmailDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.Host, config["host"], false))
		collect(helpers.AttrValueSetInt64(&o.Port, config["port"]))
		collect(helpers.AttrValueSetListString(&o.Recipients, config["recipients"], false))
		collect(helpers.AttrValueSetString(&o.Sender, config["sender"], false))
		collect(helpers.AttrValueSetInt64(&o.Timeout, config["timeout"]))
		collect(helpers.AttrValueSetBool(&o.UseSsl, config["use_ssl"]))
		collect(helpers.AttrValueSetBool(&o.UseTls, config["use_tls"]))
		collect(helpers.AttrValueSetString(&o.Username, config["username"], false))
	}
	return diags, nil
}

type notificationTemplateEmailResource = framework.GenericResource[notificationTemplateEmailTerraformModel, notificationTemplateEmailBodyRequestModel, *notificationTemplateEmailTerraformModel]

func NewNotificationTemplateEmailResource() resource.Resource {
	attrs := framework.NotificationBaseResourceAttrs()
	attrs["host"] = schema.StringAttribute{
		Description: "Host.",
		Required:    true,
	}
	attrs["password"] = schema.StringAttribute{
		Description: "Password.",
		Required:    true,
		Sensitive:   true,
	}
	attrs["port"] = schema.Int64Attribute{
		Description: "Port.",
		Required:    true,
	}
	attrs["recipients"] = schema.ListAttribute{
		Description: "Recipient List.",
		ElementType: types.StringType,
		Required:    true,
	}
	attrs["sender"] = schema.StringAttribute{
		Description: "Sender Email.",
		Required:    true,
	}
	attrs["timeout"] = schema.Int64Attribute{
		Description: "Timeout. AWX defaults this to 30 when unset.",
		Optional:    true,
		Computed:    true,
		Default:     int64default.StaticInt64(30),
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
	}
	attrs["use_ssl"] = schema.BoolAttribute{
		Description: "Use SSL.",
		Required:    true,
	}
	attrs["use_tls"] = schema.BoolAttribute{
		Description: "Use TLS.",
		Required:    true,
	}
	attrs["username"] = schema.StringAttribute{
		Description: "Username.",
		Required:    true,
	}
	return &notificationTemplateEmailResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template_email", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.ResourceCfg[notificationTemplateEmailTerraformModel, notificationTemplateEmailBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages an AWX `Email` (`email`) notification template with first-class typed configuration attributes. Equivalent to `awx_notification_template` with `notification_type = \"email\"`, but with per-field schema validation and sensitivity instead of a JSON `notification_configuration` string.",
				Attributes:          attrs,
			},
			IDAccessor: func(m *notificationTemplateEmailTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:      "id",
			Hook:       hookNotificationTemplateEmail,
			WriteOnlyPlanToState: func(plan, state *notificationTemplateEmailTerraformModel) {
				if state.NotificationType.IsNull() || state.NotificationType.IsUnknown() {
					state.NotificationType = types.StringValue("email")
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "NotificationTemplateEmail",
		},
	}
}

type notificationTemplateEmailDataSource = framework.GenericDataSource[notificationTemplateEmailDataSourceTerraformModel, *notificationTemplateEmailDataSourceTerraformModel]

func NewNotificationTemplateEmailDataSource() datasource.DataSource {
	attrs := framework.NotificationBaseDataSourceAttrs()
	attrs["host"] = dschema.StringAttribute{
		Description: "Host.",
		Computed:    true,
	}
	attrs["port"] = dschema.Int64Attribute{
		Description: "Port.",
		Computed:    true,
	}
	attrs["recipients"] = dschema.ListAttribute{
		Description: "Recipient List.",
		ElementType: types.StringType,
		Computed:    true,
	}
	attrs["sender"] = dschema.StringAttribute{
		Description: "Sender Email.",
		Computed:    true,
	}
	attrs["timeout"] = dschema.Int64Attribute{
		Description: "Timeout. AWX defaults this to 30 when unset.",
		Computed:    true,
	}
	attrs["use_ssl"] = dschema.BoolAttribute{
		Description: "Use SSL.",
		Computed:    true,
	}
	attrs["use_tls"] = dschema.BoolAttribute{
		Description: "Use TLS.",
		Computed:    true,
	}
	attrs["username"] = dschema.StringAttribute{
		Description: "Username.",
		Computed:    true,
	}
	return &notificationTemplateEmailDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template_email", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.DataSourceCfg[notificationTemplateEmailDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `Email` (`email`) notification template by ID or name.",
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
			ResourceName: "NotificationTemplateEmail",
		},
	}
}
