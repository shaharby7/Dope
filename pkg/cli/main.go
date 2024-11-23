package main

import (
	"github.com/shaharby7/Dope/pkg/cli/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.OnInitialize(cmd.InitViper)
	cmd.Execute()
}
