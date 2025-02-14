package v1

import (
	"reflect"

	"github.com/shaharby7/Dope/pkg/entities/entity"
)

var ProjectManifest = &entity.EntityTypeManifest{
	Name:            "Project",
	BindingSettings: nil,
	ValuesType:      reflect.TypeOf(ProjectConfig{}),
	CliOptions: &entity.CliOptions{
		Aliases:      []string{"project", "proj"},
		PathTemplate: "./",
	},
}

type ProjectConfig struct {
	DopeVersion  string                    `validate:"required" yaml:"dopeVersion"`
	Module       string                    `validate:"required" yaml:"module"`
	Versioning   *ProjectVersioningOptions `validate:"required" yaml:"versioning"`
	Apps         []*AppConfig              `yaml:"apps,omitempty"`
	Environments []*EnvConfig              `yaml:"environments,omitempty"`
	E2E          *E2EConfig                `yaml:"e2e,omitempty"`
}

type ProjectVersioningOptions struct {
	Granularity VERSIONING_GRANULARITY_LEVELS `validate:"required" yaml:"granularity"`
	Version     string                        `yaml:"version"`
}

type E2EConfig struct {
	Package string `validate:"required" yaml:"package"`
}

type VERSIONING_GRANULARITY_LEVELS string

const (
	VERSIONING_GRANULARITY_LEVEL_PROJECT VERSIONING_GRANULARITY_LEVELS = "project"
	VERSIONING_GRANULARITY_LEVEL_APP     VERSIONING_GRANULARITY_LEVELS = "app"
)