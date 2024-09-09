package files

import (
	"bytes"
	"fmt"
	"html/template"
	"path"

	"github.com/shaharby7/Dope/pkg/utils"
)

func generateFileByTemplate[FileData any, PathArgs any](
	basePath string,
	tmpl TEMPLATE_FILES,
	fileInput *FileGenerationInput[FileData, PathArgs],
) (*OutputFile, error) {
	filePathTemplate := path.Join(basePath, string(tmpl))
	var filePath string
	var err error
	var content string
	if fileInput != nil {
		buff := &bytes.Buffer{}
		t, err := template.New("temp").Parse(filePathTemplate)
		if err != nil {
			return nil, utils.FailedBecause(fmt.Sprintf("generate template (%s)", tmpl), err)
		}
		err = t.Execute(buff, fileInput.Params.Path)
		if err != nil {
			return nil, utils.FailedBecause(fmt.Sprintf("evaluate file name (%s)", tmpl), err)
		}
		filePath = buff.String()
	} else {
		filePath = filePathTemplate
	}
	var data FileData
	if fileInput != nil {
		data = fileInput.Data
	}
	buff := &bytes.Buffer{}
	err = getTemplate(tmpl).Execute(buff, data)
	if err != nil {
		return nil, utils.FailedBecause(fmt.Sprintf("parse template (%s)", tmpl), err)
	}
	content = buff.String()
	return &OutputFile{
		Path:    filePath,
		Content: content,
	}, nil
}
