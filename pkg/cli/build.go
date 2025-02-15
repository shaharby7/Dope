package cli

import (
	"github.com/shaharby7/Dope/pkg/build"
	bTypes "github.com/shaharby7/Dope/pkg/build/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cmdRoot.AddCommand(cmdBuild)
}

var cmdBuild = &cobra.Command{
	Use:   "build",
	Short: "build the project",
	Long:  `builds different stages of the project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := build.BuildProject(
			viper.GetString(string(CONF_VARS_DOPE_PATH)),
			viper.GetString(string(CONF_VARS_DST)),
			bTypes.BuildOptions{},
		)
		return err
	},
}
