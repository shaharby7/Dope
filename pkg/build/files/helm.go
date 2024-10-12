package files

import (
	"fmt"

	bTypes "github.com/shaharby7/Dope/pkg/build/types"
	"github.com/shaharby7/Dope/pkg/utils"
	"github.com/shaharby7/Dope/types"

	"gopkg.in/yaml.v3"
)

type helmPathArgs struct {
	App string
	Env string
}

type imageData struct {
	Registry string
	Tag      string
}

func generateHelmFiles(
	_ *types.ProjectConfig,
	metadata *bTypes.BuildMetadata,
	env string,
	appConfig *types.AppConfig,
	appEnvConfig *types.AppEnvConfig,
) ([]*OutputFile, error) {
	pathArgs := &helmPathArgs{App: appConfig.Name, Env: env}
	imageDataInput := &imageData{
		Registry: appEnvConfig.Registry,
		Tag:      metadata.GitRef,
	}
	imageFile, err := generateOutputFile(
		templateId_HELM_IMAGE,
		pathArgs,
		imageDataInput,
	)
	if err != nil {
		return nil, err
	}
	appValues, err := yaml.Marshal(appEnvConfig.Values)
	if err != nil {
		return nil, utils.FailedBecause(
			fmt.Sprintf("marshal yaml for app %s, env %s", appConfig.Name, env),
			err,
		)
	}
	valuesFile, err := generateOutputFile(
		templateId_HELM_VALUES,
		pathArgs,
		string(appValues),
	)
	if err != nil {
		return nil, err
	}
	return []*OutputFile{imageFile, valuesFile}, nil
}
