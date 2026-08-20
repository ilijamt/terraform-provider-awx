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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
)

type notificationTemplateWebhookTerraformModel struct {
	ID                     types.Int64  `tfsdk:"id" json:"id"`
	Name                   types.String `tfsdk:"name" json:"name"`
	Description            types.String `tfsdk:"description" json:"description"`
	Organization           types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType       types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages               types.String `tfsdk:"messages" json:"messages"`
	DisableSslVerification types.Bool   `tfsdk:"disable_ssl_verification" json:"-"`
	Headers                types.Map    `tfsdk:"headers" json:"-"`
	HttpMethod             types.String `tfsdk:"http_method" json:"-"`
	Password               types.String `tfsdk:"password" json:"-"`
	Url                    types.String `tfsdk:"url" json:"-"`
	Username               types.String `tfsdk:"username" json:"-"`
}

func (o *notificationTemplateWebhookTerraformModel) Clone() notificationTemplateWebhookTerraformModel {
	return *o
}

type notificationTemplateWebhookBodyRequestModel struct {
	Name                      string          `json:"name"`
	Description               string          `json:"description,omitempty"`
	Organization              int64           `json:"organization"`
	NotificationType          string          `json:"notification_type"`
	NotificationConfiguration map[string]any  `json:"notification_configuration"`
	Messages                  json.RawMessage `json:"messages,omitempty"`
}

// BodyRequest drops unset values rather than sending zero values, so AWX
// applies its own defaults.
func (o *notificationTemplateWebhookTerraformModel) BodyRequest() *notificationTemplateWebhookBodyRequestModel {
	req := &notificationTemplateWebhookBodyRequestModel{
		Name:             o.Name.ValueString(),
		Description:      o.Description.ValueString(),
		Organization:     o.Organization.ValueInt64(),
		NotificationType: "webhook",
	}
	if v := o.Messages.ValueString(); v != "" {
		req.Messages = json.RawMessage(v)
	}

	config := map[string]any{}
	if !o.DisableSslVerification.IsNull() && !o.DisableSslVerification.IsUnknown() {
		config["disable_ssl_verification"] = o.DisableSslVerification.ValueBool()
	}
	// AWX requires the key even when the map is empty.
	config["headers"] = helpers.MapAsStringMap(o.Headers, false)
	if !o.HttpMethod.IsNull() && !o.HttpMethod.IsUnknown() {
		config["http_method"] = o.HttpMethod.ValueString()
	}
	if !o.Password.IsNull() && !o.Password.IsUnknown() {
		config["password"] = o.Password.ValueString()
	}
	if !o.Url.IsNull() && !o.Url.IsUnknown() {
		config["url"] = o.Url.ValueString()
	}
	if !o.Username.IsNull() && !o.Username.IsUnknown() {
		config["username"] = o.Username.ValueString()
	}
	req.NotificationConfiguration = config
	return req
}

func (o *notificationTemplateWebhookTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetBool(&o.DisableSslVerification, config["disable_ssl_verification"]))
		collect(helpers.AttrValueSetMapString(&o.Headers, config["headers"], false))
		collect(helpers.AttrValueSetString(&o.HttpMethod, config["http_method"], false))
		collect(helpers.AttrValueSetString(&o.Password, config["password"], false))
		collect(helpers.AttrValueSetString(&o.Url, config["url"], false))
		collect(helpers.AttrValueSetString(&o.Username, config["username"], false))
	}
	return diags, nil
}

// AWX answers every secret field with `$encrypted$`, which would otherwise read
// as drift on every plan.
func hookNotificationTemplateWebhook(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *notificationTemplateWebhookTerraformModel) error {
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

type notificationTemplateWebhookDataSourceTerraformModel struct {
	ID                     types.Int64  `tfsdk:"id" json:"id"`
	Name                   types.String `tfsdk:"name" json:"name"`
	Description            types.String `tfsdk:"description" json:"description"`
	Organization           types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType       types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages               types.String `tfsdk:"messages" json:"messages"`
	DisableSslVerification types.Bool   `tfsdk:"disable_ssl_verification" json:"-"`
	Headers                types.Map    `tfsdk:"headers" json:"-"`
	HttpMethod             types.String `tfsdk:"http_method" json:"-"`
	Url                    types.String `tfsdk:"url" json:"-"`
	Username               types.String `tfsdk:"username" json:"-"`
}

func (o *notificationTemplateWebhookDataSourceTerraformModel) Clone() notificationTemplateWebhookDataSourceTerraformModel {
	return *o
}

func (o *notificationTemplateWebhookDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetBool(&o.DisableSslVerification, config["disable_ssl_verification"]))
		collect(helpers.AttrValueSetMapString(&o.Headers, config["headers"], false))
		collect(helpers.AttrValueSetString(&o.HttpMethod, config["http_method"], false))
		collect(helpers.AttrValueSetString(&o.Url, config["url"], false))
		collect(helpers.AttrValueSetString(&o.Username, config["username"], false))
	}
	return diags, nil
}

