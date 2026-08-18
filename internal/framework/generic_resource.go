package framework

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
)

// DataModel is the base constraint for all generated Terraform model types.
type DataModel[T any] interface {
	*T
	Clone() T
	UpdateFromApiData(data map[string]any) (diag.Diagnostics, error)
}

// ResourceModel extends DataModel with BodyRequest for resource types.
type ResourceModel[T any, B any] interface {
	DataModel[T]
	BodyRequest() *B
}

// HookFunc is the signature for pre-state-set hooks.
type HookFunc[T any] func(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *T) error

// WaitLifecycleCfg configures post-Create/Update polling on a generic resource.
// When non-nil and ShouldWait returns true on the plan, the framework polls
// EndpointForModel(state) for Field and blocks until it lands in
// SuccessValues, FailureValues, or the resolved timeout elapses.
type WaitLifecycleCfg[T any] struct {
	// ShouldWait reads the wait toggle from the plan. Returns false → skip.
	ShouldWait func(plan *T) bool
	// EndpointForModel returns the polling endpoint for a populated state model.
	EndpointForModel func(model *T) string
	// Field is the JSON field on the polled response to inspect.
	Field string
	// SuccessValues are terminal values that mean the wait succeeded.
	SuccessValues []string
	// FailureValues are terminal values that mean the wait failed.
	FailureValues []string
	// PollInterval is how long to sleep between polls. Zero → framework default.
	PollInterval time.Duration
	// DefaultTimeout is used when ResolveTimeout returns 0.
	DefaultTimeout time.Duration
	// ResolveTimeout pulls the right duration off the plan's timeouts block
	// (Create vs Update). Returns 0 if the user didn't set a timeout.
	ResolveTimeout func(ctx context.Context, plan *T, callee hooks.Callee) (time.Duration, diag.Diagnostics)
}

// ConfigureFunc runs once at Configure time after the client is wired up. Used
// by resources/data sources that need to look something up from the AWX API at
// startup (e.g. resolving a credential_type ID by namespace) and stash it in a
// closure for later Create/Read/Update calls. Returning a diagnostic with errors
// fails Configure and surfaces a real error to the user instead of panicking.
type ConfigureFunc func(ctx context.Context, client Requester) diag.Diagnostics

// ResourceCfg holds per-resource configuration for the generic CRUD handler.
type ResourceCfg[T any, B any] struct {
	// Schema is the Terraform resource schema.
	Schema rschema.Schema
	// Hook is called before setting state (nil if no hook).
	Hook HookFunc[T]
	// OnConfigure runs once at Configure time after the client is wired up.
	// Use it to look up values from the AWX API and cache them in a closure.
	OnConfigure ConfigureFunc
	// MutateBody runs after BodyRequest() in Create and Update. Use it to
	// inject values that are resolved at Configure time (not present on the
	// plan) into the outbound request body — e.g. the credential_type ID
	// looked up by namespace.
	MutateBody func(plan *T, body *B)
	// WriteOnlyPlanToBody copies write-only fields from plan to body request (nil if none).
	WriteOnlyPlanToBody func(plan *T, body *B)
	// WriteOnlyPlanToState copies write-only fields from plan to state (nil if none).
	WriteOnlyPlanToState func(plan, state *T)
	// CopyExtraAttributes copies non-API ("Terraform-only") attributes from
	// plan to state so they round-trip without going through UpdateFromApiData.
	// Same call site as WriteOnlyPlanToState.
	CopyExtraAttributes func(plan, state *T)
	// EmitTimeouts injects a `timeouts { create, update }` block into the
	// resource schema at Schema() time. Pairs with WaitLifecycle.ResolveTimeout.
	EmitTimeouts bool
	// WaitLifecycle, when non-nil, polls the resource after Create/Update
	// until the configured field reaches a terminal value.
	WaitLifecycle *WaitLifecycleCfg[T]
	// CreateEndpoint returns the POST target when the API creates the object at
	// a parent-scoped URL but manages it elsewhere. Create then reconciles with
	// a PATCH against the regular endpoint, because those URLs take only the
	// fields that bring the object into existence and drop the rest silently.
	CreateEndpoint func(plan *T) string
	// ImportIDParts names the schema attributes a slash-separated import ID maps
	// onto, in order. Needed when a read cannot recover every required
	// attribute: a parent id that lives only in the create URL never comes back
	// in the response. Empty means the import ID is the resource ID alone.
	ImportIDParts []string
	// IDAccessor returns the ID value from a model instance for endpoint construction (nil for NoId).
	IDAccessor func(model *T) any
	// IDKey is the schema attribute name carrying the imported ID (typically "id"). Empty when NoId.
	IDKey string
	// IDIsString true → the ID schema attribute is a string (passthrough import).
	// false → parse req.ID as int64 before setting (default for AWX numeric IDs).
	IDIsString bool
	// NoId means the resource has no ID field (settings-style). Create uses PATCH, endpoints have no ID.
	NoId bool
	// NoImport disables terraform import for this resource. Attempts return an error diagnostic.
	NoImport bool
	// SoftDelete replaces the DELETE call with a PATCH carrying this body, for
	// objects AWX retires through a state field. DELETE on an instance answers
	// 405; moving node_state to "deprovisioning" is what removes it.
	SoftDelete map[string]any
	// UnDeletable means Delete is a no-op.
	UnDeletable bool
	// ApiVersion is passed to hook functions.
	ApiVersion string
	// ResourceName is used in error messages. Defaults to TypeName if empty.
	ResourceName string
}

