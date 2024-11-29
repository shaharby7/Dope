package build

import (
	"fmt"

	"github.com/shaharby7/Dope/pkg/build/files"
	bTypes "github.com/shaharby7/Dope/pkg/build/types"
	v1 "github.com/shaharby7/Dope/pkg/config/V1"
	"github.com/shaharby7/Dope/pkg/config/entity"

	"github.com/shaharby7/Dope/pkg/config"
	configHelpers "github.com/shaharby7/Dope/pkg/config/helpers"
	"github.com/shaharby7/Dope/pkg/utils"
	fsUtils "github.com/shaharby7/Dope/pkg/utils/fs"
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
	eTree, err := config.LoadEntitiesTree(dopePath)
	if err != nil {
		return utils.FailedBecause(
			fmt.Sprintf("generate config from file (%s)", dopePath),
			err,
		)
	}
	appsList := getApplicationsList(eTree, &options)
	envsList := getEnvironmentList(eTree, &options)

	outputFiles, err := files.GenerateFiles(
		dst, eTree, metadata, appsList, envsList,
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

func getApplicationsList(eTree *config.EntitiesTree, buildOptions *bTypes.BuildOptions) []string {
	if len(buildOptions.Apps) > 0 {
		return buildOptions.Apps
	}
	return getEntitiesNamesListByType(eTree, v1.DOPE_CORE_TYPES_APP)
}

func getEnvironmentList(eTree *config.EntitiesTree, buildOptions *bTypes.BuildOptions) []string {
	if len(buildOptions.Envs) > 0 {
		return buildOptions.Apps
	}
	return getEntitiesNamesListByType(eTree, v1.DOPE_CORE_TYPES_ENV)
}

func getEntitiesNamesListByType(eTree *config.EntitiesTree, t v1.DOPE_CORE_TYPES) []string {
	apps := configHelpers.GetCoreEntitiesByType(*eTree, t)
	appsList, _ := utils.Map(apps, func(e *entity.Entity) (string, error) { return e.Name, nil })
	return utils.RemoveDuplicates(appsList)
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
