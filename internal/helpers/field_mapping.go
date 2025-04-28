package helpers

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// FieldSetter is a function type for setting model fields
type FieldSetter func(any) (diag.Diagnostics, error)

// FieldMapping defines a mapping between API field names and their setter functions
type FieldMapping struct {
	APIField string
	Setter   FieldSetter
	Data     map[string]any
}

// ApplyFieldMappings applies a list of field mappings to a data map
func ApplyFieldMappings(data map[string]any, mappings ...FieldMapping) (diags diag.Diagnostics, err error) {
	diags = make(diag.Diagnostics, 0)
	if data == nil {
		err = fmt.Errorf("data must not be nil")
		diags.AddError("nil pointer", err.Error())
		return diags, err
	}
	var me = new(multierror.Error)
	for _, mapping := range mappings {
		var fieldData = data
		if mapping.Data != nil {
			fieldData = mapping.Data
		}
		var d diag.Diagnostics
		var e error
		if d, e = mapping.Setter(fieldData[mapping.APIField]); e != nil {
			err = multierror.Append(me, e)
		}
		diags.Append(d...)
	}

	return diags, me.ErrorOrNil()
}
