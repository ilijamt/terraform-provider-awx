package helpers

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
)

func IsEmptyValue(val attr.Value) bool {
	return val.IsNull() || val.IsUnknown() || val.String() == ""
}
