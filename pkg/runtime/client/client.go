package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/shaharby7/Dope/types"
)

type ClientData struct {
	Url string
}

func ParseClientData() *ClientData {
	return &ClientData{
		Url: "http://localhost:3000", // todo: take url from env
	}
}

func CreateTypedClientCall[In any, Out any](
	clientData *ClientData,
	method string,
	path string,
	callback func(
		ctx context.Context,
		input *In,
		payload *types.ActionInputMetadata,
	) (*Out,
		*types.ActionOutputMetadata,
		error,
	)) func(ctx context.Context,
	input *In,
	payload *types.ActionInputMetadata,
) (*Out,
	*types.ActionOutputMetadata,
	error,
) {
	var fullPath string
	if path != "" {
		fullPath = clientData.Url + path
	} else {
		fullPath = clientData.Url
	}
	clientCall := func(ctx context.Context,
		input *In,
		payload *types.ActionInputMetadata,
	) (*Out,
		*types.ActionOutputMetadata,
		error,
	) {
		formatted, err := json.Marshal(input)
		if err != nil {
			return nil, nil, err
		}
		req, err := http.NewRequestWithContext(
			ctx,
			method,
			fullPath,
			bytes.NewBuffer(formatted),
		)
		if err != nil {
			return nil, nil, err
		}
		if payload != nil && payload.HTTPServer != nil {
			query := req.URL.Query()
			for _, param := range payload.HTTPServer.Params {
				query.Add(param.Key, param.Value)
			}
			req.URL.RawQuery = query.Encode()
		}
		resp, err := http.DefaultClient.Do(req)
		respHeaders := make(map[string]string, 0)
		if resp != nil {
			for key := range resp.Header {
				respHeaders[key] = resp.Header.Get(key)
			}
		}
		if err != nil {
			return nil, nil, err
		}
		outputMetadata := &types.ActionOutputMetadata{
			HTTPServer: &types.HTTPServerResponseConfig{
				StatusCode: resp.StatusCode,
				Headers:    respHeaders,
			},
		}
		var out *Out = new(Out)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, outputMetadata, err
		}
		err = json.Unmarshal(body, out)
		return out, outputMetadata, err
	}
	return clientCall
}
