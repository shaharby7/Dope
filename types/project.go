package types

// import "gopkg.in/yaml.v3"

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
	Name        string `validate:"required"`
	Description string
	Package     string `validate:"required"`
	Ref         string `validate:"required"`
	// Bind        yaml.Node
	Bind *HTTPSeverActionConfig //TODO: support dynamic config for other controller types
}

type EnvConfig struct {
	Name        string `validate:"required"`
	Description string
}
