package adminclient

import (
	"fmt"
	"github.com/shaharby7/Dope/pkg/runtime/client"
    
    "github.com/shaharby7/Dope/example/pkg/greeter"
)

var clientDomains map[string]string

func InitiateClient(domains map[string]string) {
	clientDomains = domains
}



var GET_admin__api_ugly_names_set_name = client.CreateTypedClientCall(
    "GET",
	fmt.Sprintf("%s%s", clientDomains["admin"], "/api/ugly-names/set-name"),
    greeter.Greet,
)

var GET_admin__api_ugly_names_unset_name = client.CreateTypedClientCall(
    "GET",
	fmt.Sprintf("%s%s", clientDomains["admin"], "/api/ugly-names/unset-name"),
    greeter.Greet,
)

var GET_admin__api_ugly_names_list = client.CreateTypedClientCall(
    "GET",
	fmt.Sprintf("%s%s", clientDomains["admin"], "/api/ugly-names/list"),
    greeter.Greet,
)
