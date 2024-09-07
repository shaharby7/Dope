package cmd

import (
	"github.com/shaharby7/Dope/pkg/build"
	"github.com/spf13/cobra"
)

var projPath string
var dst string

func init() {
	cmdRoot.AddCommand(cmdBuild)
	cmdBuild.Flags().StringVarP(&projPath, "path", "p", "./project.dope.yaml", "path to project file")
	cmdBuild.Flags().StringVarP(&dst, "destination", "d", "./build", "destination of the build files")
}

var cmdBuild = &cobra.Command{
	Use:   "build",
	Short: "build the project",
	Long:  `builds different stages of the project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := build.BuildProject(projPath, dst, build.BuildOptions{})
		return err
	},
}
