package build

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"github.com/shaharby7/Dope/pkg/utils"

	"github.com/shaharby7/Dope/pkg/build/files"
	"github.com/shaharby7/Dope/types"
	"gopkg.in/yaml.v3"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func BuildProject(projPath string, dst string, options BuildOptions) error {
	config, err := getConfigByPath(projPath)
	if err != nil {
		return fmt.Errorf("could not generate config from file (%s): %w", projPath, err)
	}
	appsList := getApplicationsList(config, &options)
	envsList := getEnvironmentList(config, &options)

	outputFiles, err := files.GenerateFiles(
		dst, config, appsList, envsList,
	)
	if err != nil {
		utils.FailedBecause(
			"compile files",
			err,
		)
	}
	err = writeFiles(dst, outputFiles)
	if err != nil {
		return fmt.Errorf("could not write files: %w", err)
	}
	return nil
}

func getConfigByPath(projectFile string) (*types.ProjectConfig, error) {
	config := &types.ProjectConfig{}
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

func getApplicationsList(config *types.ProjectConfig, buildOptions *BuildOptions) []string {
	if len(buildOptions.Apps) > 0 {
		return buildOptions.Apps
	}
	appsList, _ := utils.Map(config.Apps, func(app types.AppConfig) (string, error) { return app.Name, nil })
	return utils.RemoveDuplicates(appsList)
}

func getEnvironmentList(config *types.ProjectConfig, buildOptions *BuildOptions) []string {
	if len(buildOptions.Envs) > 0 {
		return buildOptions.Apps
	}
	envsList, _ := utils.Map(config.Environments, func(env types.EnvConfig) (string, error) { return env.Name, nil })
	return utils.RemoveDuplicates(envsList)
}
