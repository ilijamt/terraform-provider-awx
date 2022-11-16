package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AttrValueSetJsonString(obj *types.String, data any) (d diag.Diagnostics, err error) {
	if obj == nil {
		err = fmt.Errorf("obj is nil")
		d.AddError(
			fmt.Sprintf("nil pointer passed"),
			err.Error(),
		)
		return d, err
	}

	if data == nil {
		*obj = types.StringNull()
		return d, nil
	}

	var v []byte
	v, err = json.Marshal(data)
	*obj = types.StringValue(string(v))

	return d, err
}
