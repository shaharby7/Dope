package files

import (
	"github.com/shaharby7/Dope/pkg/utils"
	"github.com/shaharby7/Dope/types"
)

type helmPathArgs struct {
	App string
	Env string
}

func helmFilesGenGen(
	config *types.ProjectConfig,
	appConfig *types.AppConfig,
	appEnvConfig *types.AppEnvConfig,
) []iFileGenerator {
	pathArgs := &helmPathArgs{App: appConfig.Name, Env: appEnvConfig.Name}
	mainFile := newFileGenerator(
		templateId_SRC_FILE_MAIN,
		pathArgs,
		utils.Empty,
	)
	return []iFileGenerator{mainFile}
}
