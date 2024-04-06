package types

import (
	"context"
	"fmt"
	"reflect"
	"sync"

	"github.com/julienschmidt/httprouter"
)

type ControllerType string

// enums
const (
	HTTPServer ControllerType = "HTTPServer"
	Command    ControllerType = "Command"
)

// base types
type Controller[ActionConfig any] interface {
	Start(ctx context.Context, wg *sync.WaitGroup) error
	// RegisterAction(name string, action Action[any, any], config ActionConfig)
}

type Action[In any, Out any] func(
	ctx context.Context,
	input In,
	controllerPayload *ActionInputMetadata,
) (
	Out,
	*ActionOutputMetadata,
	error,
)

type TypedAction struct {
	Callback Action[any, any]
	In       reflect.Type
	Out      reflect.Type
	Config   *ActionConfig
}

func CreateTypedAction[In any, Out any](config *ActionConfig, action Action[In, Out]) *TypedAction {
	cb := func(
		ctx context.Context,
		input any,
		controllerPayload *ActionInputMetadata,
	) (
		any,
		*ActionOutputMetadata,
		error,
	) {
		in, ok := input.(In)
		if !ok {
			return nil, nil, fmt.Errorf("got incorrect input type for action")
		}
		return action(
			ctx,
			in,
			controllerPayload,
		)
	}
	return &TypedAction{
		Callback: cb,
		In:       reflect.TypeOf(action).In(1),
		Out:      reflect.TypeOf(action).Out(0),
		Config:   config,
	}
}

// configs
type ActionInputMetadata struct {
	HTTPServer *httprouter.Params
}

type ActionOutputMetadata struct {
	HTTPServer *HTTPServerResponseConfig
}

// HTTPServer specifics

type HTTPServerConfig struct {
	Port uint32
}

type HTTPSeverActionConfig struct {
	Method HTTPMethod
}

type HTTPServerResponseConfig struct {
	StatusCode int
	Headers    map[string]string
}

type HTTPMethod string

const (
	GET     HTTPMethod = "GET"
	POST    HTTPMethod = "POST"
	HEAD    HTTPMethod = "HEAD"
	PUT     HTTPMethod = "PUT"
	PATCH   HTTPMethod = "PATCH" // RFC 5789
	DELETE  HTTPMethod = "DELETE"
	CONNECT HTTPMethod = "CONNECT"
	OPTIONS HTTPMethod = "OPTIONS"
	TRACE   HTTPMethod = "TRACE"
)

// Command specifics
