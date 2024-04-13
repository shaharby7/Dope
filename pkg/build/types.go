package build

import (
	"github.com/shaharby7/Dope/types"
)

type FILES string

const (
	SRC_FILE_MAIN       FILES = "src/main.go"
	SRC_FILE_CONTROLLER FILES = "src/controllers.go"
	DOCKERFILE          FILES = "Dockerfile"
)

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

type dockerfileInput struct {
	AppName string
}
