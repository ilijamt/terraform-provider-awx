package awx

import (
	"context"
	"fmt"
)

func hookWorkflowJobTemplate(ctx context.Context, apiVersion string, source Source, callee Callee, orig, state *workflowJobTemplateTerraformModel) (err error) {
	if source == SourceResource && (state == nil || orig == nil) && (callee == CalleeUpdate || callee == CalleeRead) {
		return fmt.Errorf("state and orig required for resource")
	}
	return nil
}
