package helpers

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ListAsStringSlice(list types.List, trim bool) []string {
	return asStringSlice(list, trim)
}
