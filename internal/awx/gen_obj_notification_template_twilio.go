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

type notificationTemplateTwilioTerraformModel struct {
	ID               types.Int64  `tfsdk:"id" json:"id"`
	Name             types.String `tfsdk:"name" json:"name"`
	Description      types.String `tfsdk:"description" json:"description"`
	Organization     types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages         types.String `tfsdk:"messages" json:"messages"`
	AccountSid       types.String `tfsdk:"account_sid" json:"-"`
	AccountToken     types.String `tfsdk:"account_token" json:"-"`
	FromNumber       types.String `tfsdk:"from_number" json:"-"`
	ToNumbers        types.List   `tfsdk:"to_numbers" json:"-"`
}

func (o *notificationTemplateTwilioTerraformModel) Clone() notificationTemplateTwilioTerraformModel {
	return *o
}

type notificationTemplateTwilioBodyRequestModel struct {
	Name                      string          `json:"name"`
	Description               string          `json:"description,omitempty"`
	Organization              int64           `json:"organization"`
	NotificationType          string          `json:"notification_type"`
	NotificationConfiguration map[string]any  `json:"notification_configuration"`
	Messages                  json.RawMessage `json:"messages,omitempty"`
}

// BodyRequest drops unset values rather than sending zero values, so AWX
// applies its own defaults.
func (o *notificationTemplateTwilioTerraformModel) BodyRequest() *notificationTemplateTwilioBodyRequestModel {
	req := &notificationTemplateTwilioBodyRequestModel{
		Name:             o.Name.ValueString(),
		Description:      o.Description.ValueString(),
		Organization:     o.Organization.ValueInt64(),
		NotificationType: "twilio",
	}
	if v := o.Messages.ValueString(); v != "" {
		req.Messages = json.RawMessage(v)
	}

	config := map[string]any{}
	if !o.AccountSid.IsNull() && !o.AccountSid.IsUnknown() {
		config["account_sid"] = o.AccountSid.ValueString()
	}
	if !o.AccountToken.IsNull() && !o.AccountToken.IsUnknown() {
		config["account_token"] = o.AccountToken.ValueString()
	}
	if !o.FromNumber.IsNull() && !o.FromNumber.IsUnknown() {
		config["from_number"] = o.FromNumber.ValueString()
	}
	if !o.ToNumbers.IsNull() && !o.ToNumbers.IsUnknown() {
		config["to_numbers"] = helpers.ListAsStringSlice(o.ToNumbers, false)
	}
	req.NotificationConfiguration = config
	return req
}

func (o *notificationTemplateTwilioTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.AccountSid, config["account_sid"], false))
		collect(helpers.AttrValueSetString(&o.AccountToken, config["account_token"], false))
		collect(helpers.AttrValueSetString(&o.FromNumber, config["from_number"], false))
		collect(helpers.AttrValueSetListString(&o.ToNumbers, config["to_numbers"], false))
	}
	return diags, nil
}

// AWX answers every secret field with `$encrypted$`, which would otherwise read
// as drift on every plan.
func hookNotificationTemplateTwilio(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *notificationTemplateTwilioTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// AWX never returns this value, so state has to carry a known null.
		if orig.AccountToken.IsNull() || orig.AccountToken.IsUnknown() {
			state.AccountToken = types.StringNull()
		} else {
			state.AccountToken = orig.AccountToken
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.AccountToken, state.AccountToken); subbed {
			state.AccountToken = v
		}
	}
	return nil
}

type notificationTemplateTwilioDataSourceTerraformModel struct {
	ID               types.Int64  `tfsdk:"id" json:"id"`
	Name             types.String `tfsdk:"name" json:"name"`
	Description      types.String `tfsdk:"description" json:"description"`
	Organization     types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages         types.String `tfsdk:"messages" json:"messages"`
	AccountSid       types.String `tfsdk:"account_sid" json:"-"`
	FromNumber       types.String `tfsdk:"from_number" json:"-"`
	ToNumbers        types.List   `tfsdk:"to_numbers" json:"-"`
}

func (o *notificationTemplateTwilioDataSourceTerraformModel) Clone() notificationTemplateTwilioDataSourceTerraformModel {
	return *o
}

func (o *notificationTemplateTwilioDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.AccountSid, config["account_sid"], false))
		collect(helpers.AttrValueSetString(&o.FromNumber, config["from_number"], false))
		collect(helpers.AttrValueSetListString(&o.ToNumbers, config["to_numbers"], false))
	}
	return diags, nil
}

type notificationTemplateTwilioResource = framework.GenericResource[notificationTemplateTwilioTerraformModel, notificationTemplateTwilioBodyRequestModel, *notificationTemplateTwilioTerraformModel]

func NewNotificationTemplateTwilioResource() resource.Resource {
	attrs := framework.NotificationBaseResourceAttrs()
	attrs["account_sid"] = schema.StringAttribute{
		Description: "Account SID.",
		Required:    true,
	}
	attrs["account_token"] = schema.StringAttribute{
		Description: "Account Token.",
		Required:    true,
		Sensitive:   true,
	}
	attrs["from_number"] = schema.StringAttribute{
		Description: "Source Phone Number.",
		Required:    true,
	}
	attrs["to_numbers"] = schema.ListAttribute{
		Description: "Destination SMS Numbers.",
		ElementType: types.StringType,
		Required:    true,
	}
	return &notificationTemplateTwilioResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template_twilio", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.ResourceCfg[notificationTemplateTwilioTerraformModel, notificationTemplateTwilioBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages an AWX `Twilio` (`twilio`) notification template with first-class typed configuration attributes. Equivalent to `awx_notification_template` with `notification_type = \"twilio\"`, but with per-field schema validation and sensitivity instead of a JSON `notification_configuration` string.",
				Attributes:          attrs,
			},
			IDAccessor: func(m *notificationTemplateTwilioTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:      "id",
			Hook:       hookNotificationTemplateTwilio,
			WriteOnlyPlanToState: func(plan, state *notificationTemplateTwilioTerraformModel) {
				if state.NotificationType.IsNull() || state.NotificationType.IsUnknown() {
					state.NotificationType = types.StringValue("twilio")
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "NotificationTemplateTwilio",
		},
	}
}

type notificationTemplateTwilioDataSource = framework.GenericDataSource[notificationTemplateTwilioDataSourceTerraformModel, *notificationTemplateTwilioDataSourceTerraformModel]

func NewNotificationTemplateTwilioDataSource() datasource.DataSource {
	attrs := framework.NotificationBaseDataSourceAttrs()
	attrs["account_sid"] = dschema.StringAttribute{
		Description: "Account SID.",
		Computed:    true,
	}
	attrs["from_number"] = dschema.StringAttribute{
		Description: "Source Phone Number.",
		Computed:    true,
	}
	attrs["to_numbers"] = dschema.ListAttribute{
		Description: "Destination SMS Numbers.",
		ElementType: types.StringType,
		Computed:    true,
	}
	return &notificationTemplateTwilioDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template_twilio", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.DataSourceCfg[notificationTemplateTwilioDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `Twilio` (`twilio`) notification template by ID or name.",
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
			ResourceName: "NotificationTemplateTwilio",
		},
	}
}
