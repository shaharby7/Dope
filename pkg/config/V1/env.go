package v1

import (
	"reflect"

	"github.com/shaharby7/Dope/pkg/config/entity"
)

var EnvManifest = &entity.EntityTypeManifest{
	Name:            "Env",
	BindingSettings: nil,
	ValuesType:      reflect.TypeOf(EnvConfig{}),
	CliOptions: &entity.CliOptions{
		Aliases:      []string{"environment", "env"},
		PathTemplate: "envs/{{.Name}}",
	},
}

type EnvConfig struct {
	Providers *EnvProvidersConfig `validate:"required" yaml:"providers,omitempty"`
	Apps      []AppEnvConfig
}

type EnvProvidersConfig struct {
	Git *struct {
		Url  string `validate:"required" yaml:"url"`
		Path string `validate:"required" yaml:"path"`
		Ref  string `validate:"required" yaml:"ref"`
	} `validate:"required" yaml:"git"`
	Kubernetes *struct {
		Type K8S_PROVIDERS `validate:"required" yaml:"type"`
	}
	// Docker *struct { // TODO
	// 	Registry string `validate:"required" yaml:"registry"`
	// 	Prefix   string `validate:"required" yaml:"prefix"`
	// } `validate:"required" yaml:"docker"`
	Storage *struct {
		Managed bool   `validate:"required" yaml:"managed"`
		Url     string `yaml:"url"`
	} `validate:"required" yaml:"storage"`
	Cd *struct {
		Type    CD_PROVIDERS `validate:"required" yaml:"type"`
		Managed bool         `validate:"required" yaml:"managed"`
		Values  *any         `yaml:"values"`
	} `validate:"required" yaml:"cd"`
}

type K8S_PROVIDERS string

const (
	K8S_PROVIDERS_MINIKUBE K8S_PROVIDERS = "minikube"
)

type CD_PROVIDERS string

const (
	CD_PROVIDERS_ARGO CD_PROVIDERS = "argo-cd"
)
