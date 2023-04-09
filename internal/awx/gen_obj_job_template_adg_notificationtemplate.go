package awx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	p "path"

	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	c "github.com/ilijamt/terraform-provider-awx/internal/client"
)

var (
	_ resource.Resource              = &jobTemplateAssociateDisassociateNotificationTemplate{}
	_ resource.ResourceWithConfigure = &jobTemplateAssociateDisassociateNotificationTemplate{}
)

type jobTemplateAssociateDisassociateNotificationTemplateTerraformModel struct {
	JobTemplateID          types.Int64  `tfsdk:"job_template_id"`
	NotificationTemplateID types.Int64  `tfsdk:"notification_template_id"`
	Option                 types.String `tfsdk:"option"`
}

// NewJobTemplateAssociateDisassociateNotificationTemplateResource is a helper function to simplify the provider implementation.
func NewJobTemplateAssociateDisassociateNotificationTemplateResource() resource.Resource {
	return &jobTemplateAssociateDisassociateNotificationTemplate{}
}

type jobTemplateAssociateDisassociateNotificationTemplate struct {
	client   c.Client
	endpoint string
}

func (o *jobTemplateAssociateDisassociateNotificationTemplate) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/job_templates/%d/notification_templates_%s/"
}

func (o *jobTemplateAssociateDisassociateNotificationTemplate) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_job_template_associate_notification_template"
}

func (o *jobTemplateAssociateDisassociateNotificationTemplate) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"job_template_id": schema.Int64Attribute{
				Description: "Database ID for this JobTemplate.",
				Required:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"notification_template_id": schema.Int64Attribute{
				Description: "Database ID of the notificationtemplate to assign.",
				Required:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"option": schema.StringAttribute{
				Description: "Notification Option",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"started", "success", "error"}...),
				},
			},
		},
	}
}

func (o *jobTemplateAssociateDisassociateNotificationTemplate) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.RequiredTogether(
			path.MatchRoot("job_template_id"),
			path.MatchRoot("notification_template_id"),
			path.MatchRoot("option"),
		),
	}
}

func (o *jobTemplateAssociateDisassociateNotificationTemplate) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error

	// Retrieve values from state
	var plan, state jobTemplateAssociateDisassociateNotificationTemplateTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for association of JobTemplate
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.JobTemplateID.ValueInt64(), plan.Option.ValueString())) + "/"
	var buf bytes.Buffer
	var bodyRequest = associateDisassociateRequestModel{ID: plan.NotificationTemplateID.ValueInt64(), Disassociate: false}
	tflog.Debug(ctx, "[JobTemplate/Create/Associate] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for create of type notification_job_template", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to associate for JobTemplate on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	state.JobTemplateID = plan.JobTemplateID
	state.NotificationTemplateID = plan.NotificationTemplateID
	state.Option = plan.Option

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *jobTemplateAssociateDisassociateNotificationTemplate) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state jobTemplateAssociateDisassociateNotificationTemplateTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for disassociation of JobTemplate
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.JobTemplateID.ValueInt64(), state.Option.ValueString())) + "/"
	var buf bytes.Buffer
	var bodyRequest = associateDisassociateRequestModel{ID: state.NotificationTemplateID.ValueInt64(), Disassociate: true}
	tflog.Debug(ctx, "[JobTemplate/Delete/Disassociate] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for delete of type notification_job_template", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to disassociate for JobTemplate on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}

func (o *jobTemplateAssociateDisassociateNotificationTemplate) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
}

func (o *jobTemplateAssociateDisassociateNotificationTemplate) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
}
