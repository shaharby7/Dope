package files

import (
	"fmt"

	"github.com/shaharby7/Dope/pkg/entities"
	yamlUtils "github.com/shaharby7/Dope/pkg/utils/yaml"
	"github.com/shaharby7/Dope/types"

	v1 "github.com/shaharby7/Dope/pkg/entities/V1"
	"github.com/shaharby7/Dope/pkg/entities/entity"
	configHelpers "github.com/shaharby7/Dope/pkg/entities/helpers"

	fsUtils "github.com/shaharby7/Dope/pkg/utils/fs"

	bTypes "github.com/shaharby7/Dope/pkg/build/types"
	"github.com/shaharby7/Dope/pkg/utils"
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
	_ *entities.EntitiesTree,
	env string,
	appEntity *entity.Entity,
	appEnvEntity *entity.Entity,
) ([]*fsUtils.OutputFile, error) {
	pathArgs := &helmPathArgs{App: appEntity.Name, Env: env}
	appEnvConf := entity.GetEntityValues[v1.AppEnvConfig](appEnvEntity)
	appConf := entity.GetEntityValues[v1.AppConfig](appEntity)
	imageDataInput := &imageData{
		Registry: appEnvConf.Registry, // TODO: add registry as a provider and not form string
		Tag:      appConf.Version,
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
		appEntity,
		appEnvEntity,
	)
	if err != nil {
		return nil, err
	}
	controllersFile, err := _generateAppControllersFile(
		pathArgs,
		appEntity,
		appEnvEntity,
	)
	if err != nil {
		return nil, err
	}
	return []*fsUtils.OutputFile{
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
	app *entity.Entity,
	appEnv *entity.Entity,
) (*fsUtils.OutputFile, error) {
	appValues, err := yamlUtils.EncodeYamlWithIndent(appEnv.Values, 1)
	if err != nil {
		return nil, utils.FailedBecause(
			fmt.Sprintf("marshal yaml for app %s, env %s", app.Name, appEnv.Name),
			err,
		)
	}
	data := &valuesData{
		AppName:   app.Name,
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
	app *entity.Entity,
	appEnv *entity.Entity,
) (*fsUtils.OutputFile, error) {
	appEnvConfig := entity.GetEntityValues[v1.AppEnvConfig](appEnv)
	controllersStrings, err := utils.Map(
		appEnvConfig.Controllers,
		func(controller v1.ControllerEnvConfig) (string, error) {
			controllerConfig, _ := configHelpers.GetControllerConfig(controller.Name, app)
			addControllerDefaults(&controller, controllerConfig, appEnvConfig)
			addDopeEnvVars(&controller, controllerConfig, app)
			controllerString, err := yamlUtils.EncodeYamlWithIndent([]v1.ControllerEnvConfig{controller}, 1)
			if err != nil {
				return "", utils.FailedBecause(
					fmt.Sprintf("marshal yaml for app %s, env %s", app.Name, appEnv.Name),
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
	controllerEnvConfig *v1.ControllerEnvConfig,
	controllerConfig *v1.ControllerConfig,
	defaults *v1.AppEnvConfig,
) {
	controllerEnvConfig.PopulatedType_ = controllerConfig.Type
	addControllerEnvDefaults(controllerEnvConfig, defaults)
	if controllerEnvConfig.Resources == nil {
		controllerEnvConfig.Resources = &v1.ResourceRequirements{}
	}
	if defaults.Resources == nil {
		defaults.Resources = &v1.ResourceRequirements{}
	}
	addResourcesDefaults(controllerEnvConfig.Resources, defaults.Resources)
	if controllerEnvConfig.Replicas == 0 && defaults.Replicas != 0 {
		controllerEnvConfig.Replicas = defaults.Replicas
	}
}

func addControllerEnvDefaults(controller *v1.ControllerEnvConfig, defaults *v1.AppEnvConfig) {
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

func addResourcesDefaults(main *v1.ResourceRequirements, defaults *v1.ResourceRequirements) {
	if main.Limits == nil {
		main.Limits = &v1.ResourceList{}
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
		main.Requests = &v1.ResourceList{}
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
	controllerEnvConfig *v1.ControllerEnvConfig,
	controllerConfig *v1.ControllerConfig,
	appEntity *entity.Entity,
) {
	dopeEnvVars := []v1.EnvVar{
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
			Value: appEntity.Name,
		},
		{
			Name:  string(types.ENV_VAR_DOPE_PORT),
			Value: fmt.Sprintf("%d", types.DOPE_DEFAULT_PORT),
		},
	}
	if controllerConfig.Type == v1.CONTROLLER_TYPE_HTTPSERVER {
		dopeEnvVars = append(dopeEnvVars, v1.EnvVar{
			Name:  string(types.ENV_VAR_HTTPSERVER_PORT),
			Value: fmt.Sprintf("%d", types.HTTPSERVER_DEFAULT_PORT),
		})
	}
	controllerEnvConfig.Env = append(controllerEnvConfig.Env, dopeEnvVars...)
}

type tHelmDopeValues struct {
	Project      *v1.ProjectConfig      `yaml:"project,omitempty"`
	Build        *tHelmDopeValuesBuild  `yaml:"build,omitempty"`
	Apps         []*tHelmDopeValuesApp  `yaml:"apps,omitempty"`
	Providers    *v1.EnvProvidersConfig `yaml:"providers,omitempty"`
	ArgoCdValues *any                   `yaml:"argo-cd,omitempty"`
}

type tHelmDopeValuesBuild struct {
	Path string `yaml:"path"`
}

type tHelmDopeValuesApp struct {
	Name string `yaml:"name"`
}

func generateHelmDopeValuesFile(
	metadata *bTypes.BuildMetadata,
	eTree *entities.EntitiesTree,
	project *entity.Entity,
	env *entity.Entity,
) ([]*fsUtils.OutputFile, error) {
	envConf := entity.GetEntityValues[v1.EnvConfig](env)
	projectValues := entity.GetEntityValues[v1.ProjectConfig](project)
	appList, _ := utils.Map(
		configHelpers.GetCoreEntitiesByType(*eTree, v1.DOPE_CORE_TYPES_APP),
		func(app *entity.Entity) (*tHelmDopeValuesApp, error) {
			return &tHelmDopeValuesApp{
				Name: app.Name,
			}, nil
		},
	)
	values := &tHelmDopeValues{
		Project: projectValues,
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
	yaml, err := yamlUtils.EncodeYamlWithIndent(values, 1)
	if err != nil {
		return nil, utils.FailedBecause(
			fmt.Sprintf("marshal yaml for dope values, env %s", env.Name),
			err,
		)
	}
	f, err := generateOutputFile(
		templateId_HELM_DOPE_VALUES,
		struct{ Env string }{Env: env.Name},
		string(yaml),
	)
	return []*fsUtils.OutputFile{f}, err
}
