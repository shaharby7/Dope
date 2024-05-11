package greeter

import (
	"context"
	"fmt"
	"os"
	"slices"
	"strings"

	d "github.com/shaharby7/Dope/types"
	// "github.com/shaharby7/dopeexample/types"
)

type GreetInput struct {
	Name string `validate:"required" json:"name"`
}

type GreetOutput struct {
	Greet string `json:"greet"`
}

var UGLY_NAMES = strings.Split(os.Getenv("UGLY_NAMES"), ",")

func Greet(
	ctx context.Context, input *GreetInput, controllerPayload *d.ActionInputMetadata,
) (
	*GreetOutput,
	*d.ActionOutputMetadata,
	error,
) {
	greet := ""
	if slices.Contains(UGLY_NAMES, input.Name) {
		greet = fmt.Sprintf("I will not greet you %s!", input.Name)
	} else {
		greet = fmt.Sprintf("hello %s!", input.Name)
	}
	return &GreetOutput{
		Greet: greet,
	}, &d.ActionOutputMetadata{HTTPServer: &d.HTTPServerResponseConfig{StatusCode: 200}}, nil
}

// func CreateTypedAction[In any, Out any](config *ActionConfig, action Action[In, Out]) *TypedAction {
