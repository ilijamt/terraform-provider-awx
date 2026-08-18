package cmd

import (
	"bytes"
	"cmp"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"

	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/tools/generator/internal"
)

var fetchApiResourcesCmd = &cobra.Command{
	Use:   "fetch-api-resources [api-resource-directory]",
	Args:  cobra.ExactArgs(1),
	Short: "Generate the API resource for the AWX target",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		log.Printf("Connecting to '%s' with the username '%s'", farCfg.towerHost, farCfg.towerUsername)
		var outApiResourceDir = args[0]
		var configResource = fmt.Sprintf("%s/config.json", outApiResourceDir)

		log.Printf("Storing the data in %s directory", outApiResourceDir)

		var client = c.NewClientWithBasicAuth(farCfg.towerUsername, farCfg.towerPassword, farCfg.towerHost, "generator", farCfg.insecureSkipVerify, nil)
		var data internal.ApiResources
		var dataInfo internal.ApiResourcesInfo
		var ctx = context.Background()
		var req *http.Request

		_ = os.Mkdir(outApiResourceDir, os.ModePerm)
		_ = os.Mkdir(fmt.Sprintf("%s/payload", outApiResourceDir), os.ModePerm)

		data.Resources = make(map[string]map[string]any)
		dataInfo.Resources = make(map[string]string)
		data.CredentialTypes = make(map[string]map[string]any)
		dataInfo.CredentialTypes = make(map[string]string)

		// fetch the version of the system
		if req, err = client.NewRequest(ctx, http.MethodGet, "/api/v2/ping", nil); err != nil {
			return err
		}
		if err = func() error {
			var payload, err = client.Do(ctx, req)
			if err != nil {
				return err
			}
			if val, ok := payload["version"].(string); ok {
				data.Version = val
				dataInfo.Version = val
			}
			return nil
		}(); err != nil {
			return err
		}

		var cfg internal.Config
		if err = cfg.Load(configResource); err != nil {
			return err
		}

		// fetch the api endpoint
		var api internal.Api
		if api, err = func() (internal.Api, error) {
			var api = make(internal.Api)
			if req, err = client.NewRequest(ctx, http.MethodGet, "/api/v2", nil); err != nil {
				return nil, err
			}
			log.Printf("Processing %s endpoint", req.RequestURI)
			payload, err := client.Do(ctx, req)
			if err != nil {
				return api, err
			}

			var buf bytes.Buffer
			var enc = json.NewEncoder(&buf)
			enc.SetIndent("", "  ")
			if err = enc.Encode(payload); err != nil {
				return api, err
			}
			var apiFile = fmt.Sprintf("%s/api.json", outApiResourceDir)
			if err = json.Unmarshal(buf.Bytes(), &api); err != nil {
				return api, err
			}
			log.Printf("Storing api endpoint data in %s", apiFile)
			return api, os.WriteFile(apiFile, buf.Bytes(), os.ModePerm)
		}(); err != nil {
			return err
		}

		log.Printf("Found %d api endpoints", len(api))

		if err = func() error {
			var prev internal.ApiResourcesInfo
			if err := prev.Load(fmt.Sprintf("%s/info.json", outApiResourceDir)); err != nil {
				return nil
			}
			for _, item := range cfg.Items {
				if path, ok := prev.Resources[item.Name]; ok {
					dataInfo.Resources[item.Name] = path
				}
			}
			return nil
		}(); err != nil {
			return err
		}

		fetchFailures := func(cfg internal.Config) error {
			log.Printf("Fetching %d items", len(cfg.Items))
			// The loop tries every item before giving up, so one run names all
			// the endpoints that need attention rather than the first.
			var failures []error
			for _, item := range cfg.Items {
				if item.CredentialType != "" {
					delete(dataInfo.Resources, item.Name)
					continue
				}
				endpoint, err := metadataEndpoint(ctx, client, item)
				if err != nil {
					failures = append(failures, fmt.Errorf("%s: %w", item.Name, err))
					continue
				}
				req, _ = client.NewRequest(ctx, http.MethodOptions, endpoint, nil)
				log.Printf("Processing %s on the %s endpoint", item.Name, endpoint)
				payload, err := client.Do(ctx, req)
				if err != nil {
					failures = append(failures, fmt.Errorf("%s on %s: %w", item.Name, endpoint, err))
					continue
				}
				processed, err := internal.ResourceProcessor(item.Name, payload)
				if err != nil {
					failures = append(failures, fmt.Errorf("%s: %w", item.Name, err))
					continue
				}
				if missing := missingActions(item, processed); len(missing) > 0 {
					failures = append(failures, fmt.Errorf("%s on %s: OPTIONS returned no %s block, seed the instance first",
						item.Name, endpoint, strings.Join(missing, "/")))
					continue
				}
				data.Resources[item.Name] = processed
				dataInfo.Resources[item.Name] = strings.ToLower(fmt.Sprintf("payload/resource_%s.json", item.Name))
			}
			return errors.Join(failures...)
		}(cfg)

		// fetch all the defined credential types
		if err = func(cfg internal.Config) error {
			if req, err = client.NewRequest(ctx, http.MethodGet, "/api/v2/credential_types?managed=true&page_size=200", nil); err != nil {
				return err
			}

			var payload, err = client.Do(ctx, req)
			if err != nil {
				return err
			}

			var sr internal.SearchResults
			if err = mapstructure.Decode(payload, &sr); err != nil {
				return err
			}

			log.Printf("Fetched %d managed credential types", sr.Count)
			if sr.Count == 0 {
				log.Printf("No managed credential types to fetch")
				return nil
			}

			for _, ct := range sr.Results {
				if val, ok := ct["namespace"].(string); ok {
					data.CredentialTypes[val] = internal.NormalizeCredentialTypePayload(ct)
					dataInfo.CredentialTypes[val] = strings.ToLower(fmt.Sprintf("payload/credential_type_%s.json", val))
				}
			}
			return nil
		}(cfg); err != nil {
			return err
		}

		var buf bytes.Buffer
		var enc = json.NewEncoder(&buf)
		enc.SetIndent("", "  ")

		// Store the information regarding the payloads and defined resources/credential types
		var infoFile = fmt.Sprintf("%s/info.json", outApiResourceDir)
		if err = enc.Encode(dataInfo); err != nil {
			return err
		}
		log.Printf("Storing information data in %s", infoFile)
		if err = os.WriteFile(infoFile, buf.Bytes(), os.ModePerm); err != nil {
			return err
		}

		// store all the data regarding the resources; keyed off what was actually
		// fetched so a carried-over mapping never overwrites its payload with null
		for k := range data.Resources {
			var infoFile = fmt.Sprintf("%s/%s", outApiResourceDir, dataInfo.Resources[k])
			log.Printf("Storing resources payload data for %s in %s", k, infoFile)

			buf.Reset()
			if err = enc.Encode(data.Resources[k]); err != nil {
				return err
			}
			log.Printf("Storing resource payload data for %s in %s", k, infoFile)
			if err = os.WriteFile(infoFile, buf.Bytes(), os.ModePerm); err != nil {
				return err
			}
		}

		// store all the data regarding the credential types
		for k, v := range dataInfo.CredentialTypes {
			var infoFile = fmt.Sprintf("%s/%s", outApiResourceDir, v)
			buf.Reset()
			if err = enc.Encode(data.CredentialTypes[k]); err != nil {
				return err
			}
			log.Printf("Storing credential types payload data for %s in %s", k, infoFile)
			if err = os.WriteFile(infoFile, buf.Bytes(), os.ModePerm); err != nil {
				return err
			}
		}

		return fetchFailures

	},
}

