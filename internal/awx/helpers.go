package awx

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func extractDataIfSearchResult(in map[string]any) (out map[string]any, d diag.Diagnostics, err error) {
	// check data to see if it's a search result or not
	if val, ok := in["count"]; ok {
		var count int64

		switch val.(type) {
		case string:
			if count, err = strconv.ParseInt(val.(string), 10, 64); err != nil {
				d.AddError(
					fmt.Sprintf("Failed to convert count number in search result"),
					err.Error(),
				)
				return
			}
		case json.Number:
			if count, err = val.(json.Number).Int64(); err != nil {
				d.AddError(
					fmt.Sprintf("Failed to convert count number in search result"),
					err.Error(),
				)
				return
			}
		case int:
			count = val.(int64)
		}

		if count == 1 {
			if res, ok := in["results"].([]any); ok {
				out = res[0].(map[string]any)
			} else {
				err = fmt.Errorf("recevied: %T instead of []any", in["results"])
				d.AddError(
					fmt.Sprintf("Unexpected format for the results array"),
					err.Error(),
				)
				return
			}
		} else if count > 1 {
			err = fmt.Errorf("received %d entries, expected 1", count)
			d.AddError(
				fmt.Sprintf("More than one entry present, please refine your query"),
				err.Error(),
			)
			return
		} else {
			err = fmt.Errorf("received %d entries, expected 1", count)
			d.AddError(
				fmt.Sprintf("No entries found for the data source"),
				err.Error(),
			)
			return
		}
		return
	}
	return in, d, err
}
