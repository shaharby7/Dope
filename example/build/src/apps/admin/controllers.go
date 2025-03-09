package main

import (
	"github.com/shaharby7/Dope/pkg/runtime/controller"
	"github.com/shaharby7/Dope/types"
	v1 "github.com/shaharby7/Dope/pkg/entities/V1"
    
    "github.com/shaharby7/Dope/example/pkg/admin"
)

var controllers = map[string]types.Controller[any]{

    "listener": Controller_listener,

}


var Controller_listener = controller.NewHTTPServer(
	[]*types.TypedAction{
            
			controller.CreateTypedAction(
			&v1.ActionConfig{
				Name: "/api/ugly-names/set-names",
				ControllerBinding: map[string]string{
					"method":"POST",
				},
			},
            admin.SetUglyNames,
            ),
            
			controller.CreateTypedAction(
			&v1.ActionConfig{
				Name: "/api/ugly-names/unset-name/:name",
				ControllerBinding: map[string]string{
					"method":"DELETE",
				},
			},
            admin.RemoveUglyName,
            ),
            
			controller.CreateTypedAction(
			&v1.ActionConfig{
				Name: "/api/ugly-names/list",
				ControllerBinding: map[string]string{
					"method":"GET",
				},
			},
            admin.GetUglyNames,
            ),
            
			controller.CreateTypedAction(
			&v1.ActionConfig{
				Name: "/api/ugly-names/echo-header/:name",
				ControllerBinding: map[string]string{
					"method":"GET",
				},
			},
            admin.Echo,
            ),
            
	},
) 


