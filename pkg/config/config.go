package config

import (
	t "github.com/shaharby7/Dope/types"

	yamlUtils "github.com/shaharby7/Dope/pkg/utils/yaml"

	"fmt"
	"path/filepath"
)

func ReadConfig(dopePath string) (*t.ProjectConfig, error) {
	path, err := filepath.Abs(dopePath)
	if err != nil {
		return nil, fmt.Errorf("config validation error: %w", err)
	}
	dopeObjectsByTypes, err := generateDopeObjectsByTypes(path)
	if err != nil {
		return nil, err
	}
	return generateDopeConfigFromDopeObjByTypes(*dopeObjectsByTypes)
}

func WriteConfig(path string, config *t.DopeObjectFile[any]) error {
	return yamlUtils.WriteYaml(path, config)
}
