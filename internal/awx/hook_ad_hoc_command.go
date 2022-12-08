package awx

import (
	"context"
	"fmt"
)

func hookAdHocCommand(ctx context.Context, apiVersion string, source Source, callee Callee, orig, state *adHocCommandTerraformModel) (err error) {
	if source == SourceResource && (state == nil || orig == nil) && (callee == CalleeUpdate || callee == CalleeRead) {
		return fmt.Errorf("state and orig required for resource")
	}
	return nil
}
