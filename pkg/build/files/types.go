package files

import (
	"github.com/shaharby7/Dope/types"
)

type MainControllerFileData struct {
	Imports     []string
	Controllers []*controllerInput
}

type controllerInput struct {
	Name       string
	Identifier string
	Type       types.CONTROLLER_TYPE `validate:"required"`
	Actions    []*actionInput
}

type actionInput struct {
	Name              string
	Caller            string
	ControllerBinding *types.ControllerBinding
}
