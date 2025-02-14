package install

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/shaharby7/Dope/pkg/utils"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/getter"

	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
)

func InstallProject( //TODO: let's not hardcode everything :)
	projectDst string,
	options *InstallOptions,
) error {
	chartName := "dope"
	chartVersion := "0.1.5"
	namespace := "dope"

	// create a new Helm action configuration
	settings := cli.New()

	p, err := filepath.Abs(path.Join(projectDst, "helm/local/dope/values.yaml"))
	if err != nil {
		return err
	}
	valueOpts := &values.Options{
		ValueFiles: []string{
			p,
		},
	}

	actionConfig := new(action.Configuration)
	err = actionConfig.Init(settings.RESTClientGetter(), namespace, "", func(format string, v ...interface{}) {})

	if err != nil {
		return utils.FailedBecause("initiating helm action", err)
	}

	// create a new Helm install action
	installAction := action.NewInstall(actionConfig)

	// set the chart name, version, and namespace
	installAction.Namespace = namespace
	installAction.ReleaseName = chartName
	installAction.Version = chartVersion
	installAction.CreateNamespace = true

	cp, err := installAction.LocateChart("dope/dope", settings)

	if err != nil {
		return utils.FailedBecause("locate dope chart", err)
	}

	allSettings := getter.All(settings)
	vals, err := valueOpts.MergeValues(allSettings)

	if err != nil {
		return utils.FailedBecause("loading value files", err)
	}

	chartRequested, err := loader.Load(cp)

	if err != nil {
		return utils.FailedBecause("loading dope chart", err)
	}
	// install the chart
	release, err := installAction.Run(chartRequested, vals)
	if err != nil {
		return utils.FailedBecause("install dope chart", err)
	}
	fmt.Printf("successfully installed revision %d", release.Version)

	return nil
}
