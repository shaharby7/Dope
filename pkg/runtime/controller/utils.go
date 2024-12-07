package controller

import (
	"context"
	"fmt"
	"reflect"

	v1 "github.com/shaharby7/Dope/pkg/entities/V1"
	"github.com/shaharby7/Dope/types"
)

func CreateTypedAction[In any, Out any](config *v1.ActionConfig, action types.Action[In, Out]) *types.TypedAction {
	cb := func(
		ctx context.Context,
		input any,
		controllerPayload *types.ActionInputMetadata,
	) (
		any,
		*types.ActionOutputMetadata,
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
	return &types.TypedAction{
		Callback: cb,
		In:       reflect.TypeOf(action).In(1),
		Out:      reflect.TypeOf(action).Out(0),
		Config:   config,
	}
}
