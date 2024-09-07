package files

import (
	"github.com/shaharby7/Dope/pkg/utils"
	"github.com/shaharby7/Dope/types"
)

func generateRootFiles(
	dst string,
	_ *types.ProjectConfig,
) ([]*OutputFile, error) {
	dockerfile, err := generateFileByTemplate[any](
		dst, DOCKERFILE, nil,
	)
	if err != nil {
		return nil, utils.FailedBecause("generate dockerfile", err)
	}
	return []*OutputFile{dockerfile}, nil
}
