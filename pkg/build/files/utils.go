package files

import (
	"bytes"
	"fmt"
	"path"
)

func generateFileByTemplate[FileData any](
	basePath string,
	tmpl TEMPLATE_FILES,
	fileData *FileData,
) (*OutputFile, error) {
	filePath := path.Join(basePath, string(tmpl))
	content := &bytes.Buffer{}
	err := getTemplate(tmpl).Execute(content, fileData)
	if err != nil {
		return nil, fmt.Errorf("could not parse template (%s): %w", tmpl, err)
	}
	return &OutputFile{
		Path:    filePath,
		Content: content.String()}, nil
}
