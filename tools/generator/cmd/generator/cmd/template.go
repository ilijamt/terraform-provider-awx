package cmd

import (
	"fmt"
	"log"
	"os"
	"slices"
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
		var deprecated internal.Deprecated

		var configResource, apiResourcePath, resourcePath string
		apiResourcePath = args[0]
		tCfg.generatePath = args[1]
		resourcePath = tCfg.generatePath
		configResource = fmt.Sprintf("%s/config.json", apiResourcePath)
		_ = os.MkdirAll(fmt.Sprintf("%s/gen-model-data", apiResourcePath), os.ModePerm)
		_ = os.MkdirAll(fmt.Sprintf("%s/gen-model-data/resources", apiResourcePath), os.ModePerm)
		_ = os.MkdirAll(fmt.Sprintf("%s/gen-model-data/credentials", apiResourcePath), os.ModePerm)

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

		tpl, err = template.New("").Funcs(internal.FuncMap).ParseFS(generator.Fs(), "templates/*.tpl", "templates/terraform/*.tpl")
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
					// var data map[string]any
					var p *internal.ModelConfig
					var dr internal.Deprecated
					_, p, dr, err = internal.GenerateApiTfDefinition(tpl, cfg, item, resourcePath, item.Name, objmap)
					if err != nil {
						return err
					}
					deprecated.Resources = append(deprecated.Resources, dr.Resources...)
					deprecated.DataSources = append(deprecated.DataSources, dr.DataSources...)
					deprecated.Properties = append(deprecated.Properties, dr.Properties...)
					_ = p.Save(fmt.Sprintf("%s/gen-model-data/resources", apiResourcePath))
				} else {
					log.Printf("Missing definition for %s, skipping ...", item.Name)
				}
			} else {
				log.Printf("Skipping %s, disabled ...", item.Name)
			}
		}

		for _, item := range cfg.Credentials {
			if !item.Enabled {
				log.Printf("Skipping Credential %s, disabled ...", item.Name)
				continue
			}
			var p *internal.ModelCredential
			var objmap map[string]any
			var ok bool
			var inclDatasource bool
			if objmap, ok = apiResource.CredentialTypes[item.TypeName]; !ok {
				log.Printf("Missing definition for %s, skipping ...", item.Name)
				continue
			}
			p, inclDatasource, err = internal.GenerateApiTfCredentialDefinition(tpl, cfg, item, item.Name, resourcePath, objmap)
			if err != nil {
				log.Printf("Error generating credentials for '%s' in '%s': %v", item.Name, item.TypeName, err)
				return err
			}

			_ = p.Save(fmt.Sprintf("%s/gen-model-data/credentials", apiResourcePath))
			cfg.GeneratedApiResources = append(cfg.GeneratedApiResources, fmt.Sprintf("%sCredential", item.Name))
			if inclDatasource {
				cfg.GeneratedDataSourceResources = append(cfg.GeneratedDataSourceResources, fmt.Sprintf("%sCredential", item.Name))
			}
		}

		{
			filePath := fmt.Sprintf("%s/deprecated.md", apiResourcePath)
			if file, err := os.Create(filePath); err == nil {
				log.Printf("Storing deprecated data in %s\n", filePath)
				slices.Sort(deprecated.Resources)
				slices.Sort(deprecated.DataSources)
				err = tpl.ExecuteTemplate(file, "deprecated.md.tpl", deprecated)
				if err != nil {
					fmt.Println(err)
				}
				_ = file.Close()
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
