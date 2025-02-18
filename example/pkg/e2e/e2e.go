package e2e

import (
	"context"
	"fmt"

	"github.com/shaharby7/Dope/example/build/src/clients/myappClient"
	"github.com/shaharby7/Dope/example/pkg/greeter"
	"github.com/shaharby7/Dope/pkg/e2e"
	"github.com/stretchr/testify/assert"
)

func E2E_Example(t e2e.ITestProvider) {
	assert.True(t, true)
}

func E2E_Client_Example(t e2e.ITestProvider) {
	myappClient.InitiateClient(
		map[string]string{
			"myapp":"localhost:3000",
		},
	)
	output, _, err := myappClient.POST_myapp__api_greet(
		context.Background(),
		&greeter.GreetInput{
			Name: "Hadas",
		},
		nil,
	)
	fmt.Println(err)
	fmt.Println(output)
	if (output!= nil){
		fmt.Println(output.Greet)
	}
	// assert.Nil(t, err)
	// assert.Equal(
	// 	t,
	// 	"hello Hadas!",
	// 	output.Greet,
	// )
}
