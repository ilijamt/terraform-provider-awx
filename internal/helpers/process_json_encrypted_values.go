package helpers

import (
	"cmp"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ProcessJsonEncryptedValues(orig, cur types.String) (dirty bool, msg types.String, err error) {
	var curM, origM map[string]any
	var me *multierror.Error

	if e := json.Unmarshal([]byte(cur.ValueString()), &curM); e != nil {
		me = multierror.Append(me, fmt.Errorf("%w: inputs from new state", e))
	}

	if e := json.Unmarshal([]byte(orig.ValueString()), &origM); e != nil {
		me = multierror.Append(me, fmt.Errorf("%w: inputs from original state", e))
	}

	err = me.ErrorOrNil()
	if err != nil {
		return dirty, msg, me
	}

	for k, v := range curM {
		switch u := v.(type) {
		case string:
			if strings.Contains(u, "$encrypted$") {
				if _, ok := origM[k]; ok {
					dirty = true
					curM[k] = origM[k]
				} else {
					me = multierror.Append(me, fmt.Errorf("key %s not found in orig", k))
				}
			}
		}
	}

	if dirty {
		payload, _ := json.Marshal(curM)
		msg = types.StringValue(string(payload))
	}

	err = me.ErrorOrNil()
	return dirty, cmp.Or(msg, cur), err
}
