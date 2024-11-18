package config

import (
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"

	t "github.com/shaharby7/Dope/types"

	"github.com/shaharby7/Dope/pkg/config/helpers"
)

type tObjByType = map[t.DOPE_OBJECTS][]*sDopeObjectFile

func generateDopeObjectsByTypes(path string) (*tObjByType, error) {
	dopeObjectsByTypes := make(tObjByType, 0)
	err := filepath.Walk(
		path,
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				panic(err)
			}
			if info.IsDir() {
				return nil
			}
			dType, dObj, err := readDopeObjFile(path)
			if err != nil {
				panic(err)
			}
			dTypeList, ok := dopeObjectsByTypes[dType]
			if ok {
				dopeObjectsByTypes[dType] = append(dTypeList, dObj)
			} else {
				dopeObjectsByTypes[dType] = []*sDopeObjectFile{dObj}
			}
			return nil
		},
	)
	return &dopeObjectsByTypes, err
}

func generateDopeConfigFromDopeObjByTypes(
	dopeObjectsByTypes tObjByType,
) (*t.ProjectConfig, error) {
	projList, ok := dopeObjectsByTypes[t.DOPE_OBJECT_PROJECT]
	if !ok {
		return nil, errors.New("project has to have exactly 1 project objects defined, found 0")
	}
	if len(projList) > 1 {
		return nil, fmt.Errorf("project has to have exactly 1 project objects defined, found %d", len(projList))
	}
	proj := projList[0]
	config := proj.Values.(t.ProjectConfig)
	for _, app := range dopeObjectsByTypes[t.DOPE_OBJECT_APP] {
		a := app.Values.(t.AppConfig)
		config.Apps = append(config.Apps, &a)
	}
	for _, env := range dopeObjectsByTypes[t.DOPE_OBJECT_ENV] {
		a := env.Values.(t.EnvConfig)
		config.Environments = append(config.Environments, &a)
	}
	for _, appEnv := range dopeObjectsByTypes[t.DOPE_OBJECT_APP_ENV] {
		appEnvConfig := appEnv.Values.(t.AppEnvConfig)
		env, err := getEnvByAppEnvBinding(appEnv, config)
		if err != nil {
			return nil, err
		}
		appEnvConfig.AppName = appEnv.Binding.App
		env.Apps = append(env.Apps, appEnvConfig)
	}
	return &config, nil
}

func getEnvByAppEnvBinding(appEnv *sDopeObjectFile, config t.ProjectConfig) (*t.EnvConfig, error) {
	if appEnv.Binding == nil {
		return nil, fmt.Errorf("no binding found for AppEnv %s", appEnv.Name)
	}
	if appEnv.Binding.App == "" {
		return nil, fmt.Errorf("no app binding found for AppEnv %s", appEnv.Name)
	}
	if appEnv.Binding.Env == "" {
		return nil, fmt.Errorf("no env binding found for AppEnv %s", appEnv.Name)
	}
	env, err := helpers.GetEnvConfig(&config, appEnv.Binding.Env)
	if err != nil {
		return nil, fmt.Errorf("app %s was declared for env %s, but no such env defined", appEnv.Binding.App, appEnv.Binding.Env)
	}
	_, err = helpers.GetAppConfig(&config, appEnv.Binding.App)
	if err != nil {
		return nil, fmt.Errorf("app %s was declared for env %s, but no such app defined", appEnv.Binding.App, appEnv.Binding.Env)
	}
	return env, nil
}
