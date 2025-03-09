package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"reflect"
	"sync"

	"github.com/julienschmidt/httprouter"
	"github.com/shaharby7/Dope/pkg/utils"
	"github.com/shaharby7/Dope/types"
)

type HTTPServerConfig struct {
	Port string
}
type HTTPServerBindConfig struct {
	Method string `yaml:"method"`
}

type HTTPServer struct {
	config     HTTPServerConfig
	server     *http.Server
	router     *httprouter.Router
	middleware func(n httprouter.Handle) httprouter.Handle
}

func NewHTTPServer(actions []*types.TypedAction) *HTTPServer {

	router := httprouter.New()
	port := utils.Getenv(
		string(types.ENV_VAR_HTTPSERVER_PORT),
		fmt.Sprint(types.HTTPSERVER_DEFAULT_PORT),
	)
	config := &HTTPServerConfig{
		Port: port,
	}
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Port),
		Handler: router,
	}
	for _, action := range actions {
		actionConfig := action.Config
		handler := generateRouteHandler(action)
		method := utils.GetFromMapWithDefault(actionConfig.ControllerBinding, "method", "GET")
		router.Handle(
			method,
			actionConfig.Name,
			handler,
		)
	}
	router.Handle( //TODO: allow custom readyz/livez
		"GET",
		"/",
		func(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
			returnServerOutput(
				"OK",
				types.HTTPServerResponseConfig{StatusCode: 200},
				writer,
			)
		},
	)
	return &HTTPServer{
		config:     *config,
		server:     server,
		router:     router,
		middleware: defaultMiddleware, // TODO: allow custom middlewares
	}
}

func (httpServer *HTTPServer) Start(ctx context.Context, wg *sync.WaitGroup) error {
	listener, err := net.Listen("tcp", httpServer.server.Addr)
	if err != nil {
		return err
	}
	go func() {
		if err := httpServer.server.Serve(listener); err != nil {
			panic(err)
		}
	}()
	fmt.Printf("http server is listening on port: %s", httpServer.config.Port)
	return nil
}

func generateRouteHandler(action *types.TypedAction) httprouter.Handle {
	handler := func(writer http.ResponseWriter, r *http.Request, params httprouter.Params) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			returnServerOutput(
				fmt.Sprintf("could not read body: %s", err.Error()), types.HTTPServerResponseConfig{StatusCode: 400}, writer,
			)
			return
		}
		elem := reflect.New(action.In.Elem())
		in := elem.Interface()
		if err := json.Unmarshal([]byte(body), in); err != nil {
			returnServerOutput(
				fmt.Sprintf("could not parse body: %s", err.Error()), types.HTTPServerResponseConfig{StatusCode: 400}, writer,
			)
			return
		}
		inputMetadata := createInputMetadata(r, params)
		input := []reflect.Value{
			reflect.ValueOf(context.TODO()),
			reflect.ValueOf(in),
			reflect.ValueOf(inputMetadata),
		}
		rawOutput := reflect.ValueOf(action.Callback).Call(input)
		if len(rawOutput) == 3 && !rawOutput[2].IsNil() {
			returnServerOutput(
				fmt.Sprintf("action completed with error: %s", rawOutput[2]), types.HTTPServerResponseConfig{StatusCode: 400}, writer,
			)
			return
		}
		responseBody, err := json.Marshal(rawOutput[0].Interface())
		if err != nil {
			returnServerOutput(
				fmt.Sprintf("could not marshal output: %s", err.Error()), types.HTTPServerResponseConfig{StatusCode: 400}, writer,
			)
			return
		}
		returnServerOutput(
			string(responseBody),
			types.HTTPServerResponseConfig{StatusCode: 200}, //TODO, take from rawOutput 200
			writer,
		)
	}
	return handler
}

func defaultMiddleware(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		n(w, r, ps)
	}
}

func returnServerOutput(data string, responseConfig types.HTTPServerResponseConfig, responseWriter http.ResponseWriter) {
	if nil != responseConfig.Headers {
		for headerName, headerVal := range responseConfig.Headers {
			responseWriter.Header().Add(headerName, headerVal)
		}
	}
	if responseConfig.StatusCode == 0 {
		responseConfig.StatusCode = 200
	}
	responseWriter.WriteHeader(responseConfig.StatusCode)
	io.WriteString(responseWriter, data)
}

func createInputMetadata(r *http.Request, params httprouter.Params) *types.ActionInputMetadata {
	flattenedParams := make(map[string]string)
	for _, param := range params {
		flattenedParams[param.Key] = param.Value
	}
	inputMetadata := &types.ActionInputMetadata{
		HTTPServer: &types.HTTPServerRequestConfig{
			Params:  flattenedParams,
			Headers: r.Header,
			Query:   r.URL.Query(),
		},
	}
	return inputMetadata
}
