package types

type ProjectConfig struct {
	DopVersion   string `validate:"required" yaml:"dopVersion"`
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
	Name          string `validate:"required"`
	Description   string
	Controllables []ControllableConfig
}

type ControllableConfig struct {
	Name        string `validate:"required"`
	Description string
	Type        string `validate:"required,oneof=HttpServer"`
	Actionables []ActionableConfig
}

type ActionableConfig struct {
	Name        string `validate:"required"`
	Description string
	Package     string `validate:"required"`
	Ref         string `validate:"required"`
}

type EnvConfig struct {
	Name        string `validate:"required"`
	Description string
}
