package controllable

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/shaharby7/Dope/pkg/deployable/constants"
	"github.com/shaharby7/Dope/pkg/deployable/errorable"
)

type TServerInput = *http.Request
type HttpResponse struct {
	Data       string
	StatusCode int
	Headers    map[string]string
}
type TServerOutput *HttpResponse

type ServerControllableConfig struct{ Port int }

func NewServerControllableConfig(Port int) *ServerControllableConfig {
	return &ServerControllableConfig{Port: Port}
}

type HttpServerActionables = TActionables[TServerInput, TServerOutput]

type SHttpServerActionable[TActionableInput any, TActionableOutput any] struct {
	SActionable[TServerInput, TServerOutput, TActionableInput, TActionableOutput]
}

func NewHttpServerActionable[TActionableInput any, TActionableOutput any](
	RunActionable FInnerRunActionable[TActionableInput, TActionableOutput],
	MarshalActionableInput FMarshalActionableInput[TServerInput, TActionableInput],
	MarshalControllableOutput FMarshalControllableOutput[TActionableOutput, TServerOutput],
	OnError func(context.Context, error) (TServerOutput, error),
) *SHttpServerActionable[TActionableInput, TActionableOutput] {
	return &SHttpServerActionable[TActionableInput, TActionableOutput]{
		*NewActionable(
			RunActionable, MarshalActionableInput, MarshalControllableOutput, OnError,
		),
	}
}

type SHttpServerControllable struct {
	SControllable[TServerInput, TServerOutput]
	config ServerControllableConfig
	server *http.Server
}

func NewHttpServerControllable(
	ControllerName string,
	config ServerControllableConfig,
) *SHttpServerControllable {
	var server *http.Server
	return &SHttpServerControllable{
		*NewControllable(
			constants.HTTP_SERVER,
			ControllerName,
			func(execute func(actionableName string, input TServerInput) (TServerOutput, error)) error {
				mux := http.NewServeMux()
				mux.HandleFunc("/liveness-probe", func(responseWriter http.ResponseWriter, request *http.Request) {
					returnServerOutput(&HttpResponse{StatusCode: 200, Data: "OK"}, responseWriter)
				})
				mux.HandleFunc("/readiness-probe", func(responseWriter http.ResponseWriter, request *http.Request) {
					returnServerOutput(&HttpResponse{StatusCode: 200, Data: "OK"}, responseWriter)
				})
				mux.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
					if request.RequestURI == "/readiness-probe" || request.RequestURI == "/liveness-probe" {
						returnServerOutput(&HttpResponse{StatusCode: 200, Data: "OK"}, responseWriter)
					} else {
						output, err := execute(request.RequestURI, request)
						output = handleRouteError(err, output)
						returnServerOutput(output, responseWriter)
					}
				})
				server = &http.Server{
					Addr:    fmt.Sprintf(":%d", config.Port),
					Handler: mux,
				}
				go server.ListenAndServe() // TODO: handle errors and cancellations
				fmt.Printf("Http server is listening on port %s\n", server.Addr)
				return nil
			},
			func(ctx context.Context, err error) (TServerOutput, error) {
				return &HttpResponse{Data: "something went really wrong", StatusCode: 500}, nil
			}, //TODO
		),
		config,
		server,
	}
}

func handleRouteError(err error, output TServerOutput) TServerOutput {
	if nil != err {
		erbl, ok := err.(errorable.Errorable)
		if ok {
			switch erbl.Code() {
			case errorable.ACTIONABLE_NOT_FOUND:
				output = &HttpResponse{
					Data:       "Page not found",
					StatusCode: 404,
				}
				return output
			}
		}
		output = &HttpResponse{Data: "Server error", StatusCode: 500}
	}
	return output
}

func returnServerOutput(output TServerOutput, responseWriter http.ResponseWriter) {
	if nil != output.Headers {
		for headerName, headerVal := range output.Headers {
			responseWriter.Header().Add(headerName, headerVal)
		}
	}
	if output.StatusCode == 0 {
		output.StatusCode = 200
	}
	responseWriter.WriteHeader(output.StatusCode)
	io.WriteString(responseWriter, output.Data)
}
