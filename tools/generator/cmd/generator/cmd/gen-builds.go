package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/ilijamt/terraform-provider-awx/tools/generator/internal"
)

var genBuilds = &cobra.Command{
	Use:   "gen-builds [build-config-file (yaml)]",
	Args:  cobra.ExactArgs(1),
	Short: "Generate all the builds for the versions we manage",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var buildConfigFile = args[0]
		var buildConfig = new(internal.BuildConfig)
		if err = buildConfig.Load(buildConfigFile); err != nil {
			return err
		}
		var dryRun, _ = cmd.Flags().GetBool("dry-run")
		var cmds = []string{
			"git add versions.yaml",
			"git commit -m'chore: updated versions.yaml'",
			"git push",
		}

		for _, b := range *buildConfig {
			func(b *internal.BuildConfigVersion) {
				defer func() {
					if !dryRun {
						b.Inc()
					}
				}()
				v := b.GetBuildVersion()
				cmds = append(cmds,
					fmt.Sprintf("make generate build VERSION=%s", b.Version),
					fmt.Sprintf("git add internal docs resources/api/%s", b.Version),
					fmt.Sprintf("git commit -m'chore: generated version %s with tag v%s'", b.Version, b.GetBuildVersion()),
					fmt.Sprintf("git tag v%s", v),
					fmt.Sprintf("git push origin refs/tags/v%s", v),
				)
			}(b)
		}

		if !dryRun {
			err = buildConfig.Save(buildConfigFile)
			if err != nil {
				return err
			}
		}

		cmds = append(cmds,
			"git push",
		)

		_, _ = fmt.Fprintln(cmd.OutOrStdout(), strings.Join(cmds, "\n"))

		return nil

	},
}

func init() {
	genBuilds.Flags().BoolP("dry-run", "n", false, "Do not increase the build counters")
	rootCmd.AddCommand(genBuilds)
}
