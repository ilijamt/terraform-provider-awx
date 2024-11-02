package awx

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
)

func hookNotificationTemplate(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *notificationTemplateTerraformModel) (err error) {
	if source == hooks.SourceResource && (state == nil || orig == nil) && (callee == hooks.CalleeUpdate || callee == hooks.CalleeRead) {
		return fmt.Errorf("state and orig required for resource")
	}

	if source == hooks.SourceResource && callee == hooks.CalleeCreate {
		state.NotificationConfiguration = types.StringValue(orig.NotificationConfiguration.ValueString())
	} else if source == hooks.SourceResource && (callee == hooks.CalleeUpdate || callee == hooks.CalleeRead) {
		if !strings.Contains(state.NotificationConfiguration.ValueString(), "$encrypted$") {
			return nil
		}

		var inputs map[string]any
		if err = json.Unmarshal([]byte(state.NotificationConfiguration.ValueString()), &inputs); err != nil {
			return fmt.Errorf("%w: inputs from new state", err)
		}

		var origInputs map[string]any
		if err = json.Unmarshal([]byte(orig.NotificationConfiguration.ValueString()), &origInputs); err != nil {
			return fmt.Errorf("%w: inputs from original state", err)
		}
		var dirty = false

		for k, v := range inputs {
			switch v.(type) {
			case string:
				if strings.Contains(v.(string), "$encrypted$") {
					dirty = true
					inputs[k] = origInputs[k]
				}
			}
		}

		if dirty {
			var payload []byte
			if payload, err = json.Marshal(inputs); err != nil {
				return err
			}
			state.NotificationConfiguration = types.StringValue(string(payload))
		}
	}

	return nil
}
