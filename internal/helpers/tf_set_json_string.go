package helpers

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AttrValueSetJsonString(obj *types.String, data any, trim bool) (d diag.Diagnostics, err error) {
	if obj == nil {
		return nilObjErr()
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
		return d, nil
	}

	v, err := json.Marshal(data)
	if err == nil {
		*obj = types.StringValue(TrimAwxString(string(v)))
	}
	return d, err
}
