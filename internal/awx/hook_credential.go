package awx

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ilijamt/terraform-provider-awx/internal/hooks"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func hookCredential(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *credentialTerraformModel) (err error) {
	if source == hooks.SourceResource && (state == nil || orig == nil) && callee != hooks.CalleeDelete {
		return fmt.Errorf("state and orig required for resource")
	} else if source == hooks.SourceData && (state == nil) {
		return fmt.Errorf("state is required for data source")
	}

	if source == hooks.SourceResource && callee == hooks.CalleeCreate {
		state.Inputs = types.StringValue(orig.Inputs.ValueString())
	} else if source == hooks.SourceResource && (callee == hooks.CalleeRead || callee == hooks.CalleeUpdate) {
		if !strings.Contains(state.Inputs.ValueString(), "$encrypted$") {
			return nil
		}
		var inputs, origInputs map[string]string
		if err = json.Unmarshal([]byte(state.Inputs.ValueString()), &inputs); err != nil {
			return fmt.Errorf("%w: inputs from new state", err)
		}

		if !orig.Inputs.IsNull() {
			if err = json.Unmarshal([]byte(orig.Inputs.ValueString()), &origInputs); err != nil {
				return fmt.Errorf("%w: inputs from original state", err)
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
	}

	return nil
}
