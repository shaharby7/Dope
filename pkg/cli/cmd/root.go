package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"runtime/debug"
)

const _DEFAULT_DOPE_CONFIG_NAME = "dope_config"

var cmdPersistentFlags_verbose bool
var cmdPersistentFlags_dopePath string
var cmdPersistentFlags_dst string

type CONF_VARS string

const (
	CONF_VARS_DOPE_PATH CONF_VARS = "path"
	CONF_VARS_DST       CONF_VARS = "destination"
)

func init() {

	cmdRoot.PersistentFlags().BoolVarP(&cmdPersistentFlags_verbose, "verbose", "v", false, "Display more verbose output in console output. (default: false)")

	cmdRoot.PersistentFlags().StringVarP(&cmdPersistentFlags_dopePath, "path", "p", "./.dope", "path to dope path")
	viper.BindPFlag(string(CONF_VARS_DOPE_PATH), cmdRoot.PersistentFlags().Lookup("path"))
	viper.SetDefault(string(CONF_VARS_DOPE_PATH), ".dope")

	cmdRoot.PersistentFlags().StringVarP(&cmdPersistentFlags_dst, "destination", "d", "./build", "destination of the build files")
	viper.BindPFlag(string(CONF_VARS_DST), cmdRoot.PersistentFlags().Lookup("destination"))
	viper.SetDefault(string(CONF_VARS_DST), "./build")

	viper.SetConfigName(_DEFAULT_DOPE_CONFIG_NAME)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

}

var cmdRoot = &cobra.Command{
	Use:          "dope",
	Short:        "Dope cli helps managing projects built with dope",
	Long:         "Dope (github.com/shaharby7/Dope) is a framework designated for golang microservices architecture on kubernetes.\nThis cli is designated to ease the process of creating and maintaining projects build with Dope",
	SilenceUsage: true,
}

func Execute() {
	info, _ := debug.ReadBuildInfo()
	cmdRoot.Version = info.Main.Version // TODO: add actual version :( https://github.com/golang/go/issues/50603
	if err := cmdRoot.Execute(); err != nil {
		os.Exit(1)
	}
}

func InitViper() {
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if cmdPersistentFlags_verbose {
				fmt.Println("info: did not find config file, continue with default values and flags")
			}
		} else {
			panic(err)
		}
	}
}
