package main

import (
	"github.com/shaharby7/Dope/pkg/runtime/controller"
	"github.com/shaharby7/Dope/types"
	"github.com/shaharby7/dopeexample/pkg/greeter"
)

var controllers = map[string]types.Controller[any]{
	"Controllerserver1": Controllerserver1,
}

var Controllerserver1 = controller.NewHTTPServer(
	types.HTTPServerConfig{Port: 3000},
	[]*types.TypedAction{
		types.CreateTypedAction(
			&types.ActionConfig{
				Name: "/api/greet",
				ControllerBinding: &types.HTTPSeverActionConfig{
					Method: types.POST,
				},
			}, greeter.Greet),
	},
)