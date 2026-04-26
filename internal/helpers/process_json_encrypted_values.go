package helpers

// TODO: subsume the remaining callers (hook_credential.go,
// hook_notification_template.go) into the typed-credential pipeline and
// remove this helper. The typed-credential pattern handles the same
// `$encrypted$` drift problem at the field level rather than across an
// opaque JSON blob.

import (
	"cmp"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ProcessJsonEncryptedValues(orig, cur types.String) (dirty bool, msg types.String, err error) {
	var curM, origM map[string]any
	var errs []error

	if e := json.Unmarshal([]byte(cur.ValueString()), &curM); e != nil {
		errs = append(errs, fmt.Errorf("%w: inputs from new state", e))
	}
	if e := json.Unmarshal([]byte(orig.ValueString()), &origM); e != nil {
		errs = append(errs, fmt.Errorf("%w: inputs from original state", e))
	}
	if err = errors.Join(errs...); err != nil {
		return dirty, msg, err
	}

	for k, v := range curM {
		s, ok := v.(string)
		if !ok || !strings.Contains(s, "$encrypted$") {
			continue
		}
		if origVal, ok := origM[k]; ok {
			dirty = true
			curM[k] = origVal
		} else {
			errs = append(errs, fmt.Errorf("key %s not found in orig", k))
		}
	}

	if dirty {
		payload, _ := json.Marshal(curM)
		msg = types.StringValue(string(payload))
	}

	return dirty, cmp.Or(msg, cur), errors.Join(errs...)
}
