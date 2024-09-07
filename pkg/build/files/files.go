package files

import (
	"github.com/shaharby7/Dope/types"

	"github.com/shaharby7/Dope/pkg/utils"

	buildUtils "github.com/shaharby7/Dope/pkg/build/utils"
)

func GenerateFiles(
	dst string,
	config *types.ProjectConfig,
	appsList []string,
	envsList []string,
) ([]*OutputFile, error) {
	files, err := generateRootFiles(dst, config)
	if err != nil {
		return nil, utils.FailedBecause("generate root files", err)
	}
	for _, app := range appsList {
		appConf, err := buildUtils.GetAppConfig(config, app)
		if err != nil {
			return files, err
		}
		srcFiles, err := generateSrcFiles(
			dst,
			config,
			appConf,
		)
		if err != nil {
			return files, utils.FailedBecause("generating source files", err)
		}
		files = append(files, srcFiles...)
		for _, env := range envsList {
			appEnvConf, err := buildUtils.GetAppEnvConfig(config, env, app)
			if err != nil {
				return files, err
			}
			helmFiles, err := generateHelmFiles(
				dst,
				config,
				appConf,
				appEnvConf,
			)
			if err != nil {
				return files, utils.FailedBecause("generating helm files", err)
			}
			files = append(files, helmFiles...)
		}
	}
	return files, nil
}
