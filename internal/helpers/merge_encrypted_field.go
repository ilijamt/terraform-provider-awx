package helpers

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// MergeEncryptedField returns the original value when the current value is the
// AWX `$encrypted$` placeholder. AWX returns this placeholder for every secret
// field on read instead of the actual value, which would otherwise look like a
// constant drift to Terraform. When orig is null/empty (typical on import) the
// current placeholder is preserved so state matches what AWX returned.
//
// The bool return is true when the value was substituted, so callers can avoid
// allocating a new types.String when nothing changed.
func MergeEncryptedField(orig, cur types.String) (types.String, bool) {
	if !strings.Contains(cur.ValueString(), "$encrypted$") {
		return cur, false
	}
	if orig.IsNull() || orig.IsUnknown() || orig.ValueString() == "" {
		return cur, false
	}
	return orig, true
}
