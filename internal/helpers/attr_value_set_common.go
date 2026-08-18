package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

// coerceFloat64 also accepts a string, because DRF renders a DecimalField that
// way by default: AWX sends instance capacity_adjustment as "1.00" and ad hoc
// command elapsed as "0.000".
func coerceFloat64(data any) (val float64, ok bool, err error) {
	switch v := data.(type) {
	case json.Number:
		val, err = v.Float64()
		return val, true, err
	case float64:
		return v, true, nil
	case float32:
		return float64(v), true, nil
	case string:
		val, err = strconv.ParseFloat(v, 64)
		return val, true, err
	}
	return 0, false, nil
}

type elementDecoder func(v any) (val attr.Value, ok bool, err error)

func stringElement(trim bool) elementDecoder {
	return func(v any) (attr.Value, bool, error) {
		s, ok := v.(string)
		if !ok {
			return nil, false, nil
		}
		if trim {
			s = TrimAwxString(s)
		}
		return types.StringValue(s), true, nil
	}
}

func int64Element(v any) (attr.Value, bool, error) {
	n, ok, err := coerceInt64(v)
	if !ok || err != nil {
		return nil, ok, err
	}
	return types.Int64Value(n), true, nil
}

func decodeElements(data any, container string, decode elementDecoder) ([]attr.Value, diag.Diagnostics, error) {
	var d diag.Diagnostics

	var raw []any
	switch data := data.(type) {
	case []any:
		raw = data
	case []string:
		raw = make([]any, len(data))
		for i, v := range data {
			raw[i] = v
		}
	case []int64:
		raw = make([]any, len(data))
		for i, v := range data {
			raw[i] = v
		}
	default:
		err := fmt.Errorf("failed to decode and set %v of %T type", data, data)
		d.AddError(fmt.Sprintf("failed to decode value of type %T for %s", data, container), err.Error())
		return nil, d, err
	}

	out := make([]attr.Value, 0, len(raw))
	for _, v := range raw {
		val, ok, convErr := decode(v)
		if !ok || convErr != nil {
			err := fmt.Errorf("failed to decode %v of %T type for %s", v, v, container)
			d.AddError(fmt.Sprintf("failed to decode element of type %T for %s", v, container), err.Error())
			return nil, d, err
		}
		out = append(out, val)
	}
	return out, d, nil
}

func setListValue(obj *types.List, data any, elemType attr.Type, decode elementDecoder) (diag.Diagnostics, error) {
	if obj == nil {
		return nilObjErr()
	}
	if data == nil {
		*obj = types.ListValueMust(elemType, []attr.Value{})
		return nil, nil
	}
	// Not ListValueMust: a caller can still pass the wrong element type, and that
	// belongs in a diagnostic rather than a panic.
	if v, ok := data.(types.List); ok {
		val, d := types.ListValue(elemType, v.Elements())
		if d.HasError() {
			return d, fmt.Errorf("failed to set %v as types.List", data)
		}
		*obj = val
		return d, nil
	}

	vals, d, err := decodeElements(data, "types.List", decode)
	if err != nil {
		return d, err
	}
	*obj = types.ListValueMust(elemType, vals)
	return d, nil
}

func setSetValue(obj *types.Set, data any, elemType attr.Type, decode elementDecoder) (diag.Diagnostics, error) {
	if obj == nil {
		return nilObjErr()
	}
	if data == nil {
		*obj = types.SetValueMust(elemType, []attr.Value{})
		return nil, nil
	}
	if v, ok := data.(types.Set); ok {
		val, d := types.SetValue(elemType, v.Elements())
		if d.HasError() {
			return d, fmt.Errorf("failed to set %v as types.Set", data)
		}
		*obj = val
		return d, nil
	}

	vals, d, err := decodeElements(data, "types.Set", decode)
	if err != nil {
		return d, err
	}
	*obj = types.SetValueMust(elemType, vals)
	return d, nil
}