// GenericResource implements resource.Resource for any ResourceModel type.
type GenericResource[T any, B any, PT ResourceModel[T, B]] struct {
	ResourceBase
	Cfg ResourceCfg[T, B]
}

func (r *GenericResource[T, B, PT]) name() string {
	if r.Cfg.ResourceName != "" {
		return r.Cfg.ResourceName
	}
	return r.TypeName
}

func (r *GenericResource[T, B, PT]) endpointForModel(model *T) string {
	if r.Cfg.NoId || r.Cfg.IDAccessor == nil {
		return CleanEndpoint(r.Endpoint)
	}
	return EndpointWithID(r.Endpoint, r.Cfg.IDAccessor(model))
}

// Configure wires the client and then runs the optional OnConfigure hook.
// Shadows the ResourceBase.Configure promoted method so OnConfigure actually fires.
func (r *GenericResource[T, B, PT]) Configure(ctx context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	r.ResourceBase.Configure(ctx, request, response)
	if r.Cfg.OnConfigure == nil || r.Client == nil {
		return
	}
	response.Diagnostics.Append(r.Cfg.OnConfigure(ctx, r.Client)...)
}

// Schema returns r.Cfg.Schema, optionally injecting a `timeouts` block when
// EmitTimeouts is set so wait-lifecycle resources get user-tunable Create/Update
// timeouts without templating it per-resource.
func (r *GenericResource[T, B, PT]) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = r.Cfg.Schema
	if r.Cfg.EmitTimeouts {
		if resp.Schema.Blocks == nil {
			resp.Schema.Blocks = map[string]rschema.Block{}
		}
		resp.Schema.Blocks["timeouts"] = timeouts.Block(ctx, timeouts.Opts{
			Create: true,
			Update: true,
		})
	}
}

// runWaitLifecycle polls the resource after a successful Create or Update
// when WaitLifecycle is configured and the plan opts in via ShouldWait.
func (r *GenericResource[T, B, PT]) runWaitLifecycle(ctx context.Context, plan, state *T, callee hooks.Callee, diags *diag.Diagnostics) {
	wl := r.Cfg.WaitLifecycle
	if wl == nil || wl.ShouldWait == nil || !wl.ShouldWait(plan) {
		return
	}

	timeout := wl.DefaultTimeout
	if wl.ResolveTimeout != nil {
		resolved, d := wl.ResolveTimeout(ctx, plan, callee)
		if DiagnosticsHasError(diags, d...) {
			return
		}
		if resolved > 0 {
			timeout = resolved
		}
	}
	if timeout <= 0 {
		diags.AddError(
			fmt.Sprintf("Cannot wait for %s: no timeout configured", r.name()),
			"WaitLifecycle is enabled but neither the timeouts block nor DefaultTimeout produced a positive duration.",
		)
		return
	}

	endpoint := wl.EndpointForModel(state)
	if endpoint == "" {
		diags.AddError(
			fmt.Sprintf("Cannot wait for %s: missing resource ID", r.name()),
			"The API response did not include a usable ID, so the framework cannot construct a polling URL. Re-run with TF_LOG=DEBUG to inspect the response payload.",
		)
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("[%s/wait] polling for terminal status", r.name()), map[string]any{
		"endpoint": endpoint,
		"field":    wl.Field,
		"timeout":  timeout.String(),
	})

	waitCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	err := WaitForFieldValue(waitCtx, r.Client, WaitForFieldOpts{
		Endpoint:      endpoint,
		Field:         wl.Field,
		SuccessValues: wl.SuccessValues,
		FailureValues: wl.FailureValues,
		PollInterval:  wl.PollInterval,
	})
	if err == nil {
		return
	}

	var term *WaitTerminalError
	if errors.As(err, &term) {
		diags.AddError(
			fmt.Sprintf("%s reached terminal failure status %q on %s", r.name(), term.Status, endpoint),
			"AWX reported a non-recoverable status while waiting for the resource to become ready. Check the AWX UI for details.",
		)
		return
	}
	diags.AddError(
		fmt.Sprintf("Timed out or failed waiting for %s on %s", r.name(), endpoint),
		err.Error(),
	)
}

