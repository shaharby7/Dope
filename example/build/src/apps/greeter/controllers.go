package main

import (
	"github.com/shaharby7/Dope/pkg/runtime/controller"
	"github.com/shaharby7/Dope/types"
	v1 "github.com/shaharby7/Dope/pkg/entities/V1"
    
    "github.com/shaharby7/Dope/example/pkg/greeter"
)

var controllers = map[string]types.Controller[any]{

    "listener": Controller_listener,

}


var Controller_listener = controller.NewHTTPServer(
	[]*types.TypedAction{
            
			controller.CreateTypedAction(
			&v1.ActionConfig{
				Name: "/api/greet",
				ControllerBinding: map[string]string{
					"method":"POST",
				},
			},
            greeter.Greet,
            ),
            
	},
) 


