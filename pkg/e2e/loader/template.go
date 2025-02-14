package loader

import (
	"bytes"
	"text/template"
)

const (
	templateName = "main.go"
)

func createMainFile(mainFile *sMainFileTmplInput) (*bytes.Buffer, error) {
	var result bytes.Buffer
	err := tmpl.ExecuteTemplate(&result, templateName, mainFile)
	return &result, err

}

type sMainFileTmplInput struct {
	PkgPath         string
	TestIdentifiers []string
}

var tmpl *template.Template = template.Must(template.New(templateName).Parse(`
package main

import (
	"fmt"
	"os"

	__pgk_e2e "{{ .PkgPath }}"
	"github.com/shaharby7/Dope/pkg/e2e"
	"github.com/shaharby7/Dope/pkg/utils"
)

func main() {
	err := e2e.RunE2E(
		"./.dope",
		"./build",
		&e2e.E2eOptions{
			InstallBefore: utils.Ptr(false),
		},
		[]*e2e.E2ETestCase{
			{{- range .TestIdentifiers }}
			{
				Name: "{{ . }}",
				Func: __pgk_e2e.{{ . }},
			},
			{{- end }}
		},
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("E2E tests passed")
}
`))
