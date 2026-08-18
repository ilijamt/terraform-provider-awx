package helpers

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type elementer interface {
	Elements() []attr.Value
}

func asStringSlice(c elementer, trim bool) []string {
	out := []string{}
	for _, val := range c.Elements() {
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

// Non-int64 elements are skipped rather than coerced, since the schema pins the
// element type and anything else is a generator bug.
func asInt64Slice(c elementer) []int64 {
	out := []int64{}
	for _, val := range c.Elements() {
		if iv, ok := val.(types.Int64); ok {
			out = append(out, iv.ValueInt64())
		}
	}
	return out
}
