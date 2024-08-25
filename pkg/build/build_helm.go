package build

import "github.com/shaharby7/Dope/types"

func buildHelmFiles(dst string,
	config *types.ProjectConfig,
	appConfig *types.AppConfig,
	environmentConfig *types.EnvConfig,
) ([]*BuiltFile, error) {
	return make([]*BuiltFile, 0), nil
}
