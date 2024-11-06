package helpers

import (
	"bytes"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"

	"github.com/shaharby7/Dope/pkg/utils"
	"github.com/shaharby7/Dope/types"
)

func GetEnvConfig(config *types.ProjectConfig, envName string) (*types.EnvConfig, error) {
	ok, conf := utils.Find(
		config.Environments,
		func(e types.EnvConfig) bool {
			return e.Name == envName
		},
	)
	if ok {
		return conf, nil
	}
	return nil, fmt.Errorf("could not find configuration for env %s", envName)
}

func GetAppConfig(config *types.ProjectConfig, appName string) (*types.AppConfig, error) {
	ok, conf := utils.Find(
		config.Apps,
		func(e types.AppConfig) bool {
			return e.Name == appName
		},
	)
	if ok {
		return conf, nil
	}
	return nil, fmt.Errorf("could not find configuration for app %s", appName)
}

func GetAppEnvConfig(config *types.ProjectConfig, envName string, appName string) (*types.AppEnvConfig, error) {
	env, err := GetEnvConfig(config, envName)
	if err != nil {
		return nil, err
	}
	ok, conf := utils.Find(
		env.Apps,
		func(e types.AppEnvConfig) bool {
			return e.Name == appName
		},
	)
	if ok {
		return conf, nil
	}
	return nil, fmt.Errorf("could not find config for app %s in env %s", appName, envName)
}

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
