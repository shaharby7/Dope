package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"text/template"

	"github.com/shaharby7/Dope/pkg/utils/set"
)

type TEmpty struct{}

var Empty TEmpty = TEmpty{}

func Map[T, V any](ts []T, fn func(T) (V, error)) ([]V, error) {
	result := make([]V, len(ts))
	for i, t := range ts {
		val, err := fn(t)
		if err != nil {
			return nil, err
		}
		result[i] = val
	}
	return result, nil
}

func Find[T any](ts []T, fn func(T) bool) (bool, *T) {
	for _, t := range ts {
		if fn(t) {
			return true, &t
		}
	}
	return false, nil
}

func RemoveDuplicates[V comparable](vSlice []V) []V {
	return set.NewSet(
		set.OptionFromSlice(vSlice),
	).ToSlice()

}

func Getenv(name string, defaultVal string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultVal
	}
	return val
}

func IsEmpty[T any](t T) bool {
	return reflect.ValueOf(t).IsZero()
}

func GetFromMapWithDefault[T any](m map[string]T, key string, fallback T) T {
	val, ok := m[key]
	if !ok {
		return fallback
	}
	return val
}

func FailedBecause(failedTo string, err error) error {
	return fmt.Errorf("\tcould not %s because:%w", failedTo, err)
}

func GetGitHEADRef() (string, error) {
	out, err := exec.Command("git", "rev-parse", "HEAD").Output()
	if err != nil {
		return "", FailedBecause("getting HEAD commit", err)
	}
	commitHash := string(out)
	return commitHash, nil
}

func ApplyTemplateSafe(template *template.Template, templateName string, args any) (*bytes.Buffer, error) {
	var result bytes.Buffer
	var err error
	if IsEmpty(args) {
		err = template.ExecuteTemplate(&result, templateName, EMPTY_TEMPLATE_INPUT)
	} else {
		err = template.ExecuteTemplate(&result, templateName, args)
	}
	return &result, err
}

var EMPTY_TEMPLATE_INPUT *struct{ A string } = &struct{ A string }{A: "A"}
