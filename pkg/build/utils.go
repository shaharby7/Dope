package build

import (
	"fmt"
	"os"
	"path"
)

func createFileByTemplate[FileInput any](
	dstDir string,
	tmpl FILES,
	input *FileInput,
) error {
	filePath := path.Join(dstDir, string(tmpl))
	ensurePath(path.Dir(filePath))
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("could not open file %s: %w", tmpl, err)
	}
	defer file.Close()
	err = getTemplate(tmpl).Execute(file, input)
	if err != nil {
		return fmt.Errorf("could not parse template (%s): %w", tmpl, err)
	}
	return nil
}

func wrapSrcFileCreationError(appName string, fileName FILES, err error) error {
	return fmt.Errorf("could not could not create src file %s for app %s: %w", fileName, appName, err)
}

func ensurePath(args ...string) string {
	p := path.Join(args...)
	os.MkdirAll(p, os.ModePerm)
	return p
}
