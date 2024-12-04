package files

import (
	"fmt"
	"github.com/shaharby7/Dope/pkg/utils"

	fsUtils "github.com/shaharby7/Dope/pkg/utils/fs"
)


func generateOutputFile[TFileData any, TPathArgs any](
	templateId templateId,
	pathArgs TPathArgs,
	dataArgs TFileData,
) (*fsUtils.OutputFile, error) {
	fileTemplate := getTemplate(templateId)
	path, err := utils.ApplyTemplateSafe(&fileTemplate.PathTemplate, fileTemplate.Name, pathArgs)
	if err != nil {
		return nil, utils.FailedBecause(fmt.Sprintf("generate file path from generator %s", fileTemplate.Name), err)
	}
	data, err := utils.ApplyTemplateSafe(&fileTemplate.DataTemplate, fileTemplate.Name, dataArgs)
	if err != nil {
		return nil, utils.FailedBecause(fmt.Sprintf("generate file path from generator %s", fileTemplate.Name), err)
	}
	return &fsUtils.OutputFile{
		Path:    path.String(),
		Content: data.String(),
	}, nil
}
