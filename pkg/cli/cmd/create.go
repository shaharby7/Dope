package cmd

import (
	"fmt"
	"path"
	"slices"

	"github.com/shaharby7/Dope/pkg/utils"

	"github.com/shaharby7/Dope/pkg/config"

	cliUtils "github.com/shaharby7/Dope/pkg/cli/cmd/utils"

	"github.com/shaharby7/Dope/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		name, err := cliUtils.GetInputString("Name")
		if err != nil {
			utils.FailedBecause("getting name", err)
		}
		description, err := cliUtils.GetInputString("Description")
		if err != nil {
			utils.FailedBecause("getting description", err)
		}
		binding, err := getObjectBinding(objType)
		if err != nil {
			utils.FailedBecause("getting binding", err)
		}
		objPath := inferObjPath(
			viper.GetString(string(CONF_VARS_DOPE_PATH)),
			name,
			objType,
			binding,
		)
		values, err := getObjectValues(objType)
		if err != nil {
			utils.FailedBecause("values binding", err)
		}
		obj := &types.DopeObjectFile[any]{
			Api:         "Dope/V1",
			Type:        objType,
			Name:        name,
			Description: description,
			Binding:     binding,
			Values:      values,
		}
		fmt.Println(objPath)
		return config.WriteConfig(objPath, obj)
	},
}

func getObjectType(args []string) (types.DOPE_OBJECTS, error) {
	if (len(args)) != 1 {
		return "", fmt.Errorf("could not create command can only be executed with with argument, found %d", len(args))
	}
	switch args[0] { // TODO: use config object types
	case "app":
		return types.DOPE_OBJECT_APP, nil
	case "env":
		return types.DOPE_OBJECT_ENV, nil
	case "appenv":
		return types.DOPE_OBJECT_APP_ENV, nil
	case "project":
		return types.DOPE_OBJECT_PROJECT, nil
	default:
		return "", fmt.Errorf("could not infer object type, expected known dope object name, found:%s", args[0])
	}
}

func getObjectBinding(objType types.DOPE_OBJECTS) (*types.DopeObjectFileBinding, error) {
	res := &types.DopeObjectFileBinding{}
	if slices.Contains([]types.DOPE_OBJECTS{types.DOPE_OBJECT_APP_ENV}, objType) {
		env, err := cliUtils.GetInputString("Env binding")
		if err != nil {
			return nil, utils.FailedBecause("getting env binding", err)
		}
		res.Env = &env
	}
	if slices.Contains([]types.DOPE_OBJECTS{types.DOPE_OBJECT_APP_ENV}, objType) {
		app, err := cliUtils.GetInputString("App binding")
		if err != nil {
			return nil, utils.FailedBecause("getting app binding", err)
		}
		res.App = &app
	}
	if (&types.DopeObjectFileBinding{} == res) {
		return nil, nil
	}
	return res, nil
}

func inferObjPath(dst string, name string, objType types.DOPE_OBJECTS, binding *types.DopeObjectFileBinding) string {
	objPath := dst
	if binding.Env != nil || objType == types.DOPE_OBJECT_ENV {
		objPath = path.Join(objPath, "envs")
		if binding.Env != nil {
			objPath = path.Join(objPath, *binding.Env)
		}
		if objType == types.DOPE_OBJECT_ENV {
			objPath = path.Join(objPath, name)
		}
	}
	if binding.App != nil || objType == types.DOPE_OBJECT_APP {
		objPath = path.Join(objPath, "apps")
		if binding.App != nil {
			objPath = path.Join(objPath, *binding.App)
		}
		if objType == types.DOPE_OBJECT_APP {
			objPath = path.Join(objPath, name)
		}
	}
	fileName := ""
	fmt.Println(objType)
	if objType == types.DOPE_OBJECT_PROJECT {
		fileName = "project"
	} else {
		fileName = name
	}
	objPath = path.Join(objPath, fileName+".dope.yaml")
	return objPath
}

func getObjectValues(types.DOPE_OBJECTS) (any, error) {
	return nil, nil
}
