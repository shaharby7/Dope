package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"runtime/debug"
)

func init() {
}

var cmdRoot = &cobra.Command{
	Use:   "dope",
	Short: "Dope cli helps maenaging projects built with dope",
	Long:  "Dope (github.com/shaharby7/Dope) is a framework designated for golang microservices architecture on kubernetes.\nThis cli is designated to ease the process of creating and maintaining projects build with Dope",
}

func Execute() {
	info, _ := debug.ReadBuildInfo()
	cmdRoot.Version = info.Main.Version // TODO: add actual version :( https://github.com/golang/go/issues/50603
	if err := cmdRoot.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
