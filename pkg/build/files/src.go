package files

import (
	"fmt"
	"path"
	"strings"

	"github.com/shaharby7/Dope/pkg/entities"
	v1 "github.com/shaharby7/Dope/pkg/entities/V1"
	"github.com/shaharby7/Dope/pkg/entities/entity"
	configHelpers "github.com/shaharby7/Dope/pkg/entities/helpers"
	"github.com/shaharby7/Dope/pkg/utils"

	fsUtils "github.com/shaharby7/Dope/pkg/utils/fs"
	"github.com/shaharby7/Dope/pkg/utils/set"
)

type appPathArgs struct {
	App string
}

func generateAppFiles(config *entity.Entity, appConfig *entity.Entity) ([]*fsUtils.OutputFile, error) {
	pathArgs := &appPathArgs{App: appConfig.Name}
	mainFile, err := generateOutputFile(
		templateId_SRC_APP_MAIN,
		pathArgs,
		utils.Empty,
	)
	if err != nil {
		return nil, err
	}
	controllersFile, err := generateOutputFile(
		templateId_SRC_APP_CONTROLLERS,
		pathArgs,
		generateControllerData(
			config,
			appConfig,
		),
	)
	if err != nil {
		return nil, err
	}
	return []*fsUtils.OutputFile{mainFile, controllersFile}, nil
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

type clientPathArgs struct {
	Client string
}

type clientDataArgs struct {
	Name    string
	Imports []string
	Actions []struct {
		FlattenName string
		Name        string
		Caller      string
		Method      string
		App         string
	}
}

func generateClientFiles(
	eTree *entities.EntitiesTree,
	clientConf *entity.Entity,
) ([]*fsUtils.OutputFile, error) {
	pathArgs := &clientPathArgs{Client: clientConf.Name}
	clientFile, err := generateOutputFile(
		templateId_SRC_CLIENT_MAIN,
		pathArgs,
		generateClientData(eTree, clientConf),
	)
	if err != nil {
		return nil, err
	}
	return []*fsUtils.OutputFile{clientFile}, nil
}

func generateClientData(
	eTree *entities.EntitiesTree,
	clientConf *entity.Entity,
) *clientDataArgs {
	data := &clientDataArgs{
		Name: clientConf.Name,
	}
	imports := set.NewSet[string]()
	clientValues := entity.GetEntityValues[v1.ClientConfig](clientConf)
	project, _ := configHelpers.GetProject(eTree)
	projectValues := entity.GetEntityValues[v1.ProjectConfig](project)
	for _, app := range clientValues.Apps {
		appEntity, err := configHelpers.GetApp(eTree, app)
		if err != nil {
			continue
		}
		appConfig := entity.GetEntityValues[v1.AppConfig](appEntity)
		for _, controller := range appConfig.Controllers {
			for _, action := range controller.Actions {
				imports.Add(
					path.Join(projectValues.Module, action.Package),
				)
				data.Actions = append(data.Actions, struct {
					FlattenName string
					Name        string
					Caller      string
					Method      string
					App         string
				}{
					Name:        action.Name,
					FlattenName: flattenString(action.Name),
					Caller: fmt.Sprintf(
						"%s.%s", path.Base(action.Package), action.Ref,
					),
					Method: action.ControllerBinding["method"],
					App:    app,
				})
			}
		}
	}
	data.Imports = imports.ToSlice()
	return data
}

func flattenString(s string) string {
	s = strings.Replace(s, "/", "_", -1)
	s = strings.Replace(s, "-", "_", -1)
	s = strings.Replace(s, ":", "_", -1)
	return s
}