func (r *GenericResource[T, B, PT]) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if r.Cfg.NoId || r.Cfg.NoImport {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Resource %s does not support import.", r.name()),
			"This resource has not been configured to support `terraform import`.",
		)
		return
	}
	if len(r.Cfg.ImportIDParts) > 0 {
		r.importComposite(ctx, req, resp)
		return
	}
	idPath := path.Root(r.Cfg.IDKey)
	if r.Cfg.IDIsString {
		resource.ImportStatePassthroughID(ctx, idPath, req, resp)
		return
	}
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the ID for the %s.", req.ID, r.name()),
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, idPath, id)...)
}

func (r *GenericResource[T, B, PT]) importComposite(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != len(r.Cfg.ImportIDParts) {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to import state for %s, invalid format.", r.name()),
			fmt.Sprintf("requires the identifier to be set to <%s>, currently set to %s",
				strings.Join(r.Cfg.ImportIDParts, ">/<"), req.ID),
		)
		return
	}
	for i, attr := range r.Cfg.ImportIDParts {
		if attr == r.Cfg.IDKey && r.Cfg.IDIsString {
			resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(attr), parts[i])...)
			continue
		}
		value, err := strconv.ParseInt(parts[i], 10, 64)
		if err != nil {
			resp.Diagnostics.AddError(
				fmt.Sprintf("Unable to parse '%v' as an int64 number, please provide the %s for the %s.", parts[i], attr, r.name()),
				err.Error(),
			)
			return
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(attr), value)...)
	}
}

func (r *GenericResource[T, B, PT]) buildBody(plan *T) *B {
	bodyRequest := PT(plan).BodyRequest()
	if r.Cfg.WriteOnlyPlanToBody != nil {
		r.Cfg.WriteOnlyPlanToBody(plan, bodyRequest)
	}
	if r.Cfg.MutateBody != nil {
		r.Cfg.MutateBody(plan, bodyRequest)
	}
	return bodyRequest
}

// createAtSubEndpoint POSTs to the parent-scoped create URL and returns the id
// AWX assigned, which Create uses to address the reconciling PATCH.
func (r *GenericResource[T, B, PT]) createAtSubEndpoint(ctx context.Context, plan *T, diags *diag.Diagnostics) (any, bool) {
	if r.Cfg.IDAccessor == nil {
		diags.AddError(
			fmt.Sprintf("Cannot create %s through a separate endpoint", r.name()),
			"CreateEndpoint needs IDAccessor, since the follow-up request against the regular endpoint is addressed by the id AWX assigned.",
		)
		return nil, false
	}

	data, d := CreateUpdateRequest(ctx, r.Client, http.MethodPost, r.Cfg.CreateEndpoint(plan), r.buildBody(plan), r.name(), "create")
	if DiagnosticsHasError(diags, d...) {
		return nil, false
	}

	var created T
	d, err := PT(&created).UpdateFromApiData(data)
	diags.Append(d...)
	if err != nil || diags.HasError() {
		return nil, false
	}
	return r.Cfg.IDAccessor(&created), true
}

