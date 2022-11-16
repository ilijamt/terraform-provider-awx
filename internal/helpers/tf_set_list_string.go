package helpers

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AttrValueSetListString(obj *types.List, data any, trim bool) (d diag.Diagnostics, err error) {
	if obj == nil {
		err = fmt.Errorf("obj is nil")
		d.AddError(
			"nil pointer passed",
			err.Error(),
		)
		return d, err
	}

	if data == nil {
		*obj = types.ListValueMust(types.StringType, []attr.Value{})
		return nil, nil
	}

	switch data := data.(type) {
	case types.List:
		*obj = types.ListValueMust(types.StringType, data.Elements())
	case []any:
		var list []attr.Value
		for _, v := range data {
			if trim {
				list = append(list, types.StringValue(TrimAwxString(v.(string))))
			} else {
				list = append(list, types.StringValue(v.(string)))
			}
		}
		*obj = types.ListValueMust(types.StringType, list)
	case []string:
		var list []attr.Value
		for _, v := range data {
			if trim {
				list = append(list, types.StringValue(TrimAwxString(v)))
			} else {
				list = append(list, types.StringValue(v))
			}
		}
		*obj = types.ListValueMust(types.StringType, list)
	default:
		err = fmt.Errorf("failed to decode and set %v of %T type", data, data)
		d.AddError(
			fmt.Sprintf("failed to decode value of type %T for types.List", data),
			err.Error(),
		)
	}

	return d, err
}
