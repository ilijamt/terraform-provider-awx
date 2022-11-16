package internal

import (
	"encoding/json"
	"os"
)

type ApiResources struct {
	Version   string                    `json:"version"`
	Resources map[string]map[string]any `json:"resources"`
}

func (c *ApiResources) Load(filename string) error {
	var payload, err = os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(payload, &c)
}
