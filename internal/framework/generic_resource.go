package framework

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"

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

// ResourceCfg holds per-resource configuration for the generic CRUD handler.
type ResourceCfg[T any, B any] struct {
	// Schema is the Terraform resource schema.
	Schema rschema.Schema
	// Hook is called before setting state (nil if no hook).
	Hook HookFunc[T]
	// WriteOnlyPlanToBody copies write-only fields from plan to body request (nil if none).
	WriteOnlyPlanToBody func(plan *T, body *B)
	// WriteOnlyPlanToState copies write-only fields from plan to state (nil if none).
	WriteOnlyPlanToState func(plan, state *T)
	// IDAccessor returns the ID value from a model instance for endpoint construction (nil for NoId).
	IDAccessor func(model *T) any
	// IDKey is the schema attribute name carrying the imported ID (typically "id"). Empty when NoId.
	IDKey string
	// IDIsString true → the ID schema attribute is a string (passthrough import).
	// false → parse req.ID as int64 before setting (default for AWX numeric IDs).
	IDIsString bool
	// NoId means the resource has no ID field (settings-style). Create uses PATCH, endpoints have no ID.
	NoId bool
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

// Schema implements resource.Resource.
func (r *GenericResource[T, B, PT]) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = r.Cfg.Schema
}

// ImportState handles terraform import using IDKey + IDIsString from ResourceCfg.
func (r *GenericResource[T, B, PT]) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if r.Cfg.NoId {
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

// Create implements resource.Resource.
func (r *GenericResource[T, B, PT]) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var plan, state T
	if DiagnosticsHasError(&response.Diagnostics, request.Plan.Get(ctx, &plan)...) {
		return
	}

	bodyRequest := PT(&plan).BodyRequest()
	if r.Cfg.WriteOnlyPlanToBody != nil {
		r.Cfg.WriteOnlyPlanToBody(&plan, bodyRequest)
	}

	method := http.MethodPost
	if r.Cfg.NoId {
		method = http.MethodPatch
	}

	endpoint := CleanEndpoint(r.Endpoint)
	data, d := CreateUpdateRequest(ctx, r.Client, method, endpoint, bodyRequest, r.name(), "create")
	if DiagnosticsHasError(&response.Diagnostics, d...) {
		return
	}

	if d, err := PT(&state).UpdateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if r.Cfg.WriteOnlyPlanToState != nil {
		r.Cfg.WriteOnlyPlanToState(&plan, &state)
	}

	if r.Cfg.Hook != nil {
		if HookError(&response.Diagnostics, r.name(), r.Cfg.Hook(ctx, r.Cfg.ApiVersion, hooks.SourceResource, hooks.CalleeCreate, &plan, &state)) {
			return
		}
	}

	if DiagnosticsHasError(&response.Diagnostics, response.State.Set(ctx, &state)...) {
		return
	}
}

// Read implements resource.Resource.
func (r *GenericResource[T, B, PT]) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var state T
	if DiagnosticsHasError(&response.Diagnostics, request.State.Get(ctx, &state)...) {
		return
	}

	var orig T
	if r.Cfg.Hook != nil {
		orig = PT(&state).Clone()
	}

	endpoint := r.endpointForModel(&state)
	data, d := ReadRequest(ctx, r.Client, endpoint, r.name())
	if DiagnosticsHasError(&response.Diagnostics, d...) {
		return
	}

	if d, err := PT(&state).UpdateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if r.Cfg.Hook != nil {
		if HookError(&response.Diagnostics, r.name(), r.Cfg.Hook(ctx, r.Cfg.ApiVersion, hooks.SourceResource, hooks.CalleeRead, &orig, &state)) {
			return
		}
	}

	if DiagnosticsHasError(&response.Diagnostics, response.State.Set(ctx, &state)...) {
		return
	}
}

// Update implements resource.Resource.
func (r *GenericResource[T, B, PT]) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var plan, state T
	if DiagnosticsHasError(&response.Diagnostics, request.Plan.Get(ctx, &plan)...) {
		return
	}

	bodyRequest := PT(&plan).BodyRequest()
	if r.Cfg.WriteOnlyPlanToBody != nil {
		r.Cfg.WriteOnlyPlanToBody(&plan, bodyRequest)
	}

	endpoint := r.endpointForModel(&plan)
	data, d := CreateUpdateRequest(ctx, r.Client, http.MethodPatch, endpoint, bodyRequest, r.name(), "update")
	if DiagnosticsHasError(&response.Diagnostics, d...) {
		return
	}

	if d, err := PT(&state).UpdateFromApiData(data); err != nil {
		response.Diagnostics.Append(d...)
		return
	}

	if r.Cfg.WriteOnlyPlanToState != nil {
		r.Cfg.WriteOnlyPlanToState(&plan, &state)
	}

	if r.Cfg.Hook != nil {
		if HookError(&response.Diagnostics, r.name(), r.Cfg.Hook(ctx, r.Cfg.ApiVersion, hooks.SourceResource, hooks.CalleeUpdate, &plan, &state)) {
			return
		}
	}

	if DiagnosticsHasError(&response.Diagnostics, response.State.Set(ctx, &state)...) {
		return
	}
}

// Delete implements resource.Resource.
func (r *GenericResource[T, B, PT]) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	if r.Cfg.UnDeletable {
		return
	}

	var state T
	if DiagnosticsHasError(&response.Diagnostics, request.State.Get(ctx, &state)...) {
		return
	}

	endpoint := r.endpointForModel(&state)
	if DiagnosticsHasError(&response.Diagnostics, DeleteRequest(ctx, r.Client, endpoint, r.name())...) {
		return
	}
}
