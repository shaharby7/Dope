package controllable

import (
	"context"
	"fmt"
	"sync"

	"github.com/shaharby7/Dope/pkg/deployable/constants"
	"github.com/shaharby7/Dope/pkg/deployable/errorable"

	"github.com/google/uuid"
)

type TActionables[TControllableInput any, TControllableOutput any] map[string]Actionable[TControllableInput, TControllableOutput]
type fExecutor[TControllableInput any, TControllableOutput any] func(
	func(actionableName string, input TControllableInput) (TControllableOutput, error),
) error

type Controllable interface {
	Start(ctx context.Context, waitGroup *sync.WaitGroup) error
}

type SControllable[TControllableInput any, TControllableOutput any] struct {
	ControllerType     constants.CONTROLLABLE_TYPES
	ControllerName     string
	executor           fExecutor[TControllableInput, TControllableOutput]
	onActionFatalError func(context.Context, error) (TControllableOutput, error)
	actionables        TActionables[TControllableInput, TControllableOutput]
}

func NewControllable[TControllableInput any, TControllableOutput any](
	ControllerType constants.CONTROLLABLE_TYPES,
	ControllerName string,
	executor fExecutor[TControllableInput, TControllableOutput],
	onActionFatalError func(context.Context, error) (TControllableOutput, error),
) *SControllable[TControllableInput, TControllableOutput] {
	return &SControllable[TControllableInput, TControllableOutput]{
		ControllerType:     ControllerType,
		ControllerName:     ControllerName,
		executor:           executor,
		onActionFatalError: onActionFatalError,
		actionables:        make(TActionables[TControllableInput, TControllableOutput]),
	}
}

func (controllable *SControllable[TControllableInput, TControllableOutput]) RegisterActionable(
	name string, actionable Actionable[TControllableInput, TControllableOutput],
) {
	controllable.actionables[name] = actionable
}

func (controllable *SControllable[TControllableInput, TControllableOutput]) Start(
	ctx context.Context, waitGroup *sync.WaitGroup,
) error {
	ctx = context.WithValue(ctx, constants.CONTROLLABLE_TYPE, controllable.ControllerType)
	ctx = context.WithValue(ctx, constants.CONTROLLABLE_NAME, controllable.ControllerName)
	err := controllable.executor(
		func(actionableName string, input TControllableInput) (TControllableOutput, error) {
			return controllable.execute(ctx, actionableName, input)
		},
	)
	if nil != err {
		//TODO - think what should happen in this case? should the process be killed?
		return nil
	}
	fmt.Printf("Controllable %s of type %s is running\n", controllable.ControllerName, controllable.ControllerType)
	return nil
}

func (controllable *SControllable[TControllableInput, TControllableOutput]) execute(
	ctx context.Context, actionableName string, input TControllableInput,
) (
	TControllableOutput, error,
) {
	ctx = context.WithValue(ctx,
		constants.CTX_ID,
		uuid.NewString(),
	)
	actionableRef := controllable.actionables[actionableName]
	if nil == actionableRef {
		var zero TControllableOutput
		return zero, errorable.NewErrorable(
			errorable.ACTIONABLE_NOT_FOUND,
			fmt.Sprintf("%s actionable not found", actionableName),
		)
	}
	output, err := actionableRef.Run(ctx, input)
	if nil != err {
		return controllable.onActionFatalError(ctx, err)
	}
	return output, nil
}
