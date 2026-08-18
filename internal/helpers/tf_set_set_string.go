package helpers

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AttrValueSetSetString(obj *types.Set, data any, trim bool) (diag.Diagnostics, error) {
	return setSetValue(obj, data, types.StringType, stringElement(trim))
}
