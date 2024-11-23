package cmd

import (
	"fmt"

	"github.com/shaharby7/Dope/pkg/utils"
	"github.com/shaharby7/Dope/types"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

func init() {
	cmdRoot.AddCommand(cmdCreate)
}

var cmdCreate = &cobra.Command{
	Use:   "create",
	Short: "create new dope object",
	Long:  `use the create command to add new dope objects`,
	RunE: func(cmd *cobra.Command, args []string) error {
		objType, err := getObjectType(args)
		if err != nil {
			utils.FailedBecause("infer dope object type", err)
		}
		fmt.Print(*objType)
		return nil
	},
}

func getObjectType(args []string) (*types.DOPE_OBJECTS, error) {
	if (len(args)) != 1 {
		return nil, fmt.Errorf("create command can only be executed with with argument, found %d", len(args))
	}
	res := types.DOPE_OBJECTS(args[0])
	return &res, nil
}
