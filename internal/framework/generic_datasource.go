package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
)

// DataSourceCfg holds per-datasource configuration for the generic data source.
type DataSourceCfg[T any] struct {
	// Schema is the Terraform data source schema.
	Schema dschema.Schema
	// SearchGroups defines the search groups for this data source (nil for no-search).
	SearchGroups []SearchGroup
	// Hook is called before setting state (nil if no hook).
	Hook HookFunc[T]
	// OnConfigure runs once at Configure time after the client is wired up.
	// Use it to look up values from the AWX API and cache them in a closure.
	OnConfigure ConfigureFunc
	// ApiVersion is passed to hook functions.
	ApiVersion string
	// ResourceName is used in error messages. Defaults to TypeName if empty.
	ResourceName string
}

// GenericDataSource implements datasource.DataSource for any DataModel type.
type GenericDataSource[T any, PT DataModel[T]] struct {
	DataSourceBase
	Cfg DataSourceCfg[T]
}

func (ds *GenericDataSource[T, PT]) name() string {
	if ds.Cfg.ResourceName != "" {
		return ds.Cfg.ResourceName
	}
	return ds.TypeName
}

// Configure wires the client and then runs the optional OnConfigure hook.
// Shadows the DataSourceBase.Configure promoted method so OnConfigure actually fires.
func (ds *GenericDataSource[T, PT]) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	ds.DataSourceBase.Configure(ctx, req, resp)
	if ds.Cfg.OnConfigure == nil || ds.Client == nil {
		return
	}
	resp.Diagnostics.Append(ds.Cfg.OnConfigure(ctx, ds.Client)...)
}

func (ds *GenericDataSource[T, PT]) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ds.Cfg.Schema
}

func (ds *GenericDataSource[T, PT]) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state T
	var endpoint string

	hasSearch := len(ds.Cfg.SearchGroups) > 0

	if hasSearch {
		ep, d := EvaluateSearchGroups(ctx, req.Config, ds.Endpoint, ds.Cfg.SearchGroups)
		if DiagnosticsHasError(&resp.Diagnostics, d...) {
			return
		}
		endpoint = ep
	} else {
		endpoint = ds.Endpoint
	}

	data, d := ReadRequest(ctx, ds.Client, endpoint, ds.name())
	if DiagnosticsHasError(&resp.Diagnostics, d...) {
		return
	}

	if hasSearch {
		var err error
		data, d, err = helpers.ExtractDataIfSearchResult(data)
		resp.Diagnostics.Append(d...)
		if err != nil || resp.Diagnostics.HasError() {
			return
		}
	}

	d, err := PT(&state).UpdateFromApiData(data)
	resp.Diagnostics.Append(d...)
	if err != nil || resp.Diagnostics.HasError() {
		return
	}

	if ds.Cfg.Hook != nil {
		if HookError(&resp.Diagnostics, ds.name(), ds.Cfg.Hook(ctx, ds.Cfg.ApiVersion, hooks.SourceData, hooks.CalleeRead, nil, &state)) {
			return
		}
	}

	if DiagnosticsHasError(&resp.Diagnostics, resp.State.Set(ctx, &state)...) {
		return
	}
}
