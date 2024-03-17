package types

import (
	"context"
)

type HTTPServerOutConfig struct {
	StatusCode uint
	Headers    map[string]string
}

type StdoutOutputConfig struct{}

type ActionableOutputConfig struct {
	HTTPServer HTTPServerOutConfig
	Stdout     StdoutOutputConfig
}

type ActionableOutput[OutputData any] struct {
	Data   OutputData
	Config ActionableOutputConfig
}

type Actionable[InputData any, OutputData any] func(
	ctx context.Context,
	input InputData,
) (ActionableOutput[OutputData], error)
