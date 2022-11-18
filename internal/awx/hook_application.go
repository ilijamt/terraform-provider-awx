package awx

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func hookApplication(ctx context.Context, source Source, callee Callee, orig *applicationTerraformModel, state *applicationTerraformModel) (err error) {
	if source == SourceResource && (state == nil || orig == nil) && (callee == CalleeUpdate || callee == CalleeRead) {
		return fmt.Errorf("state and orig required for resource")
	}

	if state.ClientType.Equal(types.StringValue("confidential")) &&
		source == SourceResource &&
		(callee == CalleeUpdate || callee == CalleeRead) &&
		!orig.ClientType.IsNull() {
		// copy over the `client secret` as it's only shown during creation and never again
		state.ClientSecret = orig.ClientSecret
	}
	return nil
}
