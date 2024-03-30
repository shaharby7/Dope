package main

import (
	"context"
	"sync"

	"github.com/shaharby7/Dope/pkg/runtime/controller"
	"github.com/shaharby7/dopeexample/pkg/greeter"

	"github.com/shaharby7/Dope/types"
)

func main() {
	var wg sync.WaitGroup

	greetConf := &types.ActionConfig{
		Name: "/api/greet",
		Bind: &types.HTTPSeverActionConfig{Method: types.POST},
	}
	c := controller.NewHTTPServer(
		types.HTTPServerConfig{Port: 3000},
		[]*types.TypedAction{
			types.CreateTypedAction(greetConf, greeter.Greet),
		},
	)
	wg.Add(1)
	c.Start(context.Background())
	wg.Wait()
}
