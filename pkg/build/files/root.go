package files

import (
	"github.com/shaharby7/Dope/pkg/utils"
	"github.com/shaharby7/Dope/types"
)

func rooFilesGenGen(
	_ *types.ProjectConfig,
) []iFileGenerator {
	dockerfile := newFileGenerator(
		templateId_DOCKERFILE,
		utils.Empty,
		utils.Empty,
	)
	return []iFileGenerator{dockerfile}
}
