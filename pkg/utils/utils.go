package utils

import (
	"fmt"
	"os"

	"github.com/shaharby7/Dope/types"
)

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func Find[T any](ts []T, fn func(T) bool) (bool, *T) {
	for _, t := range ts {
		if fn(t) {
			return true, &t
		}
	}
	return false, nil
}

func Getenv(name types.ENV_VARS, defaultVal string) string {
	val := os.Getenv(string(name))
	if val == "" {
		return defaultVal
	}
	return val
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
