package types

type BuildOptions struct {
	Apps    []string
	Envs    []string
	Clients []string
	Stages  []BuildStage
}

type BuildStage string

const (
	BuildStage_FILES  BuildStage = "files"
	BuildStage_DOCKER BuildStage = "docker"
)

type BuildMetadata struct {
	GitRef    string
	BuildPath string
}
