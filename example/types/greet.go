package types

type GreetInput struct {
	Name string `validate:"required"`
}

type GreetOutput struct {
	Output string
}
