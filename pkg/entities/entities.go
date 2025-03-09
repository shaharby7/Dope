package entities

import (
	"fmt"
	"io/fs"
	"path/filepath"

	v1 "github.com/shaharby7/Dope/pkg/entities/V1"
	"github.com/shaharby7/Dope/pkg/entities/entity"
)

type TypedEntitiesList []*entity.Entity

type EntitiesTree = map[entity.EntityTypeUniqueIdentifier]TypedEntitiesList

const DOPE_CORE_API = "Dope"

func init() {
	entity.LoadEntityTypeManifest(DOPE_CORE_API, v1.ProjectManifest)
	entity.LoadEntityTypeManifest(DOPE_CORE_API, v1.AppManifest)
	entity.LoadEntityTypeManifest(DOPE_CORE_API, v1.EnvManifest)
	entity.LoadEntityTypeManifest(DOPE_CORE_API, v1.AppEnvManifest)
	entity.LoadEntityTypeManifest(DOPE_CORE_API, v1.ClientManifest)
}

func LoadEntitiesTree(dopePath string) (*EntitiesTree, error) {
	path, err := filepath.Abs(dopePath)
	if err != nil {
		return nil, fmt.Errorf("config validation error: %w", err)
	}
	entitiesTree := make(EntitiesTree, 0)
	err = filepath.Walk(
		path,
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				panic(err)
			}
			if info.IsDir() {
				return nil
			}
			e, err := entity.LoadEntity(path, true)
			if err != nil {
				panic(err)
			}
			typedEntitiesList, ok := entitiesTree[e.EntityTypeUniqueIdentifier]
			if ok {
				entitiesTree[e.EntityTypeUniqueIdentifier] = append(typedEntitiesList, e)
			} else {
				entitiesTree[e.EntityTypeUniqueIdentifier] = []*entity.Entity{e}
			}
			return nil
		},
	)
	return &entitiesTree, err
}
