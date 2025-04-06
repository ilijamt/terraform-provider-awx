package hooks

import (
	"context"
	"fmt"
)

var _ Hook[any] = RequireResourceStateOrOrig

func RequireResourceStateOrOrig(ctx context.Context, apiVersion string, source Source, callee Callee, orig, state any) (err error) {
	if source == SourceResource && (state == nil || orig == nil) && (callee == CalleeUpdate || callee == CalleeRead) {
		return fmt.Errorf("state or orig required for resource")
	}
	return nil
}
