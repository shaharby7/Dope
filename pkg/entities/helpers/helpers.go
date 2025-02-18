package helpers

import (
	"fmt"

	"github.com/shaharby7/Dope/pkg/entities"
	v1 "github.com/shaharby7/Dope/pkg/entities/V1"
	"github.com/shaharby7/Dope/pkg/entities/entity"
	"github.com/shaharby7/Dope/pkg/utils"
)

func GetProject(eTree *entities.EntitiesTree) (*entity.Entity, error) {
	projects := GetCoreEntitiesByType(*eTree, v1.DOPE_CORE_TYPES_PROJECT)
	project := projects[0]
	if project == nil {
		return nil, fmt.Errorf("could not find project entity")
	}
	return project, nil
}

func GetEnv(eTree *entities.EntitiesTree, envName string) (*entity.Entity, error) {
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

func GetApp(eTree *entities.EntitiesTree, appName string) (*entity.Entity, error) {
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

func GetClient(eTree *entities.EntitiesTree, clientName string) (*entity.Entity, error) {
	ok, conf := utils.Find(
		GetCoreEntitiesByType(*eTree, v1.DOPE_CORE_TYPES_CLIENT),
		func(e *entity.Entity) bool {
			return e.Name == clientName
		},
	)
	if ok {
		return *conf, nil
	}
	return nil, fmt.Errorf("could not find configuration for app %s", clientName)
}

func GetAppEnvConfig(eTree *entities.EntitiesTree, envName string, appName string) (*entity.Entity, error) {
	ok, conf := utils.Find(
		GetCoreEntitiesByType(*eTree, v1.DOPE_CORE_TYPES_APPENV),
		func(e *entity.Entity) bool {
			return e.Binding[string(v1.DOPE_CORE_TYPES_APP)] == appName &&
				e.Binding[string(v1.DOPE_CORE_TYPES_ENV)] == envName
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

func GetCoreEntitiesByType(eTree entities.EntitiesTree, t v1.DOPE_CORE_TYPES) entities.TypedEntitiesList {
	id := entity.GenerateTypeUniqueIdentifier(entities.DOPE_CORE_API, string(t))
	res := eTree[id]
	return res
}
