package types

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
	Version     string `validate:"required"`
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
	Name                string
	Registry            string
	Controllers         []ControllerEnvConfig `yaml:"controllers,omitempty"`
	ControllersDefaults ControllerEnvConfig   `yaml:"controllerDefaults,omitempty"`
	Values              AppValues             `yaml:"values,omitempty"`
}

type ControllerEnvConfig struct {
	Name      string
	Env       []EnvVar             `yaml:"env,omitempty"`
	Replicas  uint32               `yaml:"replicas,omitempty"`
	Resources ResourceRequirements `yaml:"resources,omitempty"`
	Debug     DebugOptions         `yaml:"debug,omitempty"`
}

type AppValues struct {
	ServiceAccount   AppValuesServiceAccount `yaml:"serviceAccount"`
	ImagePullSecrets ImagePullSecret         `yaml:"imagePullSecrets"`
	Annotations      Annotations             `yaml:"annotations"`
	Labels           Labels                  `yaml:"labels"`
	SecurityContext  *SecurityContext        `yaml:"securityContext"`
	VolumeMounts     []VolumeMounts          `yaml:"volumeMounts"`
	Volumes          []Volume                `yaml:"volumes"`
	NodeSelector     NodeSelector            `yaml:"nodeSelector"`
	Affinity         *Affinity               `yaml:"Affinity"`
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

type Port uint32
