package files

import (
	"github.com/shaharby7/Dope/pkg/utils"
	"github.com/shaharby7/Dope/types"
)

func generateRootFiles(
	_ *types.ProjectConfig,
) ([]*OutputFile, error) {
	dockerfile, err := generateOutputFile(
		templateId_DOCKERFILE,
		utils.Empty,
		utils.Empty,
	)
	return []*OutputFile{dockerfile}, err
}
