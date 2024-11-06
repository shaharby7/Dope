package files

import (
	"fmt"

	"github.com/shaharby7/Dope/pkg/build/helpers"
	"github.com/shaharby7/Dope/pkg/utils"
	"github.com/shaharby7/Dope/types"
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
	valuesFile, err := _generateAppValuesFile(
		pathArgs,
		appConfig,
		appEnvConfig,
	)
	if err != nil {
		return nil, err
	}
	controllersFile, err := _generateAppControllersFile(
		pathArgs,
		appConfig,
		appEnvConfig,
	)
	if err != nil {
		return nil, err
	}
	return []*OutputFile{
		imageFile,
		valuesFile,
		controllersFile,
	}, nil
}

type valuesData struct {
	AppName   string
	AppValues string
}

func _generateAppValuesFile(
	pathArgs *helmPathArgs,
	appConfig *types.AppConfig,
	appEnvConfig *types.AppEnvConfig,
) (*OutputFile, error) {
	appValues, err := helpers.EncodeYamlWithIndent(appEnvConfig.Values, 1)
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

func _generateAppControllersFile(
	pathArgs *helmPathArgs,
	appConfig *types.AppConfig,
	appEnvConfig *types.AppEnvConfig,
) (*OutputFile, error) {
	controllersStrings, err := utils.Map(
		appEnvConfig.Controllers,
		func(controller types.ControllerEnvConfig) (string, error) {
			addControllerDefaults(&controller, &appEnvConfig.ControllersDefaults)
			controllerString, err := helpers.EncodeYamlWithIndent(controller, 1)
			if err != nil {
				return "", utils.FailedBecause(
					fmt.Sprintf("marshal yaml for app %s, env %s", appConfig.Name, appEnvConfig.Name),
					err,
				)
			}
			return string(controllerString), nil
		},
	)
	if err != nil {
		return nil, err
	}
	return generateOutputFile(
		templateId_HELM_CONTROLLERS,
		pathArgs,
		controllersStrings,
	)
}

func addControllerDefaults(
	controller *types.ControllerEnvConfig,
	defaults *types.ControllerEnvConfig,
) {
	addControllerEnvDefaults(controller, defaults)
	if controller.Resources == nil {
		controller.Resources = &types.ResourceRequirements{}
	}
	if defaults.Resources == nil {
		defaults.Resources = &types.ResourceRequirements{}
	}
	addResourcesDefaults(controller.Resources, defaults.Resources)
	if controller.Replicas == 0 && defaults.Replicas != 0 {
		controller.Replicas = defaults.Replicas
	}
}

func addControllerEnvDefaults(controller *types.ControllerEnvConfig, defaults *types.ControllerEnvConfig) {
	if defaults.Env != nil {
		for _, dEnv := range defaults.Env {
			hasNonDefault := false
			for _, e := range controller.Env {
				if e.Name == dEnv.Name {
					hasNonDefault = true
					break
				}
			}
			if !hasNonDefault {
				controller.Env = append(controller.Env, dEnv)
			}
		}
	}
}

func addResourcesDefaults(main *types.ResourceRequirements, defaults *types.ResourceRequirements) {
	if main.Limits == nil {
		main.Limits = &types.ResourceList{}
	}
	if defaults != nil {
		for defaultK, defaultVal := range *defaults.Limits {
			_, ok := (*main.Limits)[defaultK]
			if !ok {
				(*main.Limits)[defaultK] = defaultVal
			}
		}
	}
	if main.Requests == nil {
		main.Requests = &types.ResourceList{}
	}
	if defaults != nil {
		for defaultK, defaultVal := range *defaults.Requests {
			_, ok := (*main.Requests)[defaultK]
			if !ok {
				(*main.Requests)[defaultK] = defaultVal
			}
		}
	}
}
