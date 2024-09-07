package files

import (
	"github.com/shaharby7/Dope/types"
)

type TEMPLATE_FILES string

const (
	SRC_FILE_MAIN       TEMPLATE_FILES = "src/main.go"
	SRC_FILE_CONTROLLER TEMPLATE_FILES = "src/controllers.go"
	DOCKERFILE          TEMPLATE_FILES = "Dockerfile"
)

type OutputFile struct {
	Path    string
	Content string
}

type mainFileInput struct {
}

type controllerFileInput struct {
	Imports     []string
	Controllers []*controllerInput
}

type controllerInput struct {
	Name       string
	Identifier string
	Type       types.ControllerType `validate:"required"`
	Actions    []*actionInput
}

type actionInput struct {
	Name              string
	Caller            string
	ControllerBinding *types.ControllerBinding
}
