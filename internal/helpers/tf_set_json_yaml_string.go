package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"gopkg.in/yaml.v3"
)

func AttrValueSetJsonYamlString(obj *types.String, data any, trim bool) (d diag.Diagnostics, err error) {
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
		var outYaml = map[string]any{}
		var outJson = map[string]any{}
		var isYaml = yaml.Unmarshal([]byte(val), &outYaml) == nil
		var isJson = json.Unmarshal([]byte(val), &outJson) == nil

		if !isYaml && !isJson {
			err = fmt.Errorf("payload is neither json nor yaml")
			d.AddError(
				"invalid payload",
				err.Error(),
			)
			return d, err
		}

		if isYaml {
			// convert it to JSON from YAML
			var payload []byte
			payload, err = json.Marshal(outYaml)
			val = string(payload)
		}

		if trim {
			*obj = types.StringValue(TrimAwxString(val))
		} else {
			*obj = types.StringValue(val)
		}
	} else {
		// nothing to do here as we are already in a structure
		// converting the structure would give it to us in the proper format
		var v []byte
		if v, err = json.Marshal(data); err == nil {
			*obj = types.StringValue(TrimAwxString(string(v)))
		}
	}

	return d, err
}
