package helpers

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// AttrBoolPointer returns nil for null or unknown so `,omitempty` drops the
// field. Unknown matters: an Optional+Computed bool with no default is unknown
// at plan time on create, and ValueBoolPointer would report that as false.
func AttrBoolPointer(val types.Bool) *bool {
	if val.IsNull() || val.IsUnknown() {
		return nil
	}
	return val.ValueBoolPointer()
}
