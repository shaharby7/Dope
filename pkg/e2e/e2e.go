package e2e

import (
	"fmt"
	"path/filepath"

	"github.com/shaharby7/Dope/pkg/e2e/loader"
	"github.com/shaharby7/Dope/pkg/entities"
	v1 "github.com/shaharby7/Dope/pkg/entities/V1"
	"github.com/shaharby7/Dope/pkg/entities/entity"
	eHelpers "github.com/shaharby7/Dope/pkg/entities/helpers"
	"github.com/shaharby7/Dope/pkg/utils"
	fsUtils "github.com/shaharby7/Dope/pkg/utils/fs"
)

func E2EProject(
	dopePath string,
	dst string,
) error {
	e2eConfig, err := getProjectE2EConfig(dopePath)
	if err != nil {
		return err
	}
	mainFile, err := loader.Load(e2eConfig.Package)
	if err != nil {
		return utils.FailedBecause("load e2e tests", err)
	}
	return executeE2EMainFile(mainFile)
}

func getProjectE2EConfig(dopePath string) (*v1.E2EConfig, error) {
	eTree, err := entities.LoadEntitiesTree(dopePath)
	if err != nil {
		return nil, utils.FailedBecause(
			fmt.Sprintf("generate config from file (%s)", dopePath),
			err,
		)
	}
	project, err := eHelpers.GetProject(eTree)
	if err != nil {
		return nil, err
	}
	projectValues := entity.GetEntityValues[v1.ProjectConfig](project)
	e2eConfig := projectValues.E2E
	if e2eConfig == nil {
		return nil, fmt.Errorf("no e2e config found in project")
	}
	return e2eConfig, nil
}

func executeE2EMainFile(mainFile string) error {
	arbitraryPath, _ := filepath.Abs(
		fmt.Sprintf("tmptest_%s", utils.RandStringRunes(10)),
	)
	defer fsUtils.RemoveDirectory(arbitraryPath)
	fsUtils.WriteFile(arbitraryPath, &fsUtils.OutputFile{
		Path:    "main.go",
		Content: mainFile,
	})
	out, err := utils.ExecCommand(fmt.Sprintf("go run %s", arbitraryPath))
	print(string(out))
	if err != nil {
		return err
	}
	return nil
}
