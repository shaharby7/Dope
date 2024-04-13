package build

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"github.com/shaharby7/Dope/pkg/utils"
	"github.com/shaharby7/Dope/types"
	"gopkg.in/yaml.v3"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func BuildProject(projectFile string, dst string) error {
	loadTemplates()
	config, err := getConfigByPath(projectFile)
	if err != nil {
		return fmt.Errorf("could not generate config from file (%s): %w", projectFile, err)
	}
	err = buildSrcFiles(dst, config)
	if err != nil {
		return fmt.Errorf("could not build src files for: %w", err)
	}
	return nil
}

func getConfigByPath(projectFile string) (*types.ProjectConfig, error) {
	path, err := filepath.Abs(projectFile)
	if err != nil {
		return nil, fmt.Errorf("config validation error: %w", err)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("config validation error: %w", err)
	}
	var config *types.ProjectConfig
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, fmt.Errorf("config validation error: %w", err)
	}
	err = validate.Struct(config)
	if err != nil {
		return nil, fmt.Errorf("config validation error: %w", err)
	}
	return config, nil
}

func buildSrcFiles(dst string, config *types.ProjectConfig) error {
	os.MkdirAll(dst, os.ModePerm)
	for _, appConfig := range config.Apps {
		var appDst = filepath.Join(dst, appConfig.Name)
		os.MkdirAll(appDst, os.ModePerm)
		err := createFileByTemplate(appDst, SRC_FILE_MAIN, &mainFileInput{})
		if err != nil {
			return wrapSrcFileCreationError(appConfig.Name, SRC_FILE_CONTROLLER, err)
		}
		err = createFileByTemplate(appDst, SRC_FILE_CONTROLLER, generateControllerInput(
			config,
			appConfig,
		))
		if err != nil {
			return wrapSrcFileCreationError(appConfig.Name, SRC_FILE_CONTROLLER, err)
		}
	}
	return nil
}

func generateControllerInput(
	config *types.ProjectConfig,
	appConfig types.AppConfig,
) *controllerFileInput {
	imports := utils.NewSet[string]()
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
				ControllerBinding: actionConfig.ControllerBinding,
			}
			controller.Actions = append(controller.Actions, action)
		}
		controllers = append(controllers, controller)
	}

	return &controllerFileInput{
		Imports:     imports.ToSlice(),
		Controllers: controllers,
	}
}

func createFileByTemplate[FileInput any](appDst string, tmpl SRC_FILES, input *FileInput) error {
	file, err := os.Create(filepath.Join(appDst, string(tmpl)))
	if err != nil {
		return fmt.Errorf("could not open file %s: %w", tmpl, err)
	}
	defer file.Close()
	err = getTemplate(string(tmpl)).Execute(file, input)
	if err != nil {
		return fmt.Errorf("could not parse template (%s): %w", tmpl, err)
	}
	return nil
}

func wrapSrcFileCreationError(appName string, fileName SRC_FILES, err error) error {
	return fmt.Errorf("could not could not create src file %s for app %s: %w", fileName, appName, err)
}
