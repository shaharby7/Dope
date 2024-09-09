package files

import (
	"fmt"
	"path"

	"github.com/shaharby7/Dope/pkg/utils/set"
	"github.com/shaharby7/Dope/types"
)

func generateSrcFiles(dst string, config *types.ProjectConfig, appConfig *types.AppConfig) ([]*OutputFile, error) {
	appDst := path.Join(dst, appConfig.Name)
	mainFile, err := generateFileByTemplate(
		appDst,
		SRC_FILE_MAIN,
		&FileGenerationInput[*MainFileData, MainFileArgPath]{
			Params: struct{ Path MainFileArgPath }{
				Path: MainFileArgPath{App: appConfig.Name},
			},
			Data: nil,
		},
	)
	if err != nil {
		return nil, err
	}
	controllerFile, err := generateFileByTemplate(
		appDst,
		SRC_FILE_CONTROLLER,
		&FileGenerationInput[*MainControllerFileData, MainFileArgPath]{
			Params: struct{ Path MainFileArgPath }{
				Path: MainFileArgPath{App: appConfig.Name},
			},
			Data: generateControllerData(
				config,
				appConfig,
			),
		},
	)
	if err != nil {
		return nil, err
	}
	return []*OutputFile{mainFile, controllerFile}, nil
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
