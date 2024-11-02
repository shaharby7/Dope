package files

import (
	"fmt"

	"github.com/shaharby7/Dope/pkg/utils"
	"github.com/shaharby7/Dope/types"

	"gopkg.in/yaml.v3"
)

type helmPathArgs struct {
	App string
	Env string
}

type imageData struct {
	Registry string
	Tag      string
}

func generateHelmFiles(
	_ *types.ProjectConfig,
	env string,
	appConfig *types.AppConfig,
	appEnvConfig *types.AppEnvConfig,
) ([]*OutputFile, error) {
	pathArgs := &helmPathArgs{App: appConfig.Name, Env: env}
	imageDataInput := &imageData{
		Registry: appEnvConfig.Registry,
		Tag:      appConfig.Version,
	}
	imageFile, err := generateOutputFile(
		templateId_HELM_IMAGE,
		pathArgs,
		imageDataInput,
	)
	if err != nil {
		return nil, err
	}
	valuesFile, err := generateValuesFile(pathArgs, appConfig, appEnvConfig)
	if err != nil {
		return nil, err
	}
	return []*OutputFile{imageFile, valuesFile}, nil
}

type valuesData struct {
	AppName   string
	AppValues string
}

func generateValuesFile(
	pathArgs *helmPathArgs,
	appConfig *types.AppConfig,
	appEnvConfig *types.AppEnvConfig,
) (*OutputFile, error) {
	appValues, err := yaml.Marshal(appEnvConfig.Values)
	if err != nil {
		return nil, utils.FailedBecause(
			fmt.Sprintf("marshal yaml for app %s, env %s", appConfig.Name, appEnvConfig.Name),
			err,
		)
	}
	data := &valuesData{
		AppName:   appConfig.Name,
		AppValues: string(appValues),
	}
	valuesFile, err := generateOutputFile(
		templateId_HELM_VALUES,
		pathArgs,
		data,
	)
	if err != nil {
		return nil, err
	}
	return valuesFile, nil
}

// appName: {{ .AppName }}

// serviceAccount:
//   {{ .ServiceAccount }}

// imagePullSecrets: []

// pod:
//   annotations: {}
//   labels: {}
//   securityContext: {}

// volumes: []

// volumeMounts: []

// nodeSelector: {}

// tolerations: []

// affinity: {}
