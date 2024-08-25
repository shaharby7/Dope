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

func init() {
	loadTemplates()
}

func BuildProject(projPath string, dst string) error {
	config, err := getConfigByPath(projPath)
	if err != nil {
		return fmt.Errorf("could not generate config from file (%s): %w", projPath, err)
	}
	for _, appConfig := range config.Apps {
		srcFiles, err := buildSrcFiles(dst, &config, &appConfig)
		if err != nil {
			return fmt.Errorf("could not build src files for: %w", err)
		}
		helmFiles := make([]*BuiltFile, 0)
		for _, environmentConfig := range config.Environments {
			envHelmFiles, err := buildHelmFiles(dst, &config, &appConfig, &environmentConfig)
			if err != nil {
				return fmt.Errorf("could not build helm files because: %w", err)
			}
			helmFiles = append(helmFiles, envHelmFiles...)
		}
		writeFiles(append(srcFiles, helmFiles...))
	}
	return nil
}

func getConfigByPath(projectFile string) (types.ProjectConfig, error) {
	config := types.ProjectConfig{}
	path, err := filepath.Abs(projectFile)
	if err != nil {
		return config, fmt.Errorf("config validation error: %w", err)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("config validation error: %w", err)
	}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return config, fmt.Errorf("config validation error: %w", err)
	}
	err = validate.Struct(config)
	if err != nil {
		return config, fmt.Errorf("config validation error: %w", err)
	}
	return config, nil
}
