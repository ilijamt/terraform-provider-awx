package internal

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// MergeConfig writes apiDir/config.json from configDir and the apiDir overlay.
func MergeConfig(configDir, apiDir string) error {
	log.Printf("Merging config from '%s' with the overlay in '%s'", configDir, apiDir)

	var merged, err = buildMergedConfig(configDir, apiDir)
	if err != nil {
		return err
	}

	payload, err := encodeJSONDocument(merged)
	if err != nil {
		return err
	}

	var target = filepath.Join(apiDir, "config.json")
	log.Printf("Storing the merged config in %s", target)
	return os.WriteFile(target, payload, 0o644)
}

func buildMergedConfig(configDir, apiDir string) (*jsonObject, error) {
	var sharedDefault, err = loadJSONObject(filepath.Join(configDir, "default.json"))
	if err != nil {
		return nil, err
	}

	apiDefault, err := loadJSONObject(filepath.Join(apiDir, "config", "default.json"))
	if err != nil {
		return nil, err
	}

	types, err := loadTypeConfigs(filepath.Join(configDir, "types"))
	if err != nil {
		return nil, err
	}

	// Most versions carry no type overlay.
	var overrides = newJSONObject()
	var overlayDir = filepath.Join(apiDir, "config", "types")
	if _, err = os.Stat(overlayDir); err == nil {
		if overrides, err = loadTypeConfigs(overlayDir); err != nil {
			return nil, err
		}
	} else if !errors.Is(err, fs.ErrNotExist) {
		return nil, err
	}

	for _, name := range overrides.Keys() {
		override, _ := overrides.Get(name)
		if shared, ok := types.Get(name); ok {
			override = mergeJSON(shared, override)
		}
		types.Set(name, override)
	}

	var merged = newJSONObject()
	for _, source := range []*jsonObject{sharedDefault, apiDefault} {
		for _, key := range source.Keys() {
			value, _ := source.Get(key)
			merged.Set(key, value)
		}
	}
	merged.Set("items", types.Values())
	return merged, nil
}

func loadTypeConfigs(dir string) (*jsonObject, error) {
	var entries, err = os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var configs = newJSONObject()
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		var path = filepath.Join(dir, entry.Name())
		config, err := loadJSONObject(path)
		if err != nil {
			return nil, err
		}

		name, _ := config.Get("type_name")
		typeName, ok := name.(string)
		if !ok || typeName == "" {
			return nil, fmt.Errorf("%s: no type_name, it would not reach the generated provider", path)
		}
		configs.Set(typeName, config)
	}
	return configs, nil
}

func loadJSONObject(path string) (*jsonObject, error) {
	var payload, err = os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	document, err := decodeJSONDocument(payload)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}

	obj, ok := document.(*jsonObject)
	if !ok {
		return nil, fmt.Errorf("%s: holds %T, expected an object", path, document)
	}
	return obj, nil
}

// A list on the dst side is appended to rather than replaced.
func mergeJSON(dst, src any) any {
	if list, ok := dst.([]any); ok {
		if items, ok := src.([]any); ok {
			return append(slices.Clone(list), items...)
		}
		return append(slices.Clone(list), src)
	}

	dstObj, dstIsObj := dst.(*jsonObject)
	srcObj, srcIsObj := src.(*jsonObject)
	if !dstIsObj || !srcIsObj {
		return src
	}

	for _, key := range srcObj.Keys() {
		value, _ := srcObj.Get(key)
		if existing, ok := dstObj.Get(key); ok {
			value = mergeJSON(existing, value)
		}
		dstObj.Set(key, value)
	}
	return dstObj
}
