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

func writeFiles(files []*files.OutputFile) error {
	for _, file := range files {
		err := writeFile(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeFile(file *files.OutputFile) error {
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
