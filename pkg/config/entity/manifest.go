// Package entity allows to declare dope entity types using the `entity.EntityTypeManifest` struct, that would be parsed into `entity.Entity`
// Dope will know how to parse entities declared by .dope.yaml files and apply it to the project
package entity

import (
	"fmt"
	"reflect"
)

// `entity.EntityTypeManifest` is am abstraction of the information required for Dope to work with each Entity type.
// Project, App, Env etc. all implements a EntityTypeManifest (see github.com/shaharby7/Dope/pkg/config/V1)
type EntityTypeManifest struct {
	// Entity name
	Name string `validate:"required" yaml:"name"`
	// Defines what other entities can/must this entity be bind to
	BindingSettings *BindingSettings `yaml:"bindingSettings,omitempty"`
	// Defines what are the values this entity expects to receive
	ValuesType reflect.Type `validate:"required"`
	// instructs the cli how it interact with the entity type (for example, for the "create" command)
	CliOptions *CliOptions `yaml:"cliOptions"`
}

// Defines what other entities can/must this entity be bind to
type BindingSettings struct {
	// Dope build will throw error if the entity would not be bind to other entity of this type
	Must []string
	// Dope build will not throw error if the entity would be bind to other entity of this type
	Might []string
}

type CliOptions struct {
	// Would be used to refer by cli tool
	Aliases []string `validate:"required" yaml:"aliases"`
	// Would be used to compose a default path in the `.dope` dir for resource creation
	PathTemplate string
}

type EntityTypeUniqueIdentifier string

var KnownEntityTypes map[EntityTypeUniqueIdentifier]*EntityTypeManifest = map[EntityTypeUniqueIdentifier]*EntityTypeManifest{}

func LoadEntityTypeManifest(api string, entityManifest *EntityTypeManifest) {
	id := GenerateTypeUniqueIdentifier(api, entityManifest.Name)
	_, ok := KnownEntityTypes[id]
	if ok {
		panic(fmt.Errorf("cannot load `entity.EntityTypeManifest` %s twice", id))
	}
	KnownEntityTypes[id] = entityManifest
	for _, aliases := range entityManifest.CliOptions.Aliases {
		uniqueAlias := GenerateTypeUniqueIdentifier(api, aliases)
		_, ok := KnownEntityTypes[uniqueAlias]
		if ok {
			panic(fmt.Errorf("aliases collision for entity %s", uniqueAlias))
		}
		KnownEntityTypes[uniqueAlias] = entityManifest
	}
}

func GetEntityTypeManifest(api string, name string) (*EntityTypeManifest, error) {
	id := GenerateTypeUniqueIdentifier(api, name)
	e, ok := KnownEntityTypes[id]
	if !ok {
		return nil, fmt.Errorf("entity %s not found", id)
	}
	return e, nil
}

func GenerateTypeUniqueIdentifier(api string, name string) EntityTypeUniqueIdentifier {
	key := fmt.Sprintf("%s::%s", api, name)
	return EntityTypeUniqueIdentifier(key)
}
