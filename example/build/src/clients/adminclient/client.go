package adminclient

import (
	"github.com/shaharby7/Dope/pkg/runtime/client"
    
    "github.com/shaharby7/Dope/example/pkg/admin"
)


var clientData = client.ParseClientData()



var POST_admin__api_ugly_names_set_names = client.CreateTypedClientCall(
	clientData,
    "POST",
	"/api/ugly-names/set-names",
    admin.SetUglyNames,
)

var DEL_admin__api_ugly_names_unset_name = client.CreateTypedClientCall(
	clientData,
    "DEL",
	"/api/ugly-names/unset-name",
    admin.RemoveUglyName,
)

var GET_admin__api_ugly_names_list = client.CreateTypedClientCall(
	clientData,
    "GET",
	"/api/ugly-names/list",
    admin.GetUglyNames,
)

