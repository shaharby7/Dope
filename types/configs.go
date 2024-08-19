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
	Package           string             `validate:"required"`
	Ref               string             `validate:"required"`
	ControllerBinding *ControllerBinding `yaml:"controllerBinding"`
}

type EnvConfig struct {
	Name        string `validate:"required"`
	Description string
	Provider    string `validate:"required"`
}

// environments:
//   - name: local
//     provider: minikube
//     branch: example
//     apps:
//       - name: "app1"
//         repository: docker.io/shaharby7/app1-local
//         env:
//           UGLY_NAMES: "shahar,danny"
//         resources:
//           requests:
//             cpu: 1
//             memory: 2GB
//           limits:
//             cpu: 2
//             memory: 4GB
//     dope-essentials:
//       argo-cd:
//         enabled: true
//         redis-ha:
//           enabled: false
//         controller:
//           replicas: 1
//         server:
//           replicas: 1
//         repoServer:
//           replicas: 1
//         applicationSet:
//           replicaCount: 1
//       argo-workflows:
//         enabled: true
//         server:
//           extraArgs: [--auth-mode=server]
//           clusterWorkflowTemplates:
//             enabled: true
//         controller:
//           workflowNamespaces:
//             - dope
//           rbac:
//             create: true
//           clusterWorkflowTemplates:
//             enabled: true
//       dope-ci:
//         enabled: true
