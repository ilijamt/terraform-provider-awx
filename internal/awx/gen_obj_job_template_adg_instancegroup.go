package awx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	p "path"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/internal/models"
)

var (
	_ resource.Resource                = &jobTemplateAssociateDisassociateInstanceGroup{}
	_ resource.ResourceWithConfigure   = &jobTemplateAssociateDisassociateInstanceGroup{}
	_ resource.ResourceWithImportState = &jobTemplateAssociateDisassociateInstanceGroup{}
)

type jobTemplateAssociateDisassociateInstanceGroupTerraformModel struct {
	JobTemplateID   types.Int64 `tfsdk:"job_template_id"`
	InstanceGroupID types.Int64 `tfsdk:"instance_group_id"`
}

// NewJobTemplateAssociateDisassociateInstanceGroupResource is a helper function to simplify the provider implementation.
func NewJobTemplateAssociateDisassociateInstanceGroupResource() resource.Resource {
	return &jobTemplateAssociateDisassociateInstanceGroup{}
}

type jobTemplateAssociateDisassociateInstanceGroup struct {
	client   c.Client
	endpoint string
}

func (o *jobTemplateAssociateDisassociateInstanceGroup) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/job_templates/%d/instance_groups/"
}

func (o *jobTemplateAssociateDisassociateInstanceGroup) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_job_template_associate_instance_group"
}

func (o *jobTemplateAssociateDisassociateInstanceGroup) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"job_template_id": schema.Int64Attribute{
				Description: "Database ID for this JobTemplate.",
				Required:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"instance_group_id": schema.Int64Attribute{
				Description: "Database ID of the instancegroup to assign.",
				Required:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (o *jobTemplateAssociateDisassociateInstanceGroup) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.RequiredTogether(
			path.MatchRoot("job_template_id"),
			path.MatchRoot("instance_group_id"),
		),
	}
}

func (o *jobTemplateAssociateDisassociateInstanceGroup) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var state jobTemplateAssociateDisassociateInstanceGroupTerraformModel
	var parts = strings.Split(request.ID, "/")
	var err error
	if len(parts) != 2 {
		err = fmt.Errorf("requires the identifier to be set to <job_template_id>/<instance_group_id>, currently set to %s", request.ID)
		response.Diagnostics.AddError(
			"Unable to import state for JobTemplate association, invalid format.",
			err.Error(),
		)
		return
	}

	var jobTemplateId, instanceGroupId int64

	jobTemplateId, err = strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the job_templateId for the JobTemplate association.", request.ID),
			err.Error(),
		)
		return
	}
	state.JobTemplateID = types.Int64Value(jobTemplateId)

	instanceGroupId, err = strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the instanceGroup_id for the JobTemplate association.", request.ID),
			err.Error(),
		)
		return
	}
	state.InstanceGroupID = types.Int64Value(instanceGroupId)

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (o *jobTemplateAssociateDisassociateInstanceGroup) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error

	// Retrieve values from state
	var plan, state jobTemplateAssociateDisassociateInstanceGroupTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for association of JobTemplate
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.JobTemplateID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = models.AssociateDisassociateRequestModel{ID: plan.InstanceGroupID.ValueInt64(), Disassociate: false}
	tflog.Debug(ctx, "[JobTemplate/Create/Associate] Making a request", map[string]any{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for create of type 'default'", endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to associate for JobTemplate on %s with a payload of %#v", endpoint, bodyRequest),
			err.Error(),
		)
		return
	}

	state.JobTemplateID = plan.JobTemplateID
	state.InstanceGroupID = plan.InstanceGroupID

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *jobTemplateAssociateDisassociateInstanceGroup) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state jobTemplateAssociateDisassociateInstanceGroupTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for disassociation of JobTemplate
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.JobTemplateID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = models.AssociateDisassociateRequestModel{ID: state.InstanceGroupID.ValueInt64(), Disassociate: true}
	tflog.Debug(ctx, "[JobTemplate/Delete/Disassociate] Making a request", map[string]any{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for JobTemplate on %s for delete of type 'default'", o.endpoint),
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

func (o *jobTemplateAssociateDisassociateInstanceGroup) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
}

func (o *jobTemplateAssociateDisassociateInstanceGroup) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
}
