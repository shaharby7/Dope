package admin

import (
	"context"
	"errors"

	"github.com/shaharby7/Dope/pkg/utils"
	"github.com/shaharby7/Dope/pkg/utils/set"

	d "github.com/shaharby7/Dope/types"
)

type NamesList struct {
	set.Set[string] `json:"names"`
}

func NewNamesList() *NamesList {
	return &NamesList{*set.NewSet[string]()}
}

var UGLY_NAMES_REPO *NamesList = NewNamesList()

type Response struct{ OK bool }

var SUCCESS_RESPONSE *Response = &Response{OK: true}
var SUCCESS_META *d.ActionOutputMetadata = &d.ActionOutputMetadata{HTTPServer: &d.HTTPServerResponseConfig{StatusCode: 200}}

func SetUglyNames(
	ctx context.Context,
	names *[]string,
	controllerPayload *d.ActionInputMetadata,
) (
	*Response,
	*d.ActionOutputMetadata,
	error) {
	for _, name := range *names {
		UGLY_NAMES_REPO.Add(name)
	}
	return SUCCESS_RESPONSE, SUCCESS_META, nil
}

func RemoveUglyName(
	ctx context.Context,
	_ *utils.TEmpty,
	controllerPayload *d.ActionInputMetadata,
) (
	*Response,
	*d.ActionOutputMetadata,
	error) {
	serverPayload := controllerPayload.HTTPServer
	if serverPayload == nil {
		return nil, nil, errors.New("did not get server payload")
	}
	params := serverPayload.Params
	UGLY_NAMES_REPO.Remove(params["name"])
	return SUCCESS_RESPONSE, SUCCESS_META, nil
}

func GetUglyNames(
	ctx context.Context,
	_ *utils.TEmpty,
	controllerPayload *d.ActionInputMetadata,
) (*[]string,
	*d.ActionOutputMetadata,
	error) {
	names := UGLY_NAMES_REPO.ToSlice()
	return &names, nil, nil
}

func Echo(
	ctx context.Context,
	_ *utils.TEmpty,
	controllerPayload *d.ActionInputMetadata,
) (*d.HTTPServerRequestConfig,
	*d.ActionOutputMetadata,
	error) {
	return controllerPayload.HTTPServer, &d.ActionOutputMetadata{HTTPServer: &d.HTTPServerResponseConfig{
		Headers: controllerPayload.HTTPServer.Params,
	}}, nil
}
