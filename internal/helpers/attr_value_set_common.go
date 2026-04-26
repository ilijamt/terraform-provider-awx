package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// nilObjErr is the shared error path for AttrValueSet* helpers when the caller
// passes a nil destination pointer. Centralised so the diagnostic wording
// changes in one place.
func nilObjErr() (diag.Diagnostics, error) {
	var d diag.Diagnostics
	err := fmt.Errorf("obj is nil")
	d.AddError("nil pointer passed", err.Error())
	return d, err
}

// coerceInt64 normalises AWX-shaped numeric input. AWX always returns
// json.Number on the wire; int/int64 fallbacks cover values that callers may
// pass in directly. ok=false means the type wasn't recognised.
func coerceInt64(data any) (val int64, ok bool, err error) {
	switch v := data.(type) {
	case json.Number:
		val, err = v.Int64()
		return val, true, err
	case int64:
		return v, true, nil
	case int:
		return int64(v), true, nil
	}
	return 0, false, nil
}

func coerceFloat64(data any) (val float64, ok bool, err error) {
	switch v := data.(type) {
	case json.Number:
		val, err = v.Float64()
		return val, true, err
	case float64:
		return v, true, nil
	case float32:
		return float64(v), true, nil
	}
	return 0, false, nil
}
