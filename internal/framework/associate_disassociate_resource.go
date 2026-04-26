package framework

import (
	"context"
	"fmt"
	"net/http"
	p "path"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/models"
)

// AssociateDisassociateConfig configures the per-resource bits of an
// AssociateDisassociateResource. The resource handles a parent ↔ child
// association via AWX's POST {associate|disassociate} pattern.
type AssociateDisassociateConfig struct {
	// TypeName is the Terraform type name suffix (e.g. "host_associate_group").
	TypeName string
	// Endpoint is the API path with %d for the parent ID, plus optional %s
	// for the option flavor (notification templates).
	Endpoint string
	// ParentName is the human-readable parent resource name (e.g. "Host").
	ParentName string
	// ParentIDAttr is the tfsdk attribute name for the parent ID (e.g. "host_id").
	ParentIDAttr string
	// ChildName is the human-readable child resource name (e.g. "Group").
	ChildName string
	// ChildIDAttr is the tfsdk attribute name for the child ID (e.g. "group_id").
	ChildIDAttr string
	// AssociateType is one of "" (default), "notification_job_template", or
	// "notification_job_workflow_template". The notification flavors require
	// an option attribute with a flavor-specific OneOf set.
	AssociateType string
	// Deprecated marks the resource as deprecated.
	Deprecated bool
}

// optionValues returns the OneOf set for the option attribute, keyed by
// AssociateType. Returns nil for non-notification flavors.
func (c AssociateDisassociateConfig) optionValues() []string {
	switch c.AssociateType {
	case "notification_job_template":
		return []string{"started", "success", "error"}
	case "notification_job_workflow_template":
		return []string{"approval", "started", "success", "error"}
	default:
		return nil
	}
}

func (c AssociateDisassociateConfig) hasOption() bool {
	return c.AssociateType == "notification_job_template" || c.AssociateType == "notification_job_workflow_template"
}

var (
	_ resource.Resource                = (*AssociateDisassociateResource)(nil)
	_ resource.ResourceWithConfigure   = (*AssociateDisassociateResource)(nil)
	_ resource.ResourceWithImportState = (*AssociateDisassociateResource)(nil)
)

// AssociateDisassociateResource is the generic implementation backing every
// <parent>_associate_<child> resource. The per-resource constructors in
// internal/awx are thin wrappers that supply an AssociateDisassociateConfig.
type AssociateDisassociateResource struct {
	ResourceBase
	cfg AssociateDisassociateConfig
}

// NewAssociateDisassociateResource constructs an AssociateDisassociateResource.
func NewAssociateDisassociateResource(cfg AssociateDisassociateConfig) resource.Resource {
	return &AssociateDisassociateResource{
		ResourceBase: ResourceBase{
			ProviderBase: ProviderBase{TypeName: cfg.TypeName, Endpoint: cfg.Endpoint},
		},
		cfg: cfg,
	}
}

// Schema defines the schema for the resource.
func (o *AssociateDisassociateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	attrs := map[string]schema.Attribute{
		o.cfg.ParentIDAttr: schema.Int64Attribute{
			Description: fmt.Sprintf("Database ID for this %s.", o.cfg.ParentName),
			Required:    true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.RequiresReplace(),
			},
		},
		o.cfg.ChildIDAttr: schema.Int64Attribute{
			Description: fmt.Sprintf("Database ID of the %s to assign.", strings.ToLower(o.cfg.ChildName)),
			Required:    true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.RequiresReplace(),
			},
		},
	}
	if o.cfg.hasOption() {
		attrs["option"] = schema.StringAttribute{
			Description: "Notification Option",
			Required:    true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
			Validators: []validator.String{
				stringvalidator.OneOf(o.cfg.optionValues()...),
			},
		}
	}

	s := schema.Schema{Attributes: attrs}
	if o.cfg.Deprecated {
		s.DeprecationMessage = "This resource has been deprecated and will be removed in a future release."
	}
	resp.Schema = s
}

// ConfigValidators implements resource.ResourceWithConfigValidators.
func (o *AssociateDisassociateResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	paths := []path.Expression{
		path.MatchRoot(o.cfg.ParentIDAttr),
		path.MatchRoot(o.cfg.ChildIDAttr),
	}
	if o.cfg.hasOption() {
		paths = append(paths, path.MatchRoot("option"))
	}
	return []resource.ConfigValidator{resourcevalidator.RequiredTogether(paths...)}
}

