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
	"github.com/ilijamt/terraform-provider-awx/internal/models"
)

var (
	_ resource.Resource              = &workflowJobTemplateAssociateDisassociateNotificationTemplate{}
	_ resource.ResourceWithConfigure = &workflowJobTemplateAssociateDisassociateNotificationTemplate{}
)

type workflowJobTemplateAssociateDisassociateNotificationTemplateTerraformModel struct {
	WorkflowJobTemplateID  types.Int64  `tfsdk:"workflow_job_template_id"`
	NotificationTemplateID types.Int64  `tfsdk:"notification_template_id"`
	Option                 types.String `tfsdk:"option"`
}

// NewWorkflowJobTemplateAssociateDisassociateNotificationTemplateResource is a helper function to simplify the provider implementation.
func NewWorkflowJobTemplateAssociateDisassociateNotificationTemplateResource() resource.Resource {
	return &workflowJobTemplateAssociateDisassociateNotificationTemplate{}
}

type workflowJobTemplateAssociateDisassociateNotificationTemplate struct {
	client   c.Client
	endpoint string
}

func (o *workflowJobTemplateAssociateDisassociateNotificationTemplate) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/workflow_job_templates/%d/notification_templates_%s/"
}

func (o *workflowJobTemplateAssociateDisassociateNotificationTemplate) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_workflow_job_template_associate_notification_template"
}

func (o *workflowJobTemplateAssociateDisassociateNotificationTemplate) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"workflow_job_template_id": schema.Int64Attribute{
				Description: "Database ID for this WorkflowJobTemplate.",
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
					stringvalidator.OneOf([]string{"approval", "started", "success", "error"}...),
				},
			},
		},
	}
}

func (o *workflowJobTemplateAssociateDisassociateNotificationTemplate) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.RequiredTogether(
			path.MatchRoot("workflow_job_template_id"),
			path.MatchRoot("notification_template_id"),
			path.MatchRoot("option"),
		),
	}
}

func (o *workflowJobTemplateAssociateDisassociateNotificationTemplate) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error

	// Retrieve values from state
	var plan, state workflowJobTemplateAssociateDisassociateNotificationTemplateTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for association of WorkflowJobTemplate
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.WorkflowJobTemplateID.ValueInt64(), plan.Option.ValueString())) + "/"
	var buf bytes.Buffer
	var bodyRequest = models.AssociateDisassociateRequestModel{ID: plan.NotificationTemplateID.ValueInt64(), Disassociate: false}
	tflog.Debug(ctx, "[WorkflowJobTemplate/Create/Associate] Making a request", map[string]any{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for WorkflowJobTemplate on %s for create of type notification_job_workflow_template", endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to associate for WorkflowJobTemplate on %s with a payload of %#v", endpoint, bodyRequest),
			err.Error(),
		)
		return
	}

	state.WorkflowJobTemplateID = plan.WorkflowJobTemplateID
	state.NotificationTemplateID = plan.NotificationTemplateID
	state.Option = plan.Option

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *workflowJobTemplateAssociateDisassociateNotificationTemplate) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state workflowJobTemplateAssociateDisassociateNotificationTemplateTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for disassociation of WorkflowJobTemplate
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.WorkflowJobTemplateID.ValueInt64(), state.Option.ValueString())) + "/"
	var buf bytes.Buffer
	var bodyRequest = models.AssociateDisassociateRequestModel{ID: state.NotificationTemplateID.ValueInt64(), Disassociate: true}
	tflog.Debug(ctx, "[WorkflowJobTemplate/Delete/Disassociate] Making a request", map[string]any{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for WorkflowJobTemplate on %s for delete of type notification_job_workflow_template", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to disassociate for WorkflowJobTemplate on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}

func (o *workflowJobTemplateAssociateDisassociateNotificationTemplate) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
}

func (o *workflowJobTemplateAssociateDisassociateNotificationTemplate) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
}
