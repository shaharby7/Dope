package e2e

import (
	"context"
	"fmt"

	"github.com/shaharby7/Dope/example/build/src/clients/adminclient"
	"github.com/shaharby7/Dope/example/pkg/greeter"
	"github.com/shaharby7/Dope/pkg/e2e"
	"github.com/stretchr/testify/assert"
)

func E2E_Example(t e2e.ITestProvider) {
	assert.True(t, true)
}

func E2E_Client_Example(t e2e.ITestProvider) {
	adminclient.InitiateClient(
		map[string]string{
			"myapp":"localhost:3000",
		},
	)
	output, _, err := adminclient.GET_admin__api_ugly_names_list(
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
