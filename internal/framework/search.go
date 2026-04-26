package framework

import (
	"context"
	"fmt"
	"net/url"
	p "path"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SearchField describes one field in a search group.
type SearchField struct {
	Name      string // Attribute name in the Terraform schema (e.g., "id", "name").
	Type      string // "int64" or "string".
	URLEscape bool   // Whether to apply url.PathEscape (typically true for strings).
}

// SearchGroup describes a named search group with its URL pattern.
type SearchGroup struct {
	Name      string        // Group identifier (e.g., "by_id", "by_name").
	URLSuffix string        // fmt pattern appended to the endpoint (e.g., "%d/", "?name__exact=%s").
	Fields    []SearchField // Attributes that must all be set for this group to match.
}

// EvaluateSearchGroups checks which search group is active based on config attributes
// and returns the resolved endpoint. Returns error diagnostics if no group matches.
func EvaluateSearchGroups(ctx context.Context, config tfsdk.Config, baseEndpoint string, groups []SearchGroup) (string, diag.Diagnostics) {
	var diags diag.Diagnostics

	for _, group := range groups {
		params := []any{baseEndpoint}
		allSet := true

		for _, field := range group.Fields {
			var v attr.Value
			switch field.Type {
			case "int64":
				var tv types.Int64
				config.GetAttribute(ctx, path.Root(field.Name), &tv)
				v = tv
				if !tv.IsNull() && !tv.IsUnknown() {
					params = append(params, tv.ValueInt64())
				}
			case "string":
				var tv types.String
				config.GetAttribute(ctx, path.Root(field.Name), &tv)
				v = tv
				if !tv.IsNull() && !tv.IsUnknown() {
					val := tv.ValueString()
					if field.URLEscape {
						val = url.PathEscape(val)
					}
					params = append(params, val)
				}
			}

			if v == nil || v.IsNull() || v.IsUnknown() {
				allSet = false
				break
			}
		}

		if allSet {
			endpoint := p.Clean(fmt.Sprintf("%s/"+group.URLSuffix, params...))
			return endpoint, diags
		}
	}

	diags.AddError(
		"missing configuration for one of the predefined search groups",
		"",
	)
	return "", diags
}
