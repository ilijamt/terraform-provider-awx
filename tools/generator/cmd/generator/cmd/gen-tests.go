package cmd

import (
	"fmt"
	"log"
	"text/template"

	"github.com/ilijamt/terraform-provider-awx/tools/generator"
	"github.com/ilijamt/terraform-provider-awx/tools/generator/internal"
	"github.com/spf13/cobra"
)

// genTestsCmd represents the base command when called without any subcommands
var genTestsCmd = &cobra.Command{
	Use:   "gen-tests [config-resource] [api-resource-path] [destination] [tests-payload-location]",
	Args:  cobra.ExactArgs(4),
	Short: "Generate tests for the provider",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var tpl *template.Template

		var configResource, apiResourcePath, resourcePath, testPayloadLocation string
		configResource = args[0]
		apiResourcePath = args[1]
		tCfg.generatePath = args[2]
		testPayloadLocation = args[3]
		resourcePath = tCfg.generatePath

		if err = cfg.Load(configResource); err != nil {
			return err
		}

		var apiResourceInfo = new(internal.ApiResourcesInfo)
		if err = apiResourceInfo.Load(fmt.Sprintf("%s/info.json", apiResourcePath)); err != nil {
			return err
		}

		var apiResource = new(internal.ApiResources)
		if err = apiResource.Load(apiResourcePath, *apiResourceInfo); err != nil {
			return err
		}

		log.Printf("Processing '%s' in '%s' resource from '%s' with test data in '%s'\n", apiResourcePath, resourcePath, configResource, testPayloadLocation)

		tpl, err = template.New("").Funcs(internal.FuncMap).ParseFS(generator.Fs(), "templates/terraform/*.tpl")
		if err != nil {
			return err
		}
		tpl.Name()

		for _, item := range cfg.Items {
			if item.Enabled {
				log.Printf("Generating empty test files for %s ...", item.Name)
			} else {
				log.Printf("Skipping %s, disabled ...", item.Name)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(genTestsCmd)
}
