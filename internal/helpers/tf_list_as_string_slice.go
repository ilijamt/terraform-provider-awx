package helpers

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ListAsStringSlice converts a Terraform types.List into a []string suitable
// for an AWX request body. Elements that are types.String are unwrapped via
// ValueString(); anything else falls back to its String() form. When trim is
// true, each value is normalized through TrimAwxString.
func ListAsStringSlice(list types.List, trim bool) []string {
	out := []string{}
	for _, val := range list.Elements() {
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
