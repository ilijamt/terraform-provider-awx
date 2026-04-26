package helpers

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AttrValueSetListString(obj *types.List, data any, trim bool) (d diag.Diagnostics, err error) {
	if obj == nil {
		return nilObjErr()
	}

	if data == nil {
		*obj = types.ListValueMust(types.StringType, []attr.Value{})
		return nil, nil
	}

	maybeTrim := func(s string) string {
		if trim {
			return TrimAwxString(s)
		}
		return s
	}

	switch data := data.(type) {
	case types.List:
		*obj = types.ListValueMust(types.StringType, data.Elements())
	case []any:
		list := make([]attr.Value, 0, len(data))
		for _, v := range data {
			list = append(list, types.StringValue(maybeTrim(v.(string))))
		}
		*obj = types.ListValueMust(types.StringType, list)
	case []string:
		list := make([]attr.Value, 0, len(data))
		for _, v := range data {
			list = append(list, types.StringValue(maybeTrim(v)))
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
