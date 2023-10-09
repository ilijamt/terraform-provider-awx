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

func (c *ApiResources) Load(data string, info ApiResourcesInfo) error {
	c.Version = info.Version

	// decode and load all the resources information
	c.Resources = make(map[string]map[string]any)
	for k, v := range info.Resources {
		var payload, err = os.ReadFile(fmt.Sprintf("%s/%s", data, v))
		if err != nil {
			return err
		}
		var obj map[string]any
		if err = json.Unmarshal(payload, &obj); err != nil {
			return err
		}
		c.Resources[k] = obj
	}

	// decode and load all the credential types
	c.CredentialTypes = make(map[string]map[string]any)
	for k, v := range info.CredentialTypes {
		var payload, err = os.ReadFile(fmt.Sprintf("%s/%s", data, v))
		if err != nil {
			return err
		}
		var obj map[string]any
		if err = json.Unmarshal(payload, &obj); err != nil {
			return err
		}
		c.CredentialTypes[k] = obj
	}

	return nil
}

type ApiResourcesInfo struct {
	Version         string            `json:"version"`
	Resources       map[string]string `json:"resources"`
	CredentialTypes map[string]string `json:"credential_types"`
}

func (c *ApiResourcesInfo) Load(filename string) error {
	var payload, err = os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(payload, &c)
}