// ImportState supports importing existing associations as <parent_id>/<child_id>.
// Notification flavors carry an extra option field and are not importable.
func (o *AssociateDisassociateResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	if o.cfg.hasOption() {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to import state for %s association.", o.cfg.ParentName),
			"This association type does not support import.",
		)
		return
	}

	parts := strings.Split(request.ID, "/")
	if len(parts) != 2 {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to import state for %s association, invalid format.", o.cfg.ParentName),
			fmt.Sprintf("requires the identifier to be set to <%s>/<%s>, currently set to %s", o.cfg.ParentIDAttr, o.cfg.ChildIDAttr, request.ID),
		)
		return
	}

	parentID, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the %s for the %s association.", request.ID, o.cfg.ParentIDAttr, o.cfg.ParentName),
			err.Error(),
		)
		return
	}

	childID, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		response.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the %s for the %s association.", request.ID, o.cfg.ChildIDAttr, o.cfg.ParentName),
			err.Error(),
		)
		return
	}

	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root(o.cfg.ParentIDAttr), types.Int64Value(parentID))...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root(o.cfg.ChildIDAttr), types.Int64Value(childID))...)
}

// Create issues the associate request and copies plan into state.
func (o *AssociateDisassociateResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	parentID, childID, option, ok := o.readIDs(ctx, &request.Plan, &response.Diagnostics)
	if !ok {
		return
	}
	if !o.sendAssoc(ctx, parentID, childID, option, false, &response.Diagnostics) {
		return
	}

	if DiagnosticsHasError(&response.Diagnostics, response.State.SetAttribute(ctx, path.Root(o.cfg.ParentIDAttr), types.Int64Value(parentID))...) {
		return
	}
	if DiagnosticsHasError(&response.Diagnostics, response.State.SetAttribute(ctx, path.Root(o.cfg.ChildIDAttr), types.Int64Value(childID))...) {
		return
	}
	if o.cfg.hasOption() {
		if DiagnosticsHasError(&response.Diagnostics, response.State.SetAttribute(ctx, path.Root("option"), types.StringValue(option))...) {
			return
		}
	}
}

// Delete issues the disassociate request.
func (o *AssociateDisassociateResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	parentID, childID, option, ok := o.readIDs(ctx, &request.State, &response.Diagnostics)
	if !ok {
		return
	}
	o.sendAssoc(ctx, parentID, childID, option, true, &response.Diagnostics)
}

// Read is a no-op — these resources hold no AWX-side state worth refreshing.
func (o *AssociateDisassociateResource) Read(_ context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
}

// Update is a no-op — every attribute uses RequiresReplace.
func (o *AssociateDisassociateResource) Update(_ context.Context, _ resource.UpdateRequest, _ *resource.UpdateResponse) {
}

// attributeReader is the GetAttribute interface shared by tfsdk.Plan and tfsdk.State.
type attributeReader interface {
	GetAttribute(ctx context.Context, p path.Path, target any) diag.Diagnostics
}

// readIDs pulls parent/child IDs (and option, when applicable) from a plan or state.
func (o *AssociateDisassociateResource) readIDs(ctx context.Context, src attributeReader, diags *diag.Diagnostics) (int64, int64, string, bool) {
	var parentID, childID types.Int64
	if DiagnosticsHasError(diags, src.GetAttribute(ctx, path.Root(o.cfg.ParentIDAttr), &parentID)...) {
		return 0, 0, "", false
	}
	if DiagnosticsHasError(diags, src.GetAttribute(ctx, path.Root(o.cfg.ChildIDAttr), &childID)...) {
		return 0, 0, "", false
	}

	var option string
	if o.cfg.hasOption() {
		var optionVal types.String
		if DiagnosticsHasError(diags, src.GetAttribute(ctx, path.Root("option"), &optionVal)...) {
			return 0, 0, "", false
		}
		option = optionVal.ValueString()
	}
	return parentID.ValueInt64(), childID.ValueInt64(), option, true
}

// sendAssoc builds and sends the associate/disassociate POST.
func (o *AssociateDisassociateResource) sendAssoc(ctx context.Context, parentID, childID int64, option string, disassociate bool, diags *diag.Diagnostics) bool {
	args := []any{parentID}
	if o.cfg.hasOption() {
		args = append(args, option)
	}
	endpoint := p.Clean(fmt.Sprintf(o.Endpoint, args...)) + "/"

	op := "associate"
	if disassociate {
		op = "disassociate"
	}

	body := models.AssociateDisassociateRequestModel{ID: childID, Disassociate: disassociate}
	_, d := CreateUpdateRequest(ctx, o.Client, http.MethodPost, endpoint, body, o.cfg.ParentName, op)
	return !DiagnosticsHasError(diags, d...)
}

// Compile-time check that tfsdk.Plan and tfsdk.State satisfy our attributeReader.
var (
	_ attributeReader = (*tfsdk.Plan)(nil)
	_ attributeReader = (*tfsdk.State)(nil)
)
