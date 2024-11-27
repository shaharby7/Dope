package files

import (
	"github.com/shaharby7/Dope/pkg/utils"

	fsUtils "github.com/shaharby7/Dope/pkg/utils/fs"
	"github.com/shaharby7/Dope/types"

)

func generateRootFiles(
	_ *types.ProjectConfig,
) ([]*fsUtils.OutputFile, error) {
	dockerfile, err := generateOutputFile(
		templateId_DOCKERFILE,
		utils.Empty,
		utils.Empty,
	)
	return []*fsUtils.OutputFile{dockerfile}, err
}
