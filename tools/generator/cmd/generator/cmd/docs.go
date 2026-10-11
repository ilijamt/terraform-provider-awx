package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ilijamt/terraform-provider-awx/tools/generator/internal"
)

var docsSubcategoryCmd = &cobra.Command{
	Use:   "docs-subcategory [docs-dir]",
	Args:  cobra.ExactArgs(1),
	Short: "Set the registry subcategory on each generated resource and data source page",
	RunE: func(cmd *cobra.Command, args []string) error {
		return internal.SetDocsSubcategories(args[0])
	},
}

func init() {
	rootCmd.AddCommand(docsSubcategoryCmd)
}
