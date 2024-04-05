package build

type SRC_FILES string

const (
	SRC_FILE_MAIN       SRC_FILES = "main.go"
	SRC_FILE_CONTROLLER SRC_FILES = "controller.go"
)

type mainFileInput struct {
}

type controllerFileInput struct {
}
