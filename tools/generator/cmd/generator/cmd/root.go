package cmd

import (
	"github.com/ilijamt/terraform-provider-awx/tools/generator/internal"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "generators",
	Short: "Generates various stuff for the terraform provider",
}

var cfg internal.Config

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}
