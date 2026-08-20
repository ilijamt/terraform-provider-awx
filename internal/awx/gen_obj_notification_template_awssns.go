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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
)

type notificationTemplateAwssnsTerraformModel struct {
	ID                 types.Int64  `tfsdk:"id" json:"id"`
	Name               types.String `tfsdk:"name" json:"name"`
	Description        types.String `tfsdk:"description" json:"description"`
	Organization       types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType   types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages           types.String `tfsdk:"messages" json:"messages"`
	AwsAccessKeyId     types.String `tfsdk:"aws_access_key_id" json:"-"`
	AwsRegion          types.String `tfsdk:"aws_region" json:"-"`
	AwsSecretAccessKey types.String `tfsdk:"aws_secret_access_key" json:"-"`
	AwsSessionToken    types.String `tfsdk:"aws_session_token" json:"-"`
	SnsTopicArn        types.String `tfsdk:"sns_topic_arn" json:"-"`
}

func (o *notificationTemplateAwssnsTerraformModel) Clone() notificationTemplateAwssnsTerraformModel {
	return *o
}

type notificationTemplateAwssnsBodyRequestModel struct {
	Name                      string          `json:"name"`
	Description               string          `json:"description,omitempty"`
	Organization              int64           `json:"organization"`
	NotificationType          string          `json:"notification_type"`
	NotificationConfiguration map[string]any  `json:"notification_configuration"`
	Messages                  json.RawMessage `json:"messages,omitempty"`
}

// BodyRequest drops unset values rather than sending zero values, so AWX
// applies its own defaults.
func (o *notificationTemplateAwssnsTerraformModel) BodyRequest() *notificationTemplateAwssnsBodyRequestModel {
	req := &notificationTemplateAwssnsBodyRequestModel{
		Name:             o.Name.ValueString(),
		Description:      o.Description.ValueString(),
		Organization:     o.Organization.ValueInt64(),
		NotificationType: "awssns",
	}
	if v := o.Messages.ValueString(); v != "" {
		req.Messages = json.RawMessage(v)
	}

	config := map[string]any{}
	if !o.AwsAccessKeyId.IsNull() && !o.AwsAccessKeyId.IsUnknown() {
		config["aws_access_key_id"] = o.AwsAccessKeyId.ValueString()
	}
	if !o.AwsRegion.IsNull() && !o.AwsRegion.IsUnknown() {
		config["aws_region"] = o.AwsRegion.ValueString()
	}
	if !o.AwsSecretAccessKey.IsNull() && !o.AwsSecretAccessKey.IsUnknown() {
		config["aws_secret_access_key"] = o.AwsSecretAccessKey.ValueString()
	}
	if !o.AwsSessionToken.IsNull() && !o.AwsSessionToken.IsUnknown() {
		config["aws_session_token"] = o.AwsSessionToken.ValueString()
	}
	if !o.SnsTopicArn.IsNull() && !o.SnsTopicArn.IsUnknown() {
		config["sns_topic_arn"] = o.SnsTopicArn.ValueString()
	}
	req.NotificationConfiguration = config
	return req
}

func (o *notificationTemplateAwssnsTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.AwsAccessKeyId, config["aws_access_key_id"], false))
		collect(helpers.AttrValueSetString(&o.AwsRegion, config["aws_region"], false))
		collect(helpers.AttrValueSetString(&o.AwsSecretAccessKey, config["aws_secret_access_key"], false))
		collect(helpers.AttrValueSetString(&o.AwsSessionToken, config["aws_session_token"], false))
		collect(helpers.AttrValueSetString(&o.SnsTopicArn, config["sns_topic_arn"], false))
	}
	return diags, nil
}

// AWX answers every secret field with `$encrypted$`, which would otherwise read
// as drift on every plan.
func hookNotificationTemplateAwssns(_ context.Context, _ string, source hooks.Source, callee hooks.Callee, orig, state *notificationTemplateAwssnsTerraformModel) error {
	if source != hooks.SourceResource {
		return nil
	}

	if callee == hooks.CalleeCreate {
		// AWX never returns this value, so state has to carry a known null.
		if orig.AwsSecretAccessKey.IsNull() || orig.AwsSecretAccessKey.IsUnknown() {
			state.AwsSecretAccessKey = types.StringNull()
		} else {
			state.AwsSecretAccessKey = orig.AwsSecretAccessKey
		}
		if orig.AwsSessionToken.IsNull() || orig.AwsSessionToken.IsUnknown() {
			state.AwsSessionToken = types.StringNull()
		} else {
			state.AwsSessionToken = orig.AwsSessionToken
		}
		return nil
	}

	if callee == hooks.CalleeRead || callee == hooks.CalleeUpdate {
		if v, subbed := helpers.MergeEncryptedField(orig.AwsSecretAccessKey, state.AwsSecretAccessKey); subbed {
			state.AwsSecretAccessKey = v
		}
		if v, subbed := helpers.MergeEncryptedField(orig.AwsSessionToken, state.AwsSessionToken); subbed {
			state.AwsSessionToken = v
		}
	}
	return nil
}

