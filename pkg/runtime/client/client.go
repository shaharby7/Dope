package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/shaharby7/Dope/types"
)

func CreateTypedClientCall[In any, Out any](
	method string,
	url string,
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
			url,
			bytes.NewBuffer(formatted),
		)

		if err != nil {
			return nil, nil, err
		}
		if payload != nil && payload.HTTPServer != nil {
			for _, param := range payload.HTTPServer.Params {
				req.Header.Set(param.Key, param.Value)
			}
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
		var out *Out
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, outputMetadata, err
		}
		err = json.Unmarshal(body, out)
		return out, outputMetadata, err
	}
	return clientCall
}
