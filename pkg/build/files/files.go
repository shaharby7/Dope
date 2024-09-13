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
	generators := rooFilesGenGen(config)
	for _, app := range appsList {
		appConf, err := buildUtils.GetAppConfig(config, app)
		if err != nil {
			return nil, utils.FailedBecause("get app config", err)
		}
		srcFiles := srcFilesGenGen(
			config,
			appConf,
		)
		generators = append(generators, srcFiles...)
		for _, env := range envsList {
			appEnvConf, err := buildUtils.GetAppEnvConfig(config, env, app)
			if err != nil {
				return nil, err
			}
			helmFiles := helmFilesGenGen(
				config,
				appConf,
				appEnvConf,
			)
			generators = append(generators, helmFiles...)
		}
	}
	files, err := utils.Map(generators, func(g iFileGenerator) (*OutputFile, error) {
		return g.Generate()
	})
	if err != nil {
		return nil, utils.FailedBecause("generate files", err)
	}
	return files, nil
}
