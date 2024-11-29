package helpers

import (
	"fmt"

	"github.com/shaharby7/Dope/pkg/config"
	v1 "github.com/shaharby7/Dope/pkg/config/V1"
	"github.com/shaharby7/Dope/pkg/config/entity"
	"github.com/shaharby7/Dope/pkg/utils"
)

func GetProject(eTree *config.EntitiesTree) (*entity.Entity, error) {
	projects := GetCoreEntitiesByType(*eTree, v1.DOPE_CORE_TYPES_PROJECT)
	project := projects[0]
	if project == nil {
		return nil, fmt.Errorf("could not find project config")
	}
	return project, nil
}

func GetEnv(eTree *config.EntitiesTree, envName string) (*entity.Entity, error) {
	ok, conf := utils.Find(
		GetCoreEntitiesByType(*eTree, v1.DOPE_CORE_TYPES_ENV),
		func(e *entity.Entity) bool {
			return e.Name == envName
		},
	)
	if ok {
		return *conf, nil
	}
	return nil, fmt.Errorf("could not find configuration for env %s", envName)
}

func GetApp(eTree *config.EntitiesTree, appName string) (*entity.Entity, error) {
	ok, conf := utils.Find(
		GetCoreEntitiesByType(*eTree, v1.DOPE_CORE_TYPES_APP),
		func(e *entity.Entity) bool {
			return e.Name == appName
		},
	)
	if ok {
		return *conf, nil
	}
	return nil, fmt.Errorf("could not find configuration for app %s", appName)
}

func GetAppEnvConfig(eTree *config.EntitiesTree, envName string, appName string) (*entity.Entity, error) {
	ok, conf := utils.Find(
		GetCoreEntitiesByType(*eTree, v1.DOPE_CORE_TYPES_APPENV),
		func(e *entity.Entity) bool {
			boundedToApp := false
			boundedToEnv := false
			for _, b := range e.Binding {
				switch b.Type {
				case string(v1.DOPE_CORE_TYPES_APP):
					if b.Name == appName {
						boundedToApp = true
					}
				case string(v1.DOPE_CORE_TYPES_ENV):
					if b.Name == envName {
						boundedToEnv = true
					}
				}
			}
			return boundedToApp && boundedToEnv
		},
	)
	if ok {
		return *conf, nil
	}
	return nil, fmt.Errorf("could not find config for app %s in env %s", appName, envName)
}

func GetControllerConfig(controllerName string, appEntity *entity.Entity) (*v1.ControllerConfig, error) {
	appConfig := entity.GetEntityValues[v1.AppConfig](appEntity)
	ok, conf := utils.Find(
		appConfig.Controllers,
		func(c v1.ControllerConfig) bool {
			return c.Name == controllerName
		},
	)
	if ok {
		return conf, nil
	}
	return nil, fmt.Errorf("could not find config for controller %s within app %s", controllerName, appEntity.Name)
}

func GetCoreEntitiesByType(eTree config.EntitiesTree, t v1.DOPE_CORE_TYPES) config.TypedEntitiesList {
	id := entity.GenerateTypeUniqueIdentifier(config.DOPE_CORE_API, string(t))
	res := eTree[id]
	return res
}
