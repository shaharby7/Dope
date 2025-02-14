package e2e

import (
	"context"
	"fmt"

	"github.com/shaharby7/Dope/pkg/build"
	"github.com/shaharby7/Dope/pkg/build/types"
	"github.com/shaharby7/Dope/pkg/install"
	"github.com/shaharby7/Dope/pkg/utils"
)

type ITestProvider interface {
	Errorf(format string, args ...interface{})
}

type TestProvider struct {
	Ctx   context.Context
	Name  string
	Fail  bool
	Error error
}

func NewTesting(ctx context.Context, name string) *TestProvider {
	return &TestProvider{
		Ctx:   ctx,
		Name:  name,
		Fail:  false,
		Error: nil,
	}
}

func (t *TestProvider) Errorf(format string, args ...interface{}) {
	t.Fail = true
	t.Error = fmt.Errorf(format, args...)
}

type E2ETestCase struct {
	Name string
	Func func(runner ITestProvider)
}

type E2eOptions struct {
	BuildBefore    *bool
	InstallBefore  *bool
	UninstallAfter *bool
	FailFast       *bool
}

func E2ERuntimeMain(
	dopePath string,
	dst string,
	options *E2eOptions,
	testCases []*E2ETestCase,
) error {
	ctx := context.Background()
	err := buildAndInstall(
		dopePath,
		dst,
		options,
	)
	if err != nil {
		return err
	}
	for _, testCase := range testCases {
		t := NewTesting(
			ctx,
			testCase.Name,
		)
		testCase.Func(t)
		if t.Fail {
			return fmt.Errorf("test %s failed: %w", testCase.Name, t.Error)
		}
	}
	return nil
}

func buildAndInstall(dopePath string, dst string, options *E2eOptions) error {
	var err error
	if options.BuildBefore == nil || *options.BuildBefore {
		err = build.BuildProject(dopePath, dst, types.BuildOptions{})
		if err != nil {
			return utils.FailedBecause("building project for test", err)
		}
	}
	if options.InstallBefore == nil || *options.InstallBefore {
		err = install.InstallProject(dst, &install.InstallOptions{})
		if err != nil {
			return utils.FailedBecause("installing project for test", err)
		}
	}
	return nil
}
