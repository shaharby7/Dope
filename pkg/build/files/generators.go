package files

import (
	"bytes"
	"fmt"

	"github.com/shaharby7/Dope/pkg/utils"
	"html/template"
)

var EMPTY_TEMPLATE_INPUT *struct{ A string } = &struct{ A string }{A: "A"}

type iFileGenerator interface {
	Generate() (*OutputFile, error)
}

type fileGenerator[TFileData any, TPathArgs any] struct {
	GeneratorId string
	TemplateId  templateId
	PathArgs    TPathArgs
	Data        TFileData
}

func newFileGenerator[TFileData any, TPathArgs any](
	templateId templateId,
	pathArgs TPathArgs,
	data TFileData,
) *fileGenerator[TFileData, TPathArgs] {
	return &fileGenerator[TFileData, TPathArgs]{
		TemplateId: templateId, PathArgs: pathArgs, Data: data,
	}
}

func (generator *fileGenerator[TFileData, TPathArgs]) Generate() (*OutputFile, error) {
	fileTemplate := getTemplate(generator.TemplateId)
	path, err := applyTemplateSafe(&fileTemplate.PathTemplate, generator.PathArgs)
	if err != nil {
		return nil, utils.FailedBecause(fmt.Sprintf("generate file path from generator %s", generator.GeneratorId), err)
	}
	data, err := applyTemplateSafe(&fileTemplate.DataTemplate, generator.Data)
	if err != nil {
		return nil, utils.FailedBecause(fmt.Sprintf("generate file path from generator %s", generator.GeneratorId), err)
	}
	return &OutputFile{
		Path:    path.String(),
		Content: data.String(),
	}, nil
}

func applyTemplateSafe(template *template.Template, args any) (*bytes.Buffer, error) {
	var path bytes.Buffer
	var err error
	if utils.IsEmpty(args) {
		err = template.Execute(&path, EMPTY_TEMPLATE_INPUT)
	} else {
		err = template.Execute(&path, args)
	}
	return &path, err
}
