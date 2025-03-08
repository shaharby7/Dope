package e2e

import (
	"context"
	"fmt"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/shaharby7/Dope/example/build/src/clients/adminclient"
	"github.com/shaharby7/Dope/types"
)

func TestClientExample(t *testing.T) {
	namesList := []string{"John", "Doe", "Jane"}
	confirmation, _, err := adminclient.POST_admin__api_ugly_names_set_names(
		context.Background(),
		&namesList,
		nil,
	)
	if !confirmation.OK || err != nil {
		fmt.Println("Failed to set names")
		t.Fail()
	}
	confirmation, _, err = adminclient.DEL_admin__api_ugly_names_unset_name(
		context.Background(),
		nil,
		&types.ActionInputMetadata{
			HTTPServer: &types.HTTPServerRequestConfig{
				Params: []httprouter.Param{{Key: "name", Value: "John"}},
			},
		},
	)
	if !confirmation.OK || err != nil {
		fmt.Println("Failed to delete name")
		t.Fail()
	}
	output, _, err := adminclient.GET_admin__api_ugly_names_list(
		context.Background(),
		nil,
		nil,
	)
	fmt.Println(err)
	if output != nil {
		fmt.Println(output)
	}
}
