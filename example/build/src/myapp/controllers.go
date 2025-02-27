package main

import (
	"github.com/shaharby7/Dope/pkg/runtime/controller"
	"github.com/shaharby7/Dope/types"
	v1 "github.com/shaharby7/Dope/pkg/entities/V1"
    
    "github.com/shaharby7/Dope/example/pkg/greeter"
)

var controllers = map[string]types.Controller[any]{

    "server1": Controller_server1,

}


var Controller_server1 = controller.NewHTTPServer(
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