// applyMutation is the shared spine for Create and Update: assemble the body,
// call the API, hydrate state from the response, run write-only/extra/hook
// wiring, and poll the wait-lifecycle. Returns ok=false (caller should bail)
// whenever it has appended a hard error to diags.
func (r *GenericResource[T, B, PT]) applyMutation(
	ctx context.Context,
	plan *T,
	method, endpoint, operation string,
	callee hooks.Callee,
	diags *diag.Diagnostics,
) (state T, ok bool) {
	data, d := CreateUpdateRequest(ctx, r.Client, method, endpoint, r.buildBody(plan), r.name(), operation)
	if DiagnosticsHasError(diags, d...) {
		return state, false
	}

	d, err := PT(&state).UpdateFromApiData(data)
	diags.Append(d...)
	if err != nil || diags.HasError() {
		return state, false
	}

	if r.Cfg.WriteOnlyPlanToState != nil {
		r.Cfg.WriteOnlyPlanToState(plan, &state)
	}
	if r.Cfg.CopyExtraAttributes != nil {
		r.Cfg.CopyExtraAttributes(plan, &state)
	}
	if r.Cfg.Hook != nil {
		if HookError(diags, r.name(), r.Cfg.Hook(ctx, r.Cfg.ApiVersion, hooks.SourceResource, callee, plan, &state)) {
			return state, false
		}
	}

	r.runWaitLifecycle(ctx, plan, &state, callee, diags)
	if diags.HasError() {
		return state, false
	}
	return state, true
}

func (r *GenericResource[T, B, PT]) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var plan T
	if DiagnosticsHasError(&response.Diagnostics, request.Plan.Get(ctx, &plan)...) {
		return
	}

	method := http.MethodPost
	if r.Cfg.NoId {
		method = http.MethodPatch
	}
	endpoint := CleanEndpoint(r.Endpoint)

	if r.Cfg.CreateEndpoint != nil {
		id, ok := r.createAtSubEndpoint(ctx, &plan, &response.Diagnostics)
		if !ok {
			return
		}
		method, endpoint = http.MethodPatch, EndpointWithID(r.Endpoint, id)
	}

	state, ok := r.applyMutation(ctx, &plan, method, endpoint, "create", hooks.CalleeCreate, &response.Diagnostics)
	if !ok {
		return
	}
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (r *GenericResource[T, B, PT]) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var state T
	if DiagnosticsHasError(&response.Diagnostics, request.State.Get(ctx, &state)...) {
		return
	}

	var orig *T
	if r.Cfg.Hook != nil {
		o := PT(&state).Clone()
		orig = &o
	}

	data, found, d := ReadRequestAllowMissing(ctx, r.Client, r.endpointForModel(&state), r.name())
	if DiagnosticsHasError(&response.Diagnostics, d...) {
		return
	}
	if !found {
		response.State.RemoveResource(ctx)
		return
	}

	d, err := PT(&state).UpdateFromApiData(data)
	response.Diagnostics.Append(d...)
	if err != nil || response.Diagnostics.HasError() {
		return
	}

	if r.Cfg.Hook != nil {
		if HookError(&response.Diagnostics, r.name(), r.Cfg.Hook(ctx, r.Cfg.ApiVersion, hooks.SourceResource, hooks.CalleeRead, orig, &state)) {
			return
		}
	}

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (r *GenericResource[T, B, PT]) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var plan T
	if DiagnosticsHasError(&response.Diagnostics, request.Plan.Get(ctx, &plan)...) {
		return
	}
	state, ok := r.applyMutation(ctx, &plan, http.MethodPatch, r.endpointForModel(&plan), "update", hooks.CalleeUpdate, &response.Diagnostics)
	if !ok {
		return
	}
	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (r *GenericResource[T, B, PT]) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	if r.Cfg.UnDeletable {
		return
	}

	var state T
	if DiagnosticsHasError(&response.Diagnostics, request.State.Get(ctx, &state)...) {
		return
	}

	endpoint := r.endpointForModel(&state)
	if r.Cfg.SoftDelete != nil {
		_, d := CreateUpdateRequest(ctx, r.Client, http.MethodPatch, endpoint, r.Cfg.SoftDelete, r.name(), "delete")
		response.Diagnostics.Append(d...)
		return
	}
	if DiagnosticsHasError(&response.Diagnostics, DeleteRequest(ctx, r.Client, endpoint, r.name())...) {
		return
	}
}
