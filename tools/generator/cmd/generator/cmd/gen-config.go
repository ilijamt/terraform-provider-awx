package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ilijamt/terraform-provider-awx/tools/generator/internal"
)

var genConfigCmd = &cobra.Command{
	Use:   "generate-config [config-directory] [api-resource-directory]",
	Args:  cobra.ExactArgs(2),
	Short: "Merge the shared and version specific config into [api-resource-directory]/config.json",
	RunE: func(cmd *cobra.Command, args []string) error {
		return internal.MergeConfig(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(genConfigCmd)
}
