package internal

import (
	"cmp"
	"slices"
)

// unorderedListKeys names OPTIONS metadata lists AWX assembles from a Python
// set, so the JSON order differs between two requests against the same server.
// Left alone they rewrite every payload on each download-api run and bury the
// real API drift. See awx/api/generics.py, related_search_fields: `fields = set([])`.
var unorderedListKeys = []string{"related_search_fields"}

// volatileCredentialTypeKeys are per-install timestamps on managed credential
// types. GenerateCredentialTypeTfDefinition reads only name/kind/inputs, so
// keeping these just churns 30 payloads whenever the source AWX is reinstalled.
var volatileCredentialTypeKeys = []string{"created", "modified"}

// NormalizeResourcePayload makes a fetched OPTIONS payload byte-stable across
// downloads. Mutates and returns in.
func NormalizeResourcePayload(in map[string]any) map[string]any {
	for _, key := range unorderedListKeys {
		if values, ok := in[key].([]any); ok {
			sortStringList(values)
		}
	}
	return in
}

// NormalizeCredentialTypePayload strips per-install fields from a managed
// credential type and sorts its input choices, which AWX also derives from sets
// (azure_kv cloud_name reorders between installs). Nothing downstream reads the
// choice order — GenerateCredentialTypeTfDefinition only consumes name, kind
// and the input ids/types. Mutates and returns in.
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

// sortStringList sorts values in place, leaving it untouched unless every
// element is a string — a mixed list has no meaningful ordering to impose.
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
