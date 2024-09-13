package files

import (
	"fmt"
	"path"

	"github.com/shaharby7/Dope/pkg/utils"
	"github.com/shaharby7/Dope/pkg/utils/set"
	"github.com/shaharby7/Dope/types"
)

type appPathArgs struct {
	App string
}

func srcFilesGenGen(config *types.ProjectConfig, appConfig *types.AppConfig) []iFileGenerator {
	pathArgs := &appPathArgs{App: appConfig.Name}
	mainFile := newFileGenerator(
		templateId_SRC_FILE_MAIN,
		pathArgs,
		utils.Empty,
	)
	controllerFile := newFileGenerator(
		templateId_SRC_FILE_CONTROLLER,
		pathArgs,
		generateControllerData(
			config,
			appConfig,
		),
	)
	return []iFileGenerator{mainFile, controllerFile}
}

func generateControllerData(
	config *types.ProjectConfig,
	appConfig *types.AppConfig,
) *MainControllerFileData {
	imports := set.NewSet[string]()
	controllers := []*controllerInput{}
	for _, controllerConfig := range appConfig.Controllers {
		controller := &controllerInput{
			Name:       controllerConfig.Name,
			Identifier: "Controller_" + controllerConfig.Name,
			Type:       controllerConfig.Type,
			Actions:    []*actionInput{},
		}
		for _, actionConfig := range controllerConfig.Actions {
			imports.Add(
				path.Join(config.Metadata.Module, actionConfig.Package),
			)
			action := &actionInput{
				Name: actionConfig.Name,
				Caller: fmt.Sprintf(
					"%s.%s", path.Base(actionConfig.Package), actionConfig.Ref,
				),
				ControllerBinding: &actionConfig.ControllerBinding,
			}
			controller.Actions = append(controller.Actions, action)
		}
		controllers = append(controllers, controller)
	}

	return &MainControllerFileData{
		Imports:     imports.ToSlice(),
		Controllers: controllers,
	}
}
