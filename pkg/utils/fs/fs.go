package fs

import (
	"fmt"
	"os"
	"path"
)

type OutputFile struct {
	Path    string
	Content string
}

func (o *OutputFile) WriteFile(basePath string) error {
	return WriteFile(basePath, o)
}

func EnsurePath(args ...string) string {
	p := path.Join(args...)
	os.MkdirAll(p, os.ModePerm)
	return p
}

func WriteFiles(basePath string, files []*OutputFile) error {
	for _, file := range files {
		err := WriteFile(basePath, file)
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteFile(basePath string, file *OutputFile) error {
	absPath := path.Join(basePath, file.Path)
	EnsurePath(path.Dir(absPath))
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
