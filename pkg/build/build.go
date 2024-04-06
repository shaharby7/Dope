package build

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
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
		err := createFileByTemplate(appDst, "app_main.go", SRC_FILE_MAIN, &mainFileInput{})
		if err != nil {
			return wrapSrcFileCreationError(appConfig.Name, SRC_FILE_CONTROLLER, err)
		}
		err = createFileByTemplate(appDst, "controller.go", SRC_FILE_CONTROLLER, generateControllerInput(
			appConfig,
		))
		if err != nil {
			return wrapSrcFileCreationError(appConfig.Name, SRC_FILE_CONTROLLER, err)
		}
	}
	return nil
}

func generateControllerInput(appConfig types.AppConfig) *controllerFileInput {
	return nil
}

func createFileByTemplate[FileInput any](appDst string, templateName string, dstFile SRC_FILES, input *FileInput) error {
	file, err := os.Create(filepath.Join(appDst, string(dstFile)))
	if err != nil {
		return fmt.Errorf("could not open file %s: %w", dstFile, err)
	}
	defer file.Close()
	err = getTemplate(templateName).Execute(file, input)
	if err != nil {
		return fmt.Errorf("could not parse template (%s): %w", templateName, err)
	}
	return nil
}

func wrapSrcFileCreationError(appName string, fileName SRC_FILES, err error) error {
	return fmt.Errorf("could not could not create src file %s for app %s: %w", fileName, appName, err)
}
