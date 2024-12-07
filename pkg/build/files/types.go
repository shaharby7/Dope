package files

import (
	v1 "github.com/shaharby7/Dope/pkg/entities/V1"
)

type MainControllerFileData struct {
	Imports     []string
	Controllers []*controllerInput
}

type controllerInput struct {
	Name       string
	Identifier string
	Type       v1.CONTROLLER_TYPE `validate:"required"`
	Actions    []*actionInput
}

type actionInput struct {
	Name              string
	Caller            string
	ControllerBinding *v1.ControllerBinding
}
