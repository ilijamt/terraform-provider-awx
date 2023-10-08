package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/tools/generator/internal"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"strings"
)

var fetchApiResourcesCmd = &cobra.Command{
	Use:   "fetch-api-resources [out-api-resource-directory]",
	Args:  cobra.ExactArgs(2),
	Short: "Generate the API resource for the AWX target",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		log.Printf("Connecting to '%s' with the username '%s'", farCfg.towerHost, farCfg.towerUsername)
		var configResource = args[0]
		var outApiResourceDir = args[1]
		log.Printf("Storing the data in %s directory", outApiResourceDir)

		var client = c.NewClient(farCfg.towerUsername, farCfg.towerPassword, farCfg.towerHost, "generator", farCfg.insecureSkipVerify)
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

		// fetch all the defined endpoints in the config resource file
		if err = func(cfg internal.Config) error {
			log.Printf("Fetching %d items", len(cfg.Items))
			for _, item := range cfg.Items {
				req, _ = client.NewRequest(ctx, http.MethodOptions, item.Endpoint, nil)
				log.Printf("Processing %s on the %s endpoint", item.Name, item.Endpoint)
				payload, err := client.Do(ctx, req)
				if err != nil {
					return err
				}
				data.Resources[item.Name], err = internal.ResourceProcessor(item.Name, payload)
				dataInfo.Resources[item.Name] = strings.ToLower(fmt.Sprintf("payload/resource_%s.json", item.Name))
				if err != nil {
					return err
				}
			}
			return nil
		}(cfg); err != nil {
			return err
		}

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
					data.CredentialTypes[val] = ct
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

		// store all the data regarding the resources
		for k, v := range dataInfo.Resources {
			var infoFile = fmt.Sprintf("%s/%s", outApiResourceDir, v)
			log.Printf("Storing resources payload data for %s in %s", k, infoFile)

			buf.Reset()
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

		return nil

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
