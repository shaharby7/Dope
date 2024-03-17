package build

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"github.com/shaharby7/Dope/pkg/types"
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
		file, err := os.Create(filepath.Join(appDst, "main.go"))
		if err != nil {
			return fmt.Errorf("could not create main file for app (%s): %w", appConfig.Name, err)
		}
		defer file.Close()
		input, err := createAppTemplateInput(&config.Metadata, &appConfig)
		if err != nil {
			return fmt.Errorf("invalid config for file app (%s): %w", appConfig.Name, err)
		}
		err = getTemplate("app_main.go").Execute(file, input)
		if err != nil {
			return fmt.Errorf("could not build src files for app (%s): %w", appConfig.Name, err)
		}
	}
	return nil
}

func createAppTemplateInput(metadataConfig *types.ProjectMetadataConfig, appConfig *types.AppConfig) (*appTemplateInput, error) {
	return &appTemplateInput{
		AppConfig: appConfig,
		ProjectMetadataConfig: metadataConfig,
		Imports: []string{},
	}, nil
}
