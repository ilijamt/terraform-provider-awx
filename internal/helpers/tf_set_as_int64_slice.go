package helpers

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SetAsInt64Slice(set types.Set) []int64 {
	return asInt64Slice(set)
}
