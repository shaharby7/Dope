package files

import (
	"testing"

	"github.com/shaharby7/Dope/types"
	"github.com/stretchr/testify/assert"
)

var (
	EXAMPLE_PATH        = "./"
	DST                 = "./build/"
	REPLICAS     uint32 = uint32(3)
	UGLY_NAMES          = "SHAHAR,HADAS"
	APP_CONFIG          = &types.AppConfig{
		Name:        "test-app",
		Description: "test-app-description",
		Controllers: []types.ControllersConfig{
			{
				Name:        "test-controller",
				Description: "test-controller-description",
				Type:        "HTTPServer",
				Actions: []types.ActionConfig{
					{
						Name:        "/api/greet",
						Description: "Great anyone who wants",
						Package:     "pkg/greeter",
						Ref:         "Greet",
						ControllerBinding: types.ControllerBinding{
							"Method": "POST",
						},
					},
				},
			},
		},
	}
	LOCAL_ENV_CONFIG = types.EnvConfig{
		Name:     "local",
		Provider: "minikube",
		Apps: []types.AppEnvConfig{
			{
				Registry: "docker.io/shaharby7/hi!",
				Name:     "test-app",
				Values: types.AppHelmValues{
					Env: []types.EnvVar{
						{
							Name:  "UGLY_NAMES",
							Value: UGLY_NAMES,
						},
					},
					Replicas: REPLICAS,
					Debug: types.DebugOptions{
						Enabled:   true,
						DebugPort: types.Port(4000),
					},
					Resources: types.ResourceRequirements{
						Limits: types.ResourceList{
							types.ResourceCPU: "100m",
						},
					},
				},
			},
		},
	}
	PROJECT_CONFIG = types.ProjectConfig{
		DopeVersion: "0.1.0",
		Metadata: types.ProjectMetadataConfig{
			Name:        "test",
			Version:     "0.1.0",
			Module:      "shahar.com/hi!",
			Description: "test description",
		},
		Apps: []types.AppConfig{*APP_CONFIG},
		Environments: []types.EnvConfig{
			LOCAL_ENV_CONFIG,
		},
	}
)

func TestHelmFiles(t *testing.T) {
	files, err := GenerateFiles(
		DST, &PROJECT_CONFIG, []string{}, []string{},
	)
	assert.Nil(t, err)
	assert.True(t, len(files) > 0) //TODO
}
