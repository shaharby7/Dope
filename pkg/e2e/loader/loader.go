package loader

import (
	"fmt"
	"strings"

	"github.com/shaharby7/Dope/pkg/utils"
	"go/types"
	"golang.org/x/tools/go/packages"
)

var E2E_FUNC_PREFIX string = "E2E_"

func Load(
	sourcePath string,
) (string, error) {
	sourcePkg, err := readSourcePackage(sourcePath)
	if err != nil {
		return "", err
	}
	testIdentifiers := findTestIdentifiers(sourcePkg)
	tInput := &sMainFileTmplInput{
		PkgPath:         sourcePath,
		TestIdentifiers: testIdentifiers,
	}
	main, err := createMainFile(tInput)
	if err != nil {
		return "", utils.FailedBecause(
			"could not create main file", err,
		)
	}
	return main.String(), nil
}

func readSourcePackage(e2ePath string) (*packages.Package, error) {
	cfg := packages.Config{
		Mode: packages.NeedImports | packages.NeedSyntax | packages.NeedTypes | packages.NeedDeps | packages.NeedTypesInfo,
	}
	loadedPackages, err := packages.Load(&cfg, e2ePath)
	if err != nil {
		return nil, err
	}
	if len(loadedPackages) != 1 {
		return nil, fmt.Errorf(
			"could not parse e2e package, expected exactly 1 package, got: %d",
			len(loadedPackages),
		)
	}
	sourcePkg := loadedPackages[0]
	if len(sourcePkg.Errors) > 0 {
		return nil, sourcePkg.Errors[0]
	}
	return sourcePkg, nil
}

func findTestIdentifiers(
	e2ePgk *packages.Package,
) []string {
	pgkTypes := e2ePgk.Types.Scope().Names()
	results := utils.Filter(pgkTypes, func(candidateIdentifier string) bool {
		isPrefix := strings.HasPrefix(candidateIdentifier, E2E_FUNC_PREFIX)
		if !isPrefix {
			return false
		}
		candidateFunc := e2ePgk.Types.Scope().Lookup(candidateIdentifier)
		return verifyCandidateType(candidateFunc)
	})
	return results
}

func verifyCandidateType(candidateFunc types.Object) bool {
	funcSig, ok := candidateFunc.Type().(*types.Signature)
	if !ok {
		return false
	}
	params := funcSig.Params()
	if params.Len() != 1 {
		return false
	}
	if params.At(0).Type().String() != "github.com/shaharby7/Dope/pkg/e2e.ITestProvider" {
		//TODO: probably there are better ways to verify the type
		return false
	}
	results := funcSig.Results()
	return results.Len() == 0
}
