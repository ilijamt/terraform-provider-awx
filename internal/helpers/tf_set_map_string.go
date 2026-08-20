package helpers

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AttrValueSetMapString(obj *types.Map, data any, trim bool) (diag.Diagnostics, error) {
	var d diag.Diagnostics
	if obj == nil {
		return nilObjErr()
	}

	if data == nil {
		*obj = types.MapValueMust(types.StringType, map[string]attr.Value{})
		return d, nil
	}

	// Not MapValueMust: a wrong element type belongs in a diagnostic, not a panic.
	if v, ok := data.(types.Map); ok {
		val, d := types.MapValue(types.StringType, v.Elements())
		if d.HasError() {
			return d, fmt.Errorf("failed to set %v as types.Map", data)
		}
		*obj = val
		return d, nil
	}

	raw, ok := data.(map[string]any)
	if !ok {
		err := fmt.Errorf("failed to decode and set %v of %T type", data, data)
		d.AddError(fmt.Sprintf("failed to decode value of type %T for types.Map", data), err.Error())
		return d, err
	}

	var decode = stringElement(trim)
	out := make(map[string]attr.Value, len(raw))
	for k, v := range raw {
		val, ok, convErr := decode(v)
		if !ok || convErr != nil {
			err := fmt.Errorf("failed to decode %v of %T type for types.Map", v, v)
			d.AddError(fmt.Sprintf("failed to decode element of type %T for types.Map", v), err.Error())
			return d, err
		}
		out[k] = val
	}
	*obj = types.MapValueMust(types.StringType, out)
	return d, nil
}
