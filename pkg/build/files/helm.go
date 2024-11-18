package files

import (
	"fmt"

	"github.com/shaharby7/Dope/pkg/build/helpers"

	configHelpers "github.com/shaharby7/Dope/pkg/config/helpers"

	bTypes "github.com/shaharby7/Dope/pkg/build/types"
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

func generateAppHelmFiles(
	_ *types.ProjectConfig,
	env string,
	appConfig *types.AppConfig,
	appEnvConfig *types.AppEnvConfig,
) ([]*OutputFile, error) {
	pathArgs := &helmPathArgs{App: appConfig.Name, Env: env}
	imageDataInput := &imageData{
		Registry: appEnvConfig.Registry, // TODO: add registry as a provider and not form string
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
			controllerConfig, _ := configHelpers.GetControllerConfig(controller.Name, appConfig)
			addControllerDefaults(&controller, controllerConfig, &appEnvConfig.ControllersDefaults)
			addDopeEnvVars(&controller, controllerConfig, appConfig)
			controllerString, err := helpers.EncodeYamlWithIndent([]types.ControllerEnvConfig{controller}, 1)
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
	controllerEnvConfig *types.ControllerEnvConfig,
	controllerConfig *types.ControllerConfig,
	defaults *types.ControllerEnvConfig,
) {
	controllerEnvConfig.PopulatedType_ = controllerConfig.Type
	addControllerEnvDefaults(controllerEnvConfig, defaults)
	if controllerEnvConfig.Resources == nil {
		controllerEnvConfig.Resources = &types.ResourceRequirements{}
	}
	if defaults.Resources == nil {
		defaults.Resources = &types.ResourceRequirements{}
	}
	addResourcesDefaults(controllerEnvConfig.Resources, defaults.Resources)
	if controllerEnvConfig.Replicas == 0 && defaults.Replicas != 0 {
		controllerEnvConfig.Replicas = defaults.Replicas
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
	if defaults != nil && defaults.Limits != nil {
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
	if defaults != nil && defaults.Requests != nil {
		for defaultK, defaultVal := range *defaults.Requests {
			_, ok := (*main.Requests)[defaultK]
			if !ok {
				(*main.Requests)[defaultK] = defaultVal
			}
		}
	}
}

func addDopeEnvVars(
	controllerEnvConfig *types.ControllerEnvConfig,
	controllerConfig *types.ControllerConfig,
	appConfig *types.AppConfig,
) {
	dopeEnvVars := []types.EnvVar{
		{
			Name:  string(types.ENV_VAR_CONTROLLER_TYPE),
			Value: string(controllerConfig.Type),
		},
		{
			Name:  string(types.ENV_VAR_CONTROLLER_NAME),
			Value: controllerConfig.Name,
		},
		{
			Name:  string(types.ENV_VAR_APP_NAME),
			Value: appConfig.Name,
		},
		{
			Name:  string(types.ENV_VAR_DOPE_PORT),
			Value: fmt.Sprintf("%d", types.DOPE_DEFAULT_PORT),
		},
	}
	if controllerConfig.Type == types.CONTROLLER_TYPE_HTTPSERVER {
		dopeEnvVars = append(dopeEnvVars, types.EnvVar{
			Name:  string(types.ENV_VAR_HTTPSERVER_PORT),
			Value: fmt.Sprintf("%d", types.HTTPSERVER_DEFAULT_PORT),
		})
	}
	controllerEnvConfig.Env = append(controllerEnvConfig.Env, dopeEnvVars...)
}

type tHelmDopeValues struct {
	Project      *types.ProjectConfig      `yaml:"project,omitempty"`
	Build        *tHelmDopeValuesBuild     `yaml:"build,omitempty"`
	Apps         []*tHelmDopeValuesApp     `yaml:"apps,omitempty"`
	Providers    *types.EnvProvidersConfig `yaml:"providers,omitempty"`
	ArgoCdValues *any                      `yaml:"argo-cd,omitempty"`
}

type tHelmDopeValuesBuild struct {
	Path string `yaml:"path"`
}

type tHelmDopeValuesApp struct {
	Name string `yaml:"name"`
}

func generateHelmDopeValuesFile(metadata *bTypes.BuildMetadata, config *types.ProjectConfig, envConf *types.EnvConfig) ([]*OutputFile, error) {
	appList, _ := utils.Map(
		config.Apps,
		func(app *types.AppConfig) (*tHelmDopeValuesApp, error) {
			return &tHelmDopeValuesApp{
				Name: app.Name,
			}, nil
		},
	)
	values := &tHelmDopeValues{
		Project: &types.ProjectConfig{
			Name:        config.Name,
			Description: config.Description,
			Module:      config.Module,
			Versioning:  config.Versioning,
		},
		Build: &tHelmDopeValuesBuild{
			Path: metadata.BuildPath,
		},
		Apps:      appList,
		Providers: envConf.Providers,
	}
	if envConf.Providers != nil &&
		envConf.Providers.Cd != nil &&
		envConf.Providers.Cd.Managed {
		values.ArgoCdValues = envConf.Providers.Cd.Values
	}
	yaml, err := helpers.EncodeYamlWithIndent(values, 1)
	if err != nil {
		return nil, utils.FailedBecause(
			fmt.Sprintf("marshal yaml for dope values, env %s", envConf.Name),
			err,
		)
	}
	f, err := generateOutputFile(
		templateId_HELM_DOPE_VALUES,
		struct{ Env string }{Env: envConf.Name},
		string(yaml),
	)
	return []*OutputFile{f}, err
}
