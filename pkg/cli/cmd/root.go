package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"runtime/debug"
)

var rootCmd = &cobra.Command{
	Use:   "dope",
	Short: "Dope cli helps managing projects built with dope",
	Long:  "Dope (github.com/shaharby7/Dope) is a framework designated for golang microservices architecture on kubernetes.\nThis cli is designated to ease the process of creating and maintaining projects build with Dope",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("hi!")
	},
}

func Execute() {
	rootCmd.AddCommand(cmdBuild)
	info, _ := debug.ReadBuildInfo()
	rootCmd.Version= info.Main.Version
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
