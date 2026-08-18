package helpers

import "github.com/hashicorp/terraform-plugin-framework/types"

// KnownOrNullInt64 collapses an unknown to null. A write-only attribute takes
// its state from the plan, where an unset value is unknown on create but null
// after an import, and storing 0 for the import case fails the apply.
func KnownOrNullInt64(v types.Int64) types.Int64 {
	if v.IsUnknown() {
		return types.Int64Null()
	}
	return v
}
