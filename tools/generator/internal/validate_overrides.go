package internal

import (
	"fmt"
	"sort"
)

// ValidateOverrides reports config that cannot affect the output: an override
// naming a field the API never reports, or a knob the templates ignore for that
// field's type. Dead config reads as intent and survives review, which is how
// host.json kept two removals that never removed anything.
//
// Runs before the remove_fields_* pruning, since that mutates the actions.
func ValidateOverrides(val Item, objmap map[string]any) []string {
	actions, _ := objmap["actions"].(map[string]any)
	rawRead, _ := actions[val.ApiPropertyDataKey].(map[string]any)
	rawWrite, _ := actions[val.ApiPropertyResourceKey].(map[string]any)
	if rawRead == nil && rawWrite == nil {
		return nil
	}

	// The removal lists are judged against what the API reported, but the
	// property knobs are judged against what the templates will see: after the
	// removals, and with the api_data_override_* injections in place.
	read := effectiveFields(rawRead, val.RemoveFieldsDataSource, val.ApiDataOverride)
	write := effectiveFields(rawWrite, val.RemoveFieldsResource, val.ApiDataOverride, val.ApiDataOverrideResource)

	fieldType := func(fields map[string]any, key string) (string, bool) {
		f, ok := fields[key].(map[string]any)
		if !ok {
			return "", false
		}
		t, _ := f["type"].(string)
		return t, true
	}
	// setDefaultValue clears required once a field carries a default, so the
	// payload's own flag is not what the templates end up seeing.
	isRequired := func(fields map[string]any, key string) bool {
		f, ok := fields[key].(map[string]any)
		if !ok {
			return false
		}
		if d, ok := f["default"]; ok && d != nil {
			return false
		}
		req, _ := f["required"].(bool)
		return req
	}

	var problems []string
	for key, override := range val.PropertyOverrides {
		writeType, inWrite := fieldType(write, key)
		_, inRead := fieldType(read, key)
		if !inWrite && !inRead {
			problems = append(problems, fmt.Sprintf("property_overrides %q names a field absent from both actions", key))
			continue
		}

		required := isRequired(write, key)
		if override.Nullable != nil && *override.Nullable && writeType != "boolean" {
			problems = append(problems, fmt.Sprintf("property_overrides %q sets nullable on a %s; only a bool becomes a pointer", key, cmpType(writeType, inWrite)))
		}
		if override.OmitEmpty != nil && !*override.OmitEmpty && (required || writeType == "boolean") {
			problems = append(problems, fmt.Sprintf("property_overrides %q sets omit_empty on a field that never carries it", key))
		}
		if override.UseStateForUnknown != nil && !*override.UseStateForUnknown && required {
			problems = append(problems, fmt.Sprintf("property_overrides %q turns off use_state_for_unknown on a required field, which never has it", key))
		}
	}

	for _, key := range val.RemoveFieldsResource {
		if _, ok := rawWrite[key]; !ok {
			problems = append(problems, fmt.Sprintf("remove_fields_resource %q is not in the %s action", key, val.ApiPropertyResourceKey))
		}
	}
	for _, key := range val.RemoveFieldsDataSource {
		if _, ok := rawRead[key]; !ok {
			problems = append(problems, fmt.Sprintf("remove_fields_data_source %q is not in the %s action", key, val.ApiPropertyDataKey))
		}
	}

	sort.Strings(problems)
	return problems
}

func cmpType(t string, known bool) string {
	if !known || t == "" {
		return "read-only field"
	}
	return t
}

func effectiveFields(action map[string]any, removed []string, injected ...map[string]map[string]any) map[string]any {
	out := make(map[string]any, len(action))
	for k, v := range action {
		out[k] = v
	}
	for _, k := range removed {
		delete(out, k)
	}
	for _, set := range injected {
		for k, v := range set {
			field := map[string]any{}
			if existing, ok := out[k].(map[string]any); ok {
				for ek, ev := range existing {
					field[ek] = ev
				}
			}
			for ik, iv := range v {
				field[ik] = iv
			}
			out[k] = field
		}
	}
	return out
}
