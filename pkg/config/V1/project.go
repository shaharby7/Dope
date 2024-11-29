package v1

import (
	"reflect"

	"github.com/shaharby7/Dope/pkg/config/entity"
)

var ProjectManifest = &entity.EntityTypeManifest{
	Name:            "Project",
	Aliases:         []string{"project", "proj"},
	BindingSettings: nil,
	ValuesType:      reflect.TypeOf(ProjectConfig{}),
}

type ProjectConfig struct {
	DopeVersion  string                    `validate:"required" yaml:"dopeVersion"`
	Module       string                    `validate:"required" yaml:"module"`
	Versioning   *ProjectVersioningOptions `validate:"required" yaml:"versioning"`
	Apps         []*AppConfig              `yaml:"apps,omitempty"`
	Environments []*EnvConfig              `yaml:"environments,omitempty"`
}

type ProjectVersioningOptions struct {
	Granularity VERSIONING_GRANULARITY_LEVELS `validate:"required" yaml:"granularity"`
	Version     string                        `yaml:"version"`
}

type VERSIONING_GRANULARITY_LEVELS string

const (
	VERSIONING_GRANULARITY_LEVEL_PROJECT VERSIONING_GRANULARITY_LEVELS = "project"
	VERSIONING_GRANULARITY_LEVEL_APP     VERSIONING_GRANULARITY_LEVELS = "app"
)
