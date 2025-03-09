package types

import (
	"context"
	"reflect"
	"sync"

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
	Params  map[string]string
	Query   map[string][]string
	Headers map[string][]string
}

type ActionOutputMetadata struct {
	HTTPServer *HTTPServerResponseConfig
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