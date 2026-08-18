package helpers

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AttrValueSetSetString(obj *types.Set, data any, trim bool) (d diag.Diagnostics, err error) {
	if obj == nil {
		return nilObjErr()
	}

	if data == nil {
		*obj = types.SetValueMust(types.StringType, []attr.Value{})
		return nil, nil
	}

	maybeTrim := func(s string) string {
		if trim {
			return TrimAwxString(s)
		}
		return s
	}

	switch data := data.(type) {
	case types.Set:
		*obj = types.SetValueMust(types.StringType, data.Elements())
	case []any:
		set := make([]attr.Value, 0, len(data))
		for _, v := range data {
			s, ok := v.(string)
			if !ok {
				err = fmt.Errorf("failed to decode %v of %T type as string", v, v)
				d.AddError(fmt.Sprintf("failed to decode set element of type %T for types.Set", v), err.Error())
				return d, err
			}
			set = append(set, types.StringValue(maybeTrim(s)))
		}
		*obj = types.SetValueMust(types.StringType, set)
	case []string:
		set := make([]attr.Value, 0, len(data))
		for _, v := range data {
			set = append(set, types.StringValue(maybeTrim(v)))
		}
		*obj = types.SetValueMust(types.StringType, set)
	default:
		err = fmt.Errorf("failed to decode and set %v of %T type", data, data)
		d.AddError(
			fmt.Sprintf("failed to decode value of type %T for types.Set", data),
			err.Error(),
		)
	}

	return d, err
}
