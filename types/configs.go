package types

import (
	kApiV1 "k8s.io/api/core/v1"
)

type ProjectConfig struct {
	DopeVersion  string `validate:"required" yaml:"dopeVersion"`
	Metadata     ProjectMetadataConfig
	Apps         []AppConfig
	Environments []EnvConfig
}

type ProjectMetadataConfig struct {
	Name        string `validate:"required"`
	Version     string `validate:"required"`
	Module      string `validate:"required"`
	Description string
}

type AppConfig struct {
	Name        string `validate:"required"`
	Description string
	Controllers []ControllersConfig
}

type ControllersConfig struct {
	Name        string `validate:"required"`
	Description string
	Type        ControllerType `validate:"required"`
	Actions     []ActionConfig
}

type ActionConfig struct {
	Name              string `validate:"required"`
	Description       string
	Package           string            `validate:"required"`
	Ref               string            `validate:"required"`
	ControllerBinding ControllerBinding `yaml:"controllerBinding"`
}

type EnvConfig struct {
	Name     string `validate:"required"`
	Provider string `validate:"required"`
	Apps     []AppEnvConfig
}

type AppEnvConfig struct {
	Name     string
	Registry string
	Values   AppHelmValues
}

type AppHelmValues struct {
	Env       []EnvVar
	Replicas  uint32
	Resources ResourceRequirements
	Debug     DebugOptions
}

type ResourceRequirements struct {
	Limits   ResourceList
	Requests ResourceList
}

type ResourceList map[ResourceName]ResourceQuantity

type ResourceName string

const (
	ResourceCPU              ResourceName = "cpu"
	ResourceMemory           ResourceName = "memory"
	ResourceStorage          ResourceName = "storage"
	ResourceEphemeralStorage ResourceName = "ephemeral-storage"
)

type ResourceQuantity string

type DebugOptions struct {
	Enabled bool
	Port    Port
}

type Port uint32

type EnvVar = kApiV1.EnvVar
