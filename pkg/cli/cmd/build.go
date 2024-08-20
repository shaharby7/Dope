package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdBuild = &cobra.Command{
	Use:   "build [ args ]",
	Short: "build the project",
	Long:  `build different stages of the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("starting to build...")
	},
}
