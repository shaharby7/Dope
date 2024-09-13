package build

import (
	"fmt"
	"os"
	"path"

	"github.com/shaharby7/Dope/pkg/build/files"
)

func ensurePath(args ...string) string {
	p := path.Join(args...)
	os.MkdirAll(p, os.ModePerm)
	return p
}

func writeFiles(dst string, files []*files.OutputFile) error {
	for _, file := range files {
		err := writeFile(dst, file)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeFile(dst string, file *files.OutputFile) error {
	absPath := path.Join(dst, file.Path)
	ensurePath(path.Dir(absPath))
	fileRef, err := os.Create(absPath)
	if err != nil {
		return fmt.Errorf("could not open file %s: %w", absPath, err)
	}
	defer fileRef.Close()
	_, err = fileRef.Write([]byte(file.Content))
	if err != nil {
		return fmt.Errorf("could not write file (%s): %w", absPath, err)
	}
	return nil
}
