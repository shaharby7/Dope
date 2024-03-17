package main


import (
    
	"fmt"

	"github.com/shaharby7/Dope/pkg/deployable"
	"github.com/shaharby7/Dope/pkg/deployable/loggable"

    "context"
    
)

func initiate() *deployable.Deployable {
	myLoggable := &loggable.Loggable{}

	dep, _ := deployable.NewDeployable(
		deployable.DeployableConfig{
			ProjectName:          "app1",
		},
		*myLoggable,
		func(ctx context.Context, err error) { fmt.Println(err) },
        "",
	)

	return dep
}

func main() {
	initiate().Start(context.Background())
}