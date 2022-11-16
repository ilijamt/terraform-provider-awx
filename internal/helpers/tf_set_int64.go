package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AttrValueSetInt64(obj *types.Int64, data any) (d diag.Diagnostics, err error) {
	if obj == nil {
		err = fmt.Errorf("obj is nil")
		d.AddError(
			fmt.Sprintf("nil pointer passed"),
			err.Error(),
		)
		return d, err
	}

	if val, ok := data.(json.Number); ok {
		v, err := val.Int64()
		if err != nil {
			d.AddError(
				fmt.Sprintf("failed to convert %v to int64", val),
				err.Error(),
			)
			return d, err
		}
		*obj = types.Int64Value(v)
	} else if val, ok := data.(int64); ok {
		*obj = types.Int64Value(val)
	} else if val, ok := data.(int); ok {
		*obj = types.Int64Value(int64(val))
	} else if data == nil {
		*obj = types.Int64Null()
	} else {
		err = fmt.Errorf("invalid data type: %T", data)
		d.AddError(
			fmt.Sprintf("wrong data type passed requires json.Number, int64, int"),
			err.Error(),
		)
	}

	return d, err
}
