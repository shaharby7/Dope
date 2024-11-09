package types

import (
	k8score "k8s.io/kubernetes/pkg/apis/core"
)

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
