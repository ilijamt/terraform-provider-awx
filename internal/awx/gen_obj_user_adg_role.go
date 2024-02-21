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
	_ resource.Resource                = &userAssociateDisassociateRole{}
	_ resource.ResourceWithConfigure   = &userAssociateDisassociateRole{}
	_ resource.ResourceWithImportState = &userAssociateDisassociateRole{}
)

type userAssociateDisassociateRoleTerraformModel struct {
	UserID types.Int64 `tfsdk:"user_id"`
	RoleID types.Int64 `tfsdk:"role_id"`
}

// NewUserAssociateDisassociateRoleResource is a helper function to simplify the provider implementation.
func NewUserAssociateDisassociateRoleResource() resource.Resource {
	return &userAssociateDisassociateRole{}
}

type userAssociateDisassociateRole struct {
	client   c.Client
	endpoint string
}

func (o *userAssociateDisassociateRole) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/users/%d/roles/"
}

func (o *userAssociateDisassociateRole) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_user_associate_role"
}

func (o *userAssociateDisassociateRole) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"user_id": schema.Int64Attribute{
				Description: "Database ID for this User.",
				Required:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"role_id": schema.Int64Attribute{
				Description: "Database ID of the role to assign.",
				Required:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (o *userAssociateDisassociateRole) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.RequiredTogether(
			path.MatchRoot("user_id"),
			path.MatchRoot("role_id"),
		),
	}
}

func (o *userAssociateDisassociateRole) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var state userAssociateDisassociateRoleTerraformModel
	var parts = strings.Split(request.ID, "/")
	var err error
	if len(parts) != 2 {
		err = fmt.Errorf("requires the identifier to be set to <user_id>/<role_id>, currently set to %s", request.ID)
		response.Diagnostics.AddError(
			"Unable to import state for User association, invalid format.",
			err.Error(),
		)
		return
	}

	var userId, roleId int64

	userId, err = strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the userId for the User association.", request.ID),
			err.Error(),
		)
		return
	}
	state.UserID = types.Int64Value(userId)

	roleId, err = strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the role_id for the User association.", request.ID),
			err.Error(),
		)
		return
	}
	state.RoleID = types.Int64Value(roleId)

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (o *userAssociateDisassociateRole) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error

	// Retrieve values from state
	var plan, state userAssociateDisassociateRoleTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for association of User
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.UserID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = models.AssociateDisassociateRequestModel{ID: plan.RoleID.ValueInt64(), Disassociate: false}
	tflog.Debug(ctx, "[User/Create/Associate] Making a request", map[string]any{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for User on %s for create of type ", endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to associate for User on %s with a payload of %#v", endpoint, bodyRequest),
			err.Error(),
		)
		return
	}

	state.UserID = plan.UserID
	state.RoleID = plan.RoleID

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *userAssociateDisassociateRole) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state userAssociateDisassociateRoleTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for disassociation of User
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.UserID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = models.AssociateDisassociateRequestModel{ID: state.RoleID.ValueInt64(), Disassociate: true}
	tflog.Debug(ctx, "[User/Delete/Disassociate] Making a request", map[string]any{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for User on %s for delete of type ", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to disassociate for User on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}

func (o *userAssociateDisassociateRole) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
}

func (o *userAssociateDisassociateRole) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
}
