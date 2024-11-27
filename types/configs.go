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

type ProjectConfig struct {
	Name         string                    `validate:"required" yaml:"name"`
	Description  string                    `yaml:"description"`
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

type AppConfig struct {
	Name        string `validate:"required" yaml:"name"`
	Description string `yaml:"description"`
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

type EnvConfig struct {
	Name        string              `validate:"required" yaml:"name"`
	Description string              `yaml:"description"`
	Providers   *EnvProvidersConfig `validate:"required" yaml:"providers,omitempty"`
	Apps        []AppEnvConfig
}

type AppEnvConfig struct {
	Name                string `validate:"required" yaml:"name"`
	AppName             string `validate:"required" yaml:"appName"`
	Description         string `yaml:"description"`
	Registry            string
	Controllers         []ControllerEnvConfig `yaml:"controllers,omitempty"`
	ControllersDefaults ControllerEnvConfig   `yaml:"controllersDefaults,omitempty"`
	Values              AppValues             `yaml:"values,omitempty"`
}

type ControllerEnvConfig struct {
	Name           string
	Env            []EnvVar              `yaml:"env,omitempty"`
	Replicas       uint32                `yaml:"replicas,omitempty"`
	Resources      *ResourceRequirements `yaml:"resources,omitempty"`
	Debug          *DebugOptions         `yaml:"debug,omitempty"`
	PopulatedType_ CONTROLLER_TYPE       `yaml:"type,omitempty"`
}

type AppValues struct {
	ServiceAccount   AppValuesServiceAccount `yaml:"serviceAccount"`
	ImagePullSecrets *ImagePullSecret        `yaml:"imagePullSecrets"`
	Annotations      Annotations             `yaml:"annotations"`
	Labels           Labels                  `yaml:"labels"`
	SecurityContext  *SecurityContext        `yaml:"securityContext,flow"`
	VolumeMounts     []VolumeMounts          `yaml:"volumeMounts"`
	Volumes          []Volume                `yaml:"volumes"`
	NodeSelector     NodeSelector            `yaml:"nodeSelector"`
	Affinity         *Affinity               `yaml:"Affinity,flow"`
}

type AppValuesServiceAccount struct {
	Create      bool              `yaml:"create,omitempty"`
	Automount   bool              `yaml:"automount,omitempty"`
	Annotations map[string]string `yaml:"annotations,omitempty"`
	Name        string            `yaml:"name,omitempty"`
}

type DebugOptions struct {
	Enabled bool
	Port    Port
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
	Docker *struct {
		Registry string `validate:"required" yaml:"registry"`
		Prefix   string `validate:"required" yaml:"prefix"`
	} `validate:"required" yaml:"docker"`
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
