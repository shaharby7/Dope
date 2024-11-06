package files

import (
	"bytes"
	"fmt"

	"text/template"

	"github.com/shaharby7/Dope/pkg/utils"
)

var EMPTY_TEMPLATE_INPUT *struct{ A string } = &struct{ A string }{A: "A"}

func generateOutputFile[TFileData any, TPathArgs any](
	templateId templateId,
	pathArgs TPathArgs,
	dataArgs TFileData,
) (*OutputFile, error) {
	fileTemplate := getTemplate(templateId)
	path, err := applyTemplateSafe(&fileTemplate.PathTemplate, fileTemplate.Name, pathArgs)
	if err != nil {
		return nil, utils.FailedBecause(fmt.Sprintf("generate file path from generator %s", fileTemplate.Name), err)
	}
	data, err := applyTemplateSafe(&fileTemplate.DataTemplate, fileTemplate.Name, dataArgs)
	if err != nil {
		return nil, utils.FailedBecause(fmt.Sprintf("generate file path from generator %s", fileTemplate.Name), err)
	}
	return &OutputFile{
		Path:    path.String(),
		Content: data.String(),
	}, nil
}

func applyTemplateSafe(template *template.Template, templateName string, args any) (*bytes.Buffer, error) {
	var result bytes.Buffer
	var err error
	if utils.IsEmpty(args) {
		err = template.ExecuteTemplate(&result, templateName, EMPTY_TEMPLATE_INPUT)
	} else {
		err = template.ExecuteTemplate(&result, templateName, args)
	}
	return &result, err
}
