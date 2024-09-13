package files

import (
	"embed"
	"fmt"
	"html/template"
	"path"

	"github.com/shaharby7/Dope/pkg/utils"
)

const (
	_TEMPLATES_DIRECTORY = "templates"
	_TEMPLATE_EXTENSION  = "tmpl"
)

type templateId int

const (
	templateId_SRC_FILE_MAIN templateId = iota
	templateId_SRC_FILE_CONTROLLER
	templateId_DOCKERFILE
)

var _TEMPLATES_LIST map[templateId]string = map[templateId]string{
	templateId_SRC_FILE_MAIN:       "src/{{.App}}/main.go",
	templateId_SRC_FILE_CONTROLLER: "src/{{.App}}/controllers.go",
	templateId_DOCKERFILE:          "Dockerfile",
}

type fileTemplate struct {
	TemplateId          templateId
	_pathTemplateString string
	PathTemplate        template.Template
	DataTemplate        template.Template
}

var (
	//go:embed templates/*
	osFiles embed.FS
)

func newFileTemplate(templateId templateId, pathTemplate string) *fileTemplate {
	parsedPathTemplate, err := template.New(pathTemplate).Parse(pathTemplate)
	if err != nil {
		panic(utils.FailedBecause(fmt.Sprintf("parse path template %s", pathTemplate), err))
	}
	dataTemplateFullPath := path.Join(
		_TEMPLATES_DIRECTORY, fmt.Sprintf("%s.%s", pathTemplate, _TEMPLATE_EXTENSION),
	)
	parsedFileTemplate, err := template.ParseFS(
		osFiles, dataTemplateFullPath,
	)
	if err != nil {
		panic(utils.FailedBecause(fmt.Sprintf("parse file template %s", pathTemplate), err))
	}
	return &fileTemplate{
		TemplateId:          templateId,
		_pathTemplateString: pathTemplate,
		PathTemplate:        *parsedPathTemplate,
		DataTemplate:        *parsedFileTemplate,
	}
}

var registeredTemplates map[templateId]*fileTemplate = make(map[templateId]*fileTemplate, 0)

func init() {
	for id, path := range _TEMPLATES_LIST {
		registeredTemplates[id] = newFileTemplate(id, path)
	}
}

func getTemplate(templateId templateId) *fileTemplate {
	return registeredTemplates[templateId]
}
