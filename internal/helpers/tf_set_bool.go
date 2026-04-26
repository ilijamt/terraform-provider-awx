package helpers

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AttrValueSetBool(obj *types.Bool, data any) (d diag.Diagnostics, err error) {
	if obj == nil {
		return nilObjErr()
	}

	if data == nil {
		*obj = types.BoolNull()
		return d, nil
	}

	val, ok := data.(bool)
	if !ok {
		err = fmt.Errorf("invalid data type: %T", data)
		d.AddError("wrong data type passed requires bool", err.Error())
		return d, err
	}
	*obj = types.BoolValue(val)
	return d, nil
}
