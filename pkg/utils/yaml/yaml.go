package yaml

import (
	"bytes"
	"io"

	fsUtils "github.com/shaharby7/Dope/pkg/utils/fs"

	"gopkg.in/yaml.v3"
)

func EncodeYamlWithIndent(object any, indent int) (string, error) {
	var b bytes.Buffer
	writer := io.Writer(&b)
	encoder := yaml.NewEncoder(writer)
	encoder.SetIndent(indent)
	err := encoder.Encode(object)
	if err != nil {
		return "", err
	}
	err = encoder.Close()
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func WriteYaml(path string, object any) error {
	str, err := EncodeYamlWithIndent(object, 0)
	if err != nil {
		return err
	}
	oFile := &fsUtils.OutputFile{
		Path:    "",
		Content: str,
	}
	return oFile.WriteFile(path)
}
