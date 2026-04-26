package helpers

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AttrValueSetFloat64(obj *types.Float64, data any) (d diag.Diagnostics, err error) {
	if obj == nil {
		return nilObjErr()
	}

	if data == nil {
		*obj = types.Float64Null()
		return d, nil
	}

	val, ok, err := coerceFloat64(data)
	if err != nil {
		d.AddError(fmt.Sprintf("failed to convert %v to float64", data), err.Error())
		return d, err
	}
	if !ok {
		err = fmt.Errorf("invalid data type: %T", data)
		d.AddError("wrong data type passed requires json.Number, float64, float32", err.Error())
		return d, err
	}
	*obj = types.Float64Value(val)
	return d, nil
}
