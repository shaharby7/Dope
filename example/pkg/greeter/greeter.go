package greeter

import (
	"context"
	"fmt"

	d "github.com/shaharby7/Dope/types"
	// "github.com/shaharby7/dopeexample/types"
)

type GreetInput struct {
	Name string `validate:"required" json:"name"`
}

type GreetOutput struct {
	Greet string `json:"greet"`
}

func Greet(
	ctx context.Context, input *GreetInput, controllerPayload *d.ActionInputMetadata,
) (
	*GreetOutput,
	*d.ActionOutputMetadata,
	error,
) {
	return &GreetOutput{
		Greet: fmt.Sprintf("hello %s!", input.Name),
	}, &d.ActionOutputMetadata{HTTPServer: &d.HTTPServerResponseConfig{StatusCode: 200}}, nil
}

// func CreateTypedAction[In any, Out any](config *ActionConfig, action Action[In, Out]) *TypedAction {
