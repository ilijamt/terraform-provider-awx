package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AttrValueSetString(obj *types.String, data any) (d diag.Diagnostics, err error) {
	if obj == nil {
		err = fmt.Errorf("obj is nil")
		d.AddError(
			fmt.Sprintf("nil pointer passed"),
			err.Error(),
		)
		return d, err
	}

	if data == nil {
		*obj = types.StringNull()
		return d, nil
	}

	switch data.(type) {
	case string:
		*obj = types.StringValue(TrimString(false, false, data.(string)))
	case json.Number:
		*obj = types.StringValue(data.(json.Number).String())
	default:
		err = fmt.Errorf("invalid data type: %T", data)
		d.AddError(
			fmt.Sprintf("wrong data type passed requires string, json.Number, map[string]any, []any"),
			err.Error(),
		)
	}

	return d, err
}
