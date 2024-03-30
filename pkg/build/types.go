package build

import "github.com/shaharby7/Dope/types"

type appTemplateInput struct {
	*types.ProjectMetadataConfig
	*types.AppConfig
	Imports []string
}
