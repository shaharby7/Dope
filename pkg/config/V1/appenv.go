package v1

import (
	"reflect"

	"github.com/shaharby7/Dope/pkg/config/entity"
	k8score "k8s.io/kubernetes/pkg/apis/core"
)

var AppEnvManifest = &entity.EntityTypeManifest{
	Name: "AppEnv",
	BindingSettings: &entity.BindingSettings{
		Must: []string{"Env", "App"},
	},
	ValuesType: reflect.TypeOf(AppEnvConfig{}),
	CliOptions: &entity.CliOptions{
		Aliases:      []string{"appenv"},
		PathTemplate: "envs/{{.Binding.Env}}/apps",
	},
}

type AppEnvConfig struct {
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

type Port uint32

// TODO, use "github.com/shaharby7/types/k8s"

type EnvVar struct {
	Name      string        `yaml:"name"`
	Value     string        `yaml:"value,omitempty"`
	ValueFrom *EnvVarSource `yaml:"valueFrom,omitempty"`
}

type EnvVarSource struct {
	FieldRef         *ObjectFieldSelector   `yaml:"fieldRef,omitempty"`
	ResourceFieldRef *ResourceFieldSelector `yaml:"resourceFieldRef,omitempty"`
	ConfigMapKeyRef  *ConfigMapKeySelector  `yaml:"configMapKeyRef,omitempty"`
	SecretKeyRef     *SecretKeySelector     `yaml:"secretKeyRef,omitempty"`
}

type ObjectFieldSelector struct {
	APIVersion string `yaml:"apiVersion,omitempty"`
	FieldPath  string `yaml:"fieldPath"`
}

type ResourceFieldSelector struct {
	ContainerName string           `yaml:"containerName,omitempty"`
	Resource      string           `yaml:"resource"`
	Divisor       ResourceQuantity `yaml:"divisor,omitempty"`
}

type ConfigMapKeySelector struct {
	LocalObjectReference `yaml:",inline"`
	Key                  string `yaml:"key"`
	Optional             *bool  `yaml:"optional,omitempty"`
}

type SecretKeySelector struct {
	LocalObjectReference `yaml:",inline"`
	Key                  string `yaml:"key"`
	Optional             *bool  `yaml:"optional,omitempty"`
}

type LocalObjectReference struct {
	Name string `yaml:"name,omitempty"`
}

type Annotations map[string]string

type Labels map[string]string

type ImagePullSecret struct {
	Name  string `yaml:"name"`
	Image string `yaml:"image,omitempty"`
}

type ResourceRequirements struct {
	Limits   *ResourceList `yaml:"limits,omitempty"`
	Requests *ResourceList `yaml:"requests,omitempty"`
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

type VolumeMounts = k8score.VolumeMount

type SecurityContext = k8score.SecurityContext

type NodeSelector = map[string]string

type Toleration = k8score.Toleration

type Volume = k8score.Volume

type Affinity = k8score.Affinity
