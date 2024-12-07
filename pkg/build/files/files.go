package files

import (
	bTypes "github.com/shaharby7/Dope/pkg/build/types"

	configHelpers "github.com/shaharby7/Dope/pkg/config/helpers"

	fsUtils "github.com/shaharby7/Dope/pkg/utils/fs"
	"github.com/shaharby7/Dope/pkg/utils"
	"github.com/shaharby7/Dope/types"
)

func GenerateFiles(
	dst string,
	config *types.ProjectConfig,
	metadata *bTypes.BuildMetadata,
	appsList []string,
	envsList []string,
) ([]*fsUtils.OutputFile, error) {
	files := make([]*fsUtils.OutputFile, 0)

	f, err := generateRootFiles(config)
	if err != nil {
		return nil, utils.FailedBecause("generate root files", err)
	}
	files = append(files, f...)
	for _, env := range envsList {
		envConf, err := configHelpers.GetEnvConfig(config, env)
		if err != nil {
			return nil, err
		}
		f, err := generateHelmDopeValuesFile(metadata, config, envConf)
		if err != nil {
			return nil, utils.FailedBecause("generate helm dope values file", err)
		}
		files = append(files, f...)
	}
	for _, app := range appsList {
		appConf, err := configHelpers.GetAppConfig(config, app)
		if err != nil {
			return nil, err
		}
		f, err := generateSrcFiles(config, appConf)
		if err != nil {
			return nil, utils.FailedBecause("generate src files", err)
		}
		files = append(files, f...)
		for _, env := range envsList {
			appEnvConf, err := configHelpers.GetAppEnvConfig(config, env, app)
			if err != nil {
				return nil, err
			}
			f, err := generateAppHelmFiles(config, env, appConf, appEnvConf)
			if err != nil {
				return nil, utils.FailedBecause("generate helm values files", err)
			}
			files = append(files, f...)
		}
	}
	return files, nil
}
