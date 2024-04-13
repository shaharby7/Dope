package utils

import (
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
