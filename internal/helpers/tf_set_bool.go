package helpers

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AttrValueSetBool(obj *types.Bool, data any) (d diag.Diagnostics, err error) {
	if obj == nil {
		err = fmt.Errorf("obj is nil")
		d.AddError(
			fmt.Sprintf("nil pointer passed"),
			err.Error(),
		)
		return d, err
	}

	if val, ok := data.(bool); ok {
		*obj = types.BoolValue(val)
	} else if data == nil {
		*obj = types.BoolNull()
	} else {
		err = fmt.Errorf("invalid data type: %T", data)
		d.AddError(
			fmt.Sprintf("wrong data type passed requires bool"),
			err.Error(),
		)
	}

	return d, err
}
