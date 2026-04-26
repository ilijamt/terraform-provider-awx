package helpers

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func ExtractDataIfSearchResult(in map[string]any) (out map[string]any, d diag.Diagnostics, err error) {
	raw, ok := in["count"]
	if !ok {
		return in, d, nil
	}

	count, err := parseSearchCount(raw)
	if err != nil {
		d.AddError("Failed to convert count number in search result", err.Error())
		return nil, d, err
	}

	if count > 1 {
		err = fmt.Errorf("received %d entries, expected 1", count)
		d.AddError("More than one entry present, please refine your query", err.Error())
		return nil, d, err
	}

	results, ok := in["results"].([]any)
	if !ok {
		err = fmt.Errorf("received: %T instead of []any", in["results"])
		d.AddError("Unexpected format for the results array", err.Error())
		return nil, d, err
	}
	if len(results) == 0 {
		err = fmt.Errorf("expected %d results, got 0", count)
		if count == 0 {
			d.AddError("No entries found for the data source", err.Error())
		} else {
			d.AddError("No data in the results array", err.Error())
		}
		return nil, d, err
	}

	out, ok = results[0].(map[string]any)
	if !ok {
		err = fmt.Errorf("received: %T instead of map[string]any", out)
		d.AddError("Unexpected format for the results array", err.Error())
		return nil, d, err
	}
	return out, d, nil
}

// parseSearchCount handles the AWX-shaped count field, which arrives as
// json.Number on the wire but may be a string or native integer when callers
// build maps directly. Unrecognised types fall back to 0 so the caller's
// "no entries found" branch fires with a sensible diagnostic.
func parseSearchCount(raw any) (int64, error) {
	if s, ok := raw.(string); ok {
		return strconv.ParseInt(s, 10, 64)
	}
	n, _, err := coerceInt64(raw)
	return n, err
}
