package internal

import (
	"cmp"
	"slices"
)

// unorderedListKeys names OPTIONS metadata lists AWX builds from a Python set,
// so their order differs between two requests to the same server. See
// awx/api/generics.py, related_search_fields: `fields = set([])`.
var unorderedListKeys = []string{"related_search_fields"}

// volatileCredentialTypeKeys are per-install timestamps. Credential type
// generation reads only name, kind and inputs, so these change nothing except
// the diff on every download.
var volatileCredentialTypeKeys = []string{"created", "modified"}

// NormalizeResourcePayload mutates in and returns it.
func NormalizeResourcePayload(in map[string]any) map[string]any {
	for _, key := range unorderedListKeys {
		if values, ok := in[key].([]any); ok {
			sortStringList(values)
		}
	}
	return in
}

// NormalizeCredentialTypePayload mutates in and returns it. Input choices are
// set-derived too: azure_kv cloud_name reorders between installs, and nothing
// downstream reads their order.
func NormalizeCredentialTypePayload(in map[string]any) map[string]any {
	for _, key := range volatileCredentialTypeKeys {
		delete(in, key)
	}

	inputs, ok := in["inputs"].(map[string]any)
	if !ok {
		return in
	}
	fields, ok := inputs["fields"].([]any)
	if !ok {
		return in
	}
	for _, field := range fields {
		f, ok := field.(map[string]any)
		if !ok {
			continue
		}
		if choices, ok := f["choices"].([]any); ok {
			sortStringList(choices)
		}
	}
	return in
}

// sortStringList leaves values untouched unless every element is a string,
// since a mixed list has no ordering worth imposing.
func sortStringList(values []any) {
	for _, v := range values {
		if _, ok := v.(string); !ok {
			return
		}
	}
	slices.SortFunc(values, func(a, b any) int {
		return cmp.Compare(a.(string), b.(string))
	})
}
