package awx_21_8_0

import (
	"context"
	"fmt"

	"github.com/ilijamt/terraform-provider-awx/internal/hooks"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func hookApplication(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *applicationTerraformModel) (err error) {
	if source == hooks.SourceResource && (state == nil || orig == nil) && (callee == hooks.CalleeUpdate || callee == hooks.CalleeRead) {
		return fmt.Errorf("state and orig required for resource")
	}

	if state.ClientType.Equal(types.StringValue("confidential")) &&
		source == hooks.SourceResource &&
		(callee == hooks.CalleeUpdate || callee == hooks.CalleeRead) &&
		!orig.ClientType.IsNull() {
		// copy over the `client secret` as it's only shown during creation and never again
		state.ClientSecret = orig.ClientSecret
	}
	return nil
}
