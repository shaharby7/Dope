package main

import (
	"github.com/shaharby7/Dope/pkg/cli"
	"github.com/spf13/cobra"
)

func main() {
	cobra.OnInitialize(cli.InitViper)
	cli.Execute()
}
