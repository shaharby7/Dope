package greeter

import (
	"context"
	"fmt"

	"github.com/shaharby7/dopeexample/types"
)

func Greet(ctx context.Context, input *types.GreetInput) (*types.GreetOutput, error) {
	return &types.GreetOutput{
		Output: fmt.Sprintf("hello %s!", input.Name),
	}, nil
}