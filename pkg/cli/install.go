package cli

import (
	"github.com/shaharby7/Dope/pkg/install"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cmdRoot.AddCommand(cmdInstall)
}

var cmdInstall = &cobra.Command{
	Use:   "install",
	Short: "install the project",
	Long:  `install the dope essential helm chart. Thin wrapper for the command helm install dope dope/dope -n dope -f <dst>/build/helm/local/dope/values.yaml --create-namespace`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := install.InstallProject(
			viper.GetString(string(CONF_VARS_DST)),
			&install.InstallOptions{},
		)
		return err
	},
}
