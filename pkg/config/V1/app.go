package v1

import (
	"reflect"

	"github.com/shaharby7/Dope/pkg/config/entity"
)

var AppManifest = &entity.EntityTypeManifest{
	Name:            "App",
	BindingSettings: nil,
	ValuesType:      reflect.TypeOf(AppConfig{}),
	CliOptions: &entity.CliOptions{
		Aliases:      []string{"app", "application"},
		PathTemplate: "general/apps/{{.Name}}",
	},
}

type AppConfig struct {
	Version     string `validate:"required"`
	Controllers []ControllerConfig
}

type ControllerConfig struct {
	Name        string          `validate:"required" yaml:"name"`
	Description string          `yaml:"description"`
	Type        CONTROLLER_TYPE `validate:"required"`
	Actions     []ActionConfig
}

type ActionConfig struct {
	Name              string            `validate:"required" yaml:"name"`
	Description       string            `yaml:"description"`
	Package           string            `validate:"required"`
	Ref               string            `validate:"required"`
	ControllerBinding ControllerBinding `yaml:"controllerBinding"`
}

type CONTROLLER_TYPE string

// enums
const (
	CONTROLLER_TYPE_HTTPSERVER CONTROLLER_TYPE = "HTTPServer"
	CONTROLLER_TYPE_COMMAND    CONTROLLER_TYPE = "Command"
)

type ControllerBinding map[string]string
