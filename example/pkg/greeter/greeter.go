package greeter

import (
	"context"
	"fmt"
	"os"
	"slices"
	"strings"

	d "github.com/shaharby7/Dope/types"

	"github.com/shaharby7/Dope/example/build/src/clients/adminclient"
)

type GreetInput struct {
	Name string `validate:"required" json:"name"`
}

type GreetOutput struct {
	Greet string `json:"greet"`
}

var FORBIDDEN_NAMES = strings.Split(os.Getenv("FORBIDDEN_NAMES"), ",")

func Greet(
	ctx context.Context,
	input *GreetInput,
	controllerPayload *d.ActionInputMetadata,
) (
	*GreetOutput,
	*d.ActionOutputMetadata,
	error,
) {
	if slices.Contains(FORBIDDEN_NAMES, input.Name) {
		return nil, &d.ActionOutputMetadata{HTTPServer: &d.HTTPServerResponseConfig{StatusCode: 400}}, fmt.Errorf("forbidden name")
	}
	uglyNames, _, err := adminclient.GET_admin__api_ugly_names_list(ctx, nil, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch list of ugly names: %w", err)
	}
	greet := ""
	if slices.Contains(*uglyNames, input.Name) {
		greet = fmt.Sprintf("I will not greet you %s!", input.Name)
	} else {
		greet = fmt.Sprintf("hello %s!", input.Name)
	}
	return &GreetOutput{
		Greet: greet,
	}, &d.ActionOutputMetadata{HTTPServer: &d.HTTPServerResponseConfig{StatusCode: 200}}, nil
}
