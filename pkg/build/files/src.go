package files

import (
	"fmt"
	"path"

	v1 "github.com/shaharby7/Dope/pkg/entities/V1"
	"github.com/shaharby7/Dope/pkg/entities/entity"
	"github.com/shaharby7/Dope/pkg/utils"

	fsUtils "github.com/shaharby7/Dope/pkg/utils/fs"
	"github.com/shaharby7/Dope/pkg/utils/set"
)

type appPathArgs struct {
	App string
}

func generateSrcFiles(config *entity.Entity, appConfig *entity.Entity) ([]*fsUtils.OutputFile, error) {
	pathArgs := &appPathArgs{App: appConfig.Name}
	mainFile, err := generateOutputFile(
		templateId_SRC_FILE_MAIN,
		pathArgs,
		utils.Empty,
	)
	if err != nil {
		return nil, err
	}
	controllerFile, err := generateOutputFile(
		templateId_SRC_FILE_CONTROLLER,
		pathArgs,
		generateControllerData(
			config,
			appConfig,
		),
	)
	if err != nil {
		return nil, err
	}
	return []*fsUtils.OutputFile{mainFile, controllerFile}, nil
}

func generateControllerData(
	project *entity.Entity,
	app *entity.Entity,
) *MainControllerFileData {
	imports := set.NewSet[string]()
	controllers := []*controllerInput{}
	appConfig := entity.GetEntityValues[v1.AppConfig](app)
	config := entity.GetEntityValues[v1.ProjectConfig](project)
	for _, controllerConfig := range appConfig.Controllers {
		controller := &controllerInput{
			Name:       controllerConfig.Name,
			Identifier: "Controller_" + controllerConfig.Name,
			Type:       controllerConfig.Type,
			Actions:    []*actionInput{},
		}
		for _, actionConfig := range controllerConfig.Actions {
			imports.Add(
				path.Join(config.Module, actionConfig.Package),
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
