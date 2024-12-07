package files

import (
	"github.com/shaharby7/Dope/pkg/config"
	"github.com/shaharby7/Dope/pkg/utils"

	fsUtils "github.com/shaharby7/Dope/pkg/utils/fs"
)

func generateRootFiles(
	_ *config.EntitiesTree,
) ([]*fsUtils.OutputFile, error) {
	dockerfile, err := generateOutputFile(
		templateId_DOCKERFILE,
		utils.Empty,
		utils.Empty,
	)
	return []*fsUtils.OutputFile{dockerfile}, err
}
