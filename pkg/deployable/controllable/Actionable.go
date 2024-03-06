package controllable

import (
	"context"
)

type FInnerRunActionable[TActionableInput any, TActionableOutput any] func(context.Context, TActionableInput) (TActionableOutput, error)
type FMarshalActionableInput[TControllableInput any, TActionableInput any] func(context.Context, TControllableInput) (TActionableInput, error)
type FMarshalControllableOutput[TActionableOutput any, TControllableOutput any] func(context.Context, TActionableOutput) (TControllableOutput, error)

type Actionable[TControllableInput any, TControllableOutput any] interface {
	Run(context.Context, TControllableInput) (TControllableOutput, error)
}

type SActionable[TControllableInput any, TControllableOutput any, TActionableInput any, TActionableOutput any] struct {
	RunActionable             FInnerRunActionable[TActionableInput, TActionableOutput]
	MarshalActionableInput    FMarshalActionableInput[TControllableInput, TActionableInput]
	MarshalControllableOutput FMarshalControllableOutput[TActionableOutput, TControllableOutput]
	OnError                   func(context.Context, error) (TControllableOutput, error)
}

func NewActionable[TControllableInput any, TControllableOutput any, TActionableInput any, TActionableOutput any](
	RunActionable FInnerRunActionable[TActionableInput, TActionableOutput],
	MarshalActionableInput FMarshalActionableInput[TControllableInput, TActionableInput],
	MarshalControllableOutput FMarshalControllableOutput[TActionableOutput, TControllableOutput],
	OnError func(context.Context, error) (TControllableOutput, error),
) *SActionable[TControllableInput, TControllableOutput, TActionableInput, TActionableOutput] {
	return &SActionable[TControllableInput, TControllableOutput, TActionableInput, TActionableOutput]{
		RunActionable, MarshalActionableInput, MarshalControllableOutput, OnError,
	}
}

func (
	actionable *SActionable[TControllableInput, TControllableOutput, TActionableInput, TActionableOutput],
) Run(
	ctx context.Context,
	input TControllableInput,
) (TControllableOutput, error) {
	marshalledInput, err := actionable.MarshalActionableInput(ctx, input)
	if nil != err {
		return actionable.OnError(ctx, err)
	}
	actionOutput, err := actionable.RunActionable(ctx, marshalledInput)
	if nil != err {
		return actionable.OnError(ctx, err)
	}
	controllerOutput, err := actionable.MarshalControllableOutput(ctx, actionOutput)
	if nil != err {
		return actionable.OnError(ctx, err)
	}
	return controllerOutput, err
}
