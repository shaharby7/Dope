package entity

import (
	"fmt"
	"os"
	"reflect"

	"github.com/shaharby7/Dope/pkg/utils"
	"gopkg.in/yaml.v3"

	vp "github.com/go-playground/validator/v10"
	yamlUtils "github.com/shaharby7/Dope/pkg/utils/yaml"
)

var validator *vp.Validate

func init() {
	validator = vp.New(vp.WithRequiredStructEnabled())
}

// `entity.Entity` represents a Dope Entity that was initiated according to it's `entity.EntityTypeManifest`
// `entity.Entity` instance would usually be translated into one or more files during the build process
type Entity struct {
	// Identifies the package where the `entity.EntityTypeManifest` is defined at //TODO
	Api string `validate:"required" yaml:"api"`
	// Entity identifier
	Name string `validate:"required" yaml:"name"`
	// Descriptions of the entity, used for docs
	Description string `yaml:"description"`
	// entity.EntityTypeManifest identifier
	Type string `validate:"required" yaml:"type"`
	// Some entities can only live in the context of other entities, of other types. Binding would instruct Dope how to relate between the entities
	Binding EntityBind `yaml:"binding,omitempty"`
	// Values of the entity.
	Values any `yaml:"values,omitempty"`

	// Pointer to the `entity.EntityTypeManifest` that was inferred from the Api and Type.
	EntityTypeManifest *EntityTypeManifest
	// EntityTypeUniqueIdentifier
	EntityTypeUniqueIdentifier EntityTypeUniqueIdentifier `yaml:""`
}

func LoadEntity(path string, validate bool) (*Entity, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, utils.FailedBecause(fmt.Sprintf("read file %s", path), err)
	}
	var entity *Entity
	err = yaml.Unmarshal(content, &entity)
	if err != nil {
		return nil, utils.FailedBecause(fmt.Sprintf("unmarshal file %s", path), err)
	}
	if validate {
		err := validator.Struct(entity)
		if err != nil {
			return nil, fmt.Errorf("entity %s of type %s is not valid because: %w", entity.Name, entity.Type, err)
		}
	}
	entity.EntityTypeUniqueIdentifier = GenerateTypeUniqueIdentifier(entity.Api, entity.Type)
	return entity, nil
}

func GetEntityValues[T any](e *Entity) *T {
	res := e.Values.(T)
	return &res
}

func (e *Entity) UnmarshalYAML(node *yaml.Node) error {
	var err error
	partial := &struct {
		Api         string     `yaml:"api"`
		Name        string     `yaml:"name"`
		Type        string     `yaml:"type"`
		Description string     `yaml:"description"`
		Binding     EntityBind `yaml:"binding,omitempty"`
	}{}
	if err := node.Decode(&partial); err != nil {
		return utils.FailedBecause("infer type of entity", err)
	}
	e.EntityTypeManifest, err = GetEntityTypeManifest(partial.Api, partial.Type)
	if err != nil {
		return err
	}
	valuesStructFiled := &reflect.StructField{
		Name: "Values",
		Tag:  `yaml:"values"`,
		Type: e.EntityTypeManifest.ValuesType,
	}
	valuesStruct := reflect.StructOf([]reflect.StructField{*valuesStructFiled})
	valuesInstance := reflect.New(valuesStruct)
	err = node.Decode(valuesInstance.Interface())
	if err != nil {
		return err
	}
	e.Name = partial.Name
	e.Api = partial.Api
	e.Type = partial.Type
	e.Description = partial.Description
	e.Binding = partial.Binding
	e.Values = valuesInstance.Elem().FieldByName("Values").Interface()
	return nil
}

// `entity.EntityBind` instructs Dope to create a relation between this entity to other entity
type EntityBind map[string]string

func WriteEntity(path string, e *Entity) error {
	toWrite := struct {
		Api         string     `yaml:"api"`
		Type        string     `yaml:"type"`
		Name        string     `yaml:"name"`
		Description string     `yaml:"description"`
		Binding     EntityBind `yaml:"binding,omitempty"`
		Values      any        `yaml:"values,omitempty"`
	}{
		Api:         e.Api,
		Type:        e.Type,
		Name:        e.Name,
		Description: e.Description,
		Binding:     e.Binding,
		Values:      e.Values,
	}
	return yamlUtils.WriteYaml(path, toWrite)
}