type notificationTemplateWebhookResource = framework.GenericResource[notificationTemplateWebhookTerraformModel, notificationTemplateWebhookBodyRequestModel, *notificationTemplateWebhookTerraformModel]

func NewNotificationTemplateWebhookResource() resource.Resource {
	attrs := framework.NotificationBaseResourceAttrs()
	attrs["disable_ssl_verification"] = schema.BoolAttribute{
		Description: "Verify SSL. AWX defaults this to \"false\" when unset.",
		Optional:    true,
		Computed:    true,
		Default:     booldefault.StaticBool(false),
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["headers"] = schema.MapAttribute{
		Description: "HTTP Headers. AWX requires this key; the provider sends an empty map when unset.",
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.Map{
			mapplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["http_method"] = schema.StringAttribute{
		Description: "HTTP Method. AWX defaults this to \"POST\" when unset.",
		Optional:    true,
		Computed:    true,
		Default:     stringdefault.StaticString("POST"),
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["password"] = schema.StringAttribute{
		Description: "Password.",
		Optional:    true,
		Computed:    true,
		Default:     stringdefault.StaticString(""),
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Sensitive: true,
	}
	attrs["url"] = schema.StringAttribute{
		Description: "Target URL.",
		Required:    true,
	}
	attrs["username"] = schema.StringAttribute{
		Description: "Username.",
		Optional:    true,
		Computed:    true,
		Default:     stringdefault.StaticString(""),
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	return &notificationTemplateWebhookResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template_webhook", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.ResourceCfg[notificationTemplateWebhookTerraformModel, notificationTemplateWebhookBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages an AWX `Webhook` (`webhook`) notification template with first-class typed configuration attributes. Equivalent to `awx_notification_template` with `notification_type = \"webhook\"`, but with per-field schema validation and sensitivity instead of a JSON `notification_configuration` string.",
				Attributes:          attrs,
			},
			IDAccessor: func(m *notificationTemplateWebhookTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:      "id",
			Hook:       hookNotificationTemplateWebhook,
			WriteOnlyPlanToState: func(plan, state *notificationTemplateWebhookTerraformModel) {
				if state.NotificationType.IsNull() || state.NotificationType.IsUnknown() {
					state.NotificationType = types.StringValue("webhook")
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "NotificationTemplateWebhook",
		},
	}
}

type notificationTemplateWebhookDataSource = framework.GenericDataSource[notificationTemplateWebhookDataSourceTerraformModel, *notificationTemplateWebhookDataSourceTerraformModel]

func NewNotificationTemplateWebhookDataSource() datasource.DataSource {
	attrs := framework.NotificationBaseDataSourceAttrs()
	attrs["disable_ssl_verification"] = dschema.BoolAttribute{
		Description: "Verify SSL. AWX defaults this to \"false\" when unset.",
		Computed:    true,
	}
	attrs["headers"] = dschema.MapAttribute{
		Description: "HTTP Headers. AWX requires this key; the provider sends an empty map when unset.",
		ElementType: types.StringType,
		Computed:    true,
	}
	attrs["http_method"] = dschema.StringAttribute{
		Description: "HTTP Method. AWX defaults this to \"POST\" when unset.",
		Computed:    true,
	}
	attrs["url"] = dschema.StringAttribute{
		Description: "Target URL.",
		Computed:    true,
	}
	attrs["username"] = dschema.StringAttribute{
		Description: "Username.",
		Computed:    true,
	}
	return &notificationTemplateWebhookDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template_webhook", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.DataSourceCfg[notificationTemplateWebhookDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `Webhook` (`webhook`) notification template by ID or name.",
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
			ResourceName: "NotificationTemplateWebhook",
		},
	}
}