type notificationTemplateAwssnsDataSourceTerraformModel struct {
	ID               types.Int64  `tfsdk:"id" json:"id"`
	Name             types.String `tfsdk:"name" json:"name"`
	Description      types.String `tfsdk:"description" json:"description"`
	Organization     types.Int64  `tfsdk:"organization" json:"organization"`
	NotificationType types.String `tfsdk:"notification_type" json:"notification_type"`
	Messages         types.String `tfsdk:"messages" json:"messages"`
	AwsAccessKeyId   types.String `tfsdk:"aws_access_key_id" json:"-"`
	AwsRegion        types.String `tfsdk:"aws_region" json:"-"`
	SnsTopicArn      types.String `tfsdk:"sns_topic_arn" json:"-"`
}

func (o *notificationTemplateAwssnsDataSourceTerraformModel) Clone() notificationTemplateAwssnsDataSourceTerraformModel {
	return *o
}

func (o *notificationTemplateAwssnsDataSourceTerraformModel) UpdateFromApiData(data map[string]any) (diag.Diagnostics, error) {
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
		collect(helpers.AttrValueSetString(&o.AwsAccessKeyId, config["aws_access_key_id"], false))
		collect(helpers.AttrValueSetString(&o.AwsRegion, config["aws_region"], false))
		collect(helpers.AttrValueSetString(&o.SnsTopicArn, config["sns_topic_arn"], false))
	}
	return diags, nil
}

type notificationTemplateAwssnsResource = framework.GenericResource[notificationTemplateAwssnsTerraformModel, notificationTemplateAwssnsBodyRequestModel, *notificationTemplateAwssnsTerraformModel]

func NewNotificationTemplateAwssnsResource() resource.Resource {
	attrs := framework.NotificationBaseResourceAttrs()
	attrs["aws_access_key_id"] = schema.StringAttribute{
		Description: "Access Key ID.",
		Optional:    true,
		Computed:    true,
		Default:     stringdefault.StaticString(""),
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["aws_region"] = schema.StringAttribute{
		Description: "AWS Region.",
		Optional:    true,
		Computed:    true,
		Default:     stringdefault.StaticString(""),
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	attrs["aws_secret_access_key"] = schema.StringAttribute{
		Description: "Secret Access Key.",
		Optional:    true,
		Computed:    true,
		Default:     stringdefault.StaticString(""),
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Sensitive: true,
	}
	attrs["aws_session_token"] = schema.StringAttribute{
		Description: "Session Token.",
		Optional:    true,
		Computed:    true,
		Default:     stringdefault.StaticString(""),
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		Sensitive: true,
	}
	attrs["sns_topic_arn"] = schema.StringAttribute{
		Description: "SNS Topic ARN.",
		Optional:    true,
		Computed:    true,
		Default:     stringdefault.StaticString(""),
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
	return &notificationTemplateAwssnsResource{
		ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template_awssns", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.ResourceCfg[notificationTemplateAwssnsTerraformModel, notificationTemplateAwssnsBodyRequestModel]{
			Schema: schema.Schema{
				MarkdownDescription: "Manages an AWX `AWS SNS` (`awssns`) notification template with first-class typed configuration attributes. Equivalent to `awx_notification_template` with `notification_type = \"awssns\"`, but with per-field schema validation and sensitivity instead of a JSON `notification_configuration` string.",
				Attributes:          attrs,
			},
			IDAccessor: func(m *notificationTemplateAwssnsTerraformModel) any { return m.ID.ValueInt64() },
			IDKey:      "id",
			Hook:       hookNotificationTemplateAwssns,
			WriteOnlyPlanToState: func(plan, state *notificationTemplateAwssnsTerraformModel) {
				if state.NotificationType.IsNull() || state.NotificationType.IsUnknown() {
					state.NotificationType = types.StringValue("awssns")
				}
			},
			ApiVersion:   ApiVersion,
			ResourceName: "NotificationTemplateAwssns",
		},
	}
}

type notificationTemplateAwssnsDataSource = framework.GenericDataSource[notificationTemplateAwssnsDataSourceTerraformModel, *notificationTemplateAwssnsDataSourceTerraformModel]

func NewNotificationTemplateAwssnsDataSource() datasource.DataSource {
	attrs := framework.NotificationBaseDataSourceAttrs()
	attrs["aws_access_key_id"] = dschema.StringAttribute{
		Description: "Access Key ID.",
		Computed:    true,
	}
	attrs["aws_region"] = dschema.StringAttribute{
		Description: "AWS Region.",
		Computed:    true,
	}
	attrs["sns_topic_arn"] = dschema.StringAttribute{
		Description: "SNS Topic ARN.",
		Computed:    true,
	}
	return &notificationTemplateAwssnsDataSource{
		DataSourceBase: framework.DataSourceBase{ProviderBase: framework.ProviderBase{TypeName: "notification_template_awssns", Endpoint: "/api/v2/notification_templates/"}},
		Cfg: framework.DataSourceCfg[notificationTemplateAwssnsDataSourceTerraformModel]{
			Schema: dschema.Schema{
				MarkdownDescription: "Reads an AWX `AWS SNS` (`awssns`) notification template by ID or name.",
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
			ResourceName: "NotificationTemplateAwssns",
		},
	}
}
