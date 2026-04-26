package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"gopkg.in/yaml.v3"
)

func AttrValueSetJsonYamlString(obj *types.String, data any, trim bool) (d diag.Diagnostics, err error) {
	if obj == nil {
		return nilObjErr()
	}

	if data == nil {
		*obj = types.StringNull()
		return d, nil
	}

	if val, ok := data.(string); ok {
		var outYaml = map[string]any{}
		var outJson = map[string]any{}
		var isYaml = yaml.Unmarshal([]byte(val), &outYaml) == nil
		var isJson = json.Unmarshal([]byte(val), &outJson) == nil

		if !isYaml && !isJson {
			err = fmt.Errorf("payload is neither json nor yaml")
			d.AddError("invalid payload", err.Error())
			return d, err
		}

		if isYaml {
			payload, mErr := json.Marshal(outYaml)
			if mErr != nil {
				return d, mErr
			}
			val = string(payload)
		}

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
