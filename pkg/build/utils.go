package build

import (
	"bytes"
	"fmt"
	"os"
	"path"
)

func generateFileByTemplate[FileData any](
	basePath string,
	tmpl TEMPLATE_FILES,
	fileData *FileData,
) (*BuiltFile, error) {
	filePath := path.Join(basePath, string(tmpl))
	content := &bytes.Buffer{}
	err := getTemplate(tmpl).Execute(content, fileData)
	if err != nil {
		return nil, fmt.Errorf("could not parse template (%s): %w", tmpl, err)
	}
	return &BuiltFile{
		Path:    filePath,
		Content: content.String()}, nil
}

func ensurePath(args ...string) string {
	p := path.Join(args...)
	os.MkdirAll(p, os.ModePerm)
	return p
}

func writeFiles(files []*BuiltFile) error {
	for _, file := range files {
		err := writeFile(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeFile(file *BuiltFile) error {
	ensurePath(path.Dir(file.Path))
	fileRef, err := os.Create(file.Path)
	if err != nil {
		return fmt.Errorf("could not open file %s: %w", file.Path, err)
	}
	defer fileRef.Close()
	_, err = fileRef.Write([]byte(file.Content))
	if err != nil {
		return fmt.Errorf("could not write file (%s): %w", file.Path, err)
	}
	return nil
}
