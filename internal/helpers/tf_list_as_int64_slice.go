package helpers

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ListAsInt64Slice converts a Terraform types.List of ids into the []int64 an
// AWX request body expects. Non-int64 elements are skipped rather than coerced,
// since the schema pins the element type and anything else is a generator bug.
func ListAsInt64Slice(list types.List) []int64 {
	out := []int64{}
	for _, val := range list.Elements() {
		if iv, ok := val.(types.Int64); ok {
			out = append(out, iv.ValueInt64())
		}
	}
	return out
}
