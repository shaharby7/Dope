package install

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/shaharby7/Dope/pkg/utils"
)

func InstallProject(
	dopePath string,
	projectDst string,
	config *config,
) error {

	valuesFile, err := filepath.Abs(
		path.Join(
			projectDst,
			"helm",
			config.Environment,
			"dope",
			"values.yaml",
		),
	)
	if err != nil {
		return fmt.Errorf("failed to get absolute path for values file: %w", err)
	}
	mainCommand := "install"
	if config.Upgrade {
		mainCommand = "upgrade"
	}
	script := fmt.Sprintf(
		"helm %s %s dope/dope -n %s -f %s --create-namespace --wait",
		mainCommand,
		config.ReleaseName,
		config.Namespace,
		valuesFile,
	)
	out, err := utils.ExecCommand(script)
	print(string(out))
	if err != nil {
		return err
	}
	return nil
}

func UninstallProject(
	dopePath string,
	projectDst string,
	config *config,
) error {
	script := fmt.Sprintf(
		"helm uninstall dope -n %s",
		config.Namespace,
	)
	out, err := utils.ExecCommand(script)
	print(string(out))
	if err != nil {
		return err
	}
	return nil
}
