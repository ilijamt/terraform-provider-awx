package helpers

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AttrValueSetListInt64(obj *types.List, data any) (diag.Diagnostics, error) {
	return setListValue(obj, data, types.Int64Type, int64Element)
}
