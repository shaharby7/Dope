package e2e

import (
	"context"
	"testing"

	"github.com/shaharby7/Dope/example/build/src/clients/adminclient"
	"github.com/shaharby7/Dope/types"
	"github.com/stretchr/testify/assert"
)

func TestClientExample(t *testing.T) {
	namesList := []string{"John", "Doe", "Jane"}
	confirmation, _, err := adminclient.POST_admin__api_ugly_names_set_names(
		context.Background(),
		&namesList,
		nil,
	)
	if !confirmation.OK || err != nil {
		assert.Errorf(t, err, "Failed to set names")
	}
	confirmation, _, err = adminclient.DELETE_admin__api_ugly_names_unset_name__name(
		context.Background(),
		nil,
		&types.ActionInputMetadata{
			HTTPServer: &types.HTTPServerRequestConfig{
				Params: map[string]string{"name": "John"},
			},
		},
	)
	if !confirmation.OK || err != nil {
		assert.Errorf(t, err, "Failed to set names")
	}
	output, _, err := adminclient.GET_admin__api_ugly_names_list(
		context.Background(),
		nil,
		nil,
	)
	if err != nil {
		assert.Errorf(t, err, "Failed to set names")
	}
	assert.ElementsMatch(t, *output, []string{"Doe", "Jane"})
}

func TestEchoClient(t *testing.T) {
	output, _, err := adminclient.GET_admin__api_ugly_names_echo_header__name(
		context.Background(),
		nil,
		&types.ActionInputMetadata{
			HTTPServer: &types.HTTPServerRequestConfig{
				Params:  map[string]string{"name": "John"},
				Query:   map[string][]string{"age": {"20"}},
				Headers: map[string][]string{"X-Custom-Header": {"blahblah"}},
			},
		},
	)
	if err != nil {
		assert.Errorf(t, err, "Failed to get echo response")

	}
	assert.Contains(t, output.Headers["X-Custom-Header"], "blahblah")
	assert.Equal(t, output.Params, map[string]string{"name": "John"})
	assert.ElementsMatch(t, output.Query["age"], []string{"20"})
}
