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
	return &organizationAssociateDisassociateGalaxyCredential{}
}

type organizationAssociateDisassociateGalaxyCredential struct {
	client   c.Client
	endpoint string
}

func (o *organizationAssociateDisassociateGalaxyCredential) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	o.client = request.ProviderData.(c.Client)
	o.endpoint = "/api/v2/organizations/%d/galaxy_credentials/"
}

func (o *organizationAssociateDisassociateGalaxyCredential) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_organization_associate_galaxy_credential"
}

func (o *organizationAssociateDisassociateGalaxyCredential) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for association of Organization
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, plan.OrganizationID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = models.AssociateDisassociateRequestModel{ID: plan.GalaxyCredentialID.ValueInt64(), Disassociate: false}
	tflog.Debug(ctx, "[Organization/Create/Associate] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Organization on %s for create of type ", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to associate for Organization on %s", o.endpoint),
			err.Error(),
		)
		return
	}

	state.OrganizationID = plan.OrganizationID
	state.GalaxyCredentialID = plan.GalaxyCredentialID

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (o *organizationAssociateDisassociateGalaxyCredential) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var err error

	// Retrieve values from state
	var state organizationAssociateDisassociateGalaxyCredentialTerraformModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	// Creates a new request for disassociation of Organization
	var r *http.Request
	var endpoint = p.Clean(fmt.Sprintf(o.endpoint, state.OrganizationID.ValueInt64())) + "/"
	var buf bytes.Buffer
	var bodyRequest = models.AssociateDisassociateRequestModel{ID: state.GalaxyCredentialID.ValueInt64(), Disassociate: true}
	tflog.Debug(ctx, "[Organization/Delete/Disassociate] Making a request", map[string]interface{}{
		"payload":  bodyRequest,
		"method":   http.MethodPost,
		"endpoint": endpoint,
	})
	_ = json.NewEncoder(&buf).Encode(bodyRequest)
	if r, err = o.client.NewRequest(ctx, http.MethodPost, endpoint, &buf); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to create a new request for Organization on %s for delete of type ", o.endpoint),
			err.Error(),
		)
		return
	}

	if _, err = o.client.Do(ctx, r); err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to disassociate for Organization on %s", o.endpoint),
			err.Error(),
		)
		return
	}
}

func (o *organizationAssociateDisassociateGalaxyCredential) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
}

func (o *organizationAssociateDisassociateGalaxyCredential) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
}
