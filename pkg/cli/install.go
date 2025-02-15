package cli

import (
	"github.com/shaharby7/Dope/pkg/install"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var config = install.NewConfig()

func init() {
	cmdRoot.AddCommand(cmdInstall)
	cmdRoot.AddCommand(cmdUninstall)
	cmdInstall.Flags().BoolVarP(&config.Upgrade, "upgrade", "u", false, "make helm upgrade instead of install")
}

var cmdInstall = &cobra.Command{
	Use:   "install",
	Short: "install the project",
	Long: `install the dope essential helm chart. Thin wrapper for the command
	"helm install dope dope/dope -n dope -f <dst>/build/helm/local/dope/values.yaml --create-namespace"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := install.InstallProject(
			viper.GetString(string(CONF_VARS_DOPE_PATH)),
			viper.GetString(string(CONF_VARS_DST)),
			config,
		)
		return err
	},
}

var cmdUninstall = &cobra.Command{
	Use:   "uninstall",
	Short: "uninstall the project",
	Long: `uninstall the dope essential helm chart. Thin wrapper for the command
	"helm uninstall dope -n dope"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := install.UninstallProject(
			viper.GetString(string(CONF_VARS_DOPE_PATH)),
			viper.GetString(string(CONF_VARS_DST)),
			config,
		)
		return err
	},
}
