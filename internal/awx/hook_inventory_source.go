package awx

import (
	"context"
	"fmt"

	"github.com/ilijamt/terraform-provider-awx/internal/hooks"
)

func hookInventorySource(ctx context.Context, apiVersion string, source hooks.Source, callee hooks.Callee, orig, state *inventorySourceTerraformModel) (err error) {
	if source == hooks.SourceResource && (state == nil || orig == nil) && (callee == hooks.CalleeUpdate || callee == hooks.CalleeRead) {
		return fmt.Errorf("state and orig required for resource")
	}
	return nil
}
