package helpers

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AttrValueSetInt64(obj *types.Int64, data any) (d diag.Diagnostics, err error) {
	if obj == nil {
		return nilObjErr()
	}

	if data == nil {
		*obj = types.Int64Null()
		return d, nil
	}

	val, ok, err := coerceInt64(data)
	if err != nil {
		d.AddError(fmt.Sprintf("failed to convert %v to int64", data), err.Error())
		return d, err
	}
	if !ok {
		err = fmt.Errorf("invalid data type: %T", data)
		d.AddError("wrong data type passed requires json.Number, int64, int", err.Error())
		return d, err
	}
	*obj = types.Int64Value(val)
	return d, nil
}
