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
	templates map[TEMPLATE_FILES]*template.Template
)

func loadTemplates() {
	if templates == nil {
		templates = make(map[TEMPLATE_FILES]*template.Template)
	}
	err := fs.WalkDir(
		files,
		".",
		func(path string, dirEntry fs.DirEntry, err error) error {
			if err != nil {
				return fmt.Errorf("could not read template file (%s) because: %w", path, err)
			}
			if dirEntry.IsDir() {
				return nil
			}
			pt, err := template.ParseFS(
				files,
				path,
			)
			if err != nil {
				return fmt.Errorf("could not parse (%s):%w", path, err)
			}
			tmplName := TEMPLATE_FILES(strings.Replace(
				strings.Replace(path, extension, "", 1),
				templatesDir+"/",
				"",
				1,
			))
			templates[tmplName] = pt
			return nil
		},
	)
	if err != nil {
		panic(fmt.Errorf("could not load templates: %w", err))
	}
}

func getTemplate(tmplName TEMPLATE_FILES) *template.Template {
	tmpl := templates[tmplName]
	if tmpl == nil {
		panic(fmt.Errorf("could not find template %s", tmplName))
	}
	return tmpl
}
