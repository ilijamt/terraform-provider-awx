package helpers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

func RequiresReplaceIfPropertyUnsetInt64(ctx context.Context, ir planmodifier.Int64Request, rrifr *int64planmodifier.RequiresReplaceIfFuncResponse) {
	var requiresReplace bool
	var noInitialState = ir.StateValue.IsNull() || ir.StateValue.IsUnknown()
	if !noInitialState {
		requiresReplace = !ir.StateValue.IsNull() || !ir.StateValue.IsUnknown()
	}
	rrifr.RequiresReplace = requiresReplace
}
