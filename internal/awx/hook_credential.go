package awx

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func hookCredential(source Source, callee Callee, orig *credentialTerraformModel, state *credentialTerraformModel) (err error) {
	if source == SourceResource && (state == nil || orig == nil) && callee != CalleeDelete {
		return fmt.Errorf("state and orig required for resource")
	} else if source == SourceData && (state == nil) {
		return fmt.Errorf("state is required for data source")
	}

	if source == SourceResource && callee == CalleeCreate {
		state.Inputs = types.StringValue(orig.Inputs.ValueString())
	} else if source == SourceResource && (callee == CalleeRead || callee == CalleeUpdate) {
		if !strings.Contains(state.Inputs.ValueString(), "$encrypted$") {
			return nil
		}
		var inputs, origInputs map[string]string
		if err = json.Unmarshal([]byte(state.Inputs.ValueString()), &inputs); err != nil {
			return err
		}
		if err = json.Unmarshal([]byte(orig.Inputs.ValueString()), &origInputs); err != nil {
			return err
		}

		for k, v := range inputs {
			if strings.Contains(v, "$encrypted$") {
				inputs[k] = origInputs[k]
				var payload []byte
				if payload, err = json.Marshal(inputs); err != nil {
					return err
				}
				state.Inputs = types.StringValue(string(payload))
				break
			}
		}
	}

	return nil
}
