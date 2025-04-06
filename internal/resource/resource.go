package resource

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
)

// Updater is an interface for resources that can be updated with API data.
// Implementations should update their internal state based on the provided data.
type Updater interface {
	// UpdateWithApiData updates the resource with data received from an API.
	// The data parameter contains key-value pairs representing the resource's properties.
	// It returns diagnostics that may contain warnings or errors encountered during the update,
	// as well as any error that occurred during the update process.
	UpdateWithApiData(data map[string]any) (diags diag.Diagnostics, err error)
}

// Cloner is a generic interface for resources that can be cloned.
// The type parameter T represents the type of the resource being cloned.
type Cloner[T any] interface {
	// Clone creates and returns a deep copy of the resource.
	// The returned value is of the same type as the resource being cloned.
	Clone() T
}

// Body is a generic interface for resources that can provide a JSON request body.
type Body interface {
	json.Marshaler
}

// CallInfo contains information about a resource API call.
// It provides details about the resource name, endpoint, and type.
type CallInfo struct {
	// Name is the identifier of the resource.
	Name string `json:"name"`
	// Endpoint is the API endpoint URL for the resource.
	Endpoint string `json:"endpoint"`
	// TypeName is the type classification of the resource.
	TypeName string `json:"type_name"`
	// Source from where the action came from
	Source hooks.Source `json:"action"`
	// Callee is who called the action
	Callee hooks.Callee `json:"callee"`
}

// With creates a new CallInfo with the specified hooks.Source and hooks.Callee
func (r CallInfo) With(source hooks.Source, callee hooks.Callee) CallInfo {
	return CallInfo{
		Name:     r.Name,
		Endpoint: r.Endpoint,
		TypeName: r.TypeName,
		Source:   source,
		Callee:   callee,
	}
}
