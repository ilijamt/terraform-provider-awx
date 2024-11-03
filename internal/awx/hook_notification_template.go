package awx

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ilijamt/terraform-provider-awx/internal/helpers"
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
		if dirty, msg, err := helpers.ProcessJsonEncryptedValues(orig.NotificationConfiguration, state.NotificationConfiguration); err != nil {
			return err
		} else if dirty {
			state.NotificationConfiguration = msg
		}
	}

	return nil
}
