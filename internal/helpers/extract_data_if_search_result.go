package helpers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func ExtractDataIfSearchResult(in map[string]any) (out map[string]any, d diag.Diagnostics, err error) {
	// check data to see if it's a search result or not
	if val, ok := in["count"]; ok {
		var count int64

		switch val := val.(type) {
		case string:
			if count, err = strconv.ParseInt(val, 10, 64); err != nil {
				d.AddError(
					"Failed to convert count number in search result",
					err.Error(),
				)
				return
			}
		case json.Number:
			if count, err = val.Int64(); err != nil {
				d.AddError(
					"Failed to convert count number in search result",
					err.Error(),
				)
				return
			}
		case int:
			count = int64(val)
		case int8:
			count = int64(val)
		case int16:
			count = int64(val)
		case int32:
			count = int64(val)
		case int64:
			count = val
		case uint:
			count = int64(val)
		case uint8:
			count = int64(val)
		case uint16:
			count = int64(val)
		case uint32:
			count = int64(val)
		case uint64:
			count = int64(val)
		}

		if count == 1 {
			if res, ok := in["results"].([]any); ok {
				out = res[0].(map[string]any)
			} else {
				err = fmt.Errorf("recevied: %T instead of []any", in["results"])
				d.AddError(
					"Unexpected format for the results array",
					err.Error(),
				)
				return
			}
		} else if count > 1 {
			err = fmt.Errorf("received %d entries, expected 1", count)
			d.AddError(
				"More than one entry present, please refine your query",
				err.Error(),
			)
			return
		} else {
			err = fmt.Errorf("received %d entries, expected 1", count)
			d.AddError(
				"No entries found for the data source",
				err.Error(),
			)
			return
		}
		return
	}
	return in, d, err
}
