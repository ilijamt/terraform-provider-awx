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

type notificationTemplatePagerdutyTerraformModel struct {
	ID               types.Int64  `tfsdk:"id" json:"id"`
	Name             types.String `tfsdk:"name" json:"name"`
	Description      types.String `tfsdk:"description" json:"description"`
	Organization     types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages         types.String `tfsdk:"messages" json:"messages"`
	ClientName       types.String `tfsdk:"client_name" json:"-"`
	ServiceKey       types.String `tfsdk:"service_key" json:"-"`
	Subdomain        types.String `tfsdk:"subdomain" json:"-"`
	Token            types.String `tfsdk:"token" json:"-"`
}

func (o *notificationTemplatePagerdutyTerraformModel) Clone() notificationTemplatePagerdutyTerraformModel {
	return *o
}

type notificationTemplatePagerdutyBodyRequestModel struct {
	Name                      string          `json:"name"`
	Description               string          `json:"description,omitempty"`
	Organization              int64           `json:"organization"`
	NotificationType          string          `json:"notification_type"`
	NotificationConfiguration map[string]any  `json:"notification_configuration"`
	Messages                  json.RawMessage `json:"messages,omitempty"`
}

// BodyRequest drops unset values rather than sending zero values, so AWX
// applies its own defaults.
func (o *notificationTemplatePagerdutyTerraformModel) BodyRequest() *notificationTemplatePagerdutyBodyRequestModel {
	req := &notificationTemplatePagerdutyBodyRequestModel{
		Name:             o.Name.ValueString(),
		Description:      o.Description.ValueString(),
		Organization:     o.Organization.ValueInt64(),
		NotificationType: "pagerduty",
	}
	if v := o.Messages.ValueString(); v != "" {
		req.Messages = json.RawMessage(v)
	}

	config := map[string]any{}
	if !o.ClientName.IsNull() && !o.ClientName.IsUnknown() {
		config["client_name"] = o.ClientName.ValueString()
	}
	if !o.ServiceKey.IsNull() && !o.ServiceKey.IsUnknown() {
		config["service_key"] = o.ServiceKey.ValueString()
	}
	if !o.Subdomain.IsNull() && !o.Subdomain.IsUnknown() {
		config["subdomain"] = o.Subdomain.ValueString()
	}
	if !o.Token.IsNull() && !o.Token.IsUnknown() {
		config["token"] = o.Token.ValueString()
	}
	req.NotificationConfiguration = config
	return req
}

func (o *notificationTemplatePagerdutyTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.ClientName, config["client_name"], false))
		collect(helpers.AttrValueSetString(&o.ServiceKey, config["service_key"], false))
		collect(helpers.AttrValueSetString(&o.Subdomain, config["subdomain"], false))
		collect(helpers.AttrValueSetString(&o.Token, config["token"], false))
	}
	return diags, nil
}

// AWX answers every secret field with `$encrypted$`, which would otherwise read
// as drift on every plan.
func hookNotificationTemplatePagerduty(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *notificationTemplatePagerdutyTerraformModel) error {
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

type notificationTemplatePagerdutyDataSourceTerraformModel struct {
	ID               types.Int64  `tfsdk:"id" json:"id"`
	Name             types.String `tfsdk:"name" json:"name"`
	Description      types.String `tfsdk:"description" json:"description"`
	Organization     types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages         types.String `tfsdk:"messages" json:"messages"`
	ClientName       types.String `tfsdk:"client_name" json:"-"`
	ServiceKey       types.String `tfsdk:"service_key" json:"-"`
	Subdomain        types.String `tfsdk:"subdomain" json:"-"`
}

func (o *notificationTemplatePagerdutyDataSourceTerraformModel) Clone() notificationTemplatePagerdutyDataSourceTerraformModel {
	return *o
}

func (o *notificationTemplatePagerdutyDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.ClientName, config["client_name"], false))
		collect(helpers.AttrValueSetString(&o.ServiceKey, config["service_key"], false))
		collect(helpers.AttrValueSetString(&o.Subdomain, config["subdomain"], false))
	}
	return diags, nil
}

type notificationTemplatePagerdutyResource = framework.GenericResource[notificationTemplatePagerdutyTerraformModel, notificationTemplatePagerdutyBodyRequestModel, *notificationTemplatePagerdutyTerraformModel]

func NewNotificationTemplatePagerdutyResource() resource.Resource {
	attrs := framework.NotificationBaseResourceAttrs()
	attrs["client_name"] = schema.StringAttribute{
		Description: "Client Identifier.",
		Required:    true,
	}
	attrs["service_key"] = schema.StringAttribute{
		Description: "API Service/Integration Key.",
		Required:    true,
	}
	attrs["subdomain"] = schema.StringAttribute{
		Description: "Pagerduty subdomain.",
		Required:    true,
	}
	attrs["token"] = schema.StringAttribute{
		Description: "API Token.",
		Required:    true,
		Sensitive:   true,
	}
	return &notificationTemplatePagerdutyResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template_pagerduty", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.ResourceCfg[notificationTemplatePagerdutyTerraformModel, notificationTemplatePagerdutyBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages an AWX `Pagerduty` (`pagerduty`) notification template with first-class typed configuration attributes. Equivalent to `awx_notification_template` with `notification_type = \"pagerduty\"`, but with per-field schema validation and sensitivity instead of a JSON `notification_configuration` string.",
				Attributes:          attrs,
			},
			IDAccessor: func(m *notificationTemplatePagerdutyTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:      "id",
			Hook:       hookNotificationTemplatePagerduty,
			WriteOnlyPlanToState: func(plan, state *notificationTemplatePagerdutyTerraformModel) {
				if state.NotificationType.IsNull() || state.NotificationType.IsUnknown() {
					state.NotificationType = types.StringValue("pagerduty")
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "NotificationTemplatePagerduty",
		},
	}
}

type notificationTemplatePagerdutyDataSource = framework.GenericDataSource[notificationTemplatePagerdutyDataSourceTerraformModel, *notificationTemplatePagerdutyDataSourceTerraformModel]

func NewNotificationTemplatePagerdutyDataSource() datasource.DataSource {
	attrs := framework.NotificationBaseDataSourceAttrs()
	attrs["client_name"] = dschema.StringAttribute{
		Description: "Client Identifier.",
		Computed:    true,
	}
	attrs["service_key"] = dschema.StringAttribute{
		Description: "API Service/Integration Key.",
		Computed:    true,
	}
	attrs["subdomain"] = dschema.StringAttribute{
		Description: "Pagerduty subdomain.",
		Computed:    true,
	}
	return &notificationTemplatePagerdutyDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template_pagerduty", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.DataSourceCfg[notificationTemplatePagerdutyDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `Pagerduty` (`pagerduty`) notification template by ID or name.",
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
			ResourceName: "NotificationTemplatePagerduty",
		},
	}
}
