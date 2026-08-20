package helpers

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Non-string elements are skipped rather than coerced, since the schema pins
// the element type and anything else is a generator bug.
func MapAsStringMap(m types.Map, trim bool) map[string]string {
	out := map[string]string{}
	for k, val := range m.Elements() {
		sv, ok := val.(types.String)
		if !ok {
			continue
		}
		s := sv.ValueString()
		if trim {
			s = TrimAwxString(s)
		}
		out[k] = s
	}
	return out
}
