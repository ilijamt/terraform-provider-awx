package helpers

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AttrValueSetListInt64(obj *types.List, data any) (d diag.Diagnostics, err error) {
	if obj == nil {
		return nilObjErr()
	}

	if data == nil {
		*obj = types.ListValueMust(types.Int64Type, []attr.Value{})
		return nil, nil
	}

	switch data := data.(type) {
	case types.List:
		*obj = types.ListValueMust(types.Int64Type, data.Elements())
	case []any:
		list := make([]attr.Value, 0, len(data))
		for _, v := range data {
			val, ok, convErr := coerceInt64(v)
			if convErr != nil || !ok {
				err = fmt.Errorf("failed to decode %v of %T type as int64", v, v)
				d.AddError(fmt.Sprintf("failed to decode list element of type %T for types.List", v), err.Error())
				return d, err
			}
			list = append(list, types.Int64Value(val))
		}
		*obj = types.ListValueMust(types.Int64Type, list)
	case []int64:
		list := make([]attr.Value, 0, len(data))
		for _, v := range data {
			list = append(list, types.Int64Value(v))
		}
		*obj = types.ListValueMust(types.Int64Type, list)
	default:
		err = fmt.Errorf("failed to decode and set %v of %T type", data, data)
		d.AddError(
			fmt.Sprintf("failed to decode value of type %T for types.List", data),
			err.Error(),
		)
	}

	return d, err
}
