package helpers

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SetAsStringSlice(set types.Set, trim bool) []string {
	return asStringSlice(set, trim)
}
