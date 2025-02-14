package e2e

import (
	"testing"
	// exampleE2E "github.com/shaharby7/Dope/example/e2e"
)

func Test_RunE2E(t *testing.T) {
	E2ERuntimeMain(
		"./example/.dope",
		"./example/build",
		&E2eOptions{},
		[]*E2ETestCase{},
	)
}
