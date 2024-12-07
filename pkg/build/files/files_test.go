package files

// import (
// 	"slices"
// 	"testing"

// 	bTypes "github.com/shaharby7/Dope/pkg/build/types"
// 	v1 "github.com/shaharby7/Dope/pkg/entities/V1"
// 	"github.com/shaharby7/Dope/pkg/utils"
// 	"github.com/stretchr/testify/assert"

// 	fsUtils "github.com/shaharby7/Dope/pkg/utils/fs"
// )

// var (
// 	EXAMPLE_PATH           = "./"
// 	DST                    = "./build/"
// 	REPLICAS        uint32 = uint32(3)
// 	UGLY_NAMES             = "SHAHAR,HADAS"
// 	FORBIDDEN_NAMES        = "BOB"
// 	APP_CONFIG             = &v1.AppConfig{
// 		Name:        "test-app",
// 		Description: "test-app-description",
// 		Controllers: []v1.ControllerConfig{
// 			{
// 				Name:        "test-controller",
// 				Description: "test-controller-description",
// 				Type:        "HTTPServer",
// 				Actions: []v1.ActionConfig{
// 					{
// 						Name:        "/api/greet",
// 						Description: "Great anyone who wants",
// 						Package:     "pkg/greeter",
// 						Ref:         "Greet",
// 						ControllerBinding: v1.ControllerBinding{
// 							"Method": "POST",
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	LOCAL_ENV_CONFIG = &v1.EnvConfig{
// 		Name:      "local",
// 		Providers: &v1.EnvProvidersConfig{},
// 		Apps: []v1.AppEnvConfig{
// 			{
// 				Registry: "docker.io/shaharby7/hi!",
// 				AppName:  "test-app",
// 				ControllersDefaults: v1.ControllerEnvConfig{
// 					Env: []v1.EnvVar{
// 						{
// 							Name:  "UGLY_NAMES",
// 							Value: UGLY_NAMES,
// 						},
// 					},
// 					Replicas: REPLICAS,
// 					Debug: &v1.DebugOptions{
// 						Enabled: true,
// 						Port:    v1.Port(4000),
// 					},
// 					Resources: &v1.ResourceRequirements{
// 						Limits: &v1.ResourceList{
// 							v1.ResourceCPU: "100",
// 						},
// 					},
// 				},
// 				Controllers: []v1.ControllerEnvConfig{
// 					{
// 						Name: "test-controller",
// 						Env: []v1.EnvVar{
// 							{
// 								Name:  "FORBIDDEN_NAMES",
// 								Value: FORBIDDEN_NAMES,
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	PROJECT_CONFIG = v1.ProjectConfig{
// 		DopeVersion: "0.1.0",
// 		Name:        "test",
// 		Module:      "shahar.com/hi!",
// 		Description: "test description",
// 		Versioning: &v1.ProjectVersioningOptions{
// 			Granularity: v1.VERSIONING_GRANULARITY_LEVEL_APP,
// 		},
// 		Apps: []*v1.AppConfig{
// 			APP_CONFIG,
// 		},
// 		Environments: []*v1.EnvConfig{
// 			LOCAL_ENV_CONFIG,
// 		},
// 	}
// 	BUILD_METADATA = bTypes.BuildMetadata{
// 		GitRef: "123",
// 	}
// )

// func Test_HelmFiles(t *testing.T) {
// 	files, err := GenerateFiles(
// 		DST,
// 		&PROJECT_CONFIG,
// 		&BUILD_METADATA,
// 		[]string{"test-app"},
// 		[]string{"local"},
// 	)
// 	assert.Nil(t, err)
// 	paths, _ := utils.Map(files, func(file *fsUtils.OutputFile) (string, error) { return file.Path, nil })
// 	assert.True(t, slices.Contains(paths, "src/test-app/controllers.go"))
// 	assert.True(t, slices.Contains(paths, "src/test-app/main.go"))
// 	assert.True(t, slices.Contains(paths, "Dockerfile"))
// }
