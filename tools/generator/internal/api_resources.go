package internal

import (
	"encoding/json"
	"fmt"
	"golang.org/x/exp/maps"
	"os"
	"slices"
)

type Api map[string]string

func (a Api) List() []string {
	var keys = maps.Keys(a)
	slices.Sort(keys)
	return keys
}

func (a Api) Endpoint(key string) (string, error) {
	if val, ok := a[key]; ok {
		return val, nil
	}
	return "", fmt.Errorf("key not found: %s", key)
}

type ApiResources struct {
	Version         string                    `json:"version"`
	Resources       map[string]map[string]any `json:"resources"`
	CredentialTypes map[string]map[string]any `json:"credential_types"`
}

type ApiResourcesInfo struct {
	Version         string            `json:"version"`
	Resources       map[string]string `json:"resources"`
	CredentialTypes map[string]string `json:"credential_types"`
}

func (c *ApiResources) Load(filename string) error {
	var payload, err = os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(payload, &c)
}
