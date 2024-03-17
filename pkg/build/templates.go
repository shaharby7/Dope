package build

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"strings"
)

const (
	templatesDir = "templates"
	extension    = ".tmpl"
)

var (
	//go:embed templates/*
	files     embed.FS
	templates map[string]*template.Template
)

func loadTemplates() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	tmplFiles, err := fs.ReadDir(files, templatesDir)
	if err != nil {
		panic(fmt.Errorf("could not read templates dir (%s):%w", templatesDir, err))
	}
	for _, tmpl := range tmplFiles {
		if tmpl.IsDir() {
			continue
		}
		tmplName := strings.Replace(tmpl.Name(), extension, "", 1)
		pt, err := template.ParseFS(files, templatesDir+"/"+tmpl.Name())
		if err != nil {
			panic(fmt.Errorf("could not load template (%s): %w", tmplName, err))
		}
		templates[tmplName] = pt
	}
}

func getTemplate(tmplName string) *template.Template {
	tmpl := templates[tmplName]
	if tmpl == nil {
		panic(fmt.Errorf("could not find template %s", tmplName))
	}
	return tmpl
}
