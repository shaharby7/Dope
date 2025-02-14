package cmd

import (
	"github.com/shaharby7/Dope/pkg/e2e"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cmdRoot.AddCommand(cmdE2e)
}

var cmdE2e = &cobra.Command{
	Use:   "e2e",
	Short: "run e2e tests",
	Long:  `run e2e tests`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := e2e.E2EProject(
			viper.GetString(string(CONF_VARS_DOPE_PATH)),
			viper.GetString(string(CONF_VARS_DST)),
		)
		return err
	},
}
