package main

import (
	"github.com/shaharby7/Dope/pkg/runtime/controller"
	"github.com/shaharby7/Dope/types"
    
    "github.com/shaharby7/dopeexample/pkg/greeter"
)

var controllers = map[string]types.Controller[any]{

    "server1": Controller_server1,

}


var Controller_server1 = controller.NewHTTPServer(
	[]*types.TypedAction{
		controller.CreateTypedAction(
            
			&types.ActionConfig{
				Name: "/api/greet",
				ControllerBinding: &types.ControllerBinding{
					"method":"POST",
				},
			},
            greeter.Greet,
            ),
            
	},
)


