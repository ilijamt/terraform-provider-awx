package cmd

import (
	"fmt"
	"github.com/ilijamt/terraform-provider-awx/tools/generator"
	"github.com/ilijamt/terraform-provider-awx/tools/generator/internal"
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/template"
)

// templateCmd represents the base command when called without any subcommands
var templateCmd = &cobra.Command{
	Use:   "template [config-resource] [api-resource-path] [destination]",
	Args:  cobra.ExactArgs(3),
	Short: "Template all the resources for the terraform provider",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var tpl *template.Template

		var configResource, apiResourcePath, resourcePath string
		configResource = args[0]
		apiResourcePath = args[1]
		tCfg.generatePath = args[2]
		resourcePath = tCfg.generatePath

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
					err = internal.GenerateApiTfDefinition(tpl, cfg, item, resourcePath, item.Name, objmap)
					if err != nil {
						return err
					}
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
