package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	deployable "Dope/deployable"
	controllable "Dope/deployable/controllable"
	loggable "Dope/deployable/loggable"
)

func Initiate() *deployable.Deployable {
	myLoggable := &loggable.Loggable{
		Targets:    map[string]loggable.ITarget{"console": loggable.NewConsoleTarget()},
		EventTypes: map[string]struct{ Targets []string }{"info": {Targets: []string{"console"}}},
		Events:     map[string]struct{ EventTypes []string }{"my-log": {EventTypes: []string{"info"}}},
		OnError:    func(err error) { fmt.Println(err) },
	}

	dep, _ := deployable.NewDeployable(
		deployable.DeployableConfig{
			ProjectName:          "example",
			RequiredEnvVariables: []string{"PORT", "REDIS_DOMAIN", "REDIS_PORT", "SENSOR_ADDRESS"},
		},
		*myLoggable,
		func(ctx context.Context, err error) { fmt.Println(err) },
		".env", // todo - infer from deployer
	)
	server := controllable.NewHttpServerControllable(
		"controller_example",
		*controllable.NewServerControllableConfig(5000),
	)
	server.RegisterActionable(
		"/greet",
		controllable.NewHttpServerActionable[string, string](
			func(ctx context.Context, s string) (string, error) { return fmt.Sprintf("hi %s!", s), nil },
			func(ctx context.Context, r *http.Request) (string, error) {
				defer r.Body.Close()
				body, err := io.ReadAll(r.Body)
				if err != nil {
					return "", err
				}
				return string(body), nil
			},
			func(ctx context.Context, s string) (controllable.TServerOutput, error) {
				return &controllable.HttpResponse{
					Data: s,
				}, nil
			},
			func(ctx context.Context, err error) (controllable.TServerOutput, error) {
				return &controllable.HttpResponse{
					Data: "I could not read your name!",
				}, nil
			},
		),
	)

	dep.RegisterControllable(
		server,
	)

	return dep
}

func main() {
	Initiate().Start(context.Background())
}
