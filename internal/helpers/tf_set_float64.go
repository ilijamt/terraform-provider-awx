package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AttrValueSetFloat64(obj *types.Float64, data any) (d diag.Diagnostics, err error) {
	d = make(diag.Diagnostics, 0)
	if obj == nil {
		err = fmt.Errorf("obj is nil")
		d.AddError(
			"nil pointer passed",
			err.Error(),
		)
		return d, err
	}

	if data == nil {
		*obj = types.Float64Null()
	} else if val, ok := data.(json.Number); ok {
		v, err := val.Float64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to float64", val),
				err.Error(),
			)
			return d, err
		}
		*obj = types.Float64Value(v)
	} else if val, ok := data.(float64); ok {
		*obj = types.Float64Value(val)
	} else if val, ok := data.(float32); ok {
		*obj = types.Float64Value(float64(val))
	} else {
		err = fmt.Errorf("invalid data type: %T", data)
		d.AddError(
			"wrong data type passed requires json.Number, float64, float32",
			err.Error(),
		)
	}

	return d, err
}
