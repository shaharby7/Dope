package loggable

import (
	"context"
	"errors"
	"fmt"

	"Dope/deployable/constants"
	"Dope/deployable/helpers"
)

type ITarget interface {
	Log(ctx context.Context, eventName string, data string) error
	Get(ctx context.Context) (string, error)
}

type Console struct{}

func NewConsoleTarget() *Console { return &Console{} }

func (c *Console) Log(ctx context.Context, eventName string, data string) error {
	fmt.Printf("%s\t%s\t%s", ctx.Value(constants.CTX_ID), eventName, data)
	return nil
}
func (c *Console) Get(ctx context.Context) (string, error) {
	return "", errors.New("cannot get from target 'Console'")
}

type Loggable struct {
	Targets    map[string]ITarget
	EventTypes map[string]struct{ Targets []string }
	Events     map[string]struct{ EventTypes []string }
	OnError    func(error)
}

func Log(ctx context.Context, eventName string, data string) {
	go func() {
		ok := helpers.VerifyDeployableContext(ctx)
		if !ok {
			panic("Cannot log context that was not produced by deployable")
		}
		l := ctx.Value(constants.LOGGER_REF).(*Loggable)
		for _, eventTypeName := range l.Events[eventName].EventTypes {
			for _, targetName := range l.EventTypes[eventTypeName].Targets {
				err := l.Targets[targetName].Log(ctx, eventTypeName, data)
				if nil != err {
					l.OnError(err)
				}
			}
		}
	}()
}

type ConsoleTarget struct{}

func (target *ConsoleTarget) Log(eventName string, contextId string, data *any) {
	fmt.Printf("%s\t%s::::%s", contextId, eventName, *data)
}

func (target *ConsoleTarget) Get(contextId string) (any, error) {
	return nil, errors.New("cannot get from target ConsoleTarget")
}
