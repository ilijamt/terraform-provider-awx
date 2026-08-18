package helpers

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SetAsStringSlice is the types.Set counterpart of ListAsStringSlice.
func SetAsStringSlice(set types.Set, trim bool) []string {
	out := []string{}
	for _, val := range set.Elements() {
		var s string
		if sv, ok := val.(types.String); ok {
			s = sv.ValueString()
		} else {
			s = val.String()
		}
		if trim {
			s = TrimAwxString(s)
		}
		out = append(out, s)
	}
	return out
}
