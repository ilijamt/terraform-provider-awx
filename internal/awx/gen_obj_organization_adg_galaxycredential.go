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
	"github.com/ilijamt/terraform-provider-awx/internal/framework"
	"github.com/ilijamt/terraform-provider-awx/internal/models"
)

var (
	_ resource.Resource                = &organizationAssociateDisassociateGalaxyCredential{}
	_ resource.ResourceWithConfigure   = &organizationAssociateDisassociateGalaxyCredential{}
	_ resource.ResourceWithImportState = &organizationAssociateDisassociateGalaxyCredential{}
)

type organizationAssociateDisassociateGalaxyCredentialTerraformModel struct {
	OrganizationID     types.Int64 `tfsdk:"organization_id"`
	GalaxyCredentialID types.Int64 `tfsdk:"galaxy_credential_id"`
}

// NewOrganizationAssociateDisassociateGalaxyCredentialResource is a helper function to simplify the provider implementation.
func NewOrganizationAssociateDisassociateGalaxyCredentialResource() resource.Resource {
	return &organizationAssociateDisassociateGalaxyCredential{ResourceBase: framework.ResourceBase{ProviderBase: framework.ProviderBase{TypeName: "organization_associate_galaxy_credential", Endpoint: "/api/v2/organizations/%d/galaxy_credentials/"}}}
}

type organizationAssociateDisassociateGalaxyCredential struct {
	framework.ResourceBase
}

func (o *organizationAssociateDisassociateGalaxyCredential) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		DeprecationMessage: "This resource has been deprecated and will be removed in a future release.",
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.Int64Attribute{
				Description: "Database ID for this Organization.",
				Required:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"galaxy_credential_id": schema.Int64Attribute{
				Description: "Database ID of the galaxycredential to assign.",
				Required:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (o *organizationAssociateDisassociateGalaxyCredential) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.RequiredTogether(
			path.MatchRoot("organization_id"),
			path.MatchRoot("galaxy_credential_id"),
		),
	}
}

func (o *organizationAssociateDisassociateGalaxyCredential) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	var state organizationAssociateDisassociateGalaxyCredentialTerraformModel
	var parts = strings.Split(request.ID, "/")
	var err error
	if len(parts) != 2 {
		err = fmt.Errorf("requires the identifier to be set to <organization_id>/<galaxy_credential_id>, currently set to %s", request.ID)
		response.Diagnostics.AddError(
			"Unable to import state for Organization association, invalid format.",
			err.Error(),
		)
		return
	}

	var organizationId, galaxyCredentialId int64

	organizationId, err = strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the organizationId for the Organization association.", request.ID),
			err.Error(),
		)
		return
	}
	state.OrganizationID = types.Int64Value(organizationId)

	galaxyCredentialId, err = strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the galaxyCredential_id for the Organization association.", request.ID),
			err.Error(),
		)
		return
	}
	state.GalaxyCredentialID = types.Int64Value(galaxyCredentialId)

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (o *organizationAssociateDisassociateGalaxyCredential) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var err error

	// Retrieve values from state
	var plan, state organizationAssociateDisassociateGalaxyCredentialTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.Plan.Get(ctx, &plan)...) {
		return
	}

	// Creates a new request for association of Organization
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.Endpoint, plan.OrganizationID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = models.AssociateDisassociateRequestModel{ID: plan.GalaxyCredentialID.ValueInt64(), Disassociate: false}
	tflog.Debug(ctx, "[Organization/Create/Associate] Making a request", map[string]any{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.Client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Organization on %s for create of type 'default'", endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.Client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to associate for Organization on %s with a payload of %#v", endpoint, bodyRequest),
			err.Error(),
		)
		return
	}

	state.OrganizationID = plan.OrganizationID
	state.GalaxyCredentialID = plan.GalaxyCredentialID

	if framework.DiagnosticsHasError(&response.Diagnostics, response.State.Set(ctx, &state)...) {
		return
	}
}

func (o *organizationAssociateDisassociateGalaxyCredential) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state organizationAssociateDisassociateGalaxyCredentialTerraformModel
	if framework.DiagnosticsHasError(&response.Diagnostics, request.State.Get(ctx, &state)...) {
		return
	}

	// Creates a new request for disassociation of Organization
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.Endpoint, state.OrganizationID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = models.AssociateDisassociateRequestModel{ID: state.GalaxyCredentialID.ValueInt64(), Disassociate: true}
	tflog.Debug(ctx, "[Organization/Delete/Disassociate] Making a request", map[string]any{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.Client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Organization on %s for delete of type 'default'", o.Endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.Client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to disassociate for Organization on %s", o.Endpoint),
			err.Error(),
		)
		return
	}
}

func (o *organizationAssociateDisassociateGalaxyCredential) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
}

func (o *organizationAssociateDisassociateGalaxyCredential) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
}
