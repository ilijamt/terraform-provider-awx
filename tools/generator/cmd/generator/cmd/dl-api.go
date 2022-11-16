package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	c "github.com/ilijamt/terraform-provider-awx/internal/client"
	"github.com/ilijamt/terraform-provider-awx/tools/generator/internal"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
)

var fetchApiResourcesCmd = &cobra.Command{
	Use:   "fetch-api-resources [config-resource] [out-api-resource-file]",
	Args:  cobra.ExactArgs(2),
	Short: "Generate the API resource for the AWX target",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		log.Printf("Connecting to '%s' with the username '%s'", farCfg.towerHost, farCfg.towerUsername)
		var configResource = args[0]
		var outApiResourceFile = args[1]
		log.Printf("Reading configuration from %s, and storing the data in %s", configResource, outApiResourceFile)

		var client = c.NewClient(farCfg.towerUsername, farCfg.towerPassword, farCfg.towerHost, "generator", farCfg.insecureSkipVerify)
		var data internal.ApiResources
		var ctx = context.Background()
		var req *http.Request

		data.Resources = make(map[string]map[string]any)

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
			}
			return nil
		}(); err != nil {
			return err
		}

		// fetch all the defined endpoints in the config resource file
		if err = func() error {
			var cfg internal.Config
			if err = cfg.Load(configResource); err != nil {
				return err
			}

			log.Printf("Fetching %d items", len(cfg.Items))
			for _, item := range cfg.Items {
				req, _ = client.NewRequest(ctx, http.MethodOptions, item.Endpoint, nil)
				log.Printf("Processing %s on the %s endpoint", item.Name, item.Endpoint)
				payload, err := client.Do(ctx, req)
				if err != nil {
					return err
				}
				data.Resources[item.Name], err = internal.ResourceProcessor(item.Name, payload)
				if err != nil {
					return err
				}
			}
			return nil
		}(); err != nil {
			return err
		}

		var buf bytes.Buffer
		var enc = json.NewEncoder(&buf)
		enc.SetIndent("", "  ")
		if err = enc.Encode(data); err != nil {
			return err
		}

		return os.WriteFile(outApiResourceFile, buf.Bytes(), os.ModePerm)
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
