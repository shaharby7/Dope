package loader_test

import (
	"path/filepath"
	"testing"

	"github.com/shaharby7/Dope/pkg/e2e/loader"
)

func TestBasic(t *testing.T) {
	examplePath, err := filepath.Abs("../testdata")
	if err != nil {
		t.Error(err)
	}
	_, err = loader.Load(examplePath)
	if err != nil {
		t.Error(err)
	}
}
