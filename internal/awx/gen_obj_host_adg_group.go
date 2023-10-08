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

	"github.com/ilijamt/terraform-provider-awx/internal/models"

	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	c "github.com/ilijamt/terraform-provider-awx/internal/client"
)

var (
	_ resource.Resource                = &hostAssociateDisassociateGroup{}
	_ resource.ResourceWithConfigure   = &hostAssociateDisassociateGroup{}
	_ resource.ResourceWithImportState = &hostAssociateDisassociateGroup{}
)

type hostAssociateDisassociateGroupTerraformModel struct {
	HostID  types.Int64 `tfsdk:"host_id"`
	GroupID types.Int64 `tfsdk:"group_id"`
}

// NewHostAssociateDisassociateGroupResource is a helper function to simplify the provider implementation.
func NewHostAssociateDisassociateGroupResource() resource.Resource {
	return &hostAssociateDisassociateGroup{}
}

type hostAssociateDisassociateGroup struct {
	client   c.Client
	endpoint string
}

func (o *hostAssociateDisassociateGroup) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/hosts/%d/groups/"
}

func (o *hostAssociateDisassociateGroup) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_host_associate_group"
}

func (o *hostAssociateDisassociateGroup) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host_id": schema.Int64Attribute{
				Description: "Database ID for this Host.",
				Required:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"group_id": schema.Int64Attribute{
				Description: "Database ID of the group to assign.",
				Required:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (o *hostAssociateDisassociateGroup) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.RequiredTogether(
			path.MatchRoot("host_id"),
			path.MatchRoot("group_id"),
		),
	}
}

func (o *hostAssociateDisassociateGroup) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var state hostAssociateDisassociateGroupTerraformModel
	var parts = strings.Split(request.ID, "/")
	var err error
	if len(parts) != 2 {
		err = fmt.Errorf("requires the identifier to be set to <host_id>/<group_id>, currently set to %s", request.ID)
		response.Diagnostics.AddError(
			"Unable to import state for Host association, invalid format.",
			err.Error(),
		)
		return
	}

	var hostId, groupId int64

	hostId, err = strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the hostId for the Host association.", request.ID),
			err.Error(),
		)
		return
	}
	state.HostID = types.Int64Value(hostId)

	groupId, err = strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the group_id for the Host association.", request.ID),
			err.Error(),
		)
		return
	}
	state.GroupID = types.Int64Value(groupId)

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (o *hostAssociateDisassociateGroup) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error

	// Retrieve values from state
	var plan, state hostAssociateDisassociateGroupTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for association of Host
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.HostID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = models.AssociateDisassociateRequestModel{ID: plan.GroupID.ValueInt64(), Disassociate: false}
	tflog.Debug(ctx, "[Host/Create/Associate] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Host on %s for create of type ", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to associate for Host on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	state.HostID = plan.HostID
	state.GroupID = plan.GroupID

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *hostAssociateDisassociateGroup) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state hostAssociateDisassociateGroupTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for disassociation of Host
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.HostID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = models.AssociateDisassociateRequestModel{ID: state.GroupID.ValueInt64(), Disassociate: true}
	tflog.Debug(ctx, "[Host/Delete/Disassociate] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Host on %s for delete of type ", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to disassociate for Host on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}

func (o *hostAssociateDisassociateGroup) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
}

func (o *hostAssociateDisassociateGroup) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
}
