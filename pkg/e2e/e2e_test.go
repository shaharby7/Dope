package e2e

import (
	"github.com/shaharby7/Dope/pkg/e2e/loader"
	"github.com/shaharby7/Dope/pkg/utils"
	"github.com/stretchr/testify/assert"

	"testing"
)

func Test_Loader(t *testing.T) {
	str, err := loader.Load(
		"github.com/shaharby7/Dope/pkg/e2e",
	)
	assert.Nil(t, err)
	err = executeE2EMainFile(str)
	assert.Nil(t, err)
}

func Test_Runtime(t *testing.T) {
	err := E2ERuntimeMain(
		"./testdata/.dope",
		"./testdata/build",
		&E2eOptions{
			InstallBefore: utils.Ptr(false),
		},
		[]*E2ETestCase{
			{
				Name: "E2E_Example",
				Func: E2E_Example,
			},
		},
	)
	assert.Nil(t, err)
}