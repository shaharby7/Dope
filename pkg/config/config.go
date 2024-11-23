package config

import (
	t "github.com/shaharby7/Dope/types"

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
