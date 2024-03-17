package build

import "github.com/shaharby7/Dope/pkg/types"

type appTemplateInput struct {
	*types.ProjectMetadataConfig
	*types.AppConfig
	Imports []string
}
