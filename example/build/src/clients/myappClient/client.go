package myappClient

import (
	"fmt"
	"github.com/shaharby7/Dope/pkg/runtime/client"
    
    "github.com/shaharby7/Dope/example/pkg/greeter"
)

var clientDomains map[string]string

func InitiateClient(domains map[string]string) {
	clientDomains = domains
}



var POST_myapp__api_greet = client.CreateTypedClientCall(
    "POST",
	fmt.Sprintf("%s%s", clientDomains["myapp"], "/api/greet"),
    greeter.Greet,
)
