package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AttrValueSetString(obj *types.String, data any, trim bool) (d diag.Diagnostics, err error) {
	if obj == nil {
		err = fmt.Errorf("obj is nil")
		d.AddError(
			"nil pointer passed",
			err.Error(),
		)
		return d, err
	}

	if data == nil {
		*obj = types.StringNull()
		return d, nil
	}

	switch data := data.(type) {
	case string:
		if trim {
			data = TrimAwxString(data)
		}
		*obj = types.StringValue(data)
	case json.Number:
		*obj = types.StringValue(data.String())
	default:
		err = fmt.Errorf("invalid data type: %T", data)
		d.AddError(
			"wrong data type passed requires string, json.Number, map[string]any, []any",
			err.Error(),
		)
	}

	return d, err
}
