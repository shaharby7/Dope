package types

import (
	"context"
	"reflect"
	"sync"

	"github.com/julienschmidt/httprouter"
	v1 "github.com/shaharby7/Dope/pkg/entities/V1"
)

// base types
type Controller[ActionConfig any] interface {
	Start(ctx context.Context, wg *sync.WaitGroup) error
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
	Config   *v1.ActionConfig
}

// configs
type ActionInputMetadata struct {
	HTTPServer *HTTPServerRequestConfig
}

type HTTPServerRequestConfig struct {
	Params httprouter.Params
	// Headers map[string]string // TODO
}

type ActionOutputMetadata struct {
	HTTPServer *HTTPServerResponseConfig
}

// HTTPServer specifics

// type HTTPServerConfig struct {
// 	Port string
// }

// type HTTPSeverActionConfig struct {
// 	Method HTTPMethod
// }

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
