package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/ilijamt/terraform-provider-awx/tools/generator"
	"github.com/ilijamt/terraform-provider-awx/tools/generator/internal"
)

// templateCmd represents the base command when called without any subcommands
var templateCmd = &cobra.Command{
	Use:   "template [api-resource-path] [destination]",
	Args:  cobra.ExactArgs(2),
	Short: "Template all the resources for the terraform provider",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var tpl *template.Template

		var configResource, apiResourcePath, resourcePath string
		apiResourcePath = args[0]
		tCfg.generatePath = args[1]
		resourcePath = tCfg.generatePath
		configResource = fmt.Sprintf("%s/config.json", apiResourcePath)

		if err = cfg.Load(configResource); err != nil {
			return err
		}

		_ = os.Mkdir(fmt.Sprintf("%s/docs", apiResourcePath), os.ModePerm)

		var apiResourceInfo = &internal.ApiResourcesInfo{}
		if err = apiResourceInfo.Load(fmt.Sprintf("%s/info.json", apiResourcePath)); err != nil {
			return err
		}

		var apiResource = &internal.ApiResources{}
		if err = apiResource.Load(apiResourcePath, *apiResourceInfo); err != nil {
			return err
		}

		log.Printf("Processing '%s' in '%s' resource from '%s'\n", apiResourcePath, resourcePath, configResource)

		tpl, err = template.New("").Funcs(internal.FuncMap).ParseFS(generator.Fs(), "templates/terraform/*.tpl")
		if err != nil {
			return err
		}

		for _, item := range cfg.Items {
			if item.Enabled {
				if !item.NoTerraformResource {
					cfg.GeneratedApiResources = append(cfg.GeneratedApiResources, item.Name)
					for _, adg := range item.AssociateDisassociateGroups {
						cfg.GeneratedApiResources = append(cfg.GeneratedApiResources, fmt.Sprintf("%sAssociateDisassociate%s", item.Name, adg.Type))
					}
					if item.HasSurveySpec {
						cfg.GeneratedApiResources = append(cfg.GeneratedApiResources, fmt.Sprintf("%sSurvey", item.Name))
					}
				}

				if !item.NoTerraformDataSource {
					cfg.GeneratedDataSourceResources = append(cfg.GeneratedDataSourceResources, item.Name)
					if item.HasObjectRoles {
						cfg.GeneratedDataSourceResources = append(cfg.GeneratedDataSourceResources, fmt.Sprintf("%sObjectRoles", item.Name))
					}
				}

				if objmap, ok := apiResource.Resources[item.Name]; ok {
					var data map[string]any
					data, err = internal.GenerateApiTfDefinition(tpl, cfg, item, resourcePath, item.Name, objmap)
					if err != nil {
						return err
					}
					payload, _ := json.MarshalIndent(data, "", "  ")
					genDataFile := fmt.Sprintf("%s/gen-data/%s.json", apiResourcePath, item.Name)
					log.Printf("Storing generated data for '%s' in '%s'\n", item.Name, genDataFile)
					_ = os.WriteFile(genDataFile, payload, os.ModePerm)
				} else {
					log.Printf("Missing definition for %s, skipping ...", item.Name)
				}
			} else {
				log.Printf("Skipping %s, disabled ...", item.Name)
			}
		}

		return internal.GenerateApiSourcesForProvider(tpl, cfg, resourcePath, cfg.GeneratedApiResources, cfg.GeneratedDataSourceResources)
	},
}

var tCfg struct {
	generatePath string
}

func init() {
	rootCmd.AddCommand(templateCmd)
}
