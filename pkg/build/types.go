package build

import (
	"github.com/shaharby7/Dope/types"
)

type SRC_FILES string

const (
	SRC_FILE_MAIN       SRC_FILES = "main.go"
	SRC_FILE_CONTROLLER SRC_FILES = "controllers.go"
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
