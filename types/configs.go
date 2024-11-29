package types

type DOPE_OBJECTS string

const (
	DOPE_OBJECT_PROJECT DOPE_OBJECTS = "Project"
	DOPE_OBJECT_APP     DOPE_OBJECTS = "App"
	DOPE_OBJECT_ENV     DOPE_OBJECTS = "Env"
	DOPE_OBJECT_APP_ENV DOPE_OBJECTS = "AppEnv"
)

type DopeObjectFile[Obj any] struct {
	Api         string                 `validate:"required" yaml:"api"`
	Type        DOPE_OBJECTS           `validate:"required" yaml:"type"`
	Name        string                 `validate:"required" yaml:"name"`
	Description string                 `yaml:"description,omitempty"`
	Binding     *DopeObjectFileBinding `yaml:"binding,omitempty"`
	Values      Obj                    `yaml:"values,omitempty"`
}

type DopeObjectFileBinding struct {
	Env *string `yaml:"env,omitempty"`
	App *string `yaml:"app,omitempty"`
}


