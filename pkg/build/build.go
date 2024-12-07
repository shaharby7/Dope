package build

import (
	"fmt"

	"github.com/shaharby7/Dope/pkg/build/files"
	bTypes "github.com/shaharby7/Dope/pkg/build/types"

	"github.com/shaharby7/Dope/pkg/config"
	"github.com/shaharby7/Dope/pkg/utils"
	fsUtils "github.com/shaharby7/Dope/pkg/utils/fs"
	"github.com/shaharby7/Dope/types"
)

func BuildProject(
	dopePath string,
	dst string,
	options bTypes.BuildOptions,
) error {
	metadata, err := getBuildMetadata(dst)
	if err != nil {
		return utils.FailedBecause(
			"generate build metadata",
			err,
		)
	}
	config, err := config.ReadConfig(dopePath)
	if err != nil {
		return utils.FailedBecause(
			fmt.Sprintf("generate config from file (%s)", dopePath),
			err,
		)
	}
	appsList := getApplicationsList(config, &options)
	envsList := getEnvironmentList(config, &options)

	outputFiles, err := files.GenerateFiles(
		dst, config, metadata, appsList, envsList,
	)
	if err != nil {
		return utils.FailedBecause(
			"compile files",
			err,
		)
	}
	err = fsUtils.WriteFiles(dst, outputFiles)
	if err != nil {
		return fmt.Errorf("could not write files: %w", err)
	}
	return nil
}

func getApplicationsList(config *types.ProjectConfig, buildOptions *bTypes.BuildOptions) []string {
	if len(buildOptions.Apps) > 0 {
		return buildOptions.Apps
	}
	appsList, _ := utils.Map(config.Apps, func(app *types.AppConfig) (string, error) { return app.Name, nil })
	return utils.RemoveDuplicates(appsList)
}

func getEnvironmentList(config *types.ProjectConfig, buildOptions *bTypes.BuildOptions) []string {
	if len(buildOptions.Envs) > 0 {
		return buildOptions.Apps
	}
	envsList, _ := utils.Map(config.Environments, func(env *types.EnvConfig) (string, error) { return env.Name, nil })
	return utils.RemoveDuplicates(envsList)
}

func getBuildMetadata(buildPath string) (*bTypes.BuildMetadata, error) {
	gitRef, err := utils.GetGitHEADRef()
	if err != nil {
		return nil, err
	}
	return &bTypes.BuildMetadata{
		GitRef:    gitRef,
		BuildPath: buildPath,
	}, nil
}
