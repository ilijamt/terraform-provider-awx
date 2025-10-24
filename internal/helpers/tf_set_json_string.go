package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AttrValueSetJsonString(obj *types.String, data any, trim bool) (d diag.Diagnostics, err error) {
	d = make(diag.Diagnostics, 0)
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

	if val, ok := data.(string); ok {
		if trim {
			*obj = types.StringValue(TrimAwxString(val))
		} else {
			*obj = types.StringValue(val)
		}
	} else {
		var v []byte
		if v, err = json.Marshal(data); err == nil {
			*obj = types.StringValue(TrimAwxString(string(v)))
		}
	}

	return d, err
}
