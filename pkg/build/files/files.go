package files

import (
	"github.com/shaharby7/Dope/pkg/build/helpers"
	bTypes "github.com/shaharby7/Dope/pkg/build/types"
	"github.com/shaharby7/Dope/pkg/utils"
	"github.com/shaharby7/Dope/types"
)

func GenerateFiles(
	dst string,
	config *types.ProjectConfig,
	metadata *bTypes.BuildMetadata,
	appsList []string,
	envsList []string,
) ([]*OutputFile, error) {
	files := make([]*OutputFile, 0)

	f, err := generateRootFiles(config)
	if err != nil {
		return nil, utils.FailedBecause("generate root files", err)
	}
	files = append(files, f...)
	for _, app := range appsList {
		appConf, err := helpers.GetAppConfig(config, app)
		if err != nil {
			return nil, utils.FailedBecause("get app config", nil)
		}
		f, err := generateSrcFiles(config, appConf)
		if err != nil {
			return nil, utils.FailedBecause("generate src files", err)
		}
		files = append(files, f...)
		for _, env := range envsList {
			appEnvConf, err := helpers.GetAppEnvConfig(config, env, app)
			if err != nil {
				return nil, utils.FailedBecause("get app env config", nil)
			}
			f, err := generateHelmFiles(config, metadata, env, appConf, appEnvConf)
			if err != nil {
				return nil, utils.FailedBecause("generate helm files", err)
			}
			files = append(files, f...)
		}
	}
	return files, nil
}
