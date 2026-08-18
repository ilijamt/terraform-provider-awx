package helpers

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AttrValueSetSetInt64(obj *types.Set, data any) (diag.Diagnostics, error) {
	return setSetValue(obj, data, types.Int64Type, int64Element)
}
