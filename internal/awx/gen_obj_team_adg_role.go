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
	_ resource.Resource                = &teamAssociateDisassociateRole{}
	_ resource.ResourceWithConfigure   = &teamAssociateDisassociateRole{}
	_ resource.ResourceWithImportState = &teamAssociateDisassociateRole{}
)

type teamAssociateDisassociateRoleTerraformModel struct {
	TeamID types.Int64 `tfsdk:"team_id"`
	RoleID types.Int64 `tfsdk:"role_id"`
}

// NewTeamAssociateDisassociateRoleResource is a helper function to simplify the provider implementation.
func NewTeamAssociateDisassociateRoleResource() resource.Resource {
	return &teamAssociateDisassociateRole{}
}

type teamAssociateDisassociateRole struct {
	client   c.Client
	endpoint string
}

func (o *teamAssociateDisassociateRole) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/teams/%d/roles/"
}

func (o *teamAssociateDisassociateRole) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_team_associate_role"
}

func (o *teamAssociateDisassociateRole) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"team_id": schema.Int64Attribute{
				Description: "Database ID for this Team.",
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

func (o *teamAssociateDisassociateRole) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.RequiredTogether(
			path.MatchRoot("team_id"),
			path.MatchRoot("role_id"),
		),
	}
}

func (o *teamAssociateDisassociateRole) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var state teamAssociateDisassociateRoleTerraformModel
	var parts = strings.Split(request.ID, "/")
	var err error
	if len(parts) != 2 {
		err = fmt.Errorf("requires the identifier to be set to <team_id>/<role_id>, currently set to %s", request.ID)
		response.Diagnostics.AddError(
			"Unable to import state for Team association, invalid format.",
			err.Error(),
		)
		return
	}

	var teamId, roleId int64

	teamId, err = strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the teamId for the Team association.", request.ID),
			err.Error(),
		)
		return
	}
	state.TeamID = types.Int64Value(teamId)

	roleId, err = strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the role_id for the Team association.", request.ID),
			err.Error(),
		)
		return
	}
	state.RoleID = types.Int64Value(roleId)

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (o *teamAssociateDisassociateRole) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error

	// Retrieve values from state
	var plan, state teamAssociateDisassociateRoleTerraformModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for association of Team
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.TeamID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = models.AssociateDisassociateRequestModel{ID: plan.RoleID.ValueInt64(), Disassociate: false}
	tflog.Debug(ctx, "[Team/Create/Associate] Making a request", map[string]any{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Team on %s for create of type 'default'", endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to associate for Team on %s with a payload of %#v", endpoint, bodyRequest),
			err.Error(),
		)
		return
	}

	state.TeamID = plan.TeamID
	state.RoleID = plan.RoleID

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *teamAssociateDisassociateRole) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state teamAssociateDisassociateRoleTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for disassociation of Team
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.TeamID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = models.AssociateDisassociateRequestModel{ID: state.RoleID.ValueInt64(), Disassociate: true}
	tflog.Debug(ctx, "[Team/Delete/Disassociate] Making a request", map[string]any{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Team on %s for delete of type 'default'", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to disassociate for Team on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}

func (o *teamAssociateDisassociateRole) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
}

func (o *teamAssociateDisassociateRole) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
}