var farCfg struct {
	towerHost          string
	towerUsername      string
	towerPassword      string
	insecureSkipVerify bool
}

func init() {
	fetchApiResourcesCmd.Flags().StringVar(&farCfg.towerHost, "host", "", "The host we use to connect to AWX")
	fetchApiResourcesCmd.Flags().StringVar(&farCfg.towerUsername, "username", "", "The username to connect to AWX")
	fetchApiResourcesCmd.Flags().StringVar(&farCfg.towerPassword, "password", "", "The password to connect to AWX")
	fetchApiResourcesCmd.Flags().BoolVar(&farCfg.insecureSkipVerify, "insecure-skip-verify", false, "Should we skip verification of TLS")
	_ = fetchApiResourcesCmd.MarkFlagRequired("host")
	_ = fetchApiResourcesCmd.MarkFlagRequired("username")
	_ = fetchApiResourcesCmd.MarkFlagRequired("password")
	rootCmd.AddCommand(fetchApiResourcesCmd)
}

// metadataEndpoint resolves the URL to run OPTIONS against. An item can point it
// away from its CRUD endpoint and have the id discovered from a live object.
func metadataEndpoint(ctx context.Context, client c.Client, item internal.Item) (string, error) {
	endpoint, needsDiscovery := item.MetadataUrl()
	if !needsDiscovery {
		return endpoint, nil
	}
	if item.MetadataDiscovery == nil {
		return "", fmt.Errorf("metadata_endpoint %q needs an id but no metadata_discovery is configured", endpoint)
	}

	id, err := discoverId(ctx, client, *item.MetadataDiscovery)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(endpoint, id), nil
}

// discoverId reads the id off the first result of a list endpoint. An empty list
// is an error rather than a skip: with no instance to describe, the resource
// would drop out of the download without saying so.
func discoverId(ctx context.Context, client c.Client, discovery internal.MetadataDiscovery) (int64, error) {
	req, err := client.NewRequest(ctx, http.MethodGet, discovery.Endpoint, nil)
	if err != nil {
		return 0, err
	}
	payload, err := client.Do(ctx, req)
	if err != nil {
		return 0, err
	}

	var sr internal.SearchResults
	if err = mapstructure.Decode(payload, &sr); err != nil {
		return 0, err
	}
	if len(sr.Results) == 0 {
		return 0, fmt.Errorf("%s returned no results, seed the instance first", discovery.Endpoint)
	}

	field := cmp.Or(discovery.Field, "id")
	id, ok := sr.Results[0][field]
	if !ok {
		return 0, fmt.Errorf("%s returned no %q on its first result", discovery.Endpoint, field)
	}
	num, ok := id.(json.Number)
	if !ok {
		return 0, fmt.Errorf("%s returned %q as %T, expected a number", discovery.Endpoint, field, id)
	}
	return num.Int64()
}

func missingActions(item internal.Item, payload map[string]any) []string {
	var required []string
	if !item.NoTerraformResource {
		required = append(required, item.ApiPropertyResourceKey)
	}
	if !item.NoTerraformDataSource {
		required = append(required, item.ApiPropertyDataKey)
	}

	actions, _ := payload["actions"].(map[string]any)
	var missing []string
	for _, key := range required {
		if key == "" || slices.Contains(missing, key) {
			continue
		}
		if _, ok := actions[key]; !ok {
			missing = append(missing, key)
		}
	}
	return missing
}
