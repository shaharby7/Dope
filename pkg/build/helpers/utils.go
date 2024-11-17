package helpers

import (
	"bytes"
	"io"

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
