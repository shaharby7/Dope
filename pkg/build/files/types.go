package files

import (
	"github.com/shaharby7/Dope/types"
)

type TEMPLATE_FILES string

const (
	SRC_FILE_MAIN       TEMPLATE_FILES = "src/{{.App}}/main.go"
	SRC_FILE_CONTROLLER TEMPLATE_FILES = "src/{{.App}}/controllers.go"
	DOCKERFILE          TEMPLATE_FILES = "Dockerfile"
)

// type FileInput struct {
// 	Name TEMPLATE_FILES
// 	Template tem
// }

type OutputFile struct {
	Path    string
	Content string
}

type FileGenerationInput[FileData any, PathArgs any] struct {
	Params struct {
		Path PathArgs
	}
	Data FileData
}

type MainFileData struct {
}

type MainFileArgPath struct{
	App string
}

type MainControllerFileData struct {
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
