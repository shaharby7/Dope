package files

import (
	bTypes "github.com/shaharby7/Dope/pkg/build/types"
	"github.com/shaharby7/Dope/pkg/entities"

	configHelpers "github.com/shaharby7/Dope/pkg/entities/helpers"

	"github.com/shaharby7/Dope/pkg/utils"
	fsUtils "github.com/shaharby7/Dope/pkg/utils/fs"
)

func GenerateFiles(
	dst string,
	eTree *entities.EntitiesTree,
	metadata *bTypes.BuildMetadata,
	appsList []string,
	envsList []string,
	clientsList []string,
) ([]*fsUtils.OutputFile, error) {
	files := make([]*fsUtils.OutputFile, 0)

	f, err := generateRootFiles(eTree)
	if err != nil {
		return nil, utils.FailedBecause("generate root files", err)
	}
	files = append(files, f...)

	project, err := configHelpers.GetProject(eTree)
	if err != nil {
		return nil, err
	}
	for _, env := range envsList {
		envEntity, err := configHelpers.GetEnv(eTree, env)
		if err != nil {
			return nil, err
		}
		f, err := generateHelmDopeValuesFile(metadata, eTree, project, envEntity)
		if err != nil {
			return nil, utils.FailedBecause("generate helm dope values file", err)
		}
		files = append(files, f...)
	}
	for _, client := range clientsList {
		clientConf, err := configHelpers.GetClient(eTree, client)
		if err != nil {
			return nil, err
		}
		f, err := generateClientFiles(eTree, clientConf)
		if err != nil {
			return nil, utils.FailedBecause("generate src files", err)
		}
		files = append(files, f...)
	}
	for _, app := range appsList {
		appConf, err := configHelpers.GetApp(eTree, app)
		if err != nil {
			return nil, err
		}
		f, err := generateAppFiles(project, appConf)
		if err != nil {
			return nil, utils.FailedBecause("generate src files", err)
		}
		files = append(files, f...)
		for _, env := range envsList {
			appEnvConf, err := configHelpers.GetAppEnvConfig(eTree, env, app)
			if err != nil {
				return nil, err
			}
			f, err := generateAppHelmFiles(eTree, env, appConf, appEnvConf)
			if err != nil {
				return nil, utils.FailedBecause("generate helm values files", err)
			}
			files = append(files, f...)
		}
	}
	return files, nil
}