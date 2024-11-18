package files

import (
	"slices"
	"testing"

	bTypes "github.com/shaharby7/Dope/pkg/build/types"
	"github.com/shaharby7/Dope/pkg/utils"
	"github.com/shaharby7/Dope/types"
	"github.com/stretchr/testify/assert"
)

var (
	EXAMPLE_PATH           = "./"
	DST                    = "./build/"
	REPLICAS        uint32 = uint32(3)
	UGLY_NAMES             = "SHAHAR,HADAS"
	FORBIDDEN_NAMES        = "MENASHE"
	APP_CONFIG             = &types.AppConfig{
		Name:        "test-app",
		Description: "test-app-description",
		Controllers: []types.ControllerConfig{
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
	LOCAL_ENV_CONFIG = &types.EnvConfig{
		Name:     "local",
		Provider: "minikube",
		Apps: []types.AppEnvConfig{
			{
				Registry: "docker.io/shaharby7/hi!",
				Name:     "test-app",
				ControllersDefaults: types.ControllerEnvConfig{
					Env: []types.EnvVar{
						{
							Name:  "UGLY_NAMES",
							Value: UGLY_NAMES,
						},
					},
					Replicas: REPLICAS,
					Debug: &types.DebugOptions{
						Enabled: true,
						Port:    types.Port(4000),
					},
					Resources: &types.ResourceRequirements{
						Limits: &types.ResourceList{
							types.ResourceCPU: "100",
						},
					},
				},
				Controllers: []types.ControllerEnvConfig{
					{
						Name: "test-controller",
						Env: []types.EnvVar{
							{
								Name:  "FORBIDDEN_NAMES",
								Value: FORBIDDEN_NAMES,
							},
						},
					},
				},
			},
		},
	}
	PROJECT_CONFIG = types.ProjectConfig{
		DopeVersion: "0.1.0",
		Name:        "test",
		Module:      "shahar.com/hi!",
		Description: "test description",
		Versioning: &types.ProjectVersioningOptions{
			Granularity: types.VERSIONING_GRANULARITY_LEVEL_APP,
		},
		Apps: []*types.AppConfig{
			APP_CONFIG,
		},
		Environments: []*types.EnvConfig{
			LOCAL_ENV_CONFIG,
		},
	}
	BUILD_METADATA = bTypes.BuildMetadata{
		GitRef: "123",
	}
)

func Test_HelmFiles(t *testing.T) {
	files, err := GenerateFiles(
		DST,
		&PROJECT_CONFIG,
		&BUILD_METADATA,
		[]string{"test-app"},
		[]string{"local"},
	)
	assert.Nil(t, err)
	paths, _ := utils.Map(files, func(file *OutputFile) (string, error) { return file.Path, nil })
	assert.True(t, slices.Contains(paths, "src/test-app/controllers.go"))
	assert.True(t, slices.Contains(paths, "src/test-app/main.go"))
	assert.True(t, slices.Contains(paths, "Dockerfile"))
}
