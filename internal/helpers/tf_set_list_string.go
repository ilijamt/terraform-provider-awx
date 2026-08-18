package helpers

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AttrValueSetListString(obj *types.List, data any, trim bool) (diag.Diagnostics, error) {
	return setListValue(obj, data, types.StringType, stringElement(trim))
}
